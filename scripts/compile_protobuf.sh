#!/bin/sh

protoc --dart_out=./protocol_dart -I ./protobuf code.proto
protoc --go_out=./protocol_go -I ./protobuf code.proto
