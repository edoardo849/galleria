#!/bin/bash

SRC_DIR=./api/proto
DST_DIR=./pkg/api/storage

protoc -I=$SRC_DIR $SRC_DIR/image.proto --go_out=plugins=grpc:$DST_DIR 