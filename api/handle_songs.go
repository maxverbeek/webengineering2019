package api

import (
	"net/http"

	"strconv"
	"webeng/api/repository"
)

// swagger:operation GET /songs Songs
// ---
// description: Gets a list of songs
// parameters:
//   - in: query
//     name: artist_id
//     description: Filter by artist ID
//     required: false
//     type: string
//   - in: query
//     name: year
//     description: Filter by year released.
//     required: false
//     type: integer
//   - in: query
//     name: genre
//     description: Filter by artist genre.
//     required: false
//     type: string
//   - in: query
//     name: sort
//     description: sort by; {hotttnesss}.
//     required: false
//     type: string
//   - in: query
//     name: limit
//     description: The number of songs per page.
//     required: false
//     type: integer
//   - in: query
//     name: page
//     description: Retrieves the nth page of `limit`.
//     required: false
//     type: integer
// responses:
//   200:
//     description: Yields list of songs.
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
func (s *server) handleSongs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		songId := r.URL.Query().Get("songid")
		year, _ := strconv.Atoi(r.URL.Query().Get("year"))
		//		genre := r.URL.Query()["genre"]
		//		sort := r.URL.Query()["sort"]
		//		limit := r.URL.Query()["limit"]
		//		page := r.URL.Query()["page"]

		songs := s.newdb.FindSongs(&repository.Query{
			Id: songId,
			Year: year,
		})

		response := HttpResponse{
			status:  http.StatusOK,
			payload: songs,
		}

		response.Render(w, r)
	}
}
