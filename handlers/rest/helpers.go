package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Unmarshal(r *http.Request, t interface{}) error {
	x, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(x, &t)
	return err
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	r, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(r)
}
