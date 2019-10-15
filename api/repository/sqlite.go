// SQLite implementation that uses gorm
package repository

import (
	"webeng/api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SqliteStore struct {
	db *gorm.DB
}

type dbSong struct {
	// perhaps it may be easier to use raw SQL.
	gorm.Model
	model.Song
}

func (*SqliteStore) FindSong(query *Query) Song {
	return nil
}
