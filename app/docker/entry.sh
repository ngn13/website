#!/bin/bash

# replace the API URL
find ./build -type f -exec sed "s/http:\/\/placeholder\//${API_URL//\//\\/}/g" -i "{}" \;

# start the application
bun build/index.js
