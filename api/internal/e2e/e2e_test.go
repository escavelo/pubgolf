package e2e

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"

	"github.com/fatih/color"

	"github.com/pubgolf/pubgolf/api/internal/lib/dbtest"
	"github.com/pubgolf/pubgolf/api/internal/lib/testguard"
)

const testEventKey = "test-event-key"

type LogWriter struct {
	logger *log.Logger
}

func NewPrefixLogWriter(p string) *LogWriter {
	l := log.New(log.Writer(), fmt.Sprintf("[%s] ", p), log.Flags())
	l.SetFlags(0)

	return NewLogWriter(l)
}

func NewLogWriter(l *log.Logger) *LogWriter {
	return &LogWriter{
		logger: l,
	}
}

func (lw LogWriter) Write(p []byte) (int, error) {
	lw.logger.Print(string(p))

	return len(p), nil
}

func TestMain(m *testing.M) {
	testguard.E2ETest()

	color.NoColor = false

	dbURL, dbCleanupFn := dbtest.NewURL("pubgolf-e2e", false)

	runAPIMigrator(dbURL)
	seedDB(dbURL)

	serverCleanup := runAPIServer(dbURL)

	ret := m.Run()

	serverCleanup()
	dbCleanupFn()
	os.Exit(ret)
}

func seedDB(dbURL string) {
	db, err := sql.Open("pgx", dbURL)
	guard(err, "open DB connection")

	_, err = db.Exec("INSERT INTO events (key, starts_at) VALUES ($1, NOW() + '1 day')", testEventKey)
	guard(err, "seed DB")
}

func runAPIMigrator(dbURL string) {
	migrator := exec.Command(
		"doppler", "run",
		"--project", "pubgolf-api-server",
		"--config", "e2e",
		"--preserve-env",
		"--",
		"go", "run", "../../cmd/pubgolf-api-server", "--run-migrations",
	)

	migrator.Env = append(os.Environ(), "PUBGOLF_APP_DATABASE_URL="+dbURL)

	migratorLog := NewPrefixLogWriter(color.RedString("Migrator"))
	migrator.Stdout = migratorLog
	migrator.Stderr = migratorLog
	migrator.Stdin = os.Stdin

	guard(migrator.Run(), "run API migrator")
}

func runAPIServer(dbURL string) func() {
	server := exec.Command(
		"doppler", "run",
		"--project", "pubgolf-api-server",
		"--config", "e2e",
		"--preserve-env",
		"--",
		"go", "run", "../../cmd/pubgolf-api-server",
	)

	server.Env = append(os.Environ(), "PUBGOLF_APP_DATABASE_URL="+dbURL)
	server.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	serverLog := NewPrefixLogWriter(color.BlueString("API Server"))
	server.Stdout = serverLog
	server.Stderr = serverLog
	server.Stdin = os.Stdin

	guard(server.Start(), "start API server")

	serverStarted := false

	for range 60 {
		res, _ := http.Get("http://localhost:3100/health-check") //nolint:noctx
		if res != nil {
			res.Body.Close()

			if res.StatusCode == http.StatusOK {
				serverStarted = true

				break
			}
		}

		time.Sleep(1 * time.Second)
	}

	if !serverStarted {
		log.Panicln("API server startup timed out")
	}

	return func() {
		pgid, err := syscall.Getpgid(server.Process.Pid)
		guard(err, "get process group ID for API server")
		guard(syscall.Kill(-pgid, syscall.SIGINT), "send SIGINT to API server")
	}
}

// guard logs and exits on error.
func guard(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err.Error())
	}
}
