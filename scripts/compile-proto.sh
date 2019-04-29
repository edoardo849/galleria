#!/bin/bash

SRC_DIR=./api/proto
DST_DIR=./pkg/api
TPTY_DIR=./third_party
SWAGGER_DIR=./api/swagger

protoc --proto_path=$SRC_DIR \
    --proto_path=$TPTY_DIR \
    --go_out=plugins=grpc:$DST_DIR/storage $SRC_DIR/storage.proto

protoc --proto_path=$SRC_DIR \
    --proto_path=$TPTY_DIR \
    --go_out=plugins=grpc:$DST_DIR/decode $SRC_DIR/decode.proto



# protoc --proto_path=$SRC_DIR \
#     --proto_path=$TPTY_DIR \
#     --grpc-gateway_out=logtostderr=true:$DST_DIR $SRC_DIR/image.proto

# protoc --proto_path=$SRC_DIR \
#     --proto_path=$TPTY_DIR \
#     --swagger_out=logtostderr=true:$SWAGGER_DIR $SRC_DIR/image.proto