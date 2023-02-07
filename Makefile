build:
	go build -o bin/wdb ./cmd/wunderdb/wdb.go

build-cli:
	go build -o bin/wdbctl ./cmd/wdbctl/cli.go