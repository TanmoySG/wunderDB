# Script to run wdb-server
# Usage:
#       bash ./scripts/start.sh <optional flags>
#       options:
#           - skip-tests : to skip tests
#           - skip-build : to skip build and use existing binary

BASEDIR=$(dirname "$0")

# skip unit test
if [[ ! $1 == "skip-tests" ]]; then
    go test ./... -coverpkg=./... -coverprofile ./coverage.out
    go tool cover -func ./coverage.out
fi

# skip build
if [[ ! $1 == "skip-build" ]]; then
    go build -o bin/wdb_server cmd/wunderdb/wdb.go
fi 

# reuse start script in docker directory
sh docker/start.sh
