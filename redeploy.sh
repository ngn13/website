#!/bin/bash
# redeployment script for docker-compose
docker-compose down
docker rmi website_app:latest
docker rmi website_api:latest
git pull && docker-compose up -d
