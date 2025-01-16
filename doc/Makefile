# dirs
DIRS  = $(shell find src/* -type d)
DISTDIR = dist
OUTDIRS  = $(patsubst src/%,$(DISTDIR)/%,$(DIRS))

# sources
HSRCS = $(wildcard inc/*.h)
CSRCS = $(shell find -type f -name '*.c')
OBJS  = $(patsubst ./src/%.c,./$(DISTDIR)/%.o,$(CSRCS))

# compiler flags
CFLAGS  = -O3 -fstack-protector-strong -fcf-protection=full -fstack-clash-protection
LIBS    = -lctorm -lcjson
INCLUDE = -I./inc

all: doc.elf

doc.elf: $(OBJS)
	echo $(OBJS) $(OUTDIRS)
	gcc $(CFLAGS) -o $@ $^ $(LIBS)

$(DISTDIR)/%.o: src/%.c
	@mkdir -pv $(OUTDIRS)
	gcc $(CFLAGS) $(INCLUDE) -c -o $@ $^ $(LIBS)

format:
	clang-format -i -style=file $(CSRCS) $(HSRCS)

clean:
	rm -rf $(DISTDIR)

run:
	./doc.elf

.PHONY: format clean run
