# The binary to build (just the basename).
BIN := go-swap-server
SRC_BIN := server

# This repo's root import path (under GOPATH).
PKG := github.com/me-io/go-swap

# Where to push the docker image.
REGISTRY ?= me-io

# Which architecture to build - see $(BUILD_PLATFORMS) for options.
ARCH ?= amd64
OS ?= linux

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)
#
# This version-strategy uses a manual value to set the version string
#VERSION := 1.2.3

###
### These variables should not need tweaking.
###

SRC_DIRS := cmd pkg # directories which hold app source (not vendored)

# $(OS)-$(ARCH) pairs to build binaries and containers for
BUILD_PLATFORMS := linux-amd64 linux-arm linux-arm64 linux-ppc64le freebsd-amd64 freebsd-386
CONTAINER_PLATFORMS := linux-amd64 linux-arm64 linux-ppc64le # must be a subset of BUILD_PLATFORMS

# Set default base image dynamically for each arch
ifeq ($(ARCH),amd64)
    BASEIMAGE?=alpine
endif
ifeq ($(ARCH),arm)
    BASEIMAGE?=armel/busybox
endif
ifeq ($(ARCH),arm64)
    BASEIMAGE?=aarch64/busybox
endif
ifeq ($(ARCH),ppc64le)
    BASEIMAGE?=ppc64le/busybox
endif

IMAGE := $(REGISTRY)/$(BIN)-$(OS)-$(ARCH)

BUILD_IMAGE ?= golang:1.11-alpine

# If you want to build all binaries, see the 'all-build' rule.
# If you want to build all containers, see the 'all-container' rule.
# If you want to build AND push all containers, see the 'all-push' rule.
all: build

build-%:
	@$(MAKE) --no-print-directory ARCH=$(word 2,$(subst -, ,$*)) OS=$(word 1,$(subst -, ,$*)) build

container-%:
	@$(MAKE) --no-print-directory ARCH=$(word 2,$(subst -, ,$*)) OS=$(word 1,$(subst -, ,$*)) container

push-%:
	@$(MAKE) --no-print-directory ARCH=$(word 2,$(subst -, ,$*)) OS=$(word 1,$(subst -, ,$*)) push

all-build: $(addprefix build-, $(BUILD_PLATFORMS))

all-container: $(addprefix container-, $(CONTAINER_PLATFORMS))

all-push: $(addprefix push-, $(CONTAINER_PLATFORMS))

build: bin/$(OS)-$(ARCH)/$(BIN)

bin/$(OS)-$(ARCH)/$(BIN): build-dirs
ifeq ($(filter $(OS)-$(ARCH),$(BUILD_PLATFORMS)),)
	$(error unsupported build platform $(OS)-$(ARCH) not in $(BUILD_PLATFORMS))
endif
	@echo "building: $@"
	@docker run                                                                   \
	    -ti                                                                       \
	    --rm                                                                      \
	    -u $$(id -u):$$(id -g)                                                    \
	    -v "$$(pwd)/.go:/go"                                                      \
	    -v "$$(pwd):/go/src/$(PKG)"                                               \
	    -v "$$(pwd)/bin/$(OS)-$(ARCH):/go/bin"                                    \
	    -v "$$(pwd)/bin/$(OS)-$(ARCH):/go/bin/$(OS)_$(ARCH)"                      \
	    -v "$$(pwd)/.go/std/$(OS)-$(ARCH):/usr/local/go/pkg/$(OS)_$(ARCH)_static" \
		-v "$$(pwd)/.go/cache:/.cache"                                      	  \
	    -w /go/src/$(PKG)                                                         \
	    $(BUILD_IMAGE)                                                            \
	    /bin/sh -c "                                                              \
	        ARCH=$(ARCH)                                                          \
	        OS=$(OS)                                                              \
	        VERSION=$(VERSION)                                                    \
	        PKG=$(PKG)                                                            \
	        ./scripts/build.sh                                                      \
	    "

# Example: make shell CMD="-c 'date > datefile'"
shell: build-dirs
	@echo "launching a shell in the containerized build environment"
	@docker run                                                                   \
	    -ti                                                                       \
	    --rm                                                                      \
	    -u $$(id -u):$$(id -g)                                                    \
	    -v "$$(pwd)/.go:/go"                                                      \
	    -v "$$(pwd):/go/src/$(PKG)"                                               \
	    -v "$$(pwd)/bin/$(OS)-$(ARCH):/go/bin"                                    \
	    -v "$$(pwd)/bin/$(OS)-$(ARCH):/go/bin/$(OS)_$(ARCH)"                      \
	    -v "$$(pwd)/.go/std/$(OS)-$(ARCH):/usr/local/go/pkg/$(OS)_$(ARCH)_static" \
		-v "$$(pwd)/.go/cache:/.cache"                                      	  \
	    -w /go/src/$(PKG)                                                         \
	    $(BUILD_IMAGE)                                                            \
	    /bin/sh $(CMD)

DOTFILE_IMAGE = $(subst :,_,$(subst /,_,$(IMAGE))-$(VERSION))

container: .container-$(DOTFILE_IMAGE) container-name
.container-$(DOTFILE_IMAGE): bin/$(OS)-$(ARCH)/$(BIN) Dockerfile.in
	@sed \
	    -e 's|ARG_BIN|$(BIN)|g' \
	    -e 's|ARG_SRC_BIN|$(SRC_BIN)|g' \
	    -e 's|ARG_OS|$(OS)|g' \
	    -e 's|ARG_ARCH|$(ARCH)|g' \
	    -e 's|ARG_FROM|$(BASEIMAGE)|g' \
	    Dockerfile.in > .dockerfile-$(OS)-$(ARCH)
	@docker build -t $(IMAGE):$(VERSION) -f .dockerfile-$(OS)-$(ARCH) .
	@docker images -q $(IMAGE):$(VERSION) > $@

container-name:
	@echo "container: $(IMAGE):$(VERSION)"

push: .push-$(DOTFILE_IMAGE) push-name
.push-$(DOTFILE_IMAGE): .container-$(DOTFILE_IMAGE)
ifeq ($(findstring gcr.io,$(REGISTRY)),gcr.io)
	@gcloud docker -- push $(IMAGE):$(VERSION)
else
	@docker push $(IMAGE):$(VERSION)
endif
	@docker images -q $(IMAGE):$(VERSION) > $@

push-name:
	@echo "pushed: $(IMAGE):$(VERSION)"

version:
	@echo $(VERSION)

test-local:
	go test -v ./...

test-docker: build-dirs
	@dep version >/dev/null 2>&1 || ( wget -O - https://raw.githubusercontent.com/golang/dep/master/install.sh | sh )
	@dep ensure --vendor-only
	@docker run                                                                   \
	    -ti                                                                       \
	    --rm                                                                      \
	    -u $$(id -u):$$(id -g)                                                    \
	    -v "$$(pwd)/.go:/go"                                                      \
	    -v "$$(pwd):/go/src/$(PKG)"                                               \
	    -v "$$(pwd)/bin/$(OS)-$(ARCH):/go/bin"                                          \
	    -v "$$(pwd)/.go/std/$(OS)-$(ARCH):/usr/local/go/pkg/$(OS)_$(ARCH)_static" \
		-v "$$(pwd)/.go/cache:/.cache"                                      	  \
	    -w /go/src/$(PKG)                                                         \
	    $(BUILD_IMAGE)                                                            \
	    /bin/sh -c "                                                              \
	        ./scripts/test.sh $(SRC_DIRS)                                           \
	    "

build-dirs:
	@mkdir -p bin/$(OS)-$(ARCH)
	@mkdir -p .go/cache .go/src/$(PKG) .go/pkg .go/bin .go/std/$(OS)-$(ARCH)

clean: container-clean bin-clean

container-clean:
	rm -rf .container-* .dockerfile-* .push-*

bin-clean:
	rm -rf .go bin
