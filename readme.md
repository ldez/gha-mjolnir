# Mjolnir

[![Release](https://img.shields.io/github/release/ldez/gha-mjolnir.svg?style=flat)](https://github.com/ldez/gha-mjolnir/releases)
[![Build Status](https://travis-ci.com/ldez/gha-mjolnir.svg?branch=master)](https://travis-ci.com/ldez/gha-mjolnir)
[![Docker](https://img.shields.io/badge/Docker-available-blue.svg)](https://hub.docker.com/r/ldez/gha-mjolnir/)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)

Close issues related to the merge of a pull request.

Useful:

- to close multiple issues related to a pull request.
- to close issues related to a pull request not based on the default branch (i.e. `master`).
For example, when a branch is related to version (e.g. `v1.5`, `v2.0`, ...)

## Supported Syntaxes

- prefixes (case insensitive): `close`, `closes`, `closed`, `fix`, `fixes`, `fixed`, `resolve`, `resolves`, `resolved`
- issues references separators (can be mixed): ` ` (space), `,` (period)
- prefix and issues references can be separated by: ` ` (space), `:` (colon), or both.

Examples:

```
Fixes #1,#2,#3
close #1, #2, #3
fix #1 #2 #3
resolve #1,#2 #3
Resolves: #1,#2,#3
closed : #1, #2, #3
```

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
