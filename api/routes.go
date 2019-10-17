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
	"github.com/gocarina/gocsv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"mime"
	"net/http"
	"strings"

	"webeng/api/model"
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

func (response *HttpResponse) Render(w http.ResponseWriter, r *http.Request) {
	if HasContentType(r, "application/json") {
		response.RenderJSON(w, r)
	} else if r.Header.Get("Content-Type") == "text/csv" {
		response.RenderCSV(w, r)
	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		log.Print(r.Header.Get("Content-Type"))
	}
}

func (response *HttpResponse) RenderJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.status)
	json.NewEncoder(w).Encode(response.payload)
}

func (response *HttpResponse) RenderCSV(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(response.status)

	var err error

	switch response.payload.(type) {

	// if the payload is already a slice, CSV it
	case []interface{}:
		err = gocsv.Marshal(response.payload.([]interface{}), w)
		break

	// some type assertions for single types
	case model.Song:
		err = gocsv.Marshal([]model.Song{response.payload.(model.Song)}, w)
		break
		
	case model.Artist:
		err = gocsv.Marshal([]model.Artist{response.payload.(model.Artist)}, w)
		break

	case Song:
		err = gocsv.Marshal([]Song{response.payload.(Song)}, w)
		break

	case Artist:
		err = gocsv.Marshal([]Artist{response.payload.(Artist)}, w)
		break

	// in case the payload is not a slice, we return a single object
	// so we convert it to a slice. hacky hacky ;D
	// doesn't work tho :(
	default:
		log.Println("using CSV hacky hack")
		var test []interface{} = make([]interface{}, 1)
		test[0] = response.payload
		err = gocsv.Marshal(test, w)
		break
	}

	if err != nil {
		log.Print(err)
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

		response.Render(w, r)
	}
}
