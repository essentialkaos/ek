# `EK` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/ek.v3?status.svg)](https://godoc.org/pkg.re/essentialkaos/ek.v3) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/ek)](https://goreportcard.com/report/github.com/essentialkaos/ek) [![codebeat badge](https://codebeat.co/badges/3649d737-e5b9-4465-9765-b9f4ebec60ec)](https://codebeat.co/projects/github-com-essentialkaos-ek)

Auxiliary packages for Go.

### Platform support

Currently we support Linux and Mac OS X. Some packages have stubs for Windows (for autocomplete).

### Installation

````
go get pkg.re/essentialkaos/ek.v3
````

If you want update ek to latest stable release, do:

````
go get -u pkg.re/essentialkaos/ek.v3
````

### Packages

* [`arg`](https://godoc.org/pkg.re/essentialkaos/ek.v3/arg) - Package provides methods for working with command-line arguments
* [`color`](https://godoc.org/pkg.re/essentialkaos/ek.v3/color) - Package color provides methods for working with colors
* [`cron`](https://godoc.org/pkg.re/essentialkaos/ek.v3/cron) - Package provides methods for working with cron expressions
* [`crypto`](https://godoc.org/pkg.re/essentialkaos/ek.v3/crypto) - Package with utils for working with crypto data (passwords, uuids, file hashes)
* [`csv`](https://godoc.org/pkg.re/essentialkaos/ek.v3/csv) - Package with simple (without any checks) csv parser compatible with default Go parser
* [`easing`](https://godoc.org/pkg.re/essentialkaos/ek.v3/easing) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`env`](https://godoc.org/pkg.re/essentialkaos/ek.v3/env) - Package provides methods for working with environment variables
* [`errutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/errutil) - Package provides methods for working with errors
* [`fmtc`](https://godoc.org/pkg.re/essentialkaos/ek.v3/fmtc) - Package provides methods similar to fmt for colored output
* [`fmtutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/fmtutil) - Package provides methods for output formating
* [`fsutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/fsutil) - Package provides methods for working with files in posix compatible systems (Linux / Mac OS X)
* [`httputil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/httputil) - Package provides methods for working with http request/responses
* [`jsonutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/jsonutil) - Package provides methods for working with json data
* [`knf`](https://godoc.org/pkg.re/essentialkaos/ek.v3/knf) - Package provides methods for working with configs in KNF format
* [`kv`](https://godoc.org/pkg.re/essentialkaos/ek.v3/kv) - Package provides simple key-value structs
* [`log`](https://godoc.org/pkg.re/essentialkaos/ek.v3/log) - Package with improved logger
* [`mathutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/mathutil) - Package with math utils
* [`netutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/netutil) - Package with network utils
* [`path`](https://godoc.org/pkg.re/essentialkaos/ek.v3/path) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://godoc.org/pkg.re/essentialkaos/ek.v3/pid) - Package for working with pid files
* [`pluralize`](https://godoc.org/pkg.re/essentialkaos/ek.v3/pluralize) - Package pluralize provides methods for pluralization
* [`rand`](https://godoc.org/pkg.re/essentialkaos/ek.v3/rand) - Package for generating random data
* [`req`](https://godoc.org/pkg.re/essentialkaos/ek.v3/req) - Package for working with http request
* [`signal`](https://godoc.org/pkg.re/essentialkaos/ek.v3/signal) - Package for handling signals
* [`sliceutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/sliceutil) - Package with utils for working with slices
* [`sortutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/sortutil) - Package with utils for sorting slices
* [`spellcheck`](https://godoc.org/pkg.re/essentialkaos/ek.v3/spellcheck) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`strutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/strutil) - Package provides utils for working with strings
* [`system`](https://godoc.org/pkg.re/essentialkaos/ek.v3/system) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://godoc.org/pkg.re/essentialkaos/ek.v3/terminal) - Package provides methods for working with user input
* [`timeutil`](https://godoc.org/pkg.re/essentialkaos/ek.v3/timeutil) - Package with time utils
* [`tmp`](https://godoc.org/pkg.re/essentialkaos/ek.v3/tmp) - Package provides methods for working with temporary data
* [`usage`](https://godoc.org/pkg.re/essentialkaos/ek.v3/usage) - Package provides methods for rendering info for command-line tools
* [`version`](https://godoc.org/pkg.re/essentialkaos/ek.v3/version) - Package provides methods for parsing semver version info

### Projects with EK

* [ssllabs-client](https://github.com/essentialkaos/ssllabs_client) - Pretty awesome command-line client for public SSLLabs API
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
