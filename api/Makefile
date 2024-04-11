server: *.go routes/* util/*
	go build -o server .

test:
	FRONTEND_URL=http://localhost:5173/ PASSWORD=test ./server 

.PHONY: test
