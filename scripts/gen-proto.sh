#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  protoc --go_out=${CURRENT_DIR}/protos -I /usr/local/include --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/protos --go-grpc_opt=paths=source_relative ${x}/*.proto
done
