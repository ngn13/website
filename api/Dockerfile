FROM golang:1.23.6

WORKDIR /api

RUN useradd runner -r -u 1001 -d /api
RUN chown -R runner:runner /api
USER runner

COPY *.mod    ./
COPY *.sum    ./

RUN go mod download

COPY *.go     ./
COPY Makefile ./
COPY config   ./config
COPY database ./database
COPY routes   ./routes
COPY sql      ./sql
COPY status   ./status
COPY util     ./util
COPY views    ./views

RUN make

ENTRYPOINT ["/api/api.elf"]
