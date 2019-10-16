// SQLite implementation that uses gorm
package repository

import (
	"webeng/api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SqliteStore struct {
	Db *gorm.DB
}

// has to be named song because the SQlite database is created like this
type song struct {
	gorm.Model
	model.Song
}

func (s *SqliteStore) FindSong(query *Query) model.Song {
	sng := &song{}
	sng.SongId = query.Id

	var res song
	s.Db.Where(sng).Find(&res)

	return res.Song
}

type dbArtist struct {
	gorm.Model
	model.Artist
}

func (s *SqliteStore) FindArtist(query *Query) model.Artist {
	artist := &dbArtist{}
	artist.ArtistId = query.Id

	var res dbArtist
	s.Db.Where(artist).Find(&res)

	return res.Artist
}
