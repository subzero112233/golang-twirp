package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/subzero112233/golang-twirp/rpc/stats"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	endpoint := "http://localhost:8000"
	client := stats.NewStatsServiceProtobufClient(endpoint, &http.Client{})
	header := make(http.Header)
	ctx := context.Background()

	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		return
	}

	var statz []*stats.Stats

	statz = append(statz, &stats.Stats{
		PlayerName:           "Devin Booker",
		Minutes:              44,
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
		Team:                 "Phoenix Suns",
		Opponent:             "Milwaukee_Bucks",

		GameDate: timestamppb.Now(),
	})

	respAdd, err := client.AddStats(ctx, &stats.AddStatsRequest{
		Stats: statz})

	if err != nil {
		fmt.Println("AddStats returned an error :", err)
	}
	fmt.Println(respAdd)

	respGet, err := client.GetStats(ctx, &stats.GetStatsRequest{
		PlayerName: "Nicolas Batum"})

	if err != nil {
		fmt.Println("GetStats returned an error :", err)
	}
	fmt.Println(respGet)

}
