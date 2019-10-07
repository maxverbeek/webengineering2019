all: server seeddb

server:
	go build -o server ./cmd/server

seeddb:
	go build -o seeddb ./cmd/seeddb
