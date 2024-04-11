server: *.go routes/* util/*
	go build -o server .

test:
	URL=http://localhost:5173/ PASSWORD=test ./server 

.PHONY: test
