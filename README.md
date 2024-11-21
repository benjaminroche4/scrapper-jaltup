# S c r a p p e r   -   J a l t u p

## Prerequisites

### Golang

Install Golang 1.22: [https://go.dev/dl/](https://go.dev/dl/)

- [Windows 386](https://go.dev/dl/go1.22.10.windows-386.msi)
- [Windows amd64](https://go.dev/dl/go1.22.10.windows-amd64.msi)
- [Linux 386](https://go.dev/dl/go1.22.10.linux-386.tar.gz)
- [Linux amd64](https://go.dev/dl/go1.22.10.linux-amd64.tar.gz)

### Make

To install *make* on Windows, download and unzip the following files:

- [Binaries](https://gnuwin32.sourceforge.net/downlinks/make-bin-zip.php)
- [Dependencies](https://gnuwin32.sourceforge.net/downlinks/make-dep-zip.php)

To install *make* on linux:

```sh
sudo apt-get install -y make
```

### MySQL

#### Windows

Download MSI installer for windows [ici](https://dev.mysql.com/downloads/file/?id=536356)

#### Ubuntu

```sh
sudo apt-get install -y mysql-server 
```

## Environment Variables

| NAME               | DESCRIPTION                                            | DEFAULT VALUE            |
|--------------------|--------------------------------------------------------|--------------------------|
| DATABASE_NAME      | The Database schema name to use                        | jaltup                   |
| DATABASE_HOST      | The Database machine host name or ip address           | localhost                |
| DATABASE_PORT      | The Database machine port number to connnect to        | 3306                     |
| DATABASE_USERNAME  | The Database username to use for connection            | jaltup                   |
| DATABASE_PASSWORD  | The Database password to use for connection            |                          |

## Usage

```text
NAME:
   scrapper - A new cli application

USAGE:
   scrapper <command> [options]

VERSION:
   1.0.0

DESCRIPTION:
   Jaltup Scrapper

COMMANDS:
   categories  retrieve all categories from database
   clean       clean(empty) all tables, this action is irreversible
   companies   retrieve all companies from database
   count       count rows for all tables
   lba         fetch all offers from 'la bonne alternance' source
   offers      retrieve all offers from database
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dbname value          database name (default: "jaltup") [%DATABASE_NAME%]
   --host value, -H value  database host name (default: "localhost") [%DATABASE_HOST%]
   --pass value, -p value  database password [%DATABASE_PASSWORD%]
   --port value, -P value  database port number (default: 3306) [%DATABASE_PORT%]
   --user value, -u value  database username (default: "root") [%DATABASE_USERNAME%]
   --help, -h              show help
   --version, -v           print the version
```

## Compilation

### Locally

```sh
make build
```

#### Run tests

```sh
make test-short
```

#### Run linter

```sh
make lint
```

### With **Docker**

#### Build the docker image

```sh
BUILD_VERSION="1.0.0"
docker build --tag "scrapper-jaltup" --build-arg="BUILD_VERSION=${BUILD_VERSION}" .
```

#### Start a container

```sh
 docker run --rm \
   --network=host \
   --env DATABASE_NAME="jaltup" \
   --env DATABASE_HOST="127.0.0.1" \
   --env DATABASE_PORT="3306" \
   --env DATABASE_USERNAME="jaltup" \
   --env DATABASE_PASSWORD="jaltup" \
   --name="scrapper-jaltup" \
   "scrapper-jaltup"
```
