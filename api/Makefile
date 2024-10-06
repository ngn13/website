all: server

server: *.go routes/*.go database/*.go util/*.go config/*.go
	go build -o $@ .

test:
	API_FRONTEND_URL=http://localhost:5173/ API_PASSWORD=test ./server

format:
	gofmt -s -w .

.PHONY: test format
