package api

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	status  int
	payload interface{}
}

func (s *server) render(w http.ResponseWriter, r *http.Request, response HttpResponse) {
	// TODO json only for now
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.status)
	json.NewEncoder(w).Encode(response.payload)
}

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HttpResponse{
			status:  http.StatusOK,
			payload: struct{ Message string }{"Hello index."},
		}

		s.render(w, r, response)
	}
}
