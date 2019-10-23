// Repository is responsible for storing `model` in
// a more permanent location, such as an SQL database.
package repository

import "webeng/api/model"

type Query struct {
	Id    string
	Year  int
	Genre string
	Name  string
	OtherId string

	Sort  string
	Page  int
	Limit int
}

type Store interface {
	FindSong(query *Query) model.Song
	FindSongs(query *Query) []model.Song
	FindArtist(query *Query) model.Artist
	FindArtists(query *Query) []model.Artist
}
