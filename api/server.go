package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"webeng/api/repository"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type server struct {
	router *mux.Router
	newdb  *repository.SqliteStore
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

	defer db.Close()

	db.AutoMigrate(&Song{})
	db.AutoMigrate(&Artist{})
	db.AutoMigrate(&Release{})

	server := &server{
		router: r,
		newdb:  &repository.SqliteStore{db},
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

	go srv.ListenAndServe()

	c := make(chan os.Signal, 1)

	// block until process receives SIGINT
	signal.Notify(c, os.Interrupt)
	<-c
	signal.Stop(c)

	// received signal -> shutdown server
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	return nil
}
