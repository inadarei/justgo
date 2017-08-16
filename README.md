# Jump-Start Go (JustGo)

Skeleton project for jump-starting a Go-powered microservice development with Docker and Go best-practices.

To learn more: [https://justgo.rocks](https://justgo.rocks)

## INSTALLATION 

Easiest way to create a new project skeleton is to install JustGo CLI tool. There's no necessity to install Go on your machine, since the setup provides fully functioning Go environment in a Docker container. 

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

After you have installed the CLI tool, to create a skeleton of a new project, just run:

```
> justgo
```

You can see various options by running `justgo -h`. To learn how to run a project, once it is created, see [here](https://github.com/inadarei/justgo-microservice#how-to-run-a-project-once-created)

## Contributing to the CLI Interface

If you are interested in contributing to the development of this CLI tool,
following are the instructions for setting up a dev environment:

## Prerequisites
 - [dep](https://github.com/golang/dep)

## Installation

```BASH
> git clone https://github.com/inadarei/justgo.git
> cd justgo/
> dep ensure
> ./go-wrapper install
> ./go-wrapper run <someFolderToTestInstallTo>
```

## Warning for VS Code Users

If you are using VS Code with Go tooling, you will want to change the default 
`"go.formatTool": "goreturns",` formatter to `"go.formatTool": "gofmt",` instead since the former
seems unable to properly detect the usage of uuid in the code and keeps removing the uuid package's
import statement from code, making it error-out during a build. Gofmt has no such issues.
## License

[MIT](LICENSE)
