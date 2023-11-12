FROM golang:1.21.3

WORKDIR /app

COPY *.go ./
COPY *.mod ./
COPY *.sum ./
COPY Makefile ./
COPY routes ./routes
COPY util ./util

RUN make 
ENTRYPOINT ["/app/server"]
