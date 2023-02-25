#!/bin/sh

docker volume prune -f && \
git pull && \
docker-compose up -d --build && \
docker image prune -f