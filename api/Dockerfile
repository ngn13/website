FROM golang:1.23.4

WORKDIR /api

COPY *.go     ./
COPY *.mod    ./
COPY *.sum    ./
COPY Makefile ./
COPY config   ./config
COPY database ./database
COPY routes   ./routes
COPY sql      ./sql
COPY status   ./status
COPY util     ./util
COPY views    ./views

EXPOSE 7001
RUN make

ENTRYPOINT ["/api/api.elf"]
