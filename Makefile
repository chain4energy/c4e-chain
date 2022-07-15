PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
 
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=c4e \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=c4ed \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=c4ecli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

#BUILD_FLAGS := -ldflags '$(ldflags)'

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

# include Makefile.ledger
all: install

build: go.sum
		@echo "--> Building c4ed"
		go build -o ./build/c4ed -mod=readonly $(BUILD_FLAGS)  ./cmd/c4ed

install: go.sum
		@echo "--> Installing c4ed"
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/c4ed

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -cover -mod=readonly $(PACKAGES)

