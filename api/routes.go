package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

type HttpResponse struct {
	status  int
	payload interface{}
}

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

type Song struct {
	gorm.Model
	ArtistId                     string
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

func (s *server) render(w http.ResponseWriter, r *http.Request, response HttpResponse) {
	// TODO json only for now
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.status)
	json.NewEncoder(w).Encode(response.payload)
}

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/artists", s.handleArtists())
	s.router.HandleFunc("/artists/{artist_id}", s.handleArtist())
	s.router.HandleFunc("/artists/{artist_id}/stats", s.handleArtistStats())
	s.router.HandleFunc("/songs", s.handleSongs())
	s.router.HandleFunc("/songs/{song_id}", s.handleSong())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HttpResponse{
			status:  http.StatusOK,
			payload: struct{ Message string }{"Hello index."},
		}

		s.render(w, r, response)
	}
}

func (s *server) handleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//		name := r.URL.Query()["name"]
		//		genre := r.URL.Query()["genre"]
		//		sort := r.URL.Query()["sort"]
		//		limit := r.URL.Query()["limit"]
		//		page := r.URL.Query()["page"]

		response := HttpResponse{
			status:  http.StatusOK,
			payload: struct{ Message string }{"Hello Artists."},
		}

		s.render(w, r, response)
	}
}

func (s *server) handleArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["artist_id"]

		var artist Artist

		s.db.Where(&Artist{ArtistId: id}).First(&artist)

		response := HttpResponse{
			status:  http.StatusOK,
			payload: artist,
		}

		s.render(w, r, response)
	}
}

func (s *server) handleArtistStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//		id := mux.Vars(r)["artist_id"]

		response := HttpResponse{
			status:  http.StatusOK,
			payload: struct{ Message string }{"Hello Stats."},
		}

		s.render(w, r, response)
	}
}

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
			payload: struct{ Message string }{"Hello Songs."},
		}

		s.render(w, r, response)
	}
}

func (s *server) handleSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["song_id"]

		var song Song

		s.db.Where(&Song{SongId: id}).First(&song)

		response := HttpResponse{
			status:  http.StatusOK,
			payload: song,
		}

		s.render(w, r, response)
	}
}
