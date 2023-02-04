# GoPress

OpenSource Golang Blog Service

## How to run 

### Building from the source

Service can be built from source code and executed from local environment independent of operating system type and version. Only limitation is the presence of [Golang](https://golang.org) at this platform.

#### Build prerequisites
* Go language v1.18+
* GNU Make
* Git
* Docker (only to build docker binaries)
* Linux or MacOS with Unix-like shell

#### How to build

To build binary execute following command:

```bash
make
```

This will execute series of tests to validate source code correctness. Linting will be also checked for code quality.

### Running from binary distribution

Binary version can be downloaded from Docker repository and executed immediately.

```bash
docker run -it -p 8080:8080 docker.io/mateuszmierzwinski/gopress:latest
```