
compile-proto:
	./scripts/compile-proto.sh

build-local: compile-proto
	go build -o ./cmd/storage/bin/storage ./cmd/storage
	go build -o ./cmd/api/bin/api ./cmd/api