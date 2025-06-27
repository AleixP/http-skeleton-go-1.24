#!/bin/bash

set -e
printf "Removing docker containers, volumes, images and build cache"

docker stop $(docker ps -a -q) >/dev/null 2>&1 &
docker rm -f -v $(docker ps -a -q) >/dev/null 2>&1 &
docker volume prune -a -f &
docker image rm -f $(docker images -q -a) >/dev/null 2>&1 &
docker builder prune -a -f >/dev/null 2>&1 &
printf "Done".

