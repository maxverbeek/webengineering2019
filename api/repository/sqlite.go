// SQLite implementation that uses gorm
package repository

import (
	"fmt"

	"webeng/api/model"

	"github.com/jinzhu/gorm"
)

type SqliteStore struct {
	Db *gorm.DB
}

// has to be named song because the SQlite database is created like this
type song struct {
	gorm.Model
	model.Song
}

func (s *SqliteStore) FindSong(query *Query) *model.Song {
	sng := &song{}
	sng.SongId = query.Id

	var res song

	if s.Db.Where(sng).Find(&res).RecordNotFound() {
		return nil
	}

	return &res.Song
}

func (s *SqliteStore) FindSongs(query *Query) []model.Song {
	songs := make([]model.Song, 5)

	qsong := &song{}

	qsong.SongId = query.Id
	qsong.SongYear = query.Year

	q := s.Db.Where(qsong)

	if query.Genre != "" {
		q = q.Joins("JOIN artists ON artists.artist_id = songs.artist_id")
		q = q.Where("artists.artist_terms = ?", query.Genre)
	}

	if query.Name != "" {
		q = q.Where("songs.song_title LIKE ?", fmt.Sprintf("%%%s%%", query.Name))
	}

	if query.Limit != 0 {
		q = q.Limit(query.Limit)

		if query.Page != 0 {
			// page 0 is the first page, page 1 is offset by Limit
			q = q.Offset(query.Limit * query.Page)
		}
	}

	q.Find(&songs)

	return songs
}

type artist struct {
	gorm.Model
	model.Artist
}

func (s *SqliteStore) FindArtist(query *Query) *model.Artist {
	a := &artist{}
	a.ArtistId = query.Id

	var res artist

	if s.Db.Where(a).Find(&res).RecordNotFound() {
		return nil
	}

	return &res.Artist
}

func (s *SqliteStore) FindArtists(query *Query) []model.Artist {
	artists := make([]model.Artist, 5)

	qartist := &artist{}
	qartist.ArtistId = query.Id
	qartist.ArtistTerms = query.Genre

	q := s.Db.Where(qartist)

	if query.Name != "" {
		q = q.Where("artists.artist_name LIKE ?", fmt.Sprintf("%%%s%%", query.Name))
	}

	if query.Limit != 0 {
		q = q.Limit(query.Limit)

		if query.Page != 0 {
			// page 0 is the first page, page 1 is offset by Limit
			q = q.Offset(query.Limit * query.Page)
		}
	}

	q.Find(&artists)

	return artists
}
