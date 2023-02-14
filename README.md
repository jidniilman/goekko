# [WIP] Goekko
[![Golang CI Lint](https://github.com/jidniilman/goekko/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/jidniilman/goekko/actions/workflows/golangci-lint.yml)
[![go-ci](https://github.com/jidniilman/goekko/actions/workflows/go-ci.yml/badge.svg)](https://github.com/jidniilman/goekko/actions/workflows/go-ci.yml)
[![Deployment](https://github.com/jidniilman/goekko/actions/workflows/deployment.yml/badge.svg)](https://github.com/jidniilman/goekko/actions/workflows/deployment.yml)
----
Goekko is the backend service boilerplate made and used personally with scalability and modularity in mind using Repository-Service Pattern.

There are some references and inspirations to build this boilerplate:
- [nixsolutions/golang-echo-boilerplate](https://github.com/nixsolutions/golang-echo-boilerplate)
- [triaton/go-echo-boilerplate](https://github.com/triaton/go-echo-boilerplate)

## What's included?
- This project built with Go 1.20
- Web server using [Labstack Echo 4](https://github.com/labstack/echo)
- Unit Test framework using [Testify](https://github.com/stretchr/testify)
- Containerized with Dockerfile
- Integration with GitHub Actions
- Linting with Golang CI Lint
- Optional Hasura Integration and GraphQL Client

## Code Checklist
- [ ] Refactor
- [ ] Init the repo and push to VCS
- [ ] Change some information in Readme
- [ ] Don't forget to write Unit Tests
- [ ] Edit the UNIX Executable Script `/scripts/build`
- [ ] Edit the Makefile
- [ ] Edit the Dockerfile
- [ ] Edit GitHub Actions Workflow
- [ ] Lint with Golang CI Lint
- [ ] (Optional) Write Benchmark Result

## Quickstart
Basic commands of the service:
```
# Build
make build

# Start (build and run the built binaries)
make start

# Run
make run

# Test
make test
```


## How to Use
WIP

## Coding Convention
Golang is strict by default, and we are using some guidelines from :
- [Effective Go](https://go.dev/doc/effective_go)
- [CockroachDB Engineering Standards](https://wiki.crdb.io/wiki/spaces/CRDB/pages/181371303/Go+Golang+coding+guidelines)
- [Golang CI Lint](https://golangci-lint.run/)
- GoLand IDE formatting

## Unit Tests
Mainly for our unit testing, we are using table driven test approach. But we are using little sample to do our tests.

Run all unit tests with:
```
# run normal test
go test ./...
# or
make test

# with verbose output.
go test -v ./...

# test individual packages
go test -v <package-name>
```

Package Name (main entrypoint are excluded):
```
jidniilman/goekko/internal/cli      # Handling CLI application and execution
jidniilman/goekko/internal/question # Our question model
jidniilman/goekko/pkg/command       # Our data communication protocol
jidniilman/goekko/pkg/utils         # Some useful helpers
```
To see the full package list, use 
```
go list ./...
```

## Benchmark
Only `/pkg/command` and `/pkg/utils` that have benchmarked. We can do benchmark for `/internal` but it is not necessary.

Run benchmark with:
```
go test -v ./... -bench=. -benchmem
```

## Project Layout/Structure
In this project, we are using the layout standards from [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
We removed some of unused folders such as: `/api`, `/web`, `/deployments`, `/test`, `/configs`, `/docs`, and `/build`

### `/cmd`
Main applications for this project. The directory name for each application match the name of the executable (e.g., `/cmd/goekko`).
It's common to have a small main function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

### `/internal`
Private application and library code. This is the code you don't want others importing in their applications or libraries. 
Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 release notes for more details. 

### `/pkg`
Library code that's ok to use by external applications (e.g., `/pkg/utils`). Other projects will import these libraries expecting them to work.
Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. 
The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. 

### `/bin`
Executable UNIX Scripts and Binary location.

## Project Root Files

### `.gitignore`
List of git ignored files and folders.

### `go.mod`
List of Go Modules configurations such as module name, module list, go version and more.

### `.github/workflows/*.yml`
List of GitHub Actions workflow we are using.

## Godoc
Install godoc first (godoc isn't included in version `>=1.13`)
```
go install golang.org/x/tools/cmd/godoc@latest
```

### Serve via HTTP
Run it directly with `godoc -http localhost:6060`. And open this [link](http://localhost:6060/pkg/jidniilman/goekko/?m=all)

## Golang CI Lint
Golang CI Lint is included with GitHub Actions.

But if you want to lint locally, Install Golang CI Lint first:
```
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 

# test it
golangci-lint --version
```
Run lint with:
```
golangci-lint run
```
Check [https://golangci-lint.run/](https://golangci-lint.run/) for more commands and usages.

## Docker
We provide two-stage build docker image for you.

Build docker with:
```
docker build -t jidniilman/goekko:latest .
```
See created docker images with:
```
docker images
```
Run the docker image that we previously created:
```
docker run -it jidniilman/goekko:latest
```