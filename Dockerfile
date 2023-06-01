# Build wdb binary, wdb-sidekicks based on OS and Arch
FROM --platform=$BUILDPLATFORM golang:1.19-alpine3.16 AS builder
WORKDIR /app
ARG TARGETOS TARGETARCH
COPY . /app/
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/wdb_server ./cmd/wunderdb/wdb.go
RUN apk update && apk add git && apk add bash
ARG TARGETOS TARGETARCH
RUN git clone https://github.com/TanmoySG/wdb-sidekicks 
RUN cd /app/wdb-sidekicks ; git pull
RUN cd /app/wdb-sidekicks/tools/ ; chmod +x build.sh ; bash build.sh

# Load only binary from builder to lightweight alpine base-image
FROM alpine:3.16
WORKDIR /app
RUN apk update && apk add bash jq
COPY --from=builder /bin/wdb_server /app/
COPY --from=builder /app/wdb-sidekicks/tools/tools/bin /tools/
COPY --from=builder /app/docker/ /app/
ENTRYPOINT [ "bash", "/app/start.sh", "/app" ]