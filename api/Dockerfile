FROM golang:1.23.1

WORKDIR /app

COPY *.go     ./
COPY *.mod    ./
COPY *.sum    ./
COPY Makefile ./
COPY routes   ./routes
COPY global   ./global
COPY database ./database
COPY util     ./util

EXPOSE 7001
RUN make

ARG PASSWORD
ENV PASSWORD $PASSWORD

ARG FRONTEND_URL
ENV FRONTEND_URL $FRONTEND_URL

ENTRYPOINT ["/app/server"]
