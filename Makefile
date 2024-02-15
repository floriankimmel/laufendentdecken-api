build:
	@go build -o bin/laufendentdeck-api

run: build
	@./bin/laufendentdeck-api

test:
	@go test -v ./...

review:
	@http http://localhost:8080/reviews/8ba03d9f-0cfc-4653-a1d5-0ff4e1ff02a5 | jq

event:
	@http http://localhost:8080/trailEvents/7eecfeea-5070-42f7-ba1e-0536c8a55c53 | jq
