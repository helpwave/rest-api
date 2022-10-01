# helpwave api

The official restful application programming interface of helpwave

## Environment

A `.env.example` file is provided and can be used to customize the server.
copy the file to `.env` and edit it to your needs.
```bash
cp .env.example .env
```
## Docker
Run Docker Compose to start the Database:
```bash
mkdir -p ./data/postgres
docker-compose up -d postgres
```
if you want to remove the database:
```bash
docker-compose down posgres
sudo rm -rf ./data/postgres
```

## Setup

This project uses [Golang][golang] version 1.19.1, the latest version available at the time of writing.

You can build the project using
```bash
go build rest-api
```
and run it using
```bash
go run rest-api
```
The link to the api endpoint: [http://localhost:3000](http://localhost:3000)
***
```bash
go build
```
will produce a single binary in the root called "rest-api".

### Gin

We use [gin][gin] for routing.

### Swagger

This project uses [swaggo][swaggo] to generate openapi specifications.
First install the `swag` cli using:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
Now you can run `swag init` to update `/docs`.
The server will serve the specification on `/swagger/index.html` (just `/swagger` won't work).

To add a route use the [declarative comments format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).

### GORM

To connect to postgres we use [GORM][GORM].

Keep in mind the following [conventions](https://gorm.io/docs/models.html#Conventions):
> GORM prefers convention to configuration,
> by default, GORM uses ID as primary key, pluralize struct name to
> snake_cases as table name, snake_case as column name, and uses CreatedAt,
> UpdatedAt to track creating/updating time

### migrate

Install the [`migrate`][migrate] cli using:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

New migrations can be generated using `migrate create -ext sql -dir db/migrations -seq <title>`.

Users with `bash` installed can use
```bash
./migration.sh up
```
to run a migration against the db specified in `.env`.

[golang]: https://go.dev/
[gin]: https://github.com/gin-gonic/gin/
[swaggo]: https://github.com/swaggo/
[GORM]: https://gorm.io/
[migrate]: https://github.com/golang-migrate/migrate/
