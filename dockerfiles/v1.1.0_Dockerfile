# syntax=docker/dockerfile:1

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
RUN git clone -b master-E2E-1.1.0 https://github.com/chain4energy/c4e-chain.git /chain4energy

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
