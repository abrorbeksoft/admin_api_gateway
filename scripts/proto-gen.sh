#!/bin/bash
CURRENT_DIR=$(pwd)
for i in $(find ${CURRENT_DIR}/protos/* -type d); do
#  echo $i
  protoc --proto_path=${CURRENT_DIR}/protos --go_out=plugins=grpc:${CURRENT_DIR} ${i}/*.proto
done