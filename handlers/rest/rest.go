package rest

import (
	"net/http"
	"regexp"

	"github.com/rs/zerolog/log"
	"github.com/subzero112233/golang-twirp/entity"
	"github.com/subzero112233/golang-twirp/usecase/playerstats"
)

func (h *RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rgx := "^/stats/[a-zA-Z_]{1,100}?$"
	re := regexp.MustCompile(rgx)
	if re.MatchString(r.URL.Path) {
		switch r.Method {
		case "POST":
			h.AddStats(w, r)
			return
		case "GET":
			h.GetStats(w, r)
			return
		default:
			w.Write([]byte("unsupported method"))
			return
		}
	}
	w.Write([]byte("path not found"))
}

type RestHandler struct {
	Usecase playerstats.UseCase
}

func NewStatsRestHandler(usecase playerstats.UseCase) http.Handler {
	return &RestHandler{
		Usecase: usecase,
	}
}

func (h *RestHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	data, err := h.Usecase.GetStats(r.URL.Path)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
}

func (h *RestHandler) AddStats(w http.ResponseWriter, r *http.Request) {
	var input []entity.Stats
	err := Unmarshal(r, &input)
	if err != nil {
		log.Error().Msgf("failed to unmarshal request: %v", err)
		jsonResponse(w, http.StatusBadRequest, "malformed request")
		return

	}

	err = h.Usecase.AddStats(input)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, "success")
}
