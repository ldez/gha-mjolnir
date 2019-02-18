# Mjolnir

[![Release](https://img.shields.io/github/release/ldez/gha-mjolnir.svg?style=flat)](https://github.com/ldez/gha-mjolnir/releases)
[![Build Status](https://travis-ci.org/ldez/gha-mjolnir.svg?branch=master)](https://travis-ci.org/ldez/gha-mjolnir)
[![Docker Build Status](https://img.shields.io/docker/build/ldez/gha-mjolnir.svg)](https://hub.docker.com/r/ldez/gha-mjolnir/builds/)

Closes issues related to the merge of a pull request.

Useful:

- to close multiple issues related to a pull request.
- to close issues related to a pull request not based on the default branch (i.e. `master`).
By example when a branch is related to version (e.g. `v1.5`, `v2.0`, ...)

## Usage

```hcl
workflow "Auto close issues" {
  on = "pull_request"
  resolves = ["mjolnir-issues"]
}

action "mjolnir-issues" {
  uses = "docker://ldez/gha-mjolnir"
  secrets = ["GITHUB_TOKEN"]
  args = ""
}
```
