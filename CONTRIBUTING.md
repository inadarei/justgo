# Contributing to the CLI Interface

Contributions are always welcome, no matter how large or small. Substantial
feature requests should be proposed as an
[RFC](https://github.com/apiaryio/api-blueprint-rfcs/blob/master/template.md).
Before contributing, please read the [code of
conduct](https://github.com/inadarei/justgo/blob/master/CODE_OF_CONDUCT.md).

## Setup

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
> .githooks/install.sh
> dep ensure
> go run justgo.go <someFolderToTestInstallTo>
```

### Warning for VS Code Users

If you are using VS Code with Go tooling, you will want to change the default
`"go.formatTool": "goreturns",` formatter to `"go.formatTool": "gofmt",` instead
since the former seems unable to properly detect the usage of uuid in the code
and keeps removing the uuid package's import statement from code, making it
error-out during a build. Gofmt has no such issues.

## Pull Requests

We actively welcome your pull requests.

1. Create an Issue or RFC for your contribution.
1. Fork the repo and create your branch from `master`.
1. If you've added code that should be tested, add tests.
1. Update the documentation.
1. Ensure the test suite passes.
1. Run gofmt and goimports.

## License

By contributing to Justgo, you agree that your contributions will be licensed
under its [MIT license](LICENSE).
