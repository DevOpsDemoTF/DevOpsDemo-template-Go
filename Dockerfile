FROM golang:1.13-alpine as build

WORKDIR /app

ENV CGO_ENABLED=0
RUN apk add --no-cache git
RUN go get -u github.com/jstemmer/go-junit-report

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test -v ./... 2>&1 | tee - | go-junit-report > test-results.xml
RUN go build -ldflags "-s -w" -v -o app .

FROM alpine:3.10

RUN mkdir /app && \
    addgroup -S app && adduser -S app -G app

EXPOSE 8080
EXPOSE 9102

ENV DEBUG_LEVEL "DEBUG"

COPY --from=build /app/app /app/test-results.xml /app/

USER app
WORKDIR /app
CMD ["./app"]
