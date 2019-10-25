all: server seeddb

server:
	go build -o ./bin/server ./cmd/server

seeddb:
	go build -o ./bin/seeddb ./cmd/seeddb

doc:
	swagger generate spec -w cmd/server -m -o doc/documentation.yaml
	redoc-cli bundle doc/documentation.yaml -o doc/documentation.html
.PHONY: server seeddb doc
