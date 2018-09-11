<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-ek.svg"/></a></p>

<p align="center"><a href="#platform-support">Platform support</a> • <a href="#installation">Installation</a> • <a href="#sub-packages">Sub-packages</a> • <a href="#projects-with-ek">Projects with EK</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/ek.v9"><img src="https://godoc.org/pkg.re/essentialkaos/ek.v9?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/ek"><img src="https://goreportcard.com/badge/github.com/essentialkaos/ek"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-ek"><img alt="codebeat badge" src="https://codebeat.co/badges/3649d737-e5b9-4465-9765-b9f4ebec60ec" /></a>
  <a href="https://travis-ci.org/essentialkaos/ek"><img src="https://travis-ci.org/essentialkaos/ek.svg"></a>
  <a href='https://coveralls.io/github/essentialkaos/ek'><img src='https://coveralls.io/repos/github/essentialkaos/ek/badge.svg' alt='Coverage Status' /></a>
  <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Fessentialkaos%2Fek?ref=badge_shield"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Fessentialkaos%2Fek.svg?type=shield"/></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

Auxiliary packages for Go.

### Platform support

Currently we support Linux and Mac OS X (except `system` package). Some packages have stubs for Windows (_for autocomplete_).

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.8+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/ek.v9
```

If you want to update `EK` to latest stable release, do:

```
go get -u pkg.re/essentialkaos/ek.v9
```

### Sub-packages

* [`color`](https://godoc.org/pkg.re/essentialkaos/ek.v9/color) - Package provides methods for working with colors
* [`cron`](https://godoc.org/pkg.re/essentialkaos/ek.v9/cron) - Package provides methods for working with cron expressions
* [`csv`](https://godoc.org/pkg.re/essentialkaos/ek.v9/csv) - Package with simple (without any checks) CSV parser compatible with default Go parser
* [`easing`](https://godoc.org/pkg.re/essentialkaos/ek.v9/easing) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`env`](https://godoc.org/pkg.re/essentialkaos/ek.v9/env) - Package provides methods for working with environment variables
* [`errutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/errutil) - Package provides methods for working with errors
* [`directio`](https://godoc.org/pkg.re/essentialkaos/ek.v9/directio) - Package provides methods for reading/writing files with direct io
* [`fmtc`](https://godoc.org/pkg.re/essentialkaos/ek.v9/fmtc) - Package provides methods similar to fmt for colored output
* [`fmtutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/fmtutil) - Package provides methods for output formatting
* [`fmtutil/table`](https://godoc.org/pkg.re/essentialkaos/ek.v9/fmtutil/table) - Package table contains methods and structs for rendering data as a table
* [`fsutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/fsutil) - Package provides methods for working with files on POSIX compatible systems (Linux / Mac OS X)
* [`hash`](https://godoc.org/pkg.re/essentialkaos/ek.v9/hash) - Package hash contains different hash algorithms and utilities
* [`httputil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/httputil) - Package provides methods for working with HTTP request/responses
* [`initsystem`](https://godoc.org/pkg.re/essentialkaos/ek.v9/initsystem) - Package provides methods for working with different init systems (sysv, upstart, systemd)
* [`jsonutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/jsonutil) - Package provides methods for working with JSON data
* [`knf`](https://godoc.org/pkg.re/essentialkaos/ek.v9/knf) - Package provides methods for working with configs in KNF format
* [`kv`](https://godoc.org/pkg.re/essentialkaos/ek.v9/kv) - Package provides simple key-value structs
* [`log`](https://godoc.org/pkg.re/essentialkaos/ek.v9/log) - Package with an improved logger
* [`mathutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/mathutil) - Package with math utils
* [`netutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/netutil) - Package with network utils
* [`options`](https://godoc.org/pkg.re/essentialkaos/ek.v9/options) - Package provides methods for working with command-line options
* [`passwd`](https://godoc.org/pkg.re/essentialkaos/ek.v9/passwd) - Package contains methods for working with passwords
* [`path`](https://godoc.org/pkg.re/essentialkaos/ek.v9/path) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://godoc.org/pkg.re/essentialkaos/ek.v9/pid) - Package for working with PID files
* [`pluralize`](https://godoc.org/pkg.re/essentialkaos/ek.v9/pluralize) - Package provides methods for pluralization
* [`rand`](https://godoc.org/pkg.re/essentialkaos/ek.v9/rand) - Package for generating random data
* [`req`](https://godoc.org/pkg.re/essentialkaos/ek.v9/req) - Package for working with HTTP request
* [`signal`](https://godoc.org/pkg.re/essentialkaos/ek.v9/signal) - Package for handling signals
* [`sliceutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/sliceutil) - Package with utils for working with slices
* [`sortutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/sortutil) - Package with utils for sorting slices
* [`spellcheck`](https://godoc.org/pkg.re/essentialkaos/ek.v9/spellcheck) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`strutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/strutil) - Package provides utils for working with strings
* [`system/exec`](https://godoc.org/pkg.re/essentialkaos/ek.v9/system/process) - Package provides methods for executing commands
* [`system/process`](https://godoc.org/pkg.re/essentialkaos/ek.v9/system/process) - Package provides methods for getting information about active processes
* [`system/procname`](https://godoc.org/pkg.re/essentialkaos/ek.v9/system/process) - Package provides methods for changing process name in the process tree
* [`system`](https://godoc.org/pkg.re/essentialkaos/ek.v9/system) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://godoc.org/pkg.re/essentialkaos/ek.v9/terminal) - Package provides methods for working with user input
* [`terminal/window`](https://godoc.org/pkg.re/essentialkaos/ek.v9/terminal/window) - Package provides methods for working terminal window
* [`timeutil`](https://godoc.org/pkg.re/essentialkaos/ek.v9/timeutil) - Package provides methods for working with time
* [`tmp`](https://godoc.org/pkg.re/essentialkaos/ek.v9/tmp) - Package provides methods for working with temporary data
* [`usage`](https://godoc.org/pkg.re/essentialkaos/ek.v9/usage) - Package provides methods for rendering info for command-line tools
* [`usage/update`](https://godoc.org/pkg.re/essentialkaos/ek.v9/usage/update) - Package contains update checkers for different services
* [`uuid`](https://godoc.org/pkg.re/essentialkaos/ek.v9/uuid) - Package provides methods for generating version 4 and 5 UUID's
* [`version`](https://godoc.org/pkg.re/essentialkaos/ek.v9/version) - Package provides methods for parsing semver version info

### Projects with `EK`

* [Bastion](https://github.com/essentialkaos/bastion) - Utility for temporary disabling access to server
* [bibop](https://github.com/essentialkaos/bibop) - Utility for testing command-line tools
* [Deadline](https://github.com/essentialkaos/deadline) - Simple utility for controlling application working time
* [GoHeft](https://github.com/essentialkaos/goheft) - Utility for listing sizes of all used static libraries compiled into golang binary
* [GoMakeGen](https://github.com/essentialkaos/gomakegen) - Utility for generating makefiles for golang applications
* [kikky](https://github.com/gongled/kikky) - Small HTTP server and CLI utility which renders Kickstart configurations for Anaconda
* [MDToc](https://github.com/essentialkaos/mdtoc) - Utility for generating table of contents for markdown files
* [Mockka](https://github.com/essentialkaos/mockka) - Mockka is a simple utility for mocking HTTP API's
* [perfecto](https://github.com/essentialkaos/perfecto) - Tool for checking perfectly written RPM specs
* [pkg.re Morpher](https://github.com/essentialkaos/pkgre) - Part of [pkg.re](https://pkg.re) service (_provides versioned URLs for Go_)
* [RBInstall](https://github.com/essentialkaos/rbinstall) - Utility for installing prebuilt ruby to RBEnv
* [Redis Monitor Top](https://github.com/essentialkaos/redis-monitor-top) - Tiny Redis client for aggregating stats from MONITOR flow
* [Redis CLI Monitor](https://github.com/essentialkaos/redis-cli-monitor) - Tiny redis client for renamed MONITOR commands
* [Redis Latency Monitor](https://github.com/essentialkaos/redis-latency-monitor) - Tiny Redis client for latency measurement
* [SHDoc](https://github.com/essentialkaos/shdoc) - Tool for viewing and exporting docs for shell scripts
* [Sonar](https://github.com/essentialkaos/sonar) - Utility for showing user Slack status in Atlassian Jira
* [SourceIndex](https://github.com/essentialkaos/source-index) - Utility for generating an index for source archives
* [SSLScan Client](https://github.com/essentialkaos/sslcli) - Pretty awesome command-line client for public SSLLabs API
* [swptop](https://github.com/essentialkaos/swptop) - Simple utility for viewing swap consumption of processes
* [Terrafarm](https://github.com/essentialkaos/terrafarm) - Utility for working with terraform based rpmbuilder farm
* [vgrepo](https://github.com/gongled/vgrepo) - Simple CLI tool for managing Vagrant repositories
* [Yo](https://github.com/essentialkaos/yo) - Command-line YAML processor

### Build Status

| Branch | TravisCI |
|--------|----------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=master)](https://travis-ci.org/essentialkaos/ek) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=develop)](https://travis-ci.org/essentialkaos/ek) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
