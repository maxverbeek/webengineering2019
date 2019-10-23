package api

import (
	"net/http"

	"webeng/api/repository"

	"github.com/gorilla/mux"
)

// swagger:operation GET /songs/{song_id} Song
// ---
// description: Gets a song by the given ID.
// parameters:
//   - in: path
//     name: song_id
//     description: ID of the song.
//     required: true
//     type: string
// responses:
//   200:
//     description: Yields song by ID.
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
//     description: Could not find the song by ID.
func (s *server) handleSong() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		song := s.db.FindSong(&repository.Query{Id: id})

		var response HttpResponse

		if song != nil {

			artisturl, _ := s.router.Get("artists_one").URL("artist_id", song.ArtistId)

			response = HttpResponse{
				status: http.StatusOK,
				payload: RestResponse{
					Success: true,
					Data:    song,
					Links: map[string]string{
						"self":   r.URL.RequestURI(),
						"artist": artisturl.RequestURI(),
					},
				},
			}
		} else {
			response = HttpResponse{
				status: http.StatusNotFound,
				payload: RestResponse{
					Success: false,
					Message: "song not found",
				},
			}
		}

		response.Render(w, r)
	}
}
