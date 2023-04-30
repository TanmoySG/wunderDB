build:
	go build -o bin/wdb ./cmd/wunderdb/wdb.go

build-cli:
	go build -o bin/wdbctl ./cmd/wdbctl/cli.go

run:
	go run ./cmd/wunderdb/wdb.go

gen-txlog-models:
	gojsonschema -p txlModel internal/txlogs/model/txlog.schema.json -o internal/txlogs/model/model.go

instal-dev:
	go get github.com/atombender/go-jsonschema/...