FROM golang:1.21.5

WORKDIR /app

COPY *.go ./
COPY *.mod ./
COPY *.sum ./
COPY Makefile ./
COPY routes ./routes
COPY util ./util

EXPOSE 7001 
RUN make 

ARG PASSWORD
ENV PASSWORD $PASSWORD

ENTRYPOINT ["/app/server"]
