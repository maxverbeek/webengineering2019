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
	"net/http"
)

func (s *server) routes() {

	s.router.StrictSlash(true)

	api := s.router.PathPrefix("/api/v1").Subrouter()

	api.Path("/artists").
		Methods(http.MethodGet).
		HandlerFunc(s.handleArtists()).
		Name("artists_all")

	api.Path("/artists/{artist_id}").
		Methods(http.MethodGet).
		HandlerFunc(s.handleArtist()).
		Name("artists_one")

	api.Path("/artists/{artist_id}/stats").
		Methods(http.MethodGet).
		HandlerFunc(s.handleArtistStats()).
		Name("artists_stats")

	api.Path("/songs").
		Methods(http.MethodGet).
		HandlerFunc(s.handleSongs()).
		Name("songs_all")

	api.Path("/songs/{song_id}").
		Methods(http.MethodGet).
		HandlerFunc(s.handleSong()).
		Name("songs_one")

	api.Path("/songs/{song_id}").
	    Methods(http.MethodDelete).
		HandlerFunc(s.handleDeleteSong())

	s.router.
		PathPrefix("/").
		Methods(http.MethodGet).
		Handler(http.FileServer(http.Dir("./static/")))
}
