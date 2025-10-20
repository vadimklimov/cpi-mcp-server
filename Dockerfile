# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25-alpine AS build

ARG VERSION

WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . ./
RUN CGO_ENABLED=0 \
    go build -trimpath \
    -ldflags "-s -w -X github.com/vadimklimov/cpi-mcp-server/internal/appinfo.version=$VERSION" \
    -o cpi-mcp-server

# Deployment stage
FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && addgroup -g 1000 mcp \
    && adduser -D -u 1000 -G mcp mcp

WORKDIR /app

COPY --from=build --chown=mcp:mcp /build/cpi-mcp-server ./

USER mcp
EXPOSE 8080
ENTRYPOINT ["./cpi-mcp-server"]