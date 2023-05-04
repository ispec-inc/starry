#!/bin/bash

docker network create starry || true

docker-compose up -d mysql
docker-compose run --rm dockerize -wait tcp://mysql:3306 -timeout 40s

docker-compose up -d redis
docker-compose run --rm dockerize -wait tcp://redis:6379 -timeout 40s

echo 'MySQL and Redis are ready!'
