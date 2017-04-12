# `EK` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/ek.v7?status.svg)](https://godoc.org/pkg.re/essentialkaos/ek.v7) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/ek)](https://goreportcard.com/report/github.com/essentialkaos/ek) [![codebeat badge](https://codebeat.co/badges/3649d737-e5b9-4465-9765-b9f4ebec60ec)](https://codebeat.co/projects/github-com-essentialkaos-ek) [![Coverage Status](https://coveralls.io/repos/github/essentialkaos/ek/badge.svg?branch=develop)](https://coveralls.io/github/essentialkaos/ek?branch=develop) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

Auxiliary packages for Go.

## Platform support

Currently we support Linux and Mac OS X. Some packages have stubs for Windows (_for autocomplete_).

## Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.6+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get pkg.re/essentialkaos/ek.v7
```

If you want to update `EK` to latest stable release, do:

```
go get -u pkg.re/essentialkaos/ek.v7
```

## Packages

* [`arg`](https://godoc.org/pkg.re/essentialkaos/ek.v7/arg) - Package provides methods for working with command-line arguments
* [`color`](https://godoc.org/pkg.re/essentialkaos/ek.v7/color) - Package color provides methods for working with colors
* [`cron`](https://godoc.org/pkg.re/essentialkaos/ek.v7/cron) - Package provides methods for working with cron expressions
* [`csv`](https://godoc.org/pkg.re/essentialkaos/ek.v7/csv) - Package with simple (without any checks) CSV parser compatible with default Go parser
* [`easing`](https://godoc.org/pkg.re/essentialkaos/ek.v7/easing) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`env`](https://godoc.org/pkg.re/essentialkaos/ek.v7/env) - Package provides methods for working with environment variables
* [`errutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/errutil) - Package provides methods for working with errors
* [`fmtc`](https://godoc.org/pkg.re/essentialkaos/ek.v7/fmtc) - Package provides methods similar to fmt for colored output
* [`fmtutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/fmtutil) - Package provides methods for output formatting
* [`fsutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/fsutil) - Package provides methods for working with files on POSIX compatible systems (Linux / Mac OS X)
* [`hash`](https://godoc.org/pkg.re/essentialkaos/ek.v7/hash) - Package hash contains different hash algorithms and utilities
* [`httputil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/httputil) - Package provides methods for working with HTTP request/responses
* [`jsonutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/jsonutil) - Package provides methods for working with JSON data
* [`knf`](https://godoc.org/pkg.re/essentialkaos/ek.v7/knf) - Package provides methods for working with configs in KNF format
* [`kv`](https://godoc.org/pkg.re/essentialkaos/ek.v7/kv) - Package provides simple key-value structs
* [`log`](https://godoc.org/pkg.re/essentialkaos/ek.v7/log) - Package with an improved logger
* [`mathutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/mathutil) - Package with math utils
* [`netutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/netutil) - Package with network utils
* [`passwd`](https://godoc.org/pkg.re/essentialkaos/ek.v7/passwd) - Package passwd contains methods for working with passwords
* [`path`](https://godoc.org/pkg.re/essentialkaos/ek.v7/path) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://godoc.org/pkg.re/essentialkaos/ek.v7/pid) - Package for working with PID files
* [`pluralize`](https://godoc.org/pkg.re/essentialkaos/ek.v7/pluralize) - Package pluralize provides methods for pluralization
* [`rand`](https://godoc.org/pkg.re/essentialkaos/ek.v7/rand) - Package for generating random data
* [`req`](https://godoc.org/pkg.re/essentialkaos/ek.v7/req) - Package for working with HTTP request
* [`signal`](https://godoc.org/pkg.re/essentialkaos/ek.v7/signal) - Package for handling signals
* [`sliceutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/sliceutil) - Package with utils for working with slices
* [`sortutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/sortutil) - Package with utils for sorting slices
* [`spellcheck`](https://godoc.org/pkg.re/essentialkaos/ek.v7/spellcheck) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`strutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/strutil) - Package provides utils for working with strings
* [`system/process`](https://godoc.org/pkg.re/essentialkaos/ek.v7/system/process) - Package provides methods for getting information about active processes
* [`system`](https://godoc.org/pkg.re/essentialkaos/ek.v7/system) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://godoc.org/pkg.re/essentialkaos/ek.v7/terminal) - Package provides methods for working with user input
* [`terminal/window`](https://godoc.org/pkg.re/essentialkaos/ek.v7/terminal/window) - Package provides methods for working terminal window
* [`timeutil`](https://godoc.org/pkg.re/essentialkaos/ek.v7/timeutil) - Package with time utils
* [`tmp`](https://godoc.org/pkg.re/essentialkaos/ek.v7/tmp) - Package provides methods for working with temporary data
* [`usage`](https://godoc.org/pkg.re/essentialkaos/ek.v7/usage) - Package provides methods for rendering info for command-line tools
* [`uuid`](https://godoc.org/pkg.re/essentialkaos/ek.v7/uuid) - Package provides methods for generating version 4 and 5 UUID's
* [`version`](https://godoc.org/pkg.re/essentialkaos/ek.v7/version) - Package provides methods for parsing semver version info

## Projects with `EK`

* [Deadline](https://github.com/essentialkaos/deadline) - Simple utility for controlling application working time
* [GoHeft](https://github.com/essentialkaos/goheft) - Utility for listing sizes of all used static libraries compiled into golang binary
* [GoMakeGen](https://github.com/essentialkaos/gomakegen) - Utility for generating makefiles for golang applications
* [MDToc](https://github.com/essentialkaos/mdtoc) - Utility for generating table of contents for markdown files
* [Mockka](https://github.com/essentialkaos/mockka) - Mockka is a simple utility for mocking HTTP API's
* [RBInstall](https://github.com/essentialkaos/rbinstall) - Utility for installing prebuilt ruby to RBEnv
* [redis-cli-monitor](https://github.com/essentialkaos/redis-cli-monitor) - Tiny redis client for renamed MONITOR commands
* [SHDoc](https://github.com/essentialkaos/shdoc) - Tool for viewing and exporting docs for shell scripts
* [SourceIndex](https://github.com/essentialkaos/source-index) - Utility for generating an index for source archives
* [SSLScan Client](https://github.com/essentialkaos/sslcli) - Pretty awesome command-line client for public SSLLabs API
* [Terrafarm](https://github.com/essentialkaos/terrafarm) - Utility for working with terraform based rpmbuilder farm
* [Yo](https://github.com/essentialkaos/yo) - Command-line YAML processor

## Build Status

| Branch | TravisCI |
|--------|----------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=master)](https://travis-ci.org/essentialkaos/ek) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=develop)](https://travis-ci.org/essentialkaos/ek) |

## Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

## License

[EKOL](https://essentialkaos.com/ekol)
