IMAGE_NAME=$1 
IMAGE_TAG=$2

BASEDIR=$(dirname "$0")

env GOOS=linux \
    GOARCH=amd64 \
    go build -o ./bin/wdb_docker ./cmd/wunderdb/wdb.go

docker build -t wunderdb .
docker tag wunderdb:latest $IMAGE_NAME:$IMAGE_TAG
docker push $IMAGE_NAME:$IMAGE_TAG