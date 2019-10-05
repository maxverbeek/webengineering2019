package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	port int
}

func main() {
	port := flag.Int("port", 8080, "Port to run the app on")

	flag.Parse()

	err := run(&Config{
		port: *port,
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		return
	}

	os.Exit(0)

	return
}

type server struct {
	router *mux.Router
}

func run(conf *Config) error {
	r := mux.NewRouter()

	server := &server{
		router: r,
		// todo initialise database(s) and other global crap here
	}

	// set up routes (routes.go)
	server.routes()

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", conf.port),

		// maybe remove timeouts later for telnet demo?
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on 0.0.0.0:%d", conf.port)

	return srv.ListenAndServe()
}
