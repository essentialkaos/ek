<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-ek.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/ek"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://kaos.sh/r/ek"><img src="https://kaos.sh/r/ek.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/ek"><img src="https://kaos.sh/b/3649d737-e5b9-4465-9765-b9f4ebec60ec.svg" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/ek/ci"><img src="https://kaos.sh/w/ek/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/ek/codeql"><img src="https://kaos.sh/w/ek/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg" /></a>
</p>

<p align="center"><a href="#platform-support">Platform support</a> • <a href="#installation">Installation</a> • <a href="#sub-packages">Sub-packages</a> • <a href="#projects-with-ek">Projects with EK</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Auxiliary packages for Go.

### Platform support

Currently we support Linux and macOS (_except some packages_). All packages have stubs for unsupported platforms (_for autocomplete_).

<details><summary><b>More info about stubs</b></summary><p>

> Some packages cannot be used on some platforms, like `fsutil` package, which cannot be used on Windows due to using syscalls, or `system` sub-packages which require [procfs](https://en.wikipedia.org/wiki/Procfs). But you can write code on these platforms with no problem because almost all packages have stubs with information about all constants, variables, and functions available on other platforms. So, for example, Sublime with [LSP](https://lsp.sublimetext.io) on Windows will show all information about methods available only on the Linux platform. All descriptions from stubs contain symbol ❗ at the beginning as a mark of unsupported code. Code with stubs can be compiled, but any method invocation from stubs will lead to panic.

</p></details>

### Installation

Make sure you have a working Go 1.17+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get github.com/essentialkaos/ek/v12
```

If you want to update `EK` to latest stable release, do:

```
go get -u github.com/essentialkaos/ek/v12
```

### Sub-packages

* [`ansi`](https://pkg.go.dev/github.com/essentialkaos/ek/ansi) - Package provides methods for working with ANSI/VT100 control sequences
* [`cache`](https://pkg.go.dev/github.com/essentialkaos/ek/cache) - Package provides a simple in-memory key:value cache
* [`color`](https://pkg.go.dev/github.com/essentialkaos/ek/color) - Package provides methods for working with colors
* [`cron`](https://pkg.go.dev/github.com/essentialkaos/ek/cron) - Package provides methods for working with cron expressions
* [`csv`](https://pkg.go.dev/github.com/essentialkaos/ek/csv) - Package with simple (without any checks) CSV parser compatible with default Go parser
* [`easing`](https://pkg.go.dev/github.com/essentialkaos/ek/easing) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`emoji`](https://pkg.go.dev/github.com/essentialkaos/ek/emoji) - Package provides methods for working with emojis
* [`env`](https://pkg.go.dev/github.com/essentialkaos/ek/env) - Package provides methods for working with environment variables
* [`errutil`](https://pkg.go.dev/github.com/essentialkaos/ek/errutil) - Package provides methods for working with errors
* [`events`](https://pkg.go.dev/github.com/essentialkaos/ek/events) - Package provides methods and structs for creating event-driven systems
* [`directio`](https://pkg.go.dev/github.com/essentialkaos/ek/directio) - Package provides methods for reading/writing files with direct io
* [`fmtc`](https://pkg.go.dev/github.com/essentialkaos/ek/fmtc) - Package provides methods similar to fmt for colored output
* [`fmtc/lscolors`](https://pkg.go.dev/github.com/essentialkaos/ek/fmtc/lscolors) - Package provides methods for colorizing file names based on colors from dircolors
* [`fmtutil`](https://pkg.go.dev/github.com/essentialkaos/ek/fmtutil) - Package provides methods for output formatting
* [`fmtutil/table`](https://pkg.go.dev/github.com/essentialkaos/ek/fmtutil/table) - Package contains methods and structs for rendering data in tabular format
* [`fsutil`](https://pkg.go.dev/github.com/essentialkaos/ek/fsutil) - Package provides methods for working with files on POSIX compatible systems (BSD/Linux/macOS)
* [`hash`](https://pkg.go.dev/github.com/essentialkaos/ek/hash) - Package hash contains different hash algorithms and utilities
* [`httputil`](https://pkg.go.dev/github.com/essentialkaos/ek/httputil) - Package provides methods for working with HTTP request/responses
* [`initsystem`](https://pkg.go.dev/github.com/essentialkaos/ek/initsystem) - Package provides methods for working with different init systems (sysv, upstart, systemd)
* [`jsonutil`](https://pkg.go.dev/github.com/essentialkaos/ek/jsonutil) - Package provides methods for working with JSON data
* [`knf`](https://pkg.go.dev/github.com/essentialkaos/ek/knf) - Package provides methods for working with configuration files in [KNF format](https://kaos.sh/knf-spec)
* [`log`](https://pkg.go.dev/github.com/essentialkaos/ek/log) - Package with an improved logger
* [`mathutil`](https://pkg.go.dev/github.com/essentialkaos/ek/mathutil) - Package provides some additional math methods
* [`netutil`](https://pkg.go.dev/github.com/essentialkaos/ek/netutil) - Package provides methods for working with network
* [`options`](https://pkg.go.dev/github.com/essentialkaos/ek/options) - Package provides methods for working with command-line options
* [`passwd`](https://pkg.go.dev/github.com/essentialkaos/ek/passwd) - Package contains methods for working with passwords
* [`path`](https://pkg.go.dev/github.com/essentialkaos/ek/path) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://pkg.go.dev/github.com/essentialkaos/ek/pid) - Package for working with PID files
* [`pluralize`](https://pkg.go.dev/github.com/essentialkaos/ek/pluralize) - Package provides methods for pluralization
* [`progress`](https://pkg.go.dev/github.com/essentialkaos/ek/progress) - Package provides methods and structs for creating terminal progress bar
* [`rand`](https://pkg.go.dev/github.com/essentialkaos/ek/rand) - Package for generating random data
* [`req`](https://pkg.go.dev/github.com/essentialkaos/ek/req) - Package simplify working with an HTTP requests
* [`secstr`](https://pkg.go.dev/github.com/essentialkaos/ek/secstr) - Package provides methods and structs for working with protected (secure) strings
* [`signal`](https://pkg.go.dev/github.com/essentialkaos/ek/signal) - Package provides methods for handling POSIX signals
* [`sliceutil`](https://pkg.go.dev/github.com/essentialkaos/ek/sliceutil) - Package provides methods for working with slices
* [`sortutil`](https://pkg.go.dev/github.com/essentialkaos/ek/sortutil) - Package provides methods for sorting slices
* [`spellcheck`](https://pkg.go.dev/github.com/essentialkaos/ek/spellcheck) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`spinner`](https://pkg.go.dev/github.com/essentialkaos/ek/spinner) - Package provides methods for creating spinner animation for long-running tasks
* [`strutil`](https://pkg.go.dev/github.com/essentialkaos/ek/strutil) - Package provides methods for working with strings
* [`system/exec`](https://pkg.go.dev/github.com/essentialkaos/ek/system/exec) - Package provides methods for executing commands
* [`system/process`](https://pkg.go.dev/github.com/essentialkaos/ek/system/process) - Package provides methods for gathering information about active processes
* [`system/procname`](https://pkg.go.dev/github.com/essentialkaos/ek/system/procname) - Package provides methods for changing process name in the process tree
* [`system/sensors`](https://pkg.go.dev/github.com/essentialkaos/ek/system/sensors) - Package provide methods for collecting sensors information
* [`system`](https://pkg.go.dev/github.com/essentialkaos/ek/system) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://pkg.go.dev/github.com/essentialkaos/ek/terminal) - Package provides methods for working with user input
* [`terminal/window`](https://pkg.go.dev/github.com/essentialkaos/ek/terminal/window) - Package provides methods for working terminal window
* [`timeutil`](https://pkg.go.dev/github.com/essentialkaos/ek/timeutil) - Package provides methods for working with time and date
* [`tmp`](https://pkg.go.dev/github.com/essentialkaos/ek/tmp) - Package provides methods for working with temporary data
* [`usage`](https://pkg.go.dev/github.com/essentialkaos/ek/usage) - Package usage provides methods and structs for generating usage info for command-line tools
* [`usage/update`](https://pkg.go.dev/github.com/essentialkaos/ek/usage/update) - Package contains update checkers for different services
* [`usage/completion/bash`](https://pkg.go.dev/github.com/essentialkaos/ek/usage/completion/bash) - Package provides methods for generating bash completion
* [`usage/completion/fish`](https://pkg.go.dev/github.com/essentialkaos/ek/usage/completion/fish) - Package provides methods for generating fish completion
* [`usage/completion/zsh`](https://pkg.go.dev/github.com/essentialkaos/ek/usage/completion/zsh) - Package provides methods for generating zsh completion
* [`uuid`](https://pkg.go.dev/github.com/essentialkaos/ek/uuid) - Package provides methods for generating version 4 and 5 UUID's
* [`version`](https://pkg.go.dev/github.com/essentialkaos/ek/version) - Package version provides methods for working with semver version info

### Projects with `EK`

* [aligo](https://kaos.sh/aligo) - Utility for checking and viewing Golang struct alignment info
* [Bastion](https://kaos.sh/bastion) - Utility for temporary disabling access to server
* [bibop](https://kaos.sh/bibop) - Utility for testing command-line tools
* [bop](https://kaos.sh/bop) - Utility for generating bibop tests for RPM packages
* [Deadline](https://kaos.sh/deadline) - Simple utility for controlling application working time
* [fz](https://kaos.sh/fz) - Simple tool for formatting `go-fuzz` output
* [GoHeft](https://kaos.sh/goheft) - Utility for listing sizes of all used static libraries compiled into golang binary
* [GoMakeGen](https://kaos.sh/gomakegen) - Utility for generating makefiles for golang applications
* [icecli](https://kaos.sh/icecli) - Command-line tools for Icecast
* [IMC](https://kaos.sh/imc) - Simple terminal dashboard for Icecast
* [init-exporter](https://github.com/funbox/init-exporter) - Utility for exporting services described by Procfile to init system
* [jira-reindex-runner](https://kaos.sh/jira-reindex-runner) - Application for periodical running Jira re-index process
* [knf](https://kaos.sh/knf) - Simple utility for reading values from KNF files
* [MDToc](https://kaos.sh/mdtoc) - Utility for generating table of contents for markdown files
* [Mockka](https://kaos.sh/mockka) - Mockka is a simple utility for mocking HTTP API's
* [perfecto](https://kaos.sh/perfecto) - Tool for checking perfectly written RPM specs
* [pkg.re Morpher](https://kaos.sh/pkgre) - Part of pkg.re service (_provides versioned URLs for Go_)
* [RBInstall](https://kaos.sh/rbinstall) - Utility for installing prebuilt ruby to RBEnv
* [Redis CLI Monitor](https://kaos.sh/redis-cli-monitor) - Tiny redis client for renamed MONITOR commands
* [Redis Latency Monitor](https://kaos.sh/redis-latency-monitor) - Tiny Redis client for latency measurement
* [Redis Monitor Top](https://kaos.sh/redis-monitor-top) - Tiny Redis client for aggregating stats from MONITOR flow
* [rsz](https://kaos.sh/rsz) - Simple utility for image resizing
* [scratch](https://kaos.sh/scratch) - Simple utility for generating blank files for Go apps, utilities and packages
* [SHDoc](https://kaos.sh/shdoc) - Tool for viewing and exporting docs for shell scripts
* [Sonar](https://kaos.sh/sonar) - Utility for showing user Slack status in Atlassian Jira
* [SourceIndex](https://kaos.sh/source-index) - Utility for generating an index for source archives
* [SSLScan Client](https://kaos.sh/sslcli) - Pretty awesome command-line client for public SSLLabs API
* [swptop](https://kaos.sh/swptop) - Simple utility for viewing swap consumption of processes
* [uc](https://kaos.sh/uc) - Simple utility for counting unique lines
* [updown-badge-server](https://kaos.sh/updown-badge-server) - Service for generating badges for updown.io checks
* [Yo](https://kaos.sh/yo) - Command-line YAML processor

### Build Status

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/ek/ci.svg?branch=master)](https://kaos.sh/w/ek/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/ek/ci.svg?branch=develop)](https://kaos.sh/w/ek/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
