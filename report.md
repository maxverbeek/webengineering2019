# Report of our thing

Here you will find the report of our music database adventure thingy that
caused a lot of stress over the past 2 weeks.

## Architecture

We have chosen to use Golang for the back-end, using an SQLite database.
Initially we thought that using Go in the back-end would be rewarded, since it
comes with a very generous standard library that focusses on http. We do not
think that Go was a wrong choice, and we think that for production real-life
API's Go would actually be an excellent choice given its nice performance.
However for this project we were slowed down a little bit because of the steep
learning curve.

To connect to the SQLite database we used
[gORM](https://github.com/jinzhu/gorm). While this did not allow us to write
raw SQL queries, which would have been nice at times, we still think that this
library introduced a lot of simplicity in our back-end.

### Web frameworks

As you may have noticed, we have not talked about a web framework yet. That is
because we have used none. We used an external library
[gorilla-mux](https://github.com/gorilla/mux) to assist us in the organisation
of the routes. Other than that our program uses the built in http server from
Go. This gives us an incredible performance, which is absolutely destroyed by
our inefficient SQLite database.

We have created our own set-up for the back-end in a way we think is pretty
nice. We organise our dependencies (router, database) in a server struct. This
struct is defined in `api/server.go`. This is also where we initialise a
database, and set up the routes. Note that the individual routes are defined in
`api/routes.go` for the reader's convenience. This gives a very quick overview
of where the application executes code when it receives an incoming request.

> It is noteworthy that Go's package system is a little difficult to use at
> times. Every directory is considered a package, and in order to be usable
> this package needs to be in a module. A module is a directory that contains a
> `go.mod` file, such as the root directory of this repository. That makes
> `api`, `api/repository`, `api/model` and `cmd` four different packages. It is
> common Go practice to put all the main files in the `cmd` package.

The main package of interest is the `api` package, which contains the main
`server` struct, the routes, and the individual handlers for every endpoint.
These endpoints can be found in the `handle_*.go` files.

Another important package is the `api/repository` package, which contains all
of our database interaction. We kept the database completely independent from
the business logic in our API, so that in a better world we can throw away the
SQLite database and replace it with something faster.

> Another note about Go is that it does not allow cyclic dependencies. Since
> `api/repository` is a different package than `api` we cannot import these
> packages in each other. This means that if we were to define our models in
> `api` we would not be able to import them in `api/repository`. If we were to
> define our models in `api/repository` we would hardcode the ugly SQLite-gorm
> abomination into our project, which was undesirable. This is why we made the
> models separately, in `api/model`.

The `api/model` package contains the definitions of all our data types,
including a schema of what our JSON and csv responses look like. These CSV
definitions are also how we seed the database from the given csv file (see
`cmd/seeddb/main.go`).


## Documentation

The documentation is generated using openapi (formerly known as swagger). We do this by means of defining a lot of openapi definitions in our code, above every handler for every endpoint. This makes our code look quite bloated, but it makes it very easy to generate documentation. This tool can be installed using `go install github.com/go-swagger/go-swagger/cmd/swagger` -- But worry not! You need not install this, as we have generated the documentation for you.
