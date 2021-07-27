#!/bin/bash

set -e

cmd="$@"

>&2 echo "Sleeping ..."
# Wait for database readiness
sleep 10s

>&2 echo "Running migration ..."
PATH_TO_MIGRATIONS=/go/src/gitlab.udevs.io/delever/$APP/migrations
migrate -path=$PATH_TO_MIGRATIONS -database=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable up

>&2 echo "Running tests ..."
cd /go/src/gitlab.udevs.io/delever/$APP/tests
CGO_ENABLED=0 go test ./... -v -tags integration

exec $cmd