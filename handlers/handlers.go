package handlers

import (
	"net/http"

	"github.com/subzero112233/golang-twirp/handlers/rest"
	"github.com/subzero112233/golang-twirp/handlers/twirphandler"
	"github.com/subzero112233/golang-twirp/usecase/playerstats"
)

func NewHandler(usecase playerstats.UseCase, apiType string) *http.ServeMux {
	var handler http.Handler
	if apiType == "rest" {
		handler = rest.NewStatsRestHandler(usecase)
	} else {
		handler = twirphandler.NewTwirpHandler(usecase)
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	return mux
}
