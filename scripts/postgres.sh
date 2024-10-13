#!/bin/bash

# check the available version
docker --version
# check prune necessary setups
docker system prune
# create a docker volume
docker volume create --name=postgres-data

# run the postgres compose container
docker-compose -f postgres.docker-compose.yaml down
docker-compose -f postgres.docker-compose.yaml up --build