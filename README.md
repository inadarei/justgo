# Jump-Start Go (JustGo)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/inadarei/justgo/issues)
[![Go project version](https://badge.fury.io/go/github.com%2Finadarei%2Fjustgo.svg)](https://badge.fury.io/go/github.com%2Finadarei%2Fjustgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/inadarei/justgo)](https://goreportcard.com/report/github.com/inadarei/justgo-microservice)

A helpful builder for a light-weight Go [skeleton project](https://github.com/inadarei/justgo-microservice) takes care of a lot of boilerplate in jump-starting a Go-powered microservice development with Docker and Go best-practices.

To learn more: [https://justgo.rocks](https://justgo.rocks)

### Features:

1. Isolated, project-specific Go environments in a container. No cross-project dependency issues.
2. No mess with configuring GOPATH across projects!
3. Code hot-reloading out of the box!
4. Ready to ship as a container in production, when you are done working with it.
5. Future-proof choice of [dep](https://github.com/golang/dep) for dependency-management

## INSTALLATION 

Easiest way to create a new project skeleton is to install JustGo CLI tool.
There's no necessity to install Go on your machine, since the setup provides
fully functioning Go environment in a Docker container.

If you already have Go on your machine, you can install the CLI tool with:

```
> go get github.com/inadarei/justgo
```

or you can install it using Homebrew, even if you don't have Go:

```
> brew tap inadarei/casks
> brew install justgo
```

## USAGE

After you have installed the CLI tool, to create a skeleton of a new project,
just run:

```
> justgo
```

You can see various options by running `justgo -h`. To learn how to run a
project, once it is created, see
[here](https://github.com/inadarei/justgo-microservice#how-to-run-a-project-once-created)

## Upgrade with Brew

```BASH
> brew update
> brew upgrade justgo
```

## Contributing to the CLI Interface

If you are interested in contributing to the development of this CLI tool,
following are the instructions for setting up a dev environment:

### Prerequisites

 - Latest Go version (al teast 1.8+)
 - Properly set up `$GOPATH` and `GOPATH/bin` added to `$PATH`
 - Go's Dep tool installed:
     - Install via: `go get -u github.com/golang/dep/cmd/dep`

### Dev Workspace Setup

```BASH
> cd $GOPATH/src
> mkdir -p github.com/inadarei/
> git clone https://github.com/inadarei/justgo.git
> cd justgo/
> dep ensure
> go run justgo.go <someFolderToTestInstallTo>
```

### Warning for VS Code Users

If you are using VS Code with Go tooling, you will want to change the default
`"go.formatTool": "goreturns",` formatter to `"go.formatTool": "gofmt",` instead
since the former seems unable to properly detect the usage of uuid in the code
and keeps removing the uuid package's import statement from code, making it
error-out during a build. Gofmt has no such issues.

## License

[MIT](LICENSE)
