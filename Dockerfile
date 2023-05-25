# Build wdb binary based on OS and Arch
FROM --platform=$BUILDPLATFORM golang:1.19-alpine3.16 AS builder
WORKDIR /app
ARG TARGETOS TARGETARCH
COPY . /app/
RUN  GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/wdb_docker ./cmd/wunderdb/wdb.go

# Build wdb-tools
FROM --platform=$BUILDPLATFORM golang:1.19-alpine3.16 AS tools-builder
WORKDIR /tools
RUN apk update && apk add git && apk add bash
ARG TARGETOS TARGETARCH
RUN git clone https://github.com/TanmoySG/wdb-sidekicks 
RUN cd /tools/wdb-sidekicks ; git pull
RUN cd /tools/wdb-sidekicks/tools/ ; chmod +x build.sh ; bash build.sh

# Load only binary from builder to lightweight alpine base-image
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /bin/wdb_docker /app/
COPY --from=tools-builder /tools/wdb-sidekicks/tools/tools/bin /tools/
ENTRYPOINT [ "/app/wdb_docker" ]