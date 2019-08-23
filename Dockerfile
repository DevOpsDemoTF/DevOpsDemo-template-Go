FROM golang:1.12-alpine as build

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/tebeka/go2xunit github.com/golang/dep
RUN /go/bin/dep ensure
RUN go build -ldflags "-s -w" -v .
RUN 2>&1 go test -v | tee - | /go/bin/go2xunit -output test-results.xml

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
