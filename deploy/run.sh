#!/bin/bash

if [ ! -f data.db ]; then
  touch data.db
  sudo chmod 1001:1001 data.db
fi

docker-compose build
docker-compose up -d
