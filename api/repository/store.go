// Repository is responsible for storing `model` in
// a more permanent location, such as an SQL database.
package repository

import "webeng/api/model"

type Query struct {
}

type Store interface {
	//FindSongs(query *Query) []Song // TODO
	FindSong(Query *Query) Song
	FindArtist(Query *Query) Artist
}
