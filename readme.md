# Mjolnir

[![Release](https://img.shields.io/github/release/ldez/gha-mjolnir.svg?style=flat)](https://github.com/ldez/gha-mjolnir/releases)
[![Build Status](https://github.com/ldez/gha-mjolnir/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/gha-mjolnir/actions)
[![Docker](https://img.shields.io/badge/Docker-available-blue.svg)](https://hub.docker.com/r/ldez/gha-mjolnir/)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)

Close issues related to the merge of a pull request.

Useful:

- to close multiple issues related to a pull request.
- to close issues related to a pull request not based on the default branch (ex: `master`, `main`).
  For example, when a branch is related to version (e.g. `v1.5`, `v2.0`, ...)
- To add the same milestone defined on the PR to closed issues.

## Supported Syntax

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

```yaml
name: Close issues related to a merged pull request based on master branch.

on:
  pull_request:
    types: [closed]
    branches:
      - master

jobs:
  closeIssueOnPrMergeTrigger:

    runs-on: ubuntu-latest

    steps:
      - name: Closes issues related to a merged pull request.
        uses: ldez/gha-mjolnir@v1.4.1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```
