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

func (s *SqliteStore) FindSongs(query *Query) ([]model.Song, int) {
	songs := make([]model.Song, 5)

	qsong := &song{}

	qsong.SongId = query.Id
	qsong.SongYear = query.Year

	q := s.Db.Model(&song{}).Where(qsong)

	if query.Genre != "" || query.OtherId != "" {
		q = q.Joins("JOIN artists ON artists.artist_id = songs.artist_id")

		if query.Genre != "" {
			q = q.Where("artists.artist_terms LIKE ?", fmt.Sprintf("%%%s%%", query.Genre))
		}

		if query.OtherId != "" {
			q = q.Where("artists.artist_id = ?", query.OtherId)
		}
	}

	if query.Name != "" {
		q = q.Where("songs.song_title LIKE ?", fmt.Sprintf("%%%s%%", query.Name))
	}

	var count int
	q.Count(&count)

	if query.Limit > 0 {
		q = q.Limit(query.Limit)

		if query.Page > 0 {
			// page 0 is the first page, page 1 is offset by Limit
			q = q.Offset(query.Limit * query.Page)
		}
	}

	switch query.Sort {
	case "duration",
		"hotttnesss",
		"id",
		"title",
		"tempo",
		"year":
		// TODO: find good way to change order
		q = q.Order(fmt.Sprintf("song_%s desc", query.Sort))
	}

	q.Find(&songs)

	return songs, count
}

func (s *SqliteStore) DeleteSong(query *Query) bool {
	sq := &song{}
	sq.SongId = query.Id

	var count int
	s.Db.Model(&song{}).Where(sq).Count(&count)

	if count == 1 {
		s.Db.Unscoped().Where(sq).Delete(&song{})
		return true
	}

	return false
}

func (s *SqliteStore) CreateSong(newsong *model.Song) bool {
	if s.FindSong(&Query{Id: newsong.SongId}) != nil {
		// song exists
		return false
	}

	s.Db.Create(&song{Song: *newsong})

	return true
}

func (s *SqliteStore) UpdateSong(query *Query, data map[string]interface{}) bool {
	qsong := &song{}
	qsong.SongId = query.Id

	var song song
	if s.Db.Where(qsong).Find(&song).RecordNotFound() {
		// cannot update non-existing record
		return false
	}

	fixedData := make(map[string]interface{})

	for k, v := range data {
		fixedData["song_"+k] = v
	}

	s.Db.Model(&song).Updates(fixedData)

	return true
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

func (s *SqliteStore) FindArtists(query *Query) ([]model.Artist, int) {
	artists := make([]model.Artist, 5)

	qartist := &artist{}
	qartist.ArtistId = query.Id
	qartist.ArtistTerms = query.Genre

	q := s.Db.Model(&artist{}).Where(qartist)

	if query.Name != "" {
		q = q.Where("artists.artist_name LIKE ?", fmt.Sprintf("%%%s%%", query.Name))
	}

	var count int
	q.Count(&count)

	if query.Limit > 0 {
		q = q.Limit(query.Limit)

		if query.Page > 0 {
			// page 0 is the first page, page 1 is offset by Limit
			q = q.Offset(query.Limit * query.Page)
		}
	}

	switch query.Sort {
	case "familiarity",
		"hotttnesss",
		"id",
		"name",
		"similar": // what is this?
		q = q.Order(fmt.Sprintf("artist_%s desc", query.Sort))
	}

	q.Find(&artists)

	return artists, count
}
