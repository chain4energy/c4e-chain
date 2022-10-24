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
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ -Seed=589 -v \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -Verbose=true

test-simulation-benchmark-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ -v -Seed=589 \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) \
		-timeout 24h -cpuprofile cpu.out -memprofile mem.out

open-cpu-profiler-result:
	@go tool pprof cpu.out
# HOW TO READ - https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph

open-memory-profiler-result:
	@go tool pprof mem.out