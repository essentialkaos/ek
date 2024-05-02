<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-ek.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/ek.v12"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://kaos.sh/r/ek.v12"><img src="https://kaos.sh/r/ek.v12.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/l/ek"><img src="https://kaos.sh/l/210cafc2de7bf4320649.svg" alt="Code Climate Maintainability" /></a>
  <a href="https://kaos.sh/b/ek"><img src="https://kaos.sh/b/3649d737-e5b9-4465-9765-b9f4ebec60ec.svg" alt="Codebeat badge" /></a>
  <br/>
  <a href="https://kaos.sh/w/ek/ci"><img src="https://kaos.sh/w/ek/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/ek/codeql-analysis"><img src="https://kaos.sh/w/ek/codeql-analysis.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg" /></a>
</p>

<p align="center"><a href="#platform-support">Platform support</a> • <a href="#installation">Installation</a> • <a href="#sub-packages">Sub-packages</a> • <a href="#projects-with-ek">Projects with EK</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

Auxiliary packages for Go.

### Platform support

Currently we support Linux and macOS (_except some packages_). All packages have stubs for unsupported platforms (_for autocomplete_).

> [!NOTE]
> Some packages cannot be used on some platforms, such as the `fsutil` package, which cannot be used on Windows because it uses syscalls, or `system` subpackages that require [procfs](https://en.wikipedia.org/wiki/Procfs). But you can still write code on these platforms, because almost all packages have stubs with information about all the constants, variables, and functions available on other platforms. For example, Sublime with [LSP](https://lsp.sublimetext.io) on Windows will show all the information about methods that are only available on the Linux platform. All stub descriptions contain the ❗ symbol at the beginning to indicate unsupported code. Code with stubs can be compiled, but any method call from stubs will cause panic.

### Installation

Make sure you have a working Go 1.20+ workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go get github.com/essentialkaos/ek/v12
```

If you want to update `ek` to latest stable release, do:

```
go get -u github.com/essentialkaos/ek/v12
```

> [!TIP]
> If you are using SublimeText 4 (`4075+`), we strongly recommend that you install [extended Go syntax highlighting](https://github.com/essentialkaos/blackhole-theme-sublime/blob/master/fmtc.sublime-syntax) with support for `fmtc` [tags](fmtc).

### Sub-packages

* [`ansi`](https://kaos.sh/g/ek.v12/ansi) — Package provides methods for working with ANSI/VT100 control sequences
* [`cache`](https://kaos.sh/g/ek.v12/cache) — Package provides a simple in-memory key:value cache
* [`color`](https://kaos.sh/g/ek.v12/color) — Package provides methods for working with colors
* [`cron`](https://kaos.sh/g/ek.v12/cron) — Package provides methods for working with cron expressions
* [`csv`](https://kaos.sh/g/ek.v12/csv) — Package with simple (without any checks) CSV parser compatible with default Go parser
* [`easing`](https://kaos.sh/g/ek.v12/easing) — Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* [`emoji`](https://kaos.sh/g/ek.v12/emoji) — Package provides methods for working with emojis
* [`env`](https://kaos.sh/g/ek.v12/env) — Package provides methods for working with environment variables
* [`errutil`](https://kaos.sh/g/ek.v12/errutil) — Package provides methods for working with errors
* [`events`](https://kaos.sh/g/ek.v12/events) — Package provides methods and structs for creating event-driven systems
* [`directio`](https://kaos.sh/g/ek.v12/directio) — Package provides methods for reading/writing files with direct io
* [`fmtc`](https://kaos.sh/g/ek.v12/fmtc) — Package provides methods similar to fmt for colored output
* [`fmtutil`](https://kaos.sh/g/ek.v12/fmtutil) — Package provides methods for output formatting
* [`fmtutil/barcode`](https://kaos.sh/g/ek.v12/fmtutil/barcode) — Package provides methods to generate colored representation of unique data
* [`fmtutil/panel`](https://kaos.sh/g/ek.v12/fmtutil/panel) — Package contains methods for rendering panels with text
* [`fmtutil/table`](https://kaos.sh/g/ek.v12/fmtutil/table) — Package contains methods and structs for rendering data in tabular format
* [`fsutil`](https://kaos.sh/g/ek.v12/fsutil) — Package provides methods for working with files on POSIX compatible systems (BSD/Linux/macOS)
* [`hash`](https://kaos.sh/g/ek.v12/hash) — Package hash contains different hash algorithms and utilities
* [`httputil`](https://kaos.sh/g/ek.v12/httputil) — Package provides methods for working with HTTP request/responses
* [`initsystem`](https://kaos.sh/g/ek.v12/initsystem) — Package provides methods for working with different init systems (sysv, upstart, systemd)
* [`initsystem`](https://kaos.sh/g/ek.v12/initsystem/sdnotify) — Package provides methods methods for sending [notifications to systemd](https://www.freedesktop.org/software/systemd/man/latest/sd_notify.html#Well-known%20assignments)
* [`jsonutil`](https://kaos.sh/g/ek.v12/jsonutil) — Package provides methods for working with JSON data
* [`knf`](https://kaos.sh/g/ek.v12/knf) — Package provides methods for working with configuration files in [KNF format](https://kaos.sh/knf-spec)
* [`knf/united`](https://kaos.sh/g/ek.v12/knf/united) — Package provides united configuration (_knf + options + environment variables_)
* [`log`](https://kaos.sh/g/ek.v12/log) — Package with an improved logger
* [`lock`](https://kaos.sh/g/ek.v12/lock) — Package provides methods for working with lock files
* [`lscolors`](https://kaos.sh/g/ek.v12/lscolors) — Package provides methods for colorizing file names based on colors from dircolors
* [`mathutil`](https://kaos.sh/g/ek.v12/mathutil) — Package provides some additional math methods
* [`netutil`](https://kaos.sh/g/ek.v12/netutil) — Package provides methods for working with network
* [`options`](https://kaos.sh/g/ek.v12/options) — Package provides methods for working with command-line options
* [`pager`](https://kaos.sh/g/ek.v12/pager) — Package provides methods for pager setup (more/less)
* [`passthru`](https://kaos.sh/g/ek.v12/passthru) — Package provides Reader and Writer with information about the amount of data being passed
* [`passwd`](https://kaos.sh/g/ek.v12/passwd) — Package contains methods for working with passwords
* [`path`](https://kaos.sh/g/ek.v12/path) — Package for working with paths (fully compatible with base path package)
* [`pid`](https://kaos.sh/g/ek.v12/pid) — Package for working with PID files
* [`pluralize`](https://kaos.sh/g/ek.v12/pluralize) — Package provides methods for pluralization
* [`progress`](https://kaos.sh/g/ek.v12/progress) — Package provides methods and structs for creating terminal progress bar
* [`protip`](https://kaos.sh/g/ek.v12/protip) — Package for displaying usage tips
* [`rand`](https://kaos.sh/g/ek.v12/rand) — Package for generating random data
* [`req`](https://kaos.sh/g/ek.v12/req) — Package simplify working with an HTTP requests
* [`secstr`](https://kaos.sh/g/ek.v12/secstr) — Package provides methods and structs for working with protected (secure) strings
* [`signal`](https://kaos.sh/g/ek.v12/signal) — Package provides methods for handling POSIX signals
* [`sliceutil`](https://kaos.sh/g/ek.v12/sliceutil) — Package provides methods for working with slices
* [`sortutil`](https://kaos.sh/g/ek.v12/sortutil) — Package provides methods for sorting slices
* [`spellcheck`](https://kaos.sh/g/ek.v12/spellcheck) — Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* [`spinner`](https://kaos.sh/g/ek.v12/spinner) — Package provides methods for creating spinner animation for long-running tasks
* [`strutil`](https://kaos.sh/g/ek.v12/strutil) — Package provides methods for working with strings
* [`support`](https://kaos.sh/g/ek.v12/support) — Package provides methods for collecting and printing support information about system
* [`system/container`](https://kaos.sh/g/ek.v12/system/container) — Package provides methods for checking container engine info
* [`system/exec`](https://kaos.sh/g/ek.v12/system/exec) — Package provides methods for executing commands
* [`system/process`](https://kaos.sh/g/ek.v12/system/process) — Package provides methods for gathering information about active processes
* [`system/procname`](https://kaos.sh/g/ek.v12/system/procname) — Package provides methods for changing process name in the process tree
* [`system/sensors`](https://kaos.sh/g/ek.v12/system/sensors) — Package provide methods for collecting sensors information
* [`system`](https://kaos.sh/g/ek.v12/system) — Package provides methods for working with system data (metrics/users)
* [`terminal`](https://kaos.sh/g/ek.v12/terminal) — Package provides methods for printing messages to terminal
* [`terminal/input`](https://kaos.sh/g/ek.v12/terminal/input) — Package provides methods for working with user input
* [`terminal/tty`](https://kaos.sh/g/ek.v12/terminal/tty) — Package provides methods for working with TTY
* [`timeutil`](https://kaos.sh/g/ek.v12/timeutil) — Package provides methods for working with time and date
* [`tmp`](https://kaos.sh/g/ek.v12/tmp) — Package provides methods for working with temporary data
* [`usage`](https://kaos.sh/g/ek.v12/usage) — Package usage provides methods and structs for generating usage info for command-line tools
* [`usage/update`](https://kaos.sh/g/ek.v12/usage/update) — Package contains update checkers for different services
* [`usage/completion/bash`](https://kaos.sh/g/ek.v12/usage/completion/bash) — Package provides methods for generating bash completion
* [`usage/completion/fish`](https://kaos.sh/g/ek.v12/usage/completion/fish) — Package provides methods for generating fish completion
* [`usage/completion/zsh`](https://kaos.sh/g/ek.v12/usage/completion/zsh) — Package provides methods for generating zsh completion
* [`uuid`](https://kaos.sh/g/ek.v12/uuid) — Package provides methods for generating version 4 and 5 UUID's
* [`version`](https://kaos.sh/g/ek.v12/version) — Package version provides methods for working with semver version info

### Projects with `EK`

* [aligo](https://kaos.sh/aligo) — Utility for checking and viewing Golang struct alignment info
* [artefactor](https://kaos.sh/artefactor) — Utility for downloading artefacts from GitHub
* [atlassian-cloud-backuper](https://kaos.sh/atlassian-cloud-backuper) — Tool for backuping Atlassian cloud services (_Jira and Confluence_)
* [Bastion](https://kaos.sh/bastion) — Utility for temporary disabling access to server
* [bibop](https://kaos.sh/bibop) — Utility for testing command-line tools
* [bop](https://kaos.sh/bop) — Utility for generating bibop tests for RPM packages
* [Deadline](https://kaos.sh/deadline) — Simple utility for controlling application working time
* [fz](https://kaos.sh/fz) — Simple tool for formatting `go-fuzz` output
* [GoHeft](https://kaos.sh/goheft) — Utility for listing sizes of all used static libraries compiled into golang binary
* [GoMakeGen](https://kaos.sh/gomakegen) — Utility for generating makefiles for golang applications
* [icecli](https://kaos.sh/icecli) — Command-line tools for Icecast
* [IMC](https://kaos.sh/imc) — Simple terminal dashboard for Icecast
* [init-exporter](https://github.com/funbox/init-exporter) — Utility for exporting services described by Procfile to init system
* [jira-reindex-runner](https://kaos.sh/jira-reindex-runner) — Application for periodical running Jira re-index process
* [knf](https://kaos.sh/knf) — Simple utility for reading values from KNF files
* [MDToc](https://kaos.sh/mdtoc) — Utility for generating table of contents for markdown files
* [path](https://kaos.sh/path) — Dead simple tool for working with paths
* [perfecto](https://kaos.sh/perfecto) — Tool for checking perfectly written RPM specs
* [RBInstall](https://kaos.sh/rbinstall) — Utility for installing prebuilt ruby to RBEnv
* [RDS](https://kaos.sh/rds) — Tool for Redis orchestration
* [Redis CLI Monitor](https://kaos.sh/redis-cli-monitor) — Tiny redis client for renamed MONITOR commands
* [Redis Latency Monitor](https://kaos.sh/redis-latency-monitor) — Tiny Redis client for latency measurement
* [Redis Monitor Top](https://kaos.sh/redis-monitor-top) — Tiny Redis client for aggregating stats from MONITOR flow
* [rep](https://kaos.sh/rep) — YUM repository management utility
* [rsz](https://kaos.sh/rsz) — Simple utility for image resizing
* [scratch](https://kaos.sh/scratch) — Simple utility for generating blank files for Go apps, utilities and packages
* [SHDoc](https://kaos.sh/shdoc) — Tool for viewing and exporting docs for shell scripts
* [Sonar](https://kaos.sh/sonar) — Utility for showing user Slack status in Atlassian Jira
* [SourceIndex](https://kaos.sh/source-index) — Utility for generating an index for source archives
* [SSLScan Client](https://kaos.sh/sslcli) — Pretty awesome command-line client for public SSLLabs API
* [subdy](https://kaos.sh/subdy) — CLI for subdomain.center API
* [swptop](https://kaos.sh/swptop) — Simple utility for viewing swap consumption of processes
* [uc](https://kaos.sh/uc) — Simple utility for counting unique lines
* [updown-badge-server](https://kaos.sh/updown-badge-server) — Service for generating badges for updown.io checks
* [yo](https://kaos.sh/yo) — Command-line YAML processor

### CI Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/ek/ci.svg?branch=master)](https://kaos.sh/w/ek/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/ek/ci.svg?branch=develop)](https://kaos.sh/w/ek/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
