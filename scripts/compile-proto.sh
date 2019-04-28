#!/bin/bash

SRC_DIR=./api/proto
DST_DIR=./pkg/api/storage

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/image.proto