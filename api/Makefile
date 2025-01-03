GOSRCS = $(wildcard *.go) $(wildcard */*.go)

all: api.elf

api.elf: $(GOSRCS)
	go build -o $@

run:
	API_DEBUG=true API_FRONTEND_URL=http://localhost:5173/ API_PASSWORD=test ./api.elf

format:
	gofmt -s -w .

.PHONY: test format
