## Build Image
FROM golang:1.19-alpine as build

ARG E2E_SCRIPT_NAME
RUN set -eux; apk add --no-cache ca-certificates build-base wget;
RUN apk add git
# Needed by github.com/zondax/hid
RUN apk add linux-headers

WORKDIR /chain4energy
RUN git clone -b master-1.2.0 https://github.com/chain4energy/c4e-chain.git /chain4energy

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

