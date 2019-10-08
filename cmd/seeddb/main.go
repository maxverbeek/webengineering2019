package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"log"
	"net/http"

	"github.com/gocarina/gocsv"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	ArtistFamiliarity           float64 `csv:"artist.familiarity"`
	ArtistHotttnesss            float64 `csv:"artist.hotttnesss"`
	ArtistId                    string  `csv:"artist.id"`
	ArtistLatitude              float64 `csv:"artist.latitude"`
	ArtistLocation              int     `csv:"artist.location"`
	ArtistLongitude             float64 `csv:"artist.longitude"`
	ArtistName                  string  `csv:"artist.name"`
	ArtistSimilar               float64 `csv:"artist.similar"`
	ArtistTerms                 string  `csv:"artist.terms"`
	ArtistTermsFreq             float64 `csv:"artist.terms_freq"`
	ReleaseId                   int     `csv:"release.id"`
	ReleaseName                 int     `csv:"release.name"`
	SongArtistMbtags            float64 `csv:"song.artist_mbtags"`
	SongArtistMbtagsCount       float64 `csv:"song.artist_mbtags_count"`
	SongBarsConfidence          float64 `csv:"song.bars_confidence"`
	SongBarsStart               float64 `csv:"song.bars_start"`
	SongBeatsConfidence         float64 `csv:"song.beats_confidence"`
	SongBeatsStart              float64 `csv:"song.beats_start"`
	SongDuration                float64 `csv:"song.duration"`
	SongEndFadeIn               float64 `csv:"song.end_of_fade_in"`
	SongHotttnesss              float64 `csv:"song.hotttnesss"`
	SongId                      string  `csv:"song.id"`
	SongKey                     float64 `csv:"song.key"`
	SongKeyConfidence           float64 `csv:"song.key_confidence"`
	SongLoudness                float64 `csv:"song.loudness"`
	SongMode                    int     `csv:"song.mode"`
	SongModeConfidence          float64 `csv:"song.mode_confidence"`
	SongStartFadeOut            float64 `csv:"song.start_of_fade_out"`
	SongTatumsConfidence        float64 `csv:"song.tatums_confidence"`
	SongTatumsStart             float64 `csv:"song.tatums_start"`
	SongTempo                   float64 `csv:"song.tempo"`
	SongTimeSignature           float64 `csv:"song.time_signature"`
	SongTimeSignatureConfidence float64 `csv:"song.time_signature_confidence"`
	SongTitle                   int     `csv:"song.title"`
	SongYear                    int     `csv:"song.year"`
}

func run() error {
	//db, err := gorm.Open("sqlite3", "./music.db")

	//if err != nil {
	//return err
	//}

	var (
		csv io.ReadCloser
		err error
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

	fmt.Printf("%#v\n", records[0])

	return nil
}
