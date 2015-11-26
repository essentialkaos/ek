### EK [![GoDoc](https://godoc.org/github.com/essentialkaos/ek?status.svg)](https://godoc.org/github.com/essentialkaos/ek)

Auxiliary packages for Go.

### Platform support

Currently we support Linux and Mac OS X. Some packages have stubs for Windows (for autocomplete).

### Installation

````
go get github.com/essentialkaos/ek
````

### Packages

* `arg` - Package provides methods for working with command-line arguments
* `cron` - Package provides methods for working with cron expressions
* `crypto` - Package with utils for working with crypto data (passwords, uuids, file hashes)
* `csv` - Package with simple (without any checks) csv parser compatible with default Go parser
* `easing` - Package with easing functions (Back, Bounce, Circ, Cubic, Elastic, Expo, Linear, Quad, Quint, Sine)
* `env` - Package provides methods for working with environment variables
* `errutil` - Package provides methods for working with errors
* `fmtc` - Package provides methods similar to fmt for colored output
* `fmtutil` - Package provides methods for output formating
* `fsutil` - Package provides methods for working with files in posix compatible systems (Linux / Mac OS X)
* `httputil` - Package provides methods for working with http request/responses
* `jsonutil` - Package provides methods for working with json data
* `knf` - Package provides methods for working with configs in KNF format
* `log` - Package with improved logger
* `mathutil` - Package with math utils
* `netutil` - Package with network utils
* `pid` - Package for working with pid files
* `rand` - Package for generating random data
* `req` - Package for working with http request
* `signal` - Package for handling signals
* `sliceutil` - Package with utils for working with slices
* `sortutil` - Package with utils for sorting slices
* `spellcheck` - Package provides spellcheck based on Damerau–Levenshtein distance algorithm
* `strutil` - Package provides utils for working with strings
* `system` - Package provides methods for working with system data (metrics/users)
* `terminal` - Package provides methods for working with user input
* `timeutil` - Package with time utils
* `tmp` - Package provides methods for working with temporary data
* `usage` - Package provides methods for rendering info for command-line tools
* `version` - Package provides methods for parsing semver version info
* `z7` - Package provides methods for working with 7z archives (p7zip wrapper)

### Projects with EK

* [ssllabs-client](https://github.com/essentialkaos/ssllabs_client) - Pretty awesome command-line client for public SSLLabs API
* [redis-cli-monitor](https://github.com/essentialkaos/redis-cli-monitor) - Tiny redis client for renamed MONITOR commands

### Test & Coverage Status

| Branch | TravisCI | CodeCov |
|--------|----------|---------|
| Master | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=master)](https://travis-ci.org/essentialkaos/ek) | [![codecov.io](https://codecov.io/github/essentialkaos/ek/coverage.svg?branch=master)](https://codecov.io/github/essentialkaos/ek?branch=master) |
| Develop | [![Build Status](https://travis-ci.org/essentialkaos/ek.svg?branch=develop)](https://travis-ci.org/essentialkaos/ek) | [![codecov.io](https://codecov.io/github/essentialkaos/ek/coverage.svg?branch=develop)](https://codecov.io/github/essentialkaos/ek?branch=develop) |

### License

[EKOL](https://essentialkaos.com/ekol)
