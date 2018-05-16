# Gopkg path of our project
PKG=github.com/maxmcd/c4/backend

PROTO_BUILDER_DOCKERFILE=.docker/protobuilder.Dockerfile

# Binaries are built from source in cmd/ to out in bin/
CMDS=$(notdir $(wildcard ./cmd/*))
BINS=$(patsubst %,bin/%,$(CMDS))

.PHONY: builder static container

all: builder api static

builder: $(PROTO_BUILDER_DOCKERFILE)
	docker build -f $< .docker/


container: static
	docker build . -f .docker/apiserver.Dockerfile -t apiserver

api: 
	docker run --rm -it -w /go/src/$(PKG) \
	    -v $(shell pwd):/go/src/$(PKG) \
	    $(shell docker build -f $(PROTO_BUILDER_DOCKERFILE) -q .docker/) \
	    /bin/bash -c "go generate ./pkg/..."

static: api
	docker run --rm -it -w /go/src/$(PKG) \
	    -v $(shell pwd):/go/src/$(PKG) \
	    $(shell docker build -f $(PROTO_BUILDER_DOCKERFILE) -q .docker/) \
	    /bin/bash -c "CGO_ENABLED=0 make cmds"

cmds: $(BINS)

# Pattern matching rules to build binaries from cmd/ dir to bin/ dir
bin/% : cmd/% cmd/%/*.go
	go build -o $@ ./$<
