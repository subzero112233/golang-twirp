package twirphandler

import (
	"context"
	"net/http"
	"time"

	"github.com/subzero112233/golang-twirp/entity"
	"github.com/subzero112233/golang-twirp/rpc/stats"
	"github.com/subzero112233/golang-twirp/usecase/playerstats"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TwirpHandler struct {
	Usecase playerstats.UseCase
}

func NewTwirpHandler(usecase playerstats.UseCase) http.Handler {
	t := &TwirpHandler{
		Usecase: usecase,
	}
	return stats.NewStatsServiceServer(t)
}

// errors may not be working well. check this by returning an error from the usecase
func (t *TwirpHandler) GetStats(ctx context.Context, input *stats.GetStatsRequest) (*stats.GetStatsResponse, error) {
	var stat entity.Stats
	stat.PlayerName = input.PlayerName
	data, err := t.Usecase.GetStats(stat.PlayerName)
	if err != nil {
		return &stats.GetStatsResponse{}, twirp.InternalError(err.Error())
	}

	statsSlice := convertFromEntity(data)
	return &stats.GetStatsResponse{Stats: statsSlice}, nil
}

// errors may not be working well. check this by returning an error from the usecase
func (t *TwirpHandler) AddStats(ctx context.Context, input *stats.AddStatsRequest) (*stats.AddStatsResponse, error) {
	e := convertToEntity(input.Stats)
	err := t.Usecase.AddStats(e)
	if err != nil {
		return &stats.AddStatsResponse{}, twirp.InternalError(err.Error())
	}
	return &stats.AddStatsResponse{Status: "success"}, nil
}

func convertToEntity(s []*stats.Stats) []entity.Stats {
	statz := make([]entity.Stats, len(s))
	for i, stat := range s {
		statz[i].PlayerName = stat.PlayerName
		statz[i].Minutes = stat.Minutes
		statz[i].FieldGoals = stat.FieldGoals
		statz[i].FieldGoalAttempts = stat.FieldGoalAttempts
		statz[i].ThreePointersMade = stat.ThreePointersMade
		statz[i].ThreePointerAttempts = stat.ThreePointerAttempts
		statz[i].FreeThrowsMade = stat.FreeThrowsMade
		statz[i].FreeThrowAttempts = stat.FreeThrowAttempts
		statz[i].OffensiveRebounds = stat.OffensiveRebounds
		statz[i].DefensiveRebounds = stat.DefensiveRebounds
		statz[i].Assists = stat.Assists
		statz[i].Steals = stat.Steals
		statz[i].Blocks = stat.Blocks
		statz[i].Turnovers = stat.Turnovers
		statz[i].PersonalFouls = stat.PersonalFouls
		statz[i].Points = stat.Points
		statz[i].Team = stat.Team
		statz[i].Opponent = stat.Opponent
		statz[i].GameDate = stat.GameDate.AsTime().Unix()
	}
	return statz
}

func convertFromEntity(e []entity.Stats) []*stats.Stats {
	var statz []*stats.Stats
	for _, stat := range e {
		statz = append(statz, &stats.Stats{PlayerName: stat.PlayerName,
			Minutes:              stat.Minutes,
			FieldGoals:           stat.FieldGoals,
			FieldGoalAttempts:    stat.FieldGoalAttempts,
			ThreePointersMade:    stat.ThreePointersMade,
			ThreePointerAttempts: stat.ThreePointerAttempts,
			FreeThrowsMade:       stat.FreeThrowsMade,
			FreeThrowAttempts:    stat.FreeThrowAttempts,
			OffensiveRebounds:    stat.OffensiveRebounds,
			DefensiveRebounds:    stat.DefensiveRebounds,
			Assists:              stat.Assists,
			Steals:               stat.Steals,
			Blocks:               stat.Blocks,
			Turnovers:            stat.Turnovers,
			PersonalFouls:        stat.PersonalFouls,
			Points:               stat.Points,
			Team:                 stat.Team,
			Opponent:             stat.Opponent,
			GameDate:             timestamppb.New(time.Unix(stat.GameDate, 0)),
		})
	}
	return statz
}
