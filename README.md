# slp

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/slp)](https://github.com/koki-develop/slp/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/slp/ci.yml?logo=github)](https://github.com/koki-develop/slp/actions/workflows/ci.yml)
[![Maintainability](https://img.shields.io/codeclimate/maintainability/koki-develop/slp?style=flat&logo=codeclimate)](https://codeclimate.com/github/koki-develop/slp/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/slp)](https://goreportcard.com/report/github.com/koki-develop/slp)
[![LICENSE](https://img.shields.io/github/license/koki-develop/slp)](./LICENSE)

sleep command with rich progress bar.

![demo](./docs/demo.gif)

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [LICENSE](#license)

## Installation

### Homebrew

```console
$ brew install koki-develop/tap/slp
```

### `go install`

```console
$ go install github.com/koki-develop/slp@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/slp/releases/latest).

## Usage

### Basic

```sh
$ slp [time]
# e.g.
$ slp 3
```

![](./docs/demo.gif)

### Customize Color

`--gradient` flag can be used to apply a gradient between the two colors.

```sh
# e.g.
$ slp 3 --gradient "#000000,#ffffff"
```

![](./docs/gradient.gif)

You can also set a single color with the `-color` flag.

```sh
# e.g.
$ slp 3 --color "#ff0000"
```

![](./docs/color.gif)

## LICENSE

[MIT](./LICENSE)
