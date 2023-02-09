# syntax=docker/dockerfile:1

ARG BASE_IMG_TAG=nonroot

# --------------------------------------------------------
# Build
# --------------------------------------------------------

FROM golang:1.18.2-alpine3.15 as build

RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add git
# Needed by github.com/zondax/hid
RUN apk add linux-headers

WORKDIR /chain4energy
COPY . /chain4energy

RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM gcr.io/distroless/base-debian11:${BASE_IMG_TAG}

COPY --from=build /chain4energy/build/c4ed /bin/c4ed

ENV HOME /chain4energy
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317

ENTRYPOINT ["c4ed"]
CMD [ "start" ]
