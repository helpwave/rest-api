# helpwave api

The official restful application programming interface of helpwave

## Setup

This project uses [Golang][golang] version 1.19.1, the latest version available at the time of writing.

You can build the project using `go build rest-api` and run it using `go run rest-api`.
`go build` will produce a single binary in the root called "rest-api".

## Environment

A `.env.example` file is provided and can be used to customize the server.
Use `cp .env.example .env` first.

### Gin

We use [gin][gin] for routing.

### Swagger

This project uses [swaggo][swaggo] to generate openapi specifications.
First install the `swag` cli using `go install github.com/swaggo/swag/cmd/swag@latest`.
Now you can run `swag init` to update `/docs`.
The server will serve the specification on `/swagger/index.html` (just `/swagger` won't work).

To add a route use the [declarative comments format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).

[golang]: https://go.dev/
[gin]: https://github.com/gin-gonic/gin/
[swaggo]: https://github.com/swaggo/
