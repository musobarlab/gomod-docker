### Using Go Mod with Docker

#### Getting started

Check go version, you need `go version 1.11` or higher
```shell
$ go version
$ go version go1.11.2 darwin/amd64
```

Build Docker image
```shell
$ make docker
$ docker run -p 8080:8080 gomod-docker
```