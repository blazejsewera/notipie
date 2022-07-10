# build
FROM golang:1.18-bullseye AS build-env

WORKDIR /core

# preload dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN make _build-binary-in-docker
RUN mv notipie.config.json /

# run
# bullseye-slim is currently experimental, change to bullseye if necessary
FROM debian:bullseye-slim

COPY --from=build-env /notipie /
COPY --from=build-env /notipie.config.json /

EXPOSE 8080
STOPSIGNAL SIGINT

ENTRYPOINT ["/notipie"]
