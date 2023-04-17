package admin

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bufbuild/connect-go"

	"github.com/pubgolf/pubgolf/api/internal/lib/models"
	apiv1 "github.com/pubgolf/pubgolf/api/internal/lib/proto/api/v1"
	"github.com/pubgolf/pubgolf/api/internal/lib/telemetry"
)

// CreatePlayer registers a player in the given event, returning the created player object. This method is idempotent, so if the player is already registered the request will still succeed.
func (s *Server) CreatePlayer(ctx context.Context, req *connect.Request[apiv1.CreatePlayerRequest]) (*connect.Response[apiv1.CreatePlayerResponse], error) {
	telemetry.AddRecursiveAttribute(&ctx, "event.key", req.Msg.EventKey)

	eventID, err := s.dao.EventIDByKey(ctx, req.Msg.EventKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	playerParams := models.PlayerParams{
		Name: req.Msg.PlayerData.Name,
	}
	err = playerParams.ScoringCategory.FromProtoEnum(req.Msg.PlayerData.ScoringCategory)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	player, err := s.dao.CreatePlayer(ctx, eventID, playerParams)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	cat, err := player.ScoringCategory.ProtoEnum()
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	return connect.NewResponse(&apiv1.CreatePlayerResponse{
		Player: &apiv1.Player{
			Id: player.ID.String(),
			Data: &apiv1.PlayerData{
				Name:            player.Name,
				ScoringCategory: cat,
			},
		},
	}), nil
}
