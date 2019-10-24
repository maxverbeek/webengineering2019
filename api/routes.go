// Music API
//
// This is an api for accessing an manipulating a music database.
//
// Version: 0.1.0
// basePath: /api/v1
//
// Produces:
// - application/json
// - text/csv
//
// swagger:meta
package api

import (
	"net/http"
)

func (s *server) routes() {

	s.router.StrictSlash(true)

	api := s.router.PathPrefix("/api/v1").Subrouter()

	api.Path("/artists").
		Methods(http.MethodGet).
		HandlerFunc(s.handleArtists()).
		Name("artists_all")

	api.Path("/artists/{artist_id}").
		Methods(http.MethodGet).
		HandlerFunc(s.handleArtist()).
		Name("artists_one")

	api.Path("/artists/{artist_id}/stats").
		Methods(http.MethodGet).
		HandlerFunc(s.handleArtistStats()).
		Name("artists_stats")

	api.Path("/songs").
		Methods(http.MethodGet).
		HandlerFunc(s.handleSongs()).
		Name("songs_all")

	api.Path("/songs/{song_id}").
		Methods(http.MethodGet).
		HandlerFunc(s.handleSong()).
		Name("songs_one")

	s.router.
		PathPrefix("/").
		Methods(http.MethodGet).
		Handler(http.FileServer(http.Dir("./static/")))
}

// swagger:operation GET / index
//
// ---
// responses:
//   200:
//     description: successful operation
func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HttpResponse{
			status:  http.StatusOK,
			payload: struct{ Message string }{"Hello index."},
		}

		response.Render(w, r)
	}
}
