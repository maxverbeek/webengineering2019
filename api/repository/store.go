// Repository is responsible for storing `model` in
// a more permanent location, such as an SQL database.
package repository

import "webeng/api/model"

type Query struct {
	Id string
	Year int
}

type Store interface {
	FindSongs(query *Query) []model.Song
	FindSong(Query *Query) model.Song
	FindArtist(Query *Query) model.Artist
}
