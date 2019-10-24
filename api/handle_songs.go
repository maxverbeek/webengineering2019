package api

import (
	"net/http"
	"strconv"
	"log"

	"webeng/api/model"
	"webeng/api/repository"
)

type SongWithLinks struct {
	model.Song
	Links `json:"links,omitempty"` // handlerutils.go
}

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
//     name: name
//     description: Filter by song title.
//     required: false
//     type: string
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

		year, _ := strconv.Atoi(r.URL.Query().Get("year"))
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		songs, total := s.db.FindSongs(&repository.Query{
			Id:      r.URL.Query().Get("songid"),
			OtherId: r.URL.Query().Get("artist"),
			Genre:   r.URL.Query().Get("genre"),
			Name:    r.URL.Query().Get("name"),
			Year:    year,

			Sort:  r.URL.Query().Get("sort"),
			Page:  page,
			Limit: limit,
		})
		log.Println(r.URL.Query().Get("name"))
		data := make([]SongWithLinks, len(songs))

		songroute := s.router.Get("songs_one")
		artistroute := s.router.Get("artists_one")

		for i, song := range songs {
			songlink, _ := songroute.URL("song_id", song.SongId)
			artistlink, _ := artistroute.URL("artist_id", song.ArtistId)

			data[i] = SongWithLinks{
				Song: song,
				Links: Links{
					"self": songlink.RequestURI(),
					"artist": artistlink.RequestURI(),
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
