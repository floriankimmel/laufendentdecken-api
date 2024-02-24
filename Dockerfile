FROM alpine:edge AS build
RUN apk add --no-cache --update go gcc g++
WORKDIR /app
COPY . .
RUN go test -v
RUN CGO_ENABLED=1 GOOS=linux go build -o lep-api

FROM alpine:edge
WORKDIR /app
RUN apk add --no-cache sqlite
COPY --from=build /app/lep-api /app/lep-api
COPY --from=build /app/laufendentdeckendb.db /app/laufendentdeckendb.db
ENTRYPOINT /app/lep-api
