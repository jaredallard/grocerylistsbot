# syntax=docker/dockerfile:1.0.0-experimental
FROM golang:1.14-alpine as builder

# Default to using go modules and go proxy
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

ARG VERSION
WORKDIR /src

# Ensure that we're using the latest git, openssh, and ca-certs
RUN apk add --no-cache --update openssh-client git ca-certificates

# Add go.mod and go.sum first to maximize caching
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

# Build our application
RUN go build -o /src/bin/grocerylistsbot ./cmd/...

FROM alpine:3.11
ENTRYPOINT ["/usr/local/bin/grocerylistsbot"]

# add ca-certs
RUN apk add --no-cache ca-certificates

COPY --from=builder /src/bin/grocerylistsbot /usr/local/bin/grocerylistsbot
