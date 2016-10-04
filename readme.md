# `EK` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/ek.v4?status.svg)](https://godoc.org/pkg.re/essentialkaos/ek.v4) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/ek)](https://goreportcard.com/report/github.com/essentialkaos/ek) [![codebeat badge](https://codebeat.co/badges/3649d737-e5b9-4465-9765-b9f4ebec60ec)](https://codebeat.co/projects/github-com-essentialkaos-ek)

Auxiliary packages for Go.

### Platform support

Currently we support Linux and Mac OS X. Some packages have stubs for Windows (_for autocomplete_).

### Installation

````
go get pkg.re/essentialkaos/ek.v4
````

If you want update ek to latest stable release, do:

````
go get -u pkg.re/essentialkaos/ek.v4
````

### Packages

* [`arg`](https://godoc.org/pkg.re/essentialkaos/ek.v4/arg) - Package provides methods for working with command-line arguments
* [`color`](https://godoc.org/pkg.re/essentialkaos/ek.v4/color) - Package color provides methods for working with colors
* [`cron`](https://godoc.org/pkg.re/essentialkaos/ek.v4/cron) - Package provides methods for working with cron expressions
* [`csv`](https://godoc.org/pkg.re/essentialkaos/ek.v4/csv) - Package with simple (without any checks) csv parser compatible with default Go parser
* [`easing`](https://godoc.org/pkg.re/essentialkaos/ek.v4/easing) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`env`](https://godoc.org/pkg.re/essentialkaos/ek.v4/env) - Package provides methods for working with environment variables
* [`errutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/errutil) - Package provides methods for working with errors
* [`fmtc`](https://godoc.org/pkg.re/essentialkaos/ek.v4/fmtc) - Package provides methods similar to fmt for colored output
* [`fmtutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/fmtutil) - Package provides methods for output formating
* [`fsutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/fsutil) - Package provides methods for working with files on posix compatible systems (Linux / Mac OS X)
* [`hash`](https://godoc.org/pkg.re/essentialkaos/ek.v4/hash) - Package hash contains different hash algorithms and utilities
* [`httputil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/httputil) - Package provides methods for working with http request/responses
* [`jsonutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/jsonutil) - Package provides methods for working with json data
* [`knf`](https://godoc.org/pkg.re/essentialkaos/ek.v4/knf) - Package provides methods for working with configs in KNF format
* [`kv`](https://godoc.org/pkg.re/essentialkaos/ek.v4/kv) - Package provides simple key-value structs
* [`log`](https://godoc.org/pkg.re/essentialkaos/ek.v4/log) - Package with improved logger
* [`mathutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/mathutil) - Package with math utils
* [`netutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/netutil) - Package with network utils
* [`passwd`](https://godoc.org/pkg.re/essentialkaos/ek.v4/passwd) - Package passwd contains methods for working with passwords
* [`path`](https://godoc.org/pkg.re/essentialkaos/ek.v4/path) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://godoc.org/pkg.re/essentialkaos/ek.v4/pid) - Package for working with pid files
* [`pluralize`](https://godoc.org/pkg.re/essentialkaos/ek.v4/pluralize) - Package pluralize provides methods for pluralization
* [`rand`](https://godoc.org/pkg.re/essentialkaos/ek.v4/rand) - Package for generating random data
* [`req`](https://godoc.org/pkg.re/essentialkaos/ek.v4/req) - Package for working with http request
* [`signal`](https://godoc.org/pkg.re/essentialkaos/ek.v4/signal) - Package for handling signals
* [`sliceutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/sliceutil) - Package with utils for working with slices
* [`sortutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/sortutil) - Package with utils for sorting slices
* [`spellcheck`](https://godoc.org/pkg.re/essentialkaos/ek.v4/spellcheck) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`strutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/strutil) - Package provides utils for working with strings
* [`system/process`](https://godoc.org/pkg.re/essentialkaos/ek.v4/system/process) - Package provides methods for getting information about active processes
* [`system`](https://godoc.org/pkg.re/essentialkaos/ek.v4/system) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://godoc.org/pkg.re/essentialkaos/ek.v4/terminal) - Package provides methods for working with user input
* [`timeutil`](https://godoc.org/pkg.re/essentialkaos/ek.v4/timeutil) - Package with time utils
* [`tmp`](https://godoc.org/pkg.re/essentialkaos/ek.v4/tmp) - Package provides methods for working with temporary data
* [`usage`](https://godoc.org/pkg.re/essentialkaos/ek.v4/usage) - Package provides methods for rendering info for command-line tools
* [`uuid`](https://godoc.org/pkg.re/essentialkaos/ek.v4/uuid) - Package uuid contains methods for generating version 4 and 5 UUID's
* [`version`](https://godoc.org/pkg.re/essentialkaos/ek.v4/version) - Package provides methods for parsing semver version info

### Projects with EK

* [sslcli](https://github.com/essentialkaos/sslcli) - Pretty awesome command-line client for public SSLLabs API
* [redis-cli-monitor](https://github.com/essentialkaos/redis-cli-monitor) - Tiny redis client for renamed MONITOR commands
* [shdoc](https://github.com/essentialkaos/shdoc) - Tool for viewing and exporting docs for shell scripts
* [rbinstall](https://github.com/essentialkaos/rbinstall) - Utility for installing prebuilt ruby to RBEnv
* [mockka](https://github.com/essentialkaos/mockka) - Mockka is a simple utility for mocking HTTP API's
* [terrafarm](https://github.com/essentialkaos/terrafarm) - Utility for working with terraform based rpmbuilder farm
* [mdtoc](https://github.com/essentialkaos/mdtoc) - Utility for generating table of contents for markdown files

### Test & Coverage Status

| Branch | TravisCI | CodeCov |
|--------|----------|---------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=master)](https://travis-ci.org/essentialkaos/ek) | [![codecov.io](https://codecov.io/github/essentialkaos/ek/coverage.svg?branch=master)](https://codecov.io/github/essentialkaos/ek?branch=master) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=develop)](https://travis-ci.org/essentialkaos/ek) | [![codecov.io](https://codecov.io/github/essentialkaos/ek/coverage.svg?branch=develop)](https://codecov.io/github/essentialkaos/ek?branch=develop) |

### License

[EKOL](https://essentialkaos.com/ekol)
