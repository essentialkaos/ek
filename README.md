<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-ek.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/ek.v12?docs"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://kaos.sh/r/ek"><img src="https://kaos.sh/r/ek.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/ek"><img src="https://codebeat.co/badges/3649d737-e5b9-4465-9765-b9f4ebec60ec" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/ek/ci"><img src="https://kaos.sh/w/ek/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/ek/codeql"><img src="https://kaos.sh/w/ek/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg" /></a>
</p>

<p align="center"><a href="#platform-support">Platform support</a> • <a href="#installation">Installation</a> • <a href="#sub-packages">Sub-packages</a> • <a href="#projects-with-ek">Projects with EK</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Auxiliary packages for Go.

### Platform support

Currently we support Linux and Mac OS X (except `system` package). Some packages have stubs for Windows (_for autocomplete_).

### Installation

Make sure you have a working Go 1.16+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/ek.v12
```

If you want to update `EK` to latest stable release, do:

```
go get -u pkg.re/essentialkaos/ek.v12
```

### Sub-packages

* [`cache`](https://pkg.re/essentialkaos/ek.v12/cache?docs) - Package provides simple in-memory key:value store
* [`color`](https://pkg.re/essentialkaos/ek.v12/color?docs) - Package provides methods for working with colors
* [`cron`](https://pkg.re/essentialkaos/ek.v12/cron?docs) - Package provides methods for working with cron expressions
* [`csv`](https://pkg.re/essentialkaos/ek.v12/csv?docs) - Package with simple (without any checks) CSV parser compatible with default Go parser
* [`easing`](https://pkg.re/essentialkaos/ek.v12/easing?docs) - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`emoji`](https://pkg.re/essentialkaos/ek.v12/emoji?docs) - Package provides methods for working with emojis
* [`env`](https://pkg.re/essentialkaos/ek.v12/env?docs) - Package provides methods for working with environment variables
* [`errutil`](https://pkg.re/essentialkaos/ek.v12/errutil?docs) - Package provides methods for working with errors
* [`directio`](https://pkg.re/essentialkaos/ek.v12/directio?docs) - Package provides methods for reading/writing files with direct io
* [`fmtc`](https://pkg.re/essentialkaos/ek.v12/fmtc?docs) - Package provides methods similar to fmt for colored output
* [`fmtc/lscolors`](https://pkg.re/essentialkaos/ek.v12/fmtc/lscolors?docs) - Package provides method for colorizing file names based on colors from dircolors
* [`fmtutil`](https://pkg.re/essentialkaos/ek.v12/fmtutil?docs) - Package provides methods for output formatting
* [`fmtutil/table`](https://pkg.re/essentialkaos/ek.v12/fmtutil/table?docs) - Package table contains methods and structs for rendering data as a table
* [`fsutil`](https://pkg.re/essentialkaos/ek.v12/fsutil?docs) - Package provides methods for working with files on POSIX compatible systems (Linux / Mac OS X)
* [`hash`](https://pkg.re/essentialkaos/ek.v12/hash?docs) - Package hash contains different hash algorithms and utilities
* [`httputil`](https://pkg.re/essentialkaos/ek.v12/httputil?docs) - Package provides methods for working with HTTP request/responses
* [`initsystem`](https://pkg.re/essentialkaos/ek.v12/initsystem?docs) - Package provides methods for working with different init systems (sysv, upstart, systemd)
* [`jsonutil`](https://pkg.re/essentialkaos/ek.v12/jsonutil?docs) - Package provides methods for working with JSON data
* [`knf`](https://pkg.re/essentialkaos/ek.v12/knf?docs) - Package provides methods for working with configs in KNF format
* [`log`](https://pkg.re/essentialkaos/ek.v12/log?docs) - Package with an improved logger
* [`mathutil`](https://pkg.re/essentialkaos/ek.v12/mathutil?docs) - Package with math utils
* [`netutil`](https://pkg.re/essentialkaos/ek.v12/netutil?docs) - Package with network utils
* [`options`](https://pkg.re/essentialkaos/ek.v12/options?docs) - Package provides methods for working with command-line options
* [`passwd`](https://pkg.re/essentialkaos/ek.v12/passwd?docs) - Package contains methods for working with passwords
* [`path`](https://pkg.re/essentialkaos/ek.v12/path?docs) - Package for working with paths (fully compatible with base path package)
* [`pid`](https://pkg.re/essentialkaos/ek.v12/pid?docs) - Package for working with PID files
* [`pluralize`](https://pkg.re/essentialkaos/ek.v12/pluralize?docs) - Package provides methods for pluralization
* [`progress`](https://pkg.re/essentialkaos/ek.v12/progress?docs) - Package provides methods and structs for creating terminal progress bar
* [`rand`](https://pkg.re/essentialkaos/ek.v12/rand?docs) - Package for generating random data
* [`req`](https://pkg.re/essentialkaos/ek.v12/req?docs) - Package for working with HTTP request
* [`signal`](https://pkg.re/essentialkaos/ek.v12/signal?docs) - Package for handling signals
* [`sliceutil`](https://pkg.re/essentialkaos/ek.v12/sliceutil?docs) - Package with utils for working with slices
* [`sortutil`](https://pkg.re/essentialkaos/ek.v12/sortutil?docs) - Package with utils for sorting slices
* [`spellcheck`](https://pkg.re/essentialkaos/ek.v12/spellcheck?docs) - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`spinner`](https://pkg.re/essentialkaos/ek.v12/spinner?docs) - Package provides methods for creating spinner animation for long-running tasks
* [`strutil`](https://pkg.re/essentialkaos/ek.v12/strutil?docs) - Package provides utils for working with strings
* [`system/exec`](https://pkg.re/essentialkaos/ek.v12/system/process?docs) - Package provides methods for executing commands
* [`system/process`](https://pkg.re/essentialkaos/ek.v12/system/process?docs) - Package provides methods for getting information about active processes
* [`system/procname`](https://pkg.re/essentialkaos/ek.v12/system/process?docs) - Package provides methods for changing process name in the process tree
* [`system/sensors`](https://pkg.re/essentialkaos/ek.v12/system/sensors?docs) - Package sensors provide methods for collecting sensors information
* [`system`](https://pkg.re/essentialkaos/ek.v12/system?docs) - Package provides methods for working with system data (metrics/users)
* [`terminal`](https://pkg.re/essentialkaos/ek.v12/terminal?docs) - Package provides methods for working with user input
* [`terminal/window`](https://pkg.re/essentialkaos/ek.v12/terminal/window?docs) - Package provides methods for working terminal window
* [`timeutil`](https://pkg.re/essentialkaos/ek.v12/timeutil?docs) - Package provides methods for working with time and date
* [`tmp`](https://pkg.re/essentialkaos/ek.v12/tmp?docs) - Package provides methods for working with temporary data
* [`usage`](https://pkg.re/essentialkaos/ek.v12/usage?docs) - Package provides methods for rendering info for command-line tools
* [`usage/update`](https://pkg.re/essentialkaos/ek.v12/usage/update?docs) - Package contains update checkers for different services
* [`usage/completion/bash`](https://pkg.re/essentialkaos/ek.v12/usage/completion/bash?docs) - Package bash provides methods for generating bash completion
* [`usage/completion/fish`](https://pkg.re/essentialkaos/ek.v12/usage/completion/fish?docs) - Package fish provides methods for generating fish completion
* [`usage/completion/zsh`](https://pkg.re/essentialkaos/ek.v12/usage/completion/zsh?docs) - Package zsh provides methods for generating zsh completion
* [`uuid`](https://pkg.re/essentialkaos/ek.v12/uuid?docs) - Package provides methods for generating version 4 and 5 UUID's
* [`version`](https://pkg.re/essentialkaos/ek.v12/version?docs) - Package provides methods for parsing semver version info

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
* [jira-reindex-runner](https://kaos.sh/jira-reindex-runner) - Application for periodical running Jira re-index process
* [knf](https://kaos.sh/knf) - Simple utility for reading values from KNF files
* [MDToc](https://kaos.sh/mdtoc) - Utility for generating table of contents for markdown files
* [Mockka](https://kaos.sh/mockka) - Mockka is a simple utility for mocking HTTP API's
* [perfecto](https://kaos.sh/perfecto) - Tool for checking perfectly written RPM specs
* [pkg.re Morpher](https://kaos.sh/pkgre) - Part of [pkg.re](https://pkg.re) service (_provides versioned URLs for Go_)
* [RBInstall](https://kaos.sh/rbinstall) - Utility for installing prebuilt ruby to RBEnv
* [Redis CLI Monitor](https://kaos.sh/redis-cli-monitor) - Tiny redis client for renamed MONITOR commands
* [Redis Latency Monitor](https://kaos.sh/redis-latency-monitor) - Tiny Redis client for latency measurement
* [Redis Monitor Top](https://kaos.sh/redis-monitor-top) - Tiny Redis client for aggregating stats from MONITOR flow
* [scratch](https://kaos.sh/scratch) - Simple utility for generating blank files for Go apps, utilities and packages
* [SHDoc](https://kaos.sh/shdoc) - Tool for viewing and exporting docs for shell scripts
* [Sonar](https://kaos.sh/sonar) - Utility for showing user Slack status in Atlassian Jira
* [SourceIndex](https://kaos.sh/source-index) - Utility for generating an index for source archives
* [SSLScan Client](https://kaos.sh/sslcli) - Pretty awesome command-line client for public SSLLabs API
* [swptop](https://kaos.sh/swptop) - Simple utility for viewing swap consumption of processes
* [uc](https://kaos.sh/uc) - Simple utility for counting unique lines
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
