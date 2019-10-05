# Web Engineering 2019

This repository contains our code for the music API for web engineering 2019.

## Usage

```bash
go build
./webeng -h # help
./webeng    # listen on port 8080 (default)
```

## Structural organisation

Unlike most projects this one does not (yet) have a `cmd` directory that
contains command entry points, such as `main.go`. This might possibly happen in
the future if this gets too cluttered/big. At some point documentation will end
up in the `doc` directory.

Currently this projects contains a single package (which might also change)
that is the main package. In the main function we limit ourselves to some flag
parsing, and immediately go into the `run()` function. This is because main
cannot return an error, but `run()` can (which `main` will then format and
print).

Throughout the code you will find structs, such as the `config` struct and the
`server` struct. These are the way they are in the interest of reducing global
variables.

The `server` struct holds global resources, such as database connections. This
struct is accessible by every route handler function.

### Routes

In spite of Golang's generous standard library, we outsourced the routing to
[gorilla mux](https://github.com/gorilla/mux). This decision was made because
Go's router from the standard library does not filter by request method very
nicely, while this one does.

The routes can be found, as you may have guessed, in `routes.go`. This gives
a nice overview of all the API endpoints which is convenient for debugging (and
grading). Each route has an associated `HandlerFunc`, which are created by
functions in the `server` struct. We do this so that every `HandlerFunc` has
its own closure environment, and thereby void global variables.
