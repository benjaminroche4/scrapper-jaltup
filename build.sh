#!/bin/bash

image_prefix="build"
container_prefix="container"
default_target="centos"
target="$1"
shift
if [ "$target" == "" ]; then
  target=$default_target
fi
image="$image_prefix-$target"
container="$container_prefix-$image"
verbose=1

function exitError() {
  echo "$1" 1>&2
  exit 1
}

function print() {
  if [ $verbose -eq 1 ]; then
    echo "$1" >&1
  fi
}

function requirements() {
  if ! command -v docker >/dev/null 2>&1; then
    exitError "This script requires 'docker' command."
  fi
}

function build() {
  print "Building image..."

  BUILD_VERSION=$(git describe --abbrev=0 --tags 2>/dev/null |
    sed -rn 's/([0-9]+(.[0-9]+){2}).*/\1/p')

  docker build \
    --network host \
    --tag "$image" \
    --build-arg="BUILD_VERSION=${BUILD_VERSION}" \
    --file "docker/$target/Dockerfile" .
}

function start() {
  print "Starting container: $container..."
  docker run -dit --name="$container" "$image" bash

  sleep 2
  sleeping=0
  while [ $sleeping -eq 0 ]; do
    sleep 1
    sleeping=$(docker exec "$container" ps -eaf | grep -v 'ps' | grep -c bash)
    echo -n "."
  done
  echo
}

function copy() {
  print "Copying generated runtime files..."

  mkdir bin 2>/dev/null
  docker cp "$container:/work/bin/." "bin/$target"
}

function stop() {
  id=$(docker ps --filter "name=$container" --format "{{.ID}}")

  if [ "$id" != "" ]; then
    print "Stopping container: $container..."

    if ! docker stop "$container" 1>/dev/null; then
      exitError "Failed to stop container: $container"
    fi
  fi
}

function remove() {
  id=$(docker ps -all --filter "name=$container" --format "{{.ID}}")

  if [ "$id" != "" ]; then
    print "Removing container: $container."

    if ! docker rm "$container" 1>/dev/null; then
      exitError "Failed to remove container: $container"
    fi
  fi
}

requirements

build

start

copy

stop

remove
