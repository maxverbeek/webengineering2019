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

type artist struct {
	gorm.Model
	model.Artist
}

func (s *SqliteStore) FindArtist(query *Query) model.Artist {
	a := &artist{}
	a.ArtistId = query.Id

	var res artist
	s.Db.Where(a).Find(&res)

	return res.Artist
}

func (s *SqliteStore) FindSongs(query *Query) []model.Song {
	songs := make([]model.Song, 5)

	qsong := &song{}

	qsong.SongId = query.Id
	qsong.SongYear = query.Year

	s.Db.Where(qsong).Find(&songs)

	return songs
}
