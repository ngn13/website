FROM golang:1.23.4

WORKDIR /api

COPY *.go     ./
COPY *.mod    ./
COPY *.sum    ./
COPY Makefile ./
COPY util     ./util
COPY views    ./views
COPY routes   ./routes
COPY config   ./config
COPY status   ./status
COPY database ./database

EXPOSE 7001
RUN make

ENTRYPOINT ["/api/api.elf"]
