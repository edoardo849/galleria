
compile-proto:
	./scripts/compile-proto.sh

build-local: compile-proto
	go build -o ./cmd/storage/bin/storage ./cmd/storage
	go build -o ./cmd/decode/bin/decode ./cmd/decode
	go build -o ./cmd/thumbnail/bin/thumbnail ./cmd/thumbnail
	go build -o ./cmd/api/bin/api ./cmd/api
