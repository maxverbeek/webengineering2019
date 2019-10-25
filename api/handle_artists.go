package api

import (
	"net/http"
	"strconv"

	"webeng/api/model"
	"webeng/api/repository"
)

type ArtistWithLinks struct {
	model.Artist
	Links `json:"links,omitempty"` // handlerutils.go
}

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
//     description: sort by `familiarity`, `hotttnesss`, `id`, `name`, `similar`.
//       Will always be descending.
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
//         $ref: 'artists200.json'
//       text/csv:
//         $ref: 'artists200.csv'
func (s *server) handleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		artists, total := s.db.FindArtists(&repository.Query{
			Name:  r.URL.Query().Get("name"),
			Genre: r.URL.Query().Get("genre"),
			Sort:  r.URL.Query().Get("sort"),
			Limit: limit,
			Page:  page,
		})

		data := make([]ArtistWithLinks, len(artists))

		artistroute := s.router.Get("artists_one")

		for i, artist := range artists {
			artistlink, _ := artistroute.URL("artist_id", artist.ArtistId)

			data[i] = ArtistWithLinks{
				Artist: artist,
				Links: Links{
					"self": artistlink.RequestURI(),
				},
			}
		}

		response := HttpResponse{
			status: http.StatusOK,
			payload: RestResponse{
				Data:    data,
				Success: true,
				Links:   getPaginationLinks(*r.URL, total, page, limit),
			},
		}

		response.Render(w, r)
	}
}
