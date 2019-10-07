package api

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Hello"))
	}
}
