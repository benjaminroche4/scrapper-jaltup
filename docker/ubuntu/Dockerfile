FROM ubuntu:24.10 AS base

# install updates and tools
RUN apt-get update && \
  apt-get install -y bash make gcc g++ \
  build-essential autoconf automake pkg-config bison flex \
  libboost-all-dev libevent-dev libssl-dev libtool \
  git wget curl zip gzip

FROM base AS dev
# download and install go verion 1.22.5
ENV GO_VERSION=1.22.5
RUN curl -sSLO https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz && \
  rm -rf /usr/local/go && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
  rm -rf go$GO_VERSION.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV CGO_ENABLED=0

FROM dev AS builder
WORKDIR /work
COPY . .

FROM builder AS build
ARG BUILD_VERSION
ENV CGO_ENABLED=0
RUN BUILD_VERSION=${BUILD_VERSION} make

FROM base
COPY --from=build /work/bin /work/bin
