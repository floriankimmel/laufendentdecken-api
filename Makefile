build:
	@go build -o bin/laufendentdeck-api

run: build
	@./bin/laufendentdeck-api

test:
	@go test -v ./...
