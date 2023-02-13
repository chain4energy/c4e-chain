# syntax=docker/dockerfile:1

## Build Image
FROM golang:1.19-alpine as build

RUN set -eux; apk add --no-cache ca-certificates build-base wget;
RUN apk add git
# Needed by github.com/zondax/hid
RUN apk add linux-headers

WORKDIR /chain4energy
COPY . /chain4energy

RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ubuntu:22.04

COPY --from=build /chain4energy/build/c4ed /bin/c4ed

ENV HOME /chain4energy
WORKDIR $HOME

# Docker ARGs are not expanded in ENTRYPOINT in the exec mode. At the same time,
# it is impossible to add CMD arguments when running a container in the shell mode.
# As a workaround, we create the entrypoint.sh script to bypass these issues.
RUN echo "#!/bin/bash\n${E2E_SCRIPT_NAME} \"\$@\"" >> entrypoint.sh && chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]

