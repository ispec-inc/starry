#!/usr/bin/env bash
set -eu

ROUTER_TIMEOUT=10s
ROUTER_ALLOW_ORIGINS=http://localhost:3000

MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_DATABASE=starry
MYSQL_HOST=localhost
MYSQL_PORT=13306
MYSQL_SHOW_ALL_LOG=true

REDIS_MSGBS_HOST=localhost
REDIS_MSGBS_PORT=16379

export ROUTER_TIMEOUT
export ROUTER_ALLOW_ORIGINS

export MYSQL_USER
export MYSQL_PASSWORD
export MYSQL_DATABASE
export MYSQL_HOST
export MYSQL_PORT
export MYSQL_SHOW_ALL_LOG

export REDIS_MSGBS_HOST
export REDIS_MSGBS_PORT

exec "$@"
