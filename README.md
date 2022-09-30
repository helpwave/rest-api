# helpwave api

The official restful application programming interface of helpwave

## Setup

This project uses [Golang][golang] version 1.19.1, the latest version available at the time of writing.

You can build the project using `go build rest-api` and run it using `go run rest-api`.
`go build` will produce a single binary in the root called "rest-api".

## Environment

The following environment variables can be used to customize the server:
 * `PORT` - The port on which the api will listen on (default: 3000)

### Gin

We use [gin][gin] for routing.

[golang]: https://go.dev/
[gin]: https://github.com/gin-gonic/gin/
