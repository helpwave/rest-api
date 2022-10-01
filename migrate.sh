#!/bin/bash

# dotenv
export $(egrep -v '^#' .env | xargs)

POSTGRESQL_URL="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$DB_HOST:$DB_PORT/$POSTGRES_DB?sslmode=disable"

migrate -database $POSTGRESQL_URL -path db/migrations $1
