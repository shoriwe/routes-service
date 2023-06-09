# routes-service

[![Build](https://github.com/shoriwe/routes-service/actions/workflows/build.yaml/badge.svg)](https://github.com/shoriwe/routes-service/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/shoriwe/routes-service/branch/main/graph/badge.svg?token=SMCRWGJ4C5)](https://codecov.io/gh/shoriwe/routes-service)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shoriwe/routes-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoriwe/routes-service)](https://goreportcard.com/report/github.com/shoriwe/routes-service)

## Run

You can find compiled binaries on the [Release](https://github.com/shoriwe/routes-service/releases) section.

### Go install

```shell
go install github.com/shoriwe/routes-service@latest
```

After installing the binary you can run a temp server with the database running in memory with:

```shell
routes-service 127.0.0.1:5000
```

### Docker compose

A ready to go `docker-compose.yaml` can be found at the root of the repository

```shell
docker compose -f docker-compose.yaml up -d
```

You can access the service at `127.0.0.1:5000`.

### Docker image

```shell
docker run -p 5000:5000 -d ghcr.io/shoriwe/routes-service:latest
```

## Documentation

You will find the entire documentation for this project at `docs/`.

| File                               | Description                                                                                      |
| ---------------------------------- | ------------------------------------------------------------------------------------------------ |
| [docs/README.md](docs/README.md)   | Detailed description of the documentation files.                                                 |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Rules for contributing to this repository. Including commit nomenclature, code quality and more. |
| [CHANGELOG.md](CHANGELOG.md)       | This file is autogenerated by the Build Action's pipeline, should never be modified manually.    |

## Testing

Make sure you have the testing PostgreSQL service running. You can setup one easily with:
```shell
docker compose -f testing.docker-compose.yaml up -d
```

Finally:

```go
go test -count=1 -v ./...
```

## Coverage

|                                                                           Sunburst                                                                           |                                                                         Grid                                                                         |
| :----------------------------------------------------------------------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------: |
| [![sunburst](https://codecov.io/gh/shoriwe/routes-service/branch/main/graphs/sunburst.svg?token=SMCRWGJ4C5)](https://app.codecov.io/gh/shoriwe/routes-service) | [![grid](https://codecov.io/gh/shoriwe/routes-service/branch/main/graphs/tree.svg?token=SMCRWGJ4C5)](https://app.codecov.io/gh/shoriwe/routes-service) |
