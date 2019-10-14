package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"webeng/api"

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
	api.Artist
	api.Release
	api.Song
}

func run() error {
	db, err := gorm.Open("sqlite3", "./music.db")

	if err != nil {
		return err
	}

	defer db.Close()

	db.AutoMigrate(&api.Artist{})
	db.AutoMigrate(&api.Release{})
	db.AutoMigrate(&api.Song{})

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

	artists := make(map[string]api.Artist)
	songs := make(map[string]api.Song)
	releases := make(map[int]api.Release)

	tx := db.Begin()

	for _, song := range records {
		if _, ok := artists[song.Artist.ArtistId]; !ok {
			artists[song.Artist.ArtistId] = song.Artist
			tx.Create(&song.Artist)
		}
		if _, ok := releases[song.Release.ReleaseId]; !ok {
			releases[song.Release.ReleaseId] = song.Release
			tx.Create(&song.Release)
		}

		song.Song.ArtistId = song.Artist.ArtistId
		song.Song.ReleaseId = song.Release.ReleaseId
		songs[song.SongId] = song.Song

		tx.Create(&song.Song)
	}

	tx.Commit()

	fmt.Printf("%#v\n", records[1000])

	return nil
}
