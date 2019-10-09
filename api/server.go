package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type server struct {
	router *mux.Router
	db *gorm.DB
}

type Config struct {
	Port int
}

func Run(conf *Config) error {
	r := mux.NewRouter()

	db, err := gorm.Open("sqlite3", "music.db")
	
	if err != nil {
		return err
	}

	server := &server{
		router: r,
		db: db,
	}

	// set up routes (routes.go)
	server.routes()

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", conf.Port),

		// maybe remove timeouts later for telnet demo?
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on 0.0.0.0:%d", conf.Port)

	return srv.ListenAndServe()
}
