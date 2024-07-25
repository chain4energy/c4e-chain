ARG BASE_IMG_TAG=nonroot

# --------------------------------------------------------
# Build
# --------------------------------------------------------

FROM golang:1.19-alpine as build

RUN set -eux; apk add --no-cache ca-certificates build-base wget;
RUN apk add git
# Needed by github.com/zondax/hid
RUN apk add linux-headers

WORKDIR /chain4energy
RUN git clone -b v1.4.0-rc4 https://github.com/chain4energy/c4e-chain.git /chain4energy
RUN ARCH=$(uname -m) && WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$ARCH.a \
        -O /lib/libwasmvm_muslc.a && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.$ARCH | cut -d ' ' -f 1)

RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ubuntu:22.04

COPY --from=build /chain4energy/build/c4ed /bin/c4ed

ENV HOME /chain4energy
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317

ENTRYPOINT ["c4ed"]
CMD [ "start" ]
