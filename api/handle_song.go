package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"webeng/api/repository"
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

		song := s.newdb.FindSong(&repository.Query{Id: id})

		response := HttpResponse{
			status:  http.StatusOK,
			payload: song,
		}

		response.Render(w, r)
	}
}
