#!/bin/bash

BASE_DIR="proto/"
PROTO_FILES=(model.proto request.proto)
PROTO_FILES=(${PROTO_FILES[@]/#/${BASE_DIR}})

protoc --dart_out=./protocol_dart -I. ${PROTO_FILES[@]}
protoc --go_out=./protocol_go -I. ${PROTO_FILES[@]}
