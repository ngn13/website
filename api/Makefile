all: server

server: *.go routes/*.go util/*.go global/*.go database/*.go
	go build -o $@ .

test:
	FRONTEND_URL=http://localhost:5173/ PASSWORD=test ./server

format:
	gofmt -s -w .

.PHONY: test format
