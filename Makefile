
compile-proto:
	./scripts/compile-proto.sh

build-local: compile-proto
	go build -o ./cmd/storage/bin/storage ./cmd/storage
	go build -o ./cmd/decode/bin/decode ./cmd/decode
	go build -o ./cmd/thumbnail/bin/thumbnail ./cmd/thumbnail
	go build -o ./cmd/api/bin/api ./cmd/api

stop:
	docker-compose -f ./deployments/docker-compose.yml down -v

build: compile-proto
	docker-compose -f ./deployments/docker-compose.yml build

run: compile-proto
	docker-compose -f ./deployments/docker-compose.yml up

integration-tests:
	newman run ./test/integration/Progimage.postman_collection.json \
		-e ./test/integration/Progimage.postman_environment.json