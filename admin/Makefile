SRCS   = $(wildcard *.py)
PREFIX = /usr

all:

format:
	black $(SRCS)

install:
	install -Dm755 admin.py $(PREFIX)/bin/admin_script

.PHONY: format install
