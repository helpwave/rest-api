#!/bin/bash

# dotenv
export $(egrep -v '^#' .env | xargs)

POSTGRESQL_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable"

migrate -database $POSTGRESQL_URL -path db/migrations $1
