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
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type server struct {
	service *youtube.Service
	router  *mux.Router
	db      *repository.SqliteStore
}

type Config struct {
	Port int
}

func Run(conf *Config) error {
	r := mux.NewRouter()

	gormdb, err := gorm.Open("sqlite3", "music.db")

	if err != nil {
		return err
	}

	defer gormdb.Close()

	db := &repository.SqliteStore{
		Db: gormdb,
	}

	client := &http.Client{
		Transport: &transport.APIKey{Key: "AIzaSyCDjgmJS7aKZZt-BkOOkFWpMWCnCcmAn8k"},
	}

	serv, _ := youtube.New(client)

	server := &server{
		service: serv,
		router:  r,
		db:      db,
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
