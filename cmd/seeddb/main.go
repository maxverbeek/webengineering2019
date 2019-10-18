package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"webeng/api/model"

	"github.com/gocarina/gocsv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var file string

func main() {
	flag.StringVar(&file, "f", "", "Read CSV data from `file`")
	flag.Parse()
	fmt.Println("Seeding db..")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type MusicRecord struct {
	artist
	release
	song
}

type song struct {
	gorm.Model
	model.Song
}

type artist struct {
	gorm.Model
	model.Artist
}

type release struct {
	gorm.Model
	model.Release
}

func run() error {
	db, err := gorm.Open("sqlite3", "./music.db")

	if err != nil {
		return err
	}

	defer db.Close()

	db.AutoMigrate(&song{})
	db.AutoMigrate(&artist{})
	db.AutoMigrate(&release{})

	var (
		csv io.ReadCloser
	)

	if file != "" {
		csv, err = os.Open(file)

		if err != nil {
			return err
		}

	} else {
		resp, err := http.Get(
			"https://think.cs.vt.edu/corgis/datasets/csv/music/music.csv")

		if err != nil {
			return err
		}

		fmt.Println("Received data")

		csv = resp.Body
	}

	defer csv.Close()

	records := []*MusicRecord{}

	if err = gocsv.Unmarshal(csv, &records); err != nil {
		return err
	}

	fmt.Println("Parsed data")

	sort.Slice(records, func(i, j int) bool {
		return records[i].SongYear > records[j].SongYear
	})

	artists := make(map[string]artist)
	songs := make(map[string]song)
	releases := make(map[int]release)

	tx := db.Begin()

	for _, song := range records {
		if _, ok := artists[song.Artist.ArtistId]; !ok {
			artists[song.artist.ArtistId] = song.artist
			tx.Create(&song.Artist)
		}
		if _, ok := releases[song.Release.ReleaseId]; !ok {
			releases[song.release.ReleaseId] = song.release
			tx.Create(&song.Release)
		}

		song.Song.ArtistId = song.Artist.ArtistId
		song.Song.ReleaseId = song.Release.ReleaseId
		songs[song.SongId] = song.song

		tx.Create(&song.Song)
	}

	tx.Commit()

	fmt.Printf("%#v\n", records[1000])

	return nil
}
