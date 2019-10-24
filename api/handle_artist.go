package api

import (
	"net/http"
	"strconv"

	"webeng/api/repository"

	"github.com/gorilla/mux"
	"github.com/montanaflynn/stats"
)

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

		artist := s.db.FindArtist(&repository.Query{Id: id})

		var response HttpResponse

		if artist != nil {
			songsurl, _ := s.router.Get("songs_all").URL()
			urlparam := songsurl.Query()
			urlparam.Set("artist", id)
			songsurl.RawQuery = urlparam.Encode()

			statsurl, _ := s.router.Get("artists_stats").URL("artist_id", id)

			response = HttpResponse{
				status: http.StatusOK,
				payload: RestResponse{
					Success: true,
					Data:    artist,
					Links: map[string]string{
						"self":  r.URL.RequestURI(),
						"songs": songsurl.RequestURI(),
						"stats": statsurl.RequestURI(),
					},
				},
			}
		} else {
			response = HttpResponse{
				status: http.StatusNotFound,
				payload: RestResponse{
					Success: false,
					Message: "artist not found",
				},
			}
		}

		response.Render(w, r)
	}
}

type ArtistStats struct {
	Mean float64 `json:"mean" csv:"mean"`
	Median float64 `json:"median" csv:"median"`
	StandardDeviation float64 `json:"standard_deviation" csv:"standard_deviation"`
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
		id := mux.Vars(r)["artist_id"]

		year, _ := strconv.Atoi(r.URL.Query().Get("year"))

		songs, total := s.db.FindSongs(&repository.Query{
			OtherId: id,
			Year:    year,
		})

		hotnesses := make([]float64, total)

		for idx, song := range songs {
			hotnesses[idx] = song.SongHotttnesss
		}

		mean, _ := stats.Mean(hotnesses)
		median, _ := stats.Median(hotnesses)
		stdev, _ := stats.StandardDeviation(hotnesses)

		songurl, _ := s.router.Get("songs_all").URL()
		songqueryvals := songurl.Query()
		songqueryvals.Set("artist", id)
		songurl.RawQuery = songqueryvals.Encode()

		artisturl, _ := s.router.Get("artists_one").URL("artist_id", id)

		response := HttpResponse{
			status: http.StatusOK,
			payload: RestResponse{
				Success: true,
				Data: ArtistStats{
					Mean:              mean,
					Median:            median,
					StandardDeviation: stdev,
				},
				Links: map[string]string{
					"self":   r.URL.RequestURI(),
					"artist": artisturl.RequestURI(),
					"songs":  songurl.RequestURI(),
				},
			},
		}

		response.Render(w, r)
	}
}
