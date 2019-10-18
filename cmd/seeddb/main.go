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
	model.Artist
	model.Release
	model.Song
}

type song struct {
	gorm.Model `csv:"-"`
	model.Song `csv:"gofuckyourself"`
}

type artist struct {
	gorm.Model `csv:"-"`
	model.Artist `csv:"gofuckyourself"`
}

type release struct {
	gorm.Model `csv:"-"`
	model.Release `csv:"gofuckyourself"`
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

	for _, record := range records {
		if _, ok := artists[record.Artist.ArtistId]; !ok {
			a := artist{Artist: record.Artist}
			artists[record.Artist.ArtistId] = a
			tx.Create(&a)
		}
		if _, ok := releases[record.Release.ReleaseId]; !ok {
			r := release{Release: record.Release}
			releases[record.Release.ReleaseId] = r
			tx.Create(&r)
		}

		record.Song.ArtistId = record.Artist.ArtistId
		record.Song.ReleaseId = record.Release.ReleaseId

		s := song{Song: record.Song}
		songs[record.SongId] = s

		tx.Create(&s)
	}

	tx.Commit()

	fmt.Printf("%#v\n", records[1000])

	return nil
}
