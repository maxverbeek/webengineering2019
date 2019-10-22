package api

import (
	"net/http"

	"webeng/api/repository"

	"github.com/gorilla/mux"
)

// swagger:operation GET /artists/{artist_id} Artist
// ---
// description: Gets an artist by the given ID.
// parameters:
//   - in: path
//     name: artist_id
//     description: ID of the artist.
//     required: true
//     type: string
// responses:
//   200:
//     description: Yields artist by ID.
//     examples:
//       application/json:
//         - to: do
//           when: we
//           can: auto
//           gen: this
//           stupid: shit
//       text/csv: |
//         to,do,when,we,can
//         auto,gen,this,stupid,shit
//         auto,gen,this,stupid,shit
//   404:
//     description: Could not find the Artist by ID.
func (s *server) handleArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["artist_id"]

		artist := s.db.FindArtist(&repository.Query{Id: id})

		var response HttpResponse

		if artist != nil {
			response = HttpResponse{
				status: http.StatusOK,
				payload: RestResponse{
					Success: true,
					Data:    artist,
					Links:   map[string]string{"self": r.URL.RequestURI()},
				},
			}
		} else {
			response = HttpResponse{
				status: http.StatusNotFound,
				payload: RestResponse{
					Success: false,
					Message: "artist not found",
				},
			}
		}

		response.Render(w, r)
	}
}

// swagger:operation GET /artists/{artist_id}/stats ArtistStats
// ---
// description: Gets the statistics of an artist by the given ID.
// parameters:
//   - in: path
//     name: artist_id
//     description: ID of the artist.
//     required: true
//     type: string
// responses:
//   200:
//     description: Yields artist's statistics by ID.
//     examples:
//       application/json:
//         - to: do
//           when: we
//           can: auto
//           gen: this
//           stupid: shit
//       text/csv: |
//         to,do,when,we,can
//         auto,gen,this,stupid,shit
//         auto,gen,this,stupid,shit
//   404:
//     description: Could not find the artist by ID.
func (s *server) handleArtistStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//		id := mux.Vars(r)["artist_id"]

		response := HttpResponse{
			status:  http.StatusOK,
			payload: nil,
		}

		response.Render(w, r)
	}
}
