## Build Image
FROM golang:1.19-alpine as build

ARG E2E_SCRIPT_NAME
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

RUN BUILD_TAGS=muslc LINK_STATICALLY=true E2E_SCRIPT_NAME=${E2E_SCRIPT_NAME} make build-e2e-script

## Deploy image
FROM ubuntu

# Args only last for a single build stage - renew
ARG E2E_SCRIPT_NAME

COPY --from=build /chain4energy/build/${E2E_SCRIPT_NAME} /bin/${E2E_SCRIPT_NAME}

ENV HOME /chain4energy
WORKDIR $HOME

# Docker ARGs are not expanded in ENTRYPOINT in the exec mode. At the same time,
# it is impossible to add CMD arguments when running a container in the shell mode.
# As a workaround, we create the entrypoint.sh script to bypass these issues.
RUN echo "#!/bin/bash\n${E2E_SCRIPT_NAME} \"\$@\"" >> entrypoint.sh && chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]

