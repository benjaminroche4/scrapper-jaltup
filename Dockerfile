FROM ubuntu:24.10 AS base

# install updates and tools
RUN apt-get update && apt-get install -y bash make git wget curl zip gzip

# download and install go verion 1.22.10
ENV GO_VERSION=1.22.10
RUN curl -sSLO https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz && \
  rm -rf /usr/local/go && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
  rm -rf go$GO_VERSION.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

FROM base AS builder
WORKDIR /work
COPY . .

FROM builder AS build
ARG BUILD_VERSION
RUN BUILD_VERSION=${BUILD_VERSION} make

FROM base
ARG DATABASE_NAME
ARG DATABASE_HOST
ARG DATABASE_PORT
ARG DATABASE_USERNAME
ARG DATABASE_PASSWORD

ENV DATABASE_NAME=$DATABASE_NAME
ENV DATABASE_HOST=$DATABASE_HOST
ENV DATABASE_PORT=$DATABASE_PORT
ENV DATABASE_USERNAME=$DATABASE_USERNAME
ENV DATABASE_PASSWORD=$DATABASE_PASSWORD

COPY --from=build /work/bin /opt/bin

ENTRYPOINT ["/opt/bin/scrapper", "lba"]