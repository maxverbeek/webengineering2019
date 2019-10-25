package main

import (
	"flag"
	"log"
	"os"

	"webeng/api"
)

func main() {
	port := flag.Int("port", 8080, "Port to run the app on")

	flag.Parse()

	err := api.Run(&api.Config{
		Port: *port,
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		return
	}

	os.Exit(0)

	return
}
