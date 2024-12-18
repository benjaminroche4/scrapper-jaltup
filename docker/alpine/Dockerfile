
FROM golang:1.22.5-alpine3.20 AS base

# install updates and tools
RUN apk update && \
  apk add --no-cache bash make rsync gcc g++ musl-dev binutils pkgconfig \
  autoconf automake build-base binutils cmake libgcc libtool linux-headers \
  make bison flex git wget curl zip unzip gzip 

FROM base AS builder
# setup environment
WORKDIR /work
COPY . .

FROM builder AS build
ARG BUILD_VERSION
RUN BUILD_VERSION=${BUILD_VERSION} make

FROM base
COPY --from=build /work/bin /work/bin
