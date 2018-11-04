#!/bin/sh

docker-compose \
    -f docker-compose.yml \
    rm -f --stop 1>&2 || true

docker-compose \
    -f docker-compose.yml \
    up -d --force-recreate 1>&2