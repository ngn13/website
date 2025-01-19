SERVERS = app api doc

all: $(SERVERS)
	for server in $^ ; do \
		make -C $$server ; \
	done

format:
	for server in $(SERVERS) ; do \
		make -C $$server format ; \
	done

.PHONY: format
