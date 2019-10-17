package api

import (
	"net/http"

	"webeng/api/repository"
)

// swagger:operation GET /artists Artists
// ---
// description: Gets a list of artists.
// parameters:
//   - in: query
//     name: name
//     description: Filter by name of artist.
//     required: false
//     type: string
//   - in: query
//     name: genre
//     description: Filter by artist genre.
//     required: false
//     type: string
//   - in: query
//     name: sort
//     description: sort by `hotttnesss`.
//     required: false
//     type: string
//   - in: query
//     name: limit
//     description: The number of artists per page.
//     required: false
//     type: integer
//   - in: query
//     name: page
//     description: Retrieves the nth page of `limit`.
//     required: false
//     type: integer
// responses:
//   200:
//     description: Yields list of artists.
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
func (s *server) handleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// sort := r.URL.Query()["sort"]
		// limit := r.URL.Query()["limit"]
		// page := r.URL.Query()["page"]

		artists := s.newdb.FindArtists(&repository.Query{
			Name:  r.URL.Query().Get("name"),
			Genre: r.URL.Query().Get("genre")
		})

		response := HttpResponse{
			status:  http.StatusOK,
			payload: artists,
		}

		response.Render(w, r)
	}
}
