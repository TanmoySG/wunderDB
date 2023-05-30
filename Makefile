build:
	go build -o bin/wdb_server ./cmd/wunderdb/wdb.go

build-cli:
	go build -o bin/wdbctl ./cmd/wdbctl/cli.go

build-image:
	docker build  . --tag wdb-local

run:
	go run ./cmd/wunderdb/wdb.go

gen-txlog-models:
	gojsonschema -p txlModel internal/txlogs/model/txlog.schema.json -o internal/txlogs/model/model.go

instal-dev:
	go get github.com/atombender/go-jsonschema/...

coverage:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
