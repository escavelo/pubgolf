package server

import (
	"context"

	"github.com/escavelo/pubgolf/api/lib/db"
	pg "github.com/escavelo/pubgolf/api/proto/pubgolf"
)

func (server *APIServer) GetSchedule(ctx context.Context,
	req *pg.GetScheduleRequest) (*pg.GetScheduleReply, error) {
	tx, eventID, _, err := validateAuthenticatedRequest(server, ctx, &req.EventKey)
	if err != nil {
		return nil, err
	}

	venueList, err := db.GetScheduleForEvent(tx, &eventID)
	if err != nil {
		tx.Rollback()
		return nil, temporaryServerError(err)
	}

	tx.Commit()
	return &pg.GetScheduleReply{VenueList: &venueList}, nil
}

func (server *APIServer) GetScores(ctx context.Context,
	req *pg.GetScoresRequest) (*pg.GetScoresReply, error) {
	tx, eventID, _, err := validateAuthenticatedRequest(server, ctx, &req.EventKey)
	if err != nil {
		return nil, err
	}

	bestOf9, err := db.GetScoreboardBestOf9(tx, &eventID)
	if err != nil {
		return nil, temporaryServerError(err)
	}

	bestOf5, err := db.GetScoreboardBestOf5(tx, &eventID)
	if err != nil {
		return nil, temporaryServerError(err)
	}

	incomplete, err := db.GetScoreboardIncomplete(tx, &eventID)
	if err != nil {
		return nil, temporaryServerError(err)
	}

	tx.Commit()
	return &pg.GetScoresReply{
		ScoreLists: []*pg.ScoreList{
			{
				Label:  "Best of 9",
				Scores: bestOf9,
			},
			{
				Label:  "Best of 5",
				Scores: bestOf5,
			},
			{
				Label:  "Inc",
				Scores: incomplete,
			},
		},
	}, nil
}

func (server *APIServer) GetScoresForPlayer(ctx context.Context,
	req *pg.GetScoresForPlayerRequest) (*pg.GetScoresForPlayerReply, error) {
	tx, _, _, err := validateAuthenticatedRequest(server, ctx,
		&req.EventKey)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &pg.GetScoresForPlayerReply{}, nil
}
