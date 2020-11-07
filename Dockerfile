FROM golang:1.15.3-alpine

WORKDIR /opt/code
ADD ./ /opt/code

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/hammurabi-game cmd/hammurabi/main.go
ENTRYPOINT [ "bin/hammurabi-game" ]