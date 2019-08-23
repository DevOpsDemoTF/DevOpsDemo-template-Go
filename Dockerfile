FROM golang:1.12-alpine as build

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0

RUN apk add --no-cache git
RUN go get -u github.com/tebeka/go2xunit github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go test -v . | tee - | go2xunit -output test-results.xml
RUN go build -ldflags "-s -w" -v -o app .

FROM alpine:3.9

RUN mkdir /app && \
    addgroup -S app && adduser -S app -G app

EXPOSE 8080
EXPOSE 9102

ENV DEBUG_LEVEL "DEBUG"

COPY --from=build /go/src/app/app /go/src/app/test-results.xml /app/

USER app
WORKDIR /app
CMD ["./app"]
