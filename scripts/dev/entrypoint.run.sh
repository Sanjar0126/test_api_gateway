#!/bin/bash

set -e

cmd="$@"

>&2 echo "Sleeping ..."
# Wait for database readiness
sleep 10s

>&2 echo "Running migration ..."
PATH_TO_MIGRATIONS=/go/src/gitlab.udevs.io/delever/${APP}/migrations
migrate -path=$PATH_TO_MIGRATIONS -database=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable up

>&2 echo "Running server ..."
cd /go/src/gitlab.udevs.io/delever/${APP}
CGO_ENABLED=0 go run cmd/main.go

exec $cmd