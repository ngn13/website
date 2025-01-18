GOSRCS = $(wildcard *.go) $(wildcard */*.go)

all: api.elf

api.elf: $(GOSRCS)
	go build -o $@

run:
	WEBSITE_DEBUG=true WEBSITE_PASSWORD=test ./api.elf

format:
	gofmt -s -w .

.PHONY: test format
