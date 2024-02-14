build:
	@go build -o bin/laufendentdeck-api

run: build
	@./bin/laufendentdeck-api

test:
	@go test -v ./...

example:
	@http http://localhost:8080/reviews/689f2ec-b015-4d16-a022-baba3330acb9 | jq
