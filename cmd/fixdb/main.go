package main

import (
	"flag"
	"io"
	"log"
	"os"

	"webeng/api"

	"github.com/gocarina/gocsv"
)

var file string
var songf string

func main() {
	flag.StringVar(&file, "f", "", "Read CSV data from `file`")
	flag.StringVar(&songf, "s", "", "Read CSV song data from `file`")
	flag.Parse()

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

	var (
		csv io.ReadCloser
		ssv io.ReadCloser
		err error
	)

	if file != "" {
		csv, err = os.Open(file)

		if err != nil {
			return err
		}

	} else {
		log.Fatal("Please provide csv")
	}

	defer csv.Close()

	if songf != "" {
		ssv, err = os.Open(songf)

		if err != nil {
			return err
		}

	} else {
		log.Fatal("Please provide song name csv")
	}

	defer ssv.Close()

	records := []*MusicRecord{}
	songs := []*MusicRecord{}

	if err = gocsv.Unmarshal(csv, &records); err != nil {
		return err
	}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		//return csv.NewReader(in)
		return gocsv.LazyCSVReader(in) // Allows use of quotes in CSV
	})

	if err = gocsv.Unmarshal(ssv, &songs); err != nil {
		return err
	}

	for _, record := range records {
		for _, song := range songs {
			if record.SongId == song.SongId {
				record.SongTitle = song.SongTitle
				break
			}
		}
	}

	if err = gocsv.MarshalFile(&records, os.Stdout); err != nil {
		return err
	}
	return nil
}
