package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/subzero112233/golang-twirp/handlers"
	"github.com/subzero112233/golang-twirp/repository/fakedb"
	"github.com/subzero112233/golang-twirp/usecase/playerstats"
)

var (
	ser     *playerstats.Service
	mux     *http.ServeMux
	apiType string
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()

	repo := fakedb.NewFakeDB()
	ser = playerstats.LoadService(repo, &logger)

	apiType = os.Getenv("API_TYPE")
	if apiType == "" {
		apiType = "twirp"
	}

	mux = handlers.NewHandler(ser, apiType)
	http.ListenAndServe(":8000", mux)

	logger.Debug().Msgf("started %s API", apiType)
}
