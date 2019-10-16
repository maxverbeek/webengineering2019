// Repository is responsible for storing `model` in
// a more permanent location, such as an SQL database.
package repository

import "webeng/api/model"

type Query struct {
	Id string
}

type Store interface {
	//FindSongs(query *Query) []Song // TODO
	FindSong(Query *Query) model.Song
	FindArtist(Query *Query) model.Artist
}
