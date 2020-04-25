<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-ek.svg"/></a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/ek.v11"><img src="https://godoc.org/pkg.re/essentialkaos/ek.v11?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/ek"><img src="https://goreportcard.com/badge/github.com/essentialkaos/ek"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-ek"><img alt="codebeat badge" src="https://codebeat.co/badges/3649d737-e5b9-4465-9765-b9f4ebec60ec" /></a>
  <a href="https://travis-ci.com/essentialkaos/ek"><img src="https://travis-ci.com/essentialkaos/ek.svg?branch=master"></a>
  <a href="https://coveralls.io/github/essentialkaos/ek"><img src="https://coveralls.io/repos/github/essentialkaos/ek/badge.svg" alt="Coverage Status" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

<p align="center"><a href="#platform-support">Platform support</a> • <a href="#installation">Installation</a> • <a href="#sub-packages">Sub-packages</a> • <a href="#projects-with-ek">Projects with EK</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Auxiliary packages for Go.

### Platform support

Currently we support Linux and Mac OS X (except `system` package). Some packages have stubs for Windows (_for autocomplete_).

### Installation

Before the initial install, allow git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.12+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/ek.v11
```

If you want to update `EK` to latest stable release, do:

```
go get -u pkg.re/essentialkaos/ek.v11
```

### Sub-packages

* [`cache`](https://godoc.org/pkg.re/essentialkaos/ek.v11/cache) - Package provides simple in-memory key:value store
* [`color`](https://godoc.org/pkg.re/essentialkaos/ek.v11/color) - Package provides methods for working with colors
* [`cron`](https://godoc.org/pkg.re/essentialkaos/ek.v11/cron) - Package provides methods for working with cron expressions
* [`csv`](https://godoc.org/pkg.re/essentialkaos/ek.v11/csv) - Package with simple (without any checks) CSV parser compatible with default Go parser
* [`easing`](https://godoc.org/pkg.re/essentialkaos/ek.v11/easing) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`emoji`](https://godoc.org/pkg.re/essentialkaos/ek.v11/emoji) - Package provides methods for working with emojis
* [`env`](https://godoc.org/pkg.re/essentialkaos/ek.v11/env) - Package provides methods for working with environment variables
* [`errutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/errutil) - Package provides methods for working with errors
* [`directio`](https://godoc.org/pkg.re/essentialkaos/ek.v11/directio) - Package provides methods for reading/writing files with direct io
* [`fmtc`](https://godoc.org/pkg.re/essentialkaos/ek.v11/fmtc) - Package provides methods similar to fmt for colored output
* [`fmtutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/fmtutil) - Package provides methods for output formatting
* [`fmtutil/table`](https://godoc.org/pkg.re/essentialkaos/ek.v11/fmtutil/table) - Package table contains methods and structs for rendering data as a table
* [`fsutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/fsutil) - Package provides methods for working with files on POSIX compatible systems (Linux / Mac OS X)
* [`hash`](https://godoc.org/pkg.re/essentialkaos/ek.v11/hash) - Package hash contains different hash algorithms and utilities
* [`httputil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/httputil) - Package provides methods for working with HTTP request/responses
* [`initsystem`](https://godoc.org/pkg.re/essentialkaos/ek.v11/initsystem) - Package provides methods for working with different init systems (sysv, upstart, systemd)
* [`jsonutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/jsonutil) - Package provides methods for working with JSON data
* [`knf`](https://godoc.org/pkg.re/essentialkaos/ek.v11/knf) - Package provides methods for working with configs in KNF format
* [`log`](https://godoc.org/pkg.re/essentialkaos/ek.v11/log) - Package with an improved logger
* [`mathutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/mathutil) - Package with math utils
* [`netutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/netutil) - Package with network utils
* [`options`](https://godoc.org/pkg.re/essentialkaos/ek.v11/options) - Package provides methods for working with command-line options
* [`passwd`](https://godoc.org/pkg.re/essentialkaos/ek.v11/passwd) - Package contains methods for working with passwords
* [`path`](https://godoc.org/pkg.re/essentialkaos/ek.v11/path) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://godoc.org/pkg.re/essentialkaos/ek.v11/pid) - Package for working with PID files
* [`pluralize`](https://godoc.org/pkg.re/essentialkaos/ek.v11/pluralize) - Package provides methods for pluralization
* [`rand`](https://godoc.org/pkg.re/essentialkaos/ek.v11/rand) - Package for generating random data
* [`req`](https://godoc.org/pkg.re/essentialkaos/ek.v11/req) - Package for working with HTTP request
* [`signal`](https://godoc.org/pkg.re/essentialkaos/ek.v11/signal) - Package for handling signals
* [`sliceutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/sliceutil) - Package with utils for working with slices
* [`sortutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/sortutil) - Package with utils for sorting slices
* [`spellcheck`](https://godoc.org/pkg.re/essentialkaos/ek.v11/spellcheck) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`strutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/strutil) - Package provides utils for working with strings
* [`system/exec`](https://godoc.org/pkg.re/essentialkaos/ek.v11/system/process) - Package provides methods for executing commands
* [`system/process`](https://godoc.org/pkg.re/essentialkaos/ek.v11/system/process) - Package provides methods for getting information about active processes
* [`system/procname`](https://godoc.org/pkg.re/essentialkaos/ek.v11/system/process) - Package provides methods for changing process name in the process tree
* [`system/sensors`](https://godoc.org/pkg.re/essentialkaos/ek.v11/system/sensors) - Package sensors provide methods for collecting sensors information
* [`system`](https://godoc.org/pkg.re/essentialkaos/ek.v11/system) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://godoc.org/pkg.re/essentialkaos/ek.v11/terminal) - Package provides methods for working with user input
* [`terminal/window`](https://godoc.org/pkg.re/essentialkaos/ek.v11/terminal/window) - Package provides methods for working terminal window
* [`timeutil`](https://godoc.org/pkg.re/essentialkaos/ek.v11/timeutil) - Package provides methods for working with time
* [`tmp`](https://godoc.org/pkg.re/essentialkaos/ek.v11/tmp) - Package provides methods for working with temporary data
* [`usage`](https://godoc.org/pkg.re/essentialkaos/ek.v11/usage) - Package provides methods for rendering info for command-line tools
* [`usage/update`](https://godoc.org/pkg.re/essentialkaos/ek.v11/usage/update) - Package contains update checkers for different services
* [`usage/completion/bash`](https://godoc.org/pkg.re/essentialkaos/ek.v11/usage/completion/bash) - Package bash provides methods for generating bash completion
* [`usage/completion/fish`](https://godoc.org/pkg.re/essentialkaos/ek.v11/usage/completion/fish) - Package fish provides methods for generating fish completion
* [`usage/completion/zsh`](https://godoc.org/pkg.re/essentialkaos/ek.v11/usage/completion/zsh) - Package zsh provides methods for generating zsh completion
* [`uuid`](https://godoc.org/pkg.re/essentialkaos/ek.v11/uuid) - Package provides methods for generating version 4 and 5 UUID's
* [`version`](https://godoc.org/pkg.re/essentialkaos/ek.v11/version) - Package provides methods for parsing semver version info

### Projects with `EK`

* [aligo](https://github.com/essentialkaos/aligo) - Utility for checking and viewing Golang struct alignment info
* [Bastion](https://github.com/essentialkaos/bastion) - Utility for temporary disabling access to server
* [bibop](https://github.com/essentialkaos/bibop) - Utility for testing command-line tools
* [Deadline](https://github.com/essentialkaos/deadline) - Simple utility for controlling application working time
* [fz](https://github.com/essentialkaos/fz) - Simple tool for formatting `go-fuzz` output
* [GoHeft](https://github.com/essentialkaos/goheft) - Utility for listing sizes of all used static libraries compiled into golang binary
* [GoMakeGen](https://github.com/essentialkaos/gomakegen) - Utility for generating makefiles for golang applications
* [icecli](https://github.com/essentialkaos/icecli) - Command-line tools for Icecast
* [IMC](https://github.com/essentialkaos/imc) - Simple terminal dashboard for Icecast
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
| `master` | [![Build Status](https://travis-ci.com/essentialkaos/ek.svg?branch=master)](https://travis-ci.com/essentialkaos/ek) |
| `develop` | [![Build Status](https://travis-ci.com/essentialkaos/ek.svg?branch=develop)](https://travis-ci.com/essentialkaos/ek) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
