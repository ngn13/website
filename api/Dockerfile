FROM golang:1.22.5

WORKDIR /app

COPY *.go ./
COPY *.mod ./
COPY *.sum ./
COPY Makefile ./
COPY routes ./routes
COPY util ./util

EXPOSE 7001
RUN go build

ARG PASSWORD
ENV PASSWORD $PASSWORD

ARG FRONTEND_URL
ENV FRONTEND_URL $FRONTEND_URL

ENTRYPOINT ["/app/server"]
