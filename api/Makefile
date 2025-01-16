GOSRCS = $(wildcard *.go) $(wildcard */*.go)

all: api.elf

api.elf: $(GOSRCS)
	go build -o $@

run:
	API_DEBUG=true API_APP_URL=http://localhost:7002 API_PASSWORD=test ./api.elf

format:
	gofmt -s -w .

.PHONY: test format
