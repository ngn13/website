FROM golang:1.22.4

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

ARG FRONTEND_URL 
ENV FRONTEND_URL $FRONTEND_URL

ENTRYPOINT ["/app/server"]
