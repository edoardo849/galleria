#!/bin/bash

SRC_DIR=./api/proto/v1
DST_DIR=./pkg/api/v1

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/image.proto