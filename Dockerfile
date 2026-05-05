FROM golang:1.26 AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN \
  --mount=type=cache,target=/go/pkg/mod \
  go mod download
COPY . .
RUN \
  --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 go build

FROM debian:13-slim
COPY --from=build /build/tergum /usr/local/bin/tergum
