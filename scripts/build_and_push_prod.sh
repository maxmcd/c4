#!/bin/bash
set -e

cd "$(dirname "${BASH_SOURCE}")"
cd ../backend
make all
cd ..
docker-compose -f ./docker-compose-prod.yml build
docker-compose -f ./docker-compose-prod.yml push
