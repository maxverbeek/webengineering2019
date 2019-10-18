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
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/artists", s.handleArtists())
	s.router.HandleFunc("/artists/{artist_id}", s.handleArtist())
	s.router.HandleFunc("/artists/{artist_id}/stats", s.handleArtistStats())
	s.router.HandleFunc("/songs", s.handleSongs())
	s.router.HandleFunc("/songs/{song_id}", s.handleSong())
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
