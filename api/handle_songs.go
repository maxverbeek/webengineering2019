package api

import (
	"net/http"
	"strconv"
	"encoding/json"
	"log"

	"webeng/api/model"
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
		data := make([]SongWithLinks, len(songs))

		songroute := s.router.Get("songs_one")
		artistroute := s.router.Get("artists_one")

		for i, song := range songs {
			songlink, _ := songroute.URL("song_id", song.SongId)
			artistlink, _ := artistroute.URL("artist_id", song.ArtistId)

			data[i] = SongWithLinks{
				Song: song,
				Links: Links{
					"self":   songlink.RequestURI(),
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

func (s *server) handleDeleteSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		s.db.DeleteSong(&repository.Query{
			Id: id,
		})

		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *server) handleCreateSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data model.Song

		json.NewDecoder(r.Body).Decode(&data)

		log.Printf("%+v\n", data)

		if s.db.CreateSong(&data) {

			links := make(Links)

			if data.SongId != "" {
				songlink, _ := s.router.Get("songs_one").URL("song_id", data.SongId)
				links["self"] = songlink.RequestURI()
			}

			if data.ArtistId != "" {
				artistlink, _ := s.router.Get("artists_one").URL("artist_id", data.ArtistId)
				links["artist"] = artistlink.RequestURI()
			}

			response := HttpResponse{
				status: http.StatusCreated,
				payload: RestResponse{
					Data: data,
					Success: true,
					Links: links,
				},
			}

			response.Render(w, r)
		} else {
			response := HttpResponse{
				status: http.StatusConflict,
				payload: RestResponse{
					Success: false,
					Message: "failed to create song",
				},
			}

			response.Render(w, r)
		}
	}
}

func (s *server) handleUpdateSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		var data map[string]interface{}

		json.NewDecoder(r.Body).Decode(&data)

		s.db.UpdateSong(&repository.Query{Id: id}, data)

		log.Printf("%+v\n", data)

		w.WriteHeader(http.StatusNoContent)
	}
}
