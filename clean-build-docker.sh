#!/bin/bash
# Clean build, avoid rebuilding issues with cached layers
# Use this script when writing or editing dockerfiles

go clean

docker-compose down -v  # remove existing volumes
docker-compose kill     # kill running containers
docker-compose rm -f -v # removing images and anonymous volumes

# build with ssh keys copied from ssh home directory
docker-compose build --no-cache \
    --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" \
    --build-arg ssh_pub_key="$(cat ~/.ssh/id_rsa.pub)" 

docker-compose up -V
