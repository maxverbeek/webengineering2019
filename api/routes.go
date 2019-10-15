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
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"mime"
	"net/http"
	"strings"
)

type HttpResponse struct {
	status  int
	payload interface{}
}

// The artist model
// swagger:model Artist
type Artist struct {
	gorm.Model
	ArtistFamiliarity float64 `csv:"artist.familiarity"`
	ArtistHotttnesss  float64 `csv:"artist.hotttnesss"`
	ArtistId          string  `csv:"artist.id"`
	ArtistLatitude    float64 `csv:"artist.latitude"`
	ArtistLocation    int     `csv:"artist.location"`
	ArtistLongitude   float64 `csv:"artist.longitude"`
	ArtistName        string  `csv:"artist.name"`
	ArtistSimilar     float64 `csv:"artist.similar"`
	ArtistTerms       string  `csv:"artist.terms"`
	ArtistTermsFreq   float64 `csv:"artist.terms_freq"`
}

type Release struct {
	gorm.Model
	ReleaseId   int `csv:"release.id"`
	ReleaseName int `csv:"release.name"`
}

// The Song model
// swagger:model Song
type Song struct {
	gorm.Model
	ArtistId                    string
	ReleaseId                   int
	SongArtistMbtags            float64 `csv:"song.artist_mbtags"`
	SongArtistMbtagsCount       float64 `csv:"song.artist_mbtags_count"`
	SongBarsConfidence          float64 `csv:"song.bars_confidence"`
	SongBarsStart               float64 `csv:"song.bars_start"`
	SongBeatsConfidence         float64 `csv:"song.beats_confidence"`
	SongBeatsStart              float64 `csv:"song.beats_start"`
	SongDuration                float64 `csv:"song.duration"`
	SongEndFadeIn               float64 `csv:"song.end_of_fade_in"`
	SongHotttnesss              float64 `csv:"song.hotttnesss"`
	SongId                      string  `csv:"song.id"`
	SongKey                     float64 `csv:"song.key"`
	SongKeyConfidence           float64 `csv:"song.key_confidence"`
	SongLoudness                float64 `csv:"song.loudness"`
	SongMode                    int     `csv:"song.mode"`
	SongModeConfidence          float64 `csv:"song.mode_confidence"`
	SongStartFadeOut            float64 `csv:"song.start_of_fade_out"`
	SongTatumsConfidence        float64 `csv:"song.tatums_confidence"`
	SongTatumsStart             float64 `csv:"song.tatums_start"`
	SongTempo                   float64 `csv:"song.tempo"`
	SongTimeSignature           float64 `csv:"song.time_signature"`
	SongTimeSignatureConfidence float64 `csv:"song.time_signature_confidence"`
	SongTitle                   string  `csv:"song.title"`
	SongYear                    int     `csv:"song.year"`
}

// Determine whether the request `content-type` includes a
// server-acceptable mime-type
//
// Failure should yield an HTTP 415 (`http.StatusUnsupportedMediaType`)
func HasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == "application/json"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}

func (s *server) render(w http.ResponseWriter, r *http.Request, response HttpResponse) {
	// TODO json only for now
	if HasContentType(r, "application/json") {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(response.status)
		json.NewEncoder(w).Encode(response.payload)

	} else if r.Header.Get("Content-Type") == "text/csv" {
		csv, err := gocsv.MarshalString(response.payload)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(response.status)
		w.Write([]byte(fmt.Sprintf("%v", csv)))

	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		log.Print(r.Header.Get("Content-Type"))
	}
}

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
			payload: []struct{ Message string }{struct{ Message string }{"Hello Index"}},
		}

		s.render(w, r, response)
	}
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
//     description: sort by; {hotttnesss}.
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
		//		name := r.URL.Query()["name"]
		//		genre := r.URL.Query()["genre"]
		//		sort := r.URL.Query()["sort"]
		//		limit := r.URL.Query()["limit"]
		//		page := r.URL.Query()["page"]

		response := HttpResponse{
			status:  http.StatusOK,
			payload: nil,
		}

		s.render(w, r, response)
	}
}

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

		var artist Artist

		s.db.Where(&Artist{ArtistId: id}).First(&artist)

		response := HttpResponse{
			status:  http.StatusOK,
			payload: [...]Artist{artist},
		}

		s.render(w, r, response)
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

		s.render(w, r, response)
	}
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

		// artistID := r.URL.Query()["artistid"]
		//		year := r.URL.Query()["year"]
		//		genre := r.URL.Query()["genre"]
		//		sort := r.URL.Query()["sort"]
		//		limit := r.URL.Query()["limit"]
		//		page := r.URL.Query()["page"]

		response := HttpResponse{
			status:  http.StatusOK,
			payload: nil,
		}

		s.render(w, r, response)
	}
}

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

		var song Song

		s.db.Where(&Song{SongId: id}).First(&song)

		response := HttpResponse{
			status:  http.StatusOK,
			payload: [...]Song{song},
		}

		s.render(w, r, response)
	}
}
