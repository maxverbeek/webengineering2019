package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

type Config struct {
	Port int
}

func Run(conf *Config) error {
	r := mux.NewRouter()

	server := &server{
		router: r,
		// todo initialise database(s) and other global crap here
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
