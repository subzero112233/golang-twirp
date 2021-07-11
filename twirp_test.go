package benchmark

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/subzero112233/golang-twirp/handlers"
	"github.com/subzero112233/golang-twirp/repository/fakedb"
	"github.com/subzero112233/golang-twirp/rpc/stats"
	"github.com/subzero112233/golang-twirp/usecase/playerstats"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func startTwirpServer() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()

	repo := fakedb.NewFakeDB()
	ser := playerstats.LoadService(repo, &logger)

	mux := handlers.NewHandler(ser, "twirp")
	http.ListenAndServe(":8001", mux)
}

var (
	ctx       context.Context
	rpcclient stats.StatsService
	err       error
)

func init() {
	go startTwirpServer()
	rpcclient = stats.NewStatsServiceProtobufClient("http://localhost:8001", &http.Client{})
	header := make(http.Header)
	ctx = context.Background()

	ctx, err = twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		return
	}
}

func BenchmarkTwirpAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twirpAddPlayer(rpcclient, ctx)
	}
}

func BenchmarkTwirpGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twirpGetPlayer(rpcclient, ctx)
	}
}

func twirpAddPlayer(client stats.StatsService, ctx context.Context) {
	var statz []*stats.Stats
	statz = append(statz, &stats.Stats{
		PlayerName:           "Devin Booker",
		Minutes:              44.4,
		FieldGoals:           18,
		FieldGoalAttempts:    27,
		ThreePointersMade:    6,
		ThreePointerAttempts: 10,
		FreeThrowsMade:       4,
		FreeThrowAttempts:    4,
		OffensiveRebounds:    0,
		DefensiveRebounds:    5,
		Assists:              3,
		Steals:               1,
		Blocks:               0,
		Turnovers:            3,
		PersonalFouls:        3,
		Points:               55,
		Team:                 stats.Teams_Phoenix_Suns,
		Opponent:             stats.Teams_Milwaukee_Bucks,
		GameDate:             timestamppb.Now(),
	})

	_, err := client.AddStats(ctx, &stats.AddStatsRequest{
		Stats: statz})

	if err != nil {
		fmt.Println("Twirp AddPlayer request failed with error: %w", err)
	}
}

func twirpGetPlayer(client stats.StatsService, ctx context.Context) {
	_, err := client.GetStats(ctx, &stats.GetStatsRequest{
		PlayerName: "Reggie Jackson"})

	if err != nil {
		fmt.Println("Twirp GetPlayer request failed with error: %w", err)
	}
}
