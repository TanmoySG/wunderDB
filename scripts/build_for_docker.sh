BASEDIR=$(dirname "$0")

env GOOS=linux \
    GOARCH=amd64 \
    go build -o ./bin/wdb_docker ./cmd/wunderdb/wdb.go

docker build -t wdb-test . 