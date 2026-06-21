# run server
.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go build -o bin/gobank ./cmd/server/main.go

.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test -race ./...