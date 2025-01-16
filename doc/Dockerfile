FROM ghcr.io/ngn13/ctorm:1.7

WORKDIR /doc

COPY Makefile ./
COPY docs     ./docs
COPY inc      ./inc
COPY src      ./src

EXPOSE 7003
RUN make

ENTRYPOINT ["/doc/doc.elf"]
