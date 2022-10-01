#!/bin/bash

# dotenv
export $(egrep -v '^#' .env | xargs)

POSTGRESQL_URL="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"

migrate -database $POSTGRESQL_URL -path db/migrations $1
