PACKAGES=$(shell go list ./... | grep -v '/simulation\|e2e')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true

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

ifeq (cleveldb,$(findstring cleveldb,$(C4E_BUILD_OPTIONS)))
  build_tags += gcc
else ifeq (rocksdb,$(findstring rocksdb,$(C4E_BUILD_OPTIONS)))
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=c4e \
	-X github.com/cosmos/cosmos-sdk/version.AppName=c4ed \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=c4ed \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=c4ecli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq (cleveldb,$(findstring cleveldb,$(C4E_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
else ifeq (rocksdb,$(findstring rocksdb,$(C4E_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
ifeq (,$(findstring nostrip,$(C4E_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(C4E_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

release = GOOS=$(1) GOARCH=$(2) go build -o ./build/c4ed -mod=readonly $(BUILD_FLAGS)  ./cmd/c4ed
tar = cd build && tar -cvzf c4ed_$(tag)_$(1)_$(2).tar.gz c4ed && rm c4ed


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
	@go test -coverprofile=coverage.out -mod=readonly $(PACKAGES)

release:
	@echo "--> Prepare release linux amd64"
	$(call release,linux,amd64)
	$(call tar,linux,amd64)
	@echo "--> Prepare release linux arm64"
	$(call release,linux,arm64)
	$(call tar,linux,arm64)
	@echo "--> Prepare release darwin amd64"
	$(call release,darwin,amd64)
	$(call tar,darwin,amd64)

# blockchain simulation tests

SIM_NUM_BLOCKS = 100
SIM_BLOCK_SIZE = 200
SIM_COMMIT = true
SIMAPP = ./app

test-simulation-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ -Seed=589 -v -Period=1 -PrintAllInvariants \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -Verbose=true

test-simulation-benchmark-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ -v -Seed=589 -Period=1 -PrintAllInvariants \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) \
		-timeout 24h -cpuprofile cpu.out -memprofile mem.out

test-simulation-import-export:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimTest$$ -Seed=589 -v -Period=1 -PrintAllInvariants \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -Verbose=true

open-cpu-profiler-result:
	@go tool pprof cpu.out
# HOW TO READ - https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph

open-memory-profiler-result:
	@go tool pprof mem.out

#E2E framework
#Environments description
#C4E_E2E_SKIP_UPGRADE - skip the upgrade tests
#C4E_E2E_SKIP_IBC - skip the IBC tests
#C4E_E2E_DEBUG_LOG - debug logs and print them onto the screen
#C4E_E2E_SKIP_STATE_SYNC - skip state sync test
#C4E_E2E_FORK_HEIGHT - determine if this upgrade is a fork
#C4E_E2E_SKIP_CLEANUP - skip cleaning up Docker resources in teardown
#C4E_E2E_UPGRADE_VERSION - environment variable name to determine what version we are upgrading to
#C4E_E2E_SKIP_PARAMS_CHANGE - skip params change tests

PACKAGES_E2E=$(shell go list ./... | grep '/e2e')
BUILDDIR ?= $(CURDIR)/build
E2E_UPGRADE_VERSION="v1.0.1"
E2E_SCRIPT_NAME=chain

test-e2e: C4E_E2E=True e2e-setup test-e2e-ci

run-chain: e2e-setup
	@VERSION=$(VERSION) C4E_E2E_DEBUG_LOG=True C4E_E2E_SKIP_CLEANUP=False C4E_E2E_SKIP_UPGRADE=True C4E_E2E_SKIP_IBC=True go test -mod=readonly -timeout=25m -v ./tests/e2e -testify.m ^TestAppRun

test-e2e-debug: e2e-setup
	@VERSION=$(VERSION) C4E_E2E_DEBUG_LOG=True C4E_E2E_UPGRADE_VERSION=$(E2E_UPGRADE_VERSION) C4E_E2E_SKIP_CLEANUP=False go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E) -count=1

test-e2e-short: e2e-setup
	@VERSION=$(VERSION) C4E_E2E_DEBUG_LOG=True C4E_E2E_SKIP_UPGRADE=True C4E_E2E_SKIP_IBC=True C4E_E2E_SKIP_STATE_SYNC=True C4E_E2E_SKIP_CLEANUP=False go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E) -count=1

e2e-setup: e2e-remove-resources
	@echo Finished e2e environment setup, ready to start the test

e2e-check-image-sha:
	tests/e2e/scripts/run/check_image_sha.sh

e2e-remove-resources:
	tests/e2e/scripts/run/remove_stale_resources.sh

build-e2e-script:
	mkdir -p $(BUILDDIR)
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./tests/e2e/initialization/$(E2E_SCRIPT_NAME)

# Docker commands

docker-build-debug:
	@docker build -t chain4energy:debug --build-arg BASE_IMG_TAG=debug -f Dockerfile .

docker-build-e2e-init-chain:
	@docker build -t chain4energy-e2e-chain-init:debug --build-arg E2E_SCRIPT_NAME=chain -f tests/e2e/initialization/init.Dockerfile .

docker-build-e2e-init-node:
	@docker build -t chain4energy-e2e-node-init:debug --build-arg E2E_SCRIPT_NAME=node -f tests/e2e/initialization/init.Dockerfile .
