FROM golang:1.23.2

WORKDIR /app

COPY *.go     ./
COPY *.mod    ./
COPY *.sum    ./
COPY Makefile ./
COPY routes   ./routes
COPY config   ./config
COPY database ./database
COPY util     ./util

EXPOSE 7001
RUN make

ENTRYPOINT ["/app/server"]
