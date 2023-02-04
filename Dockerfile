# Build wdb binary based on OS and Arch
FROM --platform=$BUILDPLATFORM golang:1.19-alpine3.16 AS builder
WORKDIR /app
ARG TARGETOS TARGETARCH
COPY . /app/
RUN  GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/wdb_docker ./cmd/wunderdb/wdb.go

# Load only binary from builder to lightweight alpine base-image
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /bin/wdb_docker /app/
ENTRYPOINT [ "/app/wdb_docker" ]