package benchmark

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/subzero112233/golang-twirp/handlers"
	"github.com/subzero112233/golang-twirp/repository/fakedb"
	"github.com/subzero112233/golang-twirp/usecase/playerstats"
)

var client http.Client

func startServer() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()

	repo := fakedb.NewFakeDB()
	ser := playerstats.LoadService(repo, &logger)

	mux := handlers.NewHandler(ser, "rest")
	http.ListenAndServe(":8000", mux)
}

func init() {
	go startServer()
	client = http.Client{}
}

func BenchmarkRestAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		restAddPlayer(client)
	}
}

func BenchmarkRestGet(b *testing.B) {
	client := http.Client{}
	for n := 0; n < b.N; n++ {
		restGetPlayer(client)
	}
}

func restAddPlayer(client http.Client) {
	data := []byte(`[{"player_name":"Devin Booker","minutes":44.4,"field_goals":18,"field_goal_attempts":27,"three_pointers_made":6,"three_pointer_attempts":10,"free_throws_made":4,"free_throw_attempts":4,"offensive_rebounds":0,"defensive_rebounds":5,"assists":3,"steals":1,"turnovers":3,"personal_fouls":3,"points":55,"team_name":"Phoenix Suns","opponent":"Milwaukee Bucks","game_date":1257894000}]`)

	resp, err := client.Post("http://localhost:8000/stats/reggie_jackson", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("REST AddPlayer request failed with error: %w", err)
	}
	defer resp.Body.Close()
}

func restGetPlayer(client http.Client) {
	resp, err := client.Get("http://localhost:8000/stats/reggie_jackson")
	if err != nil {
		fmt.Println("REST AddPlayer request failed with error: %w", err)

	}
	defer resp.Body.Close()
}
