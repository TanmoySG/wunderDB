# Build wdb binary, wdb-sidekicks based on OS and Arch
FROM --platform=$BUILDPLATFORM golang:1.24-alpine3.16 AS builder
WORKDIR /app
ARG TARGETOS TARGETARCH
COPY . /app/
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/wdb_server ./cmd/wunderdb/main.go

# Load only binary from builder to lightweight alpine base-image
FROM alpine:3.21
WORKDIR /app
RUN apk update && apk add --no-cache bash jq
COPY --from=builder /bin/wdb_server /app/
CMD ["/app/wdb_server"]