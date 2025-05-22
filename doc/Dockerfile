FROM ghcr.io/ngn13/ctorm:1.8.1

WORKDIR /doc

COPY Makefile ./
COPY docs     ./docs
COPY inc      ./inc
COPY src      ./src

RUN useradd runner -r -u 1001 -d /doc
RUN chown -R runner:runner /doc

USER runner
RUN make

EXPOSE 7003
ENTRYPOINT ["/doc/doc.elf"]
