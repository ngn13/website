all: server

server: *.go routes/* util/* global/* database/*
	go build -o server .

test:
	FRONTEND_URL=http://localhost:5173/ PASSWORD=test ./server 

format:
	gofmt -s -w .

.PHONY: test format
