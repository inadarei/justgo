#!/bin/sh

# link pre-commit from .git/hooks to here so every checkout has it
gitpath=`git rev-parse --git-dir`
projectpath=`dirname "$gitpath"`
ln -s -f "$projectpath/.githooks/pre-commit" "$gitpath/hooks/pre-commit"
