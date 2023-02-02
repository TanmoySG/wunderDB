# smaller size ~15MB , but no shell
# FROM scratch 

FROM alpine:3.14

WORKDIR /app

COPY bin/wdb_docker /app/wdb_docker

ENTRYPOINT [ "/app/wdb_docker" ]
