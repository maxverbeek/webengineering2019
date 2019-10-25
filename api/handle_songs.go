package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"webeng/api/model"
	"webeng/api/repository"

	"github.com/gorilla/mux"
)

// swagger:operation GET /songs/{song_id} Song Read
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
//         $ref: 'song200.json'
//       text/csv:
//         $ref: 'song200.csv'
//   404:
//     description: Could not find the song by ID.
//     examples:
//       application/json:
//         $ref: 'song404.json'
//       text/csv:
//         $ref: 'song404.csv'
func (s *server) handleSong() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		song := s.db.FindSong(&repository.Query{Id: id})

		artist := s.db.FindArtist(&repository.Query{Id: song.ArtistId})

		var response HttpResponse

		if song != nil {

			artisturl, _ := s.router.Get("artists_one").URL("artist_id", song.ArtistId)

			call := s.service.Search.List("snippet").
				Q(artist.ArtistName + " " + song.SongTitle).
				MaxResults(1)
			yresponse, err := call.Do()
			videoLink := ""

			if err != nil {
				log.Println(err)
			}

			for _, item := range yresponse.Items {
				switch item.Id.Kind {
				case "youtube#video":
					videoLink = "https://youtube.com/embed/" + item.Id.VideoId
				}
			}

			response = HttpResponse{
				status: http.StatusOK,
				payload: RestResponse{
					Success: true,
					Data:    song,
					Links: map[string]string{
						"self":          r.URL.RequestURI(),
						"artist":        artisturl.RequestURI(),
						"youtube_embed": videoLink,
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
//     description: sort by `duration`, `hotttnesss`, `id`, `title`, `tempo` or
//       `year`. Will always be descending.
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
//         $ref: 'songs200.json'
//       text/csv:
//         $ref: 'songs200.csv'
func (s *server) handleSongs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		year, _ := strconv.Atoi(r.URL.Query().Get("year"))
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		songs, total := s.db.FindSongs(&repository.Query{
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

// swagger:operation DELETE /songs/{song_id} Song Delete
// ---
// description: Attempts to Delete the song with the given ID.
// parameters:
//   - in: path
//     name: song_id
//     description: ID of the song.
//     required: true
//     type: string
// responses:
//   204:
//     description: Song was deleted.
func (s *server) handleDeleteSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		s.db.DeleteSong(&repository.Query{
			Id: id,
		})

		w.WriteHeader(http.StatusNoContent)
	}
}

// swagger:operation POST /songs Song Create
// ---
// description: Attempts to add the given song to the database. Only `id` and
//   `artist_id` are required, missing or extra fields are ignored.
// content:
//   application/json:
//     $ref: 'rickroll.json'
// responses:
//   400:
//     description: The `id` or `artist_id` is missing.
//     examples:
//       application/json:
//         $ref: 'postSong400.json'
//       text/csv:
//         $ref: 'postSong400.csv'
//   201:
//     description: Song was successfully created.
//     examples:
//       application/json:
//         $ref: 'postSong201.json'
//       text/csv:
//         $ref: 'postSong201.csv'
//   409:
//     description: Song with `id` already exists.
//     examples:
//       application/json:
//         $ref: 'postSong409.json'
//       text/csv:
//         $ref: 'postSong409.csv'
func (s *server) handleCreateSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data model.Song

		json.NewDecoder(r.Body).Decode(&data)

		if data.SongId == "" || data.ArtistId == "" {
			response := HttpResponse{
				status: http.StatusBadRequest,
				payload: RestResponse{
					Success: false,
					Message: "song ID and artist ID are mandatory",
				},
			}

			response.Render(w, r)
			return
		}

		log.Printf("%+v\n", data)

		if s.db.CreateSong(&data) {

			links := make(Links)

			songlink, _ := s.router.Get("songs_one").URL("song_id", data.SongId)
			artistlink, _ := s.router.Get("artists_one").URL("artist_id", data.ArtistId)
			links["self"] = songlink.RequestURI()
			links["artist"] = artistlink.RequestURI()

			response := HttpResponse{
				status: http.StatusCreated,
				payload: RestResponse{
					Data:    &data,
					Success: true,
					Links:   links,
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

// swagger:operation PUT /songs/{song_id} Song Update
// ---
// description: Attempts to update the given song in the database.
// parameters:
//   - in: path
//     name: song_id
//     description: ID of the song.
//     required: true
//     type: string
// content:
//   application/json:
//     $ref: 'rollrick.json'
// responses:
//   204:
//     description: Song was updated.
//   404:
//     description: Song {song_id} does not exist.
//     examples:
//       application/json:
//         $ref: putSong404.json
//       text/csv:
//         $ref: putSong404.csv
func (s *server) handleUpdateSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		var data map[string]interface{}

		json.NewDecoder(r.Body).Decode(&data)

		if !s.db.UpdateSong(&repository.Query{Id: id}, data) {
			response := HttpResponse{
				status: http.StatusNotFound,
				payload: RestResponse{
					Success: false,
					Message: "song not found",
				},
			}

			response.Render(w, r)
			return
		}

		log.Printf("%+v\n", data)

		w.WriteHeader(http.StatusNoContent)
	}
}
