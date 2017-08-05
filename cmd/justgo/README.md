## Instructions for developing the CLI Interface

If you are interested in contributing to the development of this CLI tool,
following are the instructions for setting up a dev environment:

## Installation

```BASH
> git clone https://github.com/inadarei/justgo.git
> cd justgo/cmd/justgo
> glide install
> ./go-wrapper install
> ./go-wrapper run <someFolderToTestInstallTo>
```

## Warning for VS Code Users

If you are using VS Code with Go tooling, you will want to change the default 
`"go.formatTool": "goreturns",` formatter to `"go.formatTool": "gofmt",` instead since the former
seems unable to properly detect the usage of uuid in the code and keeps removing the uuid package's
import statement from code, making it error-out during a build. Gofmt has no such issues.