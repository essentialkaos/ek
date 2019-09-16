## Changelog

### 11.0.0

* `[fsutil]` `GetPerms` renamed to `GetMode`
* `[fsutil]` Added support of checking for character and block devices (`C` and `B`)
* `[knf]` Validators moved to sub-package (_incompatible changes_)
* `[knf]` Removed useless dependencies
* `[kv]` Package removed

---

### 10.18.1

* `[strutil]` Fixed bug in `Substr` method for a situation when the index of start symbol is greater than the length of the string
* `[strutil]` Fixed bug in `Substring` method for a situation when the index of start symbol is greater than the length of the string

### 10.18.0

* `[knf]` Added `no` as a valid boolean value for `GetB`
* `[knf]` Added new validators for property type validation
* `[knf]` Code refactoring

### 10.17.0

* `[cache]` Added package which provides simple in-memory key:value store

### 10.16.0

* `[timeutil]` Added support of short durations (_milliseconds, microseconds or nanoseconds_) to `PrettyDuration` method

### 10.15.0

* `[log]` Added support of ANSII colors in the output
* `[log]` Using `uint8` for level codes instead of `int`

### 10.14.0

* `[version]` Added method `IsZero` for checking empty version struct

### 10.13.1

* `[errutil]` Method `Add` now allows adding slices with errors

### 10.13.0

* `[errutil]` Added possibility to limit the number of errors to store
* `[errutil]` Method `Add` now allows adding errors from other Errors struct
* `[sliceutil]` Using in-place deduplication in `Deduplicate` method

### 10.12.2

* `[req]` Changed default user-agent to `go-ek-req/10`

### 10.12.1

* `[usage]` Fixed bug with formatting options without short name

### 10.12.0

* `[req]` Added method `PostFile` for multipart file uploading

### 10.11.1

* `[fsutil]` Fixed bug with filtering listing data

### 10.11.0

* `[pid]` Added method `IsProcessWorks` for checking process state by PID
* `[pid]` Improved process state check
* `[pid]` Improved Mac OS X support

### 10.10.2

* `[terminal]` Reading user input now is more stdin friendly (_you can pass the input through the stdin_)

### 10.10.1

* `[usage]` Fixed bug with formatting options
* `[fmtutil/table]` More copy&paste friendly table rendering

### 10.10.0

* `[emoji]` New package for working with emojis

### 10.9.1

* `[usage/completion/bash]` Improved bash completion generation

### 10.9.0

* `[usage/completion/bash]` Added bash completion generator
* `[usage/completion/zsh]` Added zsh completion generator
* `[usage/completion/fish]` Added fish completion generator
* `[usage]` Added method `info.BoundOptions` for linking command with options
* `[csv]` Added method `reader.ReadTo` for reading CSV data into slice
* `[strutil]` Fixed bug in `Fields` method
* `[initsystem]` Added caching for initsystem usage status
* `[initsystem]` Improved service state search for SysV scripts on systems with Systemd
* `[usage]` Code refactoring

### 10.8.0

* `[strutil]` Added method `Exclude` as the faster replacement for `strings.ReplaceAll`

### 10.7.1

* `[fmtutil]` Fixed bug with formatting small float numbers using `PrettySize` method

### 10.7.0

* `[jsonutil]` Added `Write` as alias for `EncodeToFile`
* `[jsonutil]` Added `Read` as alias for `DecodeFile`
* `[jsonutil]` Added `WriteGz` for writing gzipped JSON data
* `[jsonutil]` Added `ReadGz` for reading gzipped JSON data

### 10.6.0

* `[strutil]` Improved parsing logic for the `Fields` method
* `[strutil]` Added additional supported quotation marks types

### 10.5.1

* `[initsystem]` Fixed bug with parsing systemd's `failed` ActiveState status
* `[initsystem]` Added tests for output parsers
* `[initsystem]` Code refactoring

### 10.5.0

* `[fmtc]` Added new methods `LPrintf`, `LPrintln`, `TLPrintf` and `TLPrintln`
* `[fmtc]` Fixed bug with parsing reset and modification tags (_found by go-fuzz_)
* `[fmtc]` Code refactoring

### 10.4.0

* `[fmtc]` Improved work with temporary output (`TPrintf`, `TPrintln`)

### 10.3.0

* `[fsutil]` Added method `IsReadableByUser` for checking read permission for some user
* `[fsutil]` Added method `IsWritableByUser` for checking write permission for some user
* `[fsutil]` Added method `IsExecutableByUser` for checking execution permission for some user

### 10.2.0

* `[version]` Added method `Simple()` which returns simple version
* `[version]` More usage examples added

### 10.1.0

* `[system]` Improved OS version search
* `[tmp]` Package refactoring

### 10.0.0
* `[system]` Added method `GetCPUInfo` for fetching info about CPUs from procfs
* `[fmtutil/table]` Added global variable `MaxWidth` for configuration of maximum table width
* `[system]` `FSInfo` now is `FSUsage` (_incompatible changes_)
* `[system]` `MemInfo` now is `MemUsage` (_incompatible changes_)
* `[system]` `CPUInfo` now is `CPUUsage` (_incompatible changes_)
* `[system]` `InterfaceInfo` now is `InterfaceStats` (_incompatible changes_)
* `[system]` `GetFSInfo()` now is `GetFSUsage()` (_incompatible changes_)
* `[system]` `GetMemInfo()` now is `GetMemUsage()` (_incompatible changes_)
* `[system]` `GetCPUInfo()` now is `GetCPUUsage()` (_incompatible changes_)
* `[system]` `GetInterfacesInfo()` now is `GetInterfacesStats()` (_incompatible changes_)
* `[initsystem]` `HasService()` now is `IsPresent()` (_incompatible changes_)
* `[initsystem]` `IsServiceWorks()` now is `IsWorks()` (_incompatible changes_)
* `[system]` Fixed bug with parsing CPU stats data (_found by go-fuzz_)
* `[fmtc]` Fixed bug with parsing reset and modification tags (_found by go-fuzz_)
* `[initsystem]` Fixed examples
* `[fmtc]` Fixed examples
* `[system]` Added fuzz testing
* `[cron]` Code refactoring
* `[timeutil]` Code refactoring
* `[fmtutil]` Increased code coverage (97.9% → 100.0%)
* `[fmtutil/table]` Increased code coverage (99.4% → 100.0%)
* `[knf]` Increased code coverage (99.6% → 100.0%)
* `[req]` Increased code coverage (97.1% → 100.0%)
* `[pid]` Increased code coverage (97.4% → 100.0%)
* `[system]` Increased code coverage (73.8% → 79.0%)

---

### 9.28.1

* `[initsystem]` Improved application state checking in systemd
* `[system]` Fixed typo in json tag for `CPUInfo.Average`

### 9.28.0

* `[system]` Improved memory usage calculation
* `[system]` Added `Shmem` and `SReclaimable` values to `MemInfo` struct
* `[system]` Fixed typo in json tag for `MemInfo.SwapCached`
* `[system]` Improved tests

### 9.27.0

* `[system/sensors]` Added package for collecting sensors data
* `[strutil]` Added method `Substring` for safe substring extraction
* `[strutil]` Added method `Extract` for safe substring extraction
* `[strutil]` Fixed tests and example for `Substr` method
* `[strutil]` Improved tests
* `[strutil]` Code refactoring

### 9.26.3

* `[strutil]` Optimization and improvements for `ReadField` method
* `[easing]` Code refactoring
* `[fmtutil]` Code refactoring
* `[knf]` Code refactoring
* `[log]` Code refactoring
* `[options]` Code refactoring
* `[pid]` Code refactoring
* `[req]` Code refactoring
* `[sliceutil]` Code refactoring
* `[strutil]` Code refactoring
* `[system]` Code refactoring
* `[terminal]` Code refactoring
* `[timeutil]` Code refactoring
* `[uuid]` Code refactoring

### 9.26.2

* `[fmtc]` Fixed bug with parsing `{}` and `{-}` as valid color tags
* `[fmtc]` Added fuzz testing

### 9.26.1

* `[fmtutil/table]` Fixed bug with rendering data using not-configured table

### 9.26.0

* `[sliceutil]` Added method `Index` which return index of given item in slice

### 9.25.2

* `[fmtutil]` Improved size parser

### 9.25.1

* `[fmtutil]` Fixed various bugs with processing NaN values

### 9.25.0

* `[req]` Added constants with status codes

### 9.24.0
* `[req]` Added method `String` for `Query` struct for query encoding

### 9.23.0

* `[log]` Added wrapper for compatibility with stdlib logger
* `[log]` Fixed race condition issue

### 9.22.3

* `[usage]` Fixed bug with aligning option info with Unicode symbols
* `[options]` Guess option type by default value type
* `[options]` Added check for unsupported default value type

### 9.22.2

* `[system/process]` Fixed windows stubs

### 9.22.1

* `[fsutil]` Improved `CopyDir` method

### 9.22.0

* `[fsutil]` Added method `CopyDir` for recursive directories copying
* `[fsutil]` Removed useless method `Current`
* `[fsutil]` Tests refactoring
* `[fsutil]` Code refactoring

### 9.21.0

* `[system/process]` Added new type `ProcSample` as a lightweight analog of ProcInfo for CPU usage calculation
* `[system/process]` Code refactoring
* `[system/process]` Increased code coverage (75.5% → 82.4%)
* `[system]` Code refactoring

### 9.20.1

* `[fmtutil]` Added method `PrettyPerc` for formatting values in percentages

### 9.20.0

* `[fmtc]` Added methods `Print` and `Sprintln` for better compatibility with `fmt` package
* `[fmtutil/table]` Fixed minor bug with output formatting
* `[options]` Code refactoring

### 9.19.0

* `[directio]` Added sub-package `directio` for writing/reading data with using direct IO
* `[fmtc]` 256 colors support with new tags (foreground: `{#000}`, background: `{%000}`)
* `[fmtc]` Added method `Is256ColorsSupported` for checking support of 256 color output
* `[fmtc]` Improved color tags syntax
* `[fmtc]` Added tags for resetting modificators (e.g. `{!*}`)
* `[fmtc]` Removed color tags overriding (i.e. now `{*}{r} == {r*}`)
* `[color]` Added method `RGB2Term` for converting RGB colors to terminal color codes

### 9.18.1

* `[system]` Fixed bug with extra new line symbol in user `Shell` field

### 9.18.0

* `[fmtc]` Temporary output feature moved from T struct to `TPrintf` and `TPrintln`

### 9.17.4

* Dependencies now download with initial `go get` command

### 9.17.3

* `[options]` Fixed bug with using `Bound` or `Conflict` fields for options (thx to @gongled)
* `[netutil]` Code refactoring
* `[netutil]` Increased code coverage (78.8% → 87.9%)

### 9.17.2

* `[netutil]` Improved main IP search

### 9.17.1

* `[strutil]` Added usage example for `Copy` method
* `[system/procname]` Added usage examples

### 9.17.0

* `[netutil]` Ignore TUN/TAP interfaces while searching main IP address
* `[initsystem]` Added method `IsEnabled` which return info about service autostart
* `[initsystem]` Method `GetServiceState` renamed to `IsServiceWorks`
* `[strutil]` Added method `Copy` for forced copying of strings

### 9.16.0

* `[strutil]` Improved `Fields` parsing
* `[fmtutil/table]` Added method `RenderHeaders` for forced headers rendering

### 9.15.0

* `[strutil]` Added ellipsis suffix customization
* `[strutil]` Added support of custom separators for `ReadField`
* `[req]` Closing response body after parsing data
* `[system]` Fixed bug with parsing `id` command output with empty group names
* `[system]` Fixed bug with calculating transferred bytes on active interfaces
* `[system]` Improved `id` and `getent` commands output parsing
* `[system]` Code refactoring

### 9.14.5

* `[terminal]` Fixed bug with empty title output

### 9.14.4

* `[system]` Code refactoring

### 9.14.3

* `[initsystem]` Fixed bug with checking service state in systemd

### 9.14.2

* `[system]` Fixed windows stubs
* `[system]` Fixed bug with unclosed file descriptor

### 9.14.1

* `[initsystem]` Fixed bug in SysV service state determination

### 9.14.0

* `[strutil]` Added new method `ReadField` for reading space/tab separated fields from given data
* `[system]` Code refactoring

### 9.13.0

* `[system]` Improved CPU usage calculation
* `[system/process]` Code refactoring
* `[system]` Code refactoring

### 9.12.0

* `[knf]` Added new validators: `NotLen`, `NotPrefix` and `NotSuffix`
* `[knf]` Validators code refactoring

### 9.11.2

* `[system/process]` Fixed bug with parsing CPU data
* `[system/process]` Increased code coverage (0.0% → 87.5%)
* `[usage/update]` Increased code coverage (0.0% → 80.0%)

### 9.11.1

* `[system/process]` Improved error handling in `GetInfo`

### 9.11.0

* `[system]` Improved IO utilization calculation
* `[system]` Improved network speed calculation

### 9.10.0

* `[system]` Added method `GetCPUStats` which return basic CPU info from `/proc/stat`
* `[system]` Improved IO utilization calculation

### 9.9.2

* `[initsystem]` Added stubs for windows

### 9.9.1

* Code refactoring

### 9.9.0

* `[system]` Improved disk usage calculation (now it similar to `df` command output)

### 9.8.0

* `[initsystem]` New package for working with different init systems (sysv, upstart, systemd)

### 9.7.1

* `[fmtc]` Improved utf8 support in temporary messages

### 9.7.0

* `[fmtc]` Added method `NewT` which creates a new struct for working with the temporary output
* `[fmtc]` More docs about color tags
* `[knf]` Removing trailing spaces from property values

### 9.6.0

* `[system/procname]` Added method `Replace` which replace just one argument in process command

### 9.5.0

* `[knf]` Added new getters `GetU`, `GetU64` and `GetI64`
* `[usage]` Improved API for `NewInfo` method

### 9.4.0

* `[options]` Added support of mixed options (string / bool)

### 9.3.0

* `[terminal]` Improved title rendering for `ReadAnswer` method
* `[terminal]` Simplified API for `ReadAnswer` method

### 9.2.0

* `[fmtutil]` Improved floating numbers formatting with `PrettyNum`

### 9.1.4

* `[fmtutil/table]` Fixed bug with color tags in headers when colors is disabled

### 9.1.3

* `[timeutil]` Fixed bug with formatting milliseconds
* `[timeutil]` Improved tests

### 9.1.2

* `[terminal]` Fixed bug with masking password in tmux

### 9.1.1

* `[fmtutil/table]` Fixed bug with rendering data with color tags

### 9.1.0

* `[version]` Fixed bug with version comparison
* `[version]` Added method `Int()` which return version as integer

### 9.0.0

* Package `args` renamed to `options` (_incompatible changes_)
* `[fmtutil/table]` Added new package for rendering data as a table
* `[fmtutil]` Added support of separator symbol configuration
* `[usage]` Improved output about a newer version
* `[usage]` Increased code coverage (0.0% → 100%)
* `[usage]` Code refactoring

---

### 8.0.3

* `[usage]` Improved options and commands info rendering

### 8.0.2

* Overall documentation improvements

### 8.0.1

* `[system/process]` Fixed windows stubs
* `[system]` Package refactoring
* `[fsutil]` Fixed checking empty directory on FreeBSD
* `[pid]` Fixed checking process state on FreeBSD

### 8.0.0

* `[system/process]` Added method `GetMemInfo` for obtaining information about memory consumption by process.
* `[system/process]` Added method `GetInfo` which return partial info from `/proc/[PID]/stat`.
* `[system/process]` Added method `CalculateCPUUsage` which can be used for process CPU usage calculation.
* `[system]` Methods for executing commands moved to `system/exec` package (_incompatible changes_)
* `[system]` Methods for changing process name moved to `system/procname` package (_incompatible changes_)
* `[system]` Minor improvements
* `[system]` Code refactoring
* `[system]` Increased code coverage (0.0% → 79.5%)

---

### 7.5.0

* `[errutil]` Implemented error interface (_added method_ `Error() string`)
* `[errutil]` Minor improvements
* `[system]` Fixed windows stubs

### 7.4.0

* `[fmtutil]` Added flag `SeparatorFullscreen` which enable full size separator
* `[terminal/window]` Window size detection code moved from `terminal` to `terminal/window` package
* `[terminal/window]` Fixed bug with unclosed TTY file descriptor
* `[fsutil]` Fixed bug with `fsutil.IsLink` (_method returns true for symlinks_)
* `[fsutil]` Fixed bug with `fsutil.GetSize` (_method returns 0 for non-existent files_)
* `[fsutil]` Improved input arguments checks in `fsutil.CopyFile`
* `[fsutil]` Added input arguments checks to `fsutil.MoveFile`
* `[fsutil]` Increased code coverage (49.8% → 97.9%)
* `[knf]` Increased code coverage (99.2% → 99.6%)
* `[jsonutil]` Increased code coverage (92.3% → 100%)

### 7.3.0

* `[sortutil]` Added methods `NatualLess` and `StringsNatual` for natural ordering
* `[jsonutil]` Added optional argument to `EncodeToFile` method with file permissions (0644 by default)
* `[jsonutil]` Code refactoring
* `[jsonutil]` Improved tests
* `[jsonutil]` Added usage examples

### 7.2.0

* `[knf]` Return default value for the property even if config struct is nil

### 7.1.0

* `[system]` Added methods `CalculateNetworkSpeed` and `CalculateIOUtil` for metrics calculation without blocking main thread
* `[system]` Code and examples refactoring

### 7.0.3

* `[passwd]` Fixed panic in `Check` for some rare cases
* `[fsutil]` Fixed typo
* `[pid]` Fixed typo
* `[system]` Fixed typo
* `[tmp]` Fixed typo
* `[knf]` Increased code coverage

### 7.0.2

* `[version]` Fixed bug with version comparison
* `[version]` Improved version data storing model
* `[usage]` Fixed bug with new application version checking mechanics

### 7.0.1

* `[fsutil]` Fixed windows stubs for compatibility with latest changes

### 7.0.0

* `[usage]` Added interface for different ways to check application updates
* `[usage]` Added Github update checker
* `[usage]` Moved `CommandsColorTag`, `OptionsColorTag`, `Breadcrumbs` to `Info` struct (_incompatible changes_)
* `[fsutil]` Now `ListingFilter` must be passed as value instead of pointer (_incompatible changes_)
* `[fsutil]` Added support of filtering by size for `ListingFilter`
* `[version]` Now `Parse` return value instead of pointer
* `[cron]` Improved expressions parsing
* `[version]` Added fuzz testing
* `[cron]` Added fuzz testing
* `[knf]` Added fuzz testing

---

### 6.2.1

* `[usage]` Improved working with GitHub API

### 6.2.0

* `[netutil]` Now GetIP return primary IPv4 address
* `[netutil]` Added method `GetIP6` which return main IPv6 address
* `[usage]` Showing info about latest available release on GitHub

### 6.1.0

* `[knf]` Added tabs support in indentation
* `[timeutil]` Added new sequences `%n` (_new line symbol_) and `%K` (_milliseconds_)
* `[timeutil]` Code refactoring

### 6.0.0

* `[passwd]` Much secure hash generation (now with sha512, bcrypt, and AES)
* `[system]` Improved changing process and arguments names
* `[system/process]` Fixed windows stubs

---

### 5.7.1

* `[usage]` Improved build info output
* `[system]` Improved OS version search process

### 5.7.0

* `[system/process]` `GetTree` now can return tree for custom root process
* `[system/process]` Fixed threads marking
* `[fmtutil]` Added method `CountDigits` for counting the number of digits in integer
* `[terminal]` Now `PrintWarnMessage` and `PrintErrorMessage` prints messages to stderr
* `[usage]` Added support for optional arguments in commands

### 5.6.0

* `[system]` Added `Distribution` and `Version` info to `SystemInfo` struct
* `[arg]` Added bound arguments support
* `[arg]` Added conflicts arguments support
* `[arg]` Added method `Q` for merging several arguments to string (useful for `Alias`, `Bound` and `Conflicts`)

### 5.5.0

* `[system]` Added method `CurrentTTY` which return path to current tty
* `[system]` Code refactoring

### 5.4.1

* `[fmtc]` Fixed bug with parsing tags

### 5.4.0

* `[usage]` Changed color for arguments from dark gray to light gray
* `[usage]` Added breadcrumbs output for commands and options
* `[fmtutil]` Fixed special symbols colorization in `ColorizePassword`

### 5.3.0

* `[fmtutil]` Added method `ColorizePassword` for password colorization
* `[passwd]` Improved password generation and strength check

### 5.2.1

* `[log]` Code refactoring
* `[tmp]` Added permissions customization for each temp struct

### 5.2.0

* `[terminal]` Added password mask symbol color customization
* `[terminal]` [go-linenoise](https://github.com/essentialkaos/go-linenoise) updated to v3

### 5.1.1

* `[req]` Improved `Engine` initialization routine
* `[terminal]` Fixed bug in windows stub with error variable name

### 5.1.0

* `[req]` Improved `SetUserAgent` method for appending subpackages versions

### 5.0.1

* `[usage]` Fixed examples header

### 5.0.0

* `[req]` Fixed major bug with setting method through helper methods
* `[req]` Multi-client feature (_use `req.Engine` instead `req.Request` struct methods_)
* `[crypto]` Package divided into multiple packages (`hash`, `passwd`, `uuid`)
* `[uuid]` Added UUID generation based on SHA-1 hash of namespace UUID and name (_version 5_)
* `[req]` Added different types support for `Query`
* `[knf]` Added `NotContains` validator which checks if given config property contains any value from given slice
* `[kv]` Using values instead pointers
* `[system]` Added custom duration support for `GetNetworkSpeed` and `GetIOUtil`
* `[version]` Improved version parsing
* `[system]` More logical `RunAsUser` arguments naming
* `[terminal]` Minor fixes in windows stubs
* `[netutil]` Added tests
* `[system]` Code refactoring
* Added usage examples

---

### 3.5.1

* `[usage]` Using dark gray color for license and copyright
* `[fmtutil]` Added global variable `SeparatorColorTag` for separator color customization
* `[fmtutil]` Added global variable `SeparatorTitleColorTag` for separator title color customization

### 3.5.0

* `[terminal]` Using forked [go.linenoise](https://github.com/essentialkaos/go-linenoise) package instead original
* `[terminal]` Added hints support from new version of `go.linenoise`
* `[fmtc]` Light colors tag (`-`) support
* `[usage]` Using dark gray color for option values and example description
* `[tmp]` Added `DefaultDirPerms` and `DefaultFilePerms` global variables for permissions customization
* `[tmp]` Improved error handling

### 3.4.2

* `[strutil]` Fixed bug with overflowing int in `Tail` method

### 3.4.1

* `[terminal]` Improved reading user input

### 3.4.0

* `[httputil]` Added `GetRequestAddr`, `GetRemoteAddr`, `GetRemoteHost`, `GetRemotePort` methods

### 3.3.1

* `[usage]` Fixed bug with rendering command groups
* `[terminal]` Small fixes in windows stubs

### 3.3.0

* `[system/process]` Added new package for getting information about active system processes
* `[terminal]` Fixed bug with title formating in `ReadAnswer` method

### 3.2.3

* `[terminal]` Fixed bug with title formating in `ReadUI` method

### 3.2.2

* `[req]` Added content types constants

### 3.2.1

* `[knf]` Fixed typo in tests
* `[strutil]` Removed unreachable code

### 3.2.0

* `[strutil]` Added method `Len` which returns number of symbols in string
* `[strutil]` UTF-8 support for `Substr`, `Tail`, `Head` and `Ellipsis` methods
* `[strutil]` Added some benchmarks to tests
* `[fsutil]` Fixed `GetPerm` stub for Windows
* `[fsutil]` Fixed package description

### 3.1.3

* `[req]` `RequestTimeout` set to 0 (_disabled_) by default

### 3.1.2

* `[terminal]` Fixed bug with source name file conventions
* `[system]` Fixed bug with appending real user info on MacOS X

### 3.1.1

* `[req]` Small fixes in Request struct fields types

### 3.1.0

* `[req]` Lazy transport initialization
* `[req]` Added `DialTimeout` and `RequestTimeout` variables for timeouts control

### 3.0.3

* `[system]` Removed debug output

### 3.0.2

* Added makefile with some helpful commands (`fmt`, `deps`, `test`)
* Small fixes in docs

### 3.0.1

* `[sliceutil]` Code refactoring
* `[knf]` Typo fixed
* `[terminal]` Typo fixed
* Some minor changes

### 3.0.0

* `[fmtutil]` Pluralization moved from `fmtutil` to separate package `pluralize` (_incompatible changes_)
* `[pluralize]` Brand new pluralization package with more than 140 languages support
* `[timeutil]` Improved `PrettyDuration` output
* `[system]` Now `SessionInfo` contnains full user info (`Info` struct) instead username (_incompatible changes_)
* `[timeutil]` Code refactoring
* `[system]` Code refactoring
* `[log]` Code refactoring
* `[arg]` Code refactoring

---

### 2.0.2

* `[pid]` Added method `IsWorks` which return true if process with PID from PID file is active
* `[pid]` Increased code coverage

### 2.0.1

* `[terminal]` Fixed bugs with Windows stubs
* `[signal]` Fixed bugs with Windows stubs

### 2.0.0

* `[color]` New package for working with colors
* `[usage]` Added color tags support for description
* `[terminal]` Improved reading y/n answers (_incompatible changes_)
* `[strutil]` Added method `Fields` for "smart" string splitting
* `[system]` Methods `GetUsername` and `GetGroupname` deprecated
* `[system]` Added method `GroupList` for user struct which returns slice with user groups names
* `[jsonutil]` Code refactoring
* `[usage]` Code refactoring

---

### 1.8.3

* `[signal]` Added method `Send` for sending signal to process

### 1.8.2

* `[log]` Fixed bug with logging empty strings

### 1.8.1

* `[sortutil]` Added method `VersionCompare` which can be used for custom version sorting

### 1.8.0

* `[sortutil]` Added case insensitive strings sorting
* `[sliceutil]` Added `Deduplicate` method
* `[strutil]` Added `ReplaceAll` method
* `[terminal]` method `fmtutil.GetTermSize` moved to `terminal.GetSize`
* `[timeutil]` Added method `ParseDuration` which parses duration in `1w2d3h5m6s` format

### 1.7.8

* `[terminal]` Custom prompt support
* `[terminal]` Custom masking symbols support
* `[terminal]` Code refactoring

### 1.7.7

* `[fsutil]` Fixed bug in `List` method with filtering output
* `[fsutil]` Fixed bug with `NotPerms` filtering

### 1.7.6

* `[env]` Added methods for getting env vars as string, int, and float

### 1.7.5

* `[usage]` Added docs for exported fields in About struct

### 1.7.4

* `[fsutils]` Added fs walker (bash `pushd`/`popd` analog)

### 1.7.3

* `[fsutil]` Method `ListAbsolute` ranamed to `ListToAbsolute`

### 1.7.2

* `[errutil]` Added method Chain

### 1.7.1

* `[log]` Improved min level changing

### 1.7.0

* `[fsutil]` Fixed major bug with closing file descriptor after directory listing
* `[fsutil]` Fixed major bug with closing file descriptor after counting lines in file
* `[fsutil]` Fixed major bug with closing file descriptor after checking number of files in directory

### 1.6.5

* `[fsutil]` Improved docs
* `[fsutil]` Added method (wrapper) for moving files

### 1.6.4

* `[path]` Added method IsDotfile for checking dotfile names

### 1.6.3

* `[strutil]` Added methods PrefixSize and SuffixSize

### 1.6.2

* `[fsutil]` Improved working with paths
* `[fsutil]` Added method ProperPath to windows stub

### 1.6.1

* `[path]` Fixed windows stub

### 1.6.0

* `[path]` Added package for working with paths

### 1.5.1

* `[knf]` Fixed bug in HasProp method which returns true for unset properties

### 1.5.0

* `[tmp]` Improved error handling
* `[tmp]` Changed name pattern of temporary files and directories

### 1.4.5

* `[pid]` Fixed bug with PID file creation
* `[pid]` Increased coverage

### 1.4.4

* `[errutil]` Added method Num which returns number of errors

### 1.4.3

* `[errutil]` Improved Add method

### 1.4.2

* `[fsutil]` Added method `ProperPath` which return first proper path from given slice

### 1.4.1

* `[fsutil]` Added partial FreeBSD support
* `[system]` Added partial FreeBSD support
* `[log]` Some minor fixes in tests

### 1.4.0

* `[kv]` Added package with simple key-value structs

### 1.3.3

* `[strutil]` Fixed bug in Tail method

### 1.3.2

* `[strutil]` Added method Head for subtraction first symbols from the string
* `[strutil]` Added method Tail for subtraction last symbols from the string

### 1.3.1

* Improved TravisCI build script for support pkg.re
* Added pkg.re usage

### 1.3.0

* `[system]` Fixed major bug with OS X compatibility
* `[fmtutil]` Fixed tests for OS X

### 1.2.2

* `[req]` Added flag for marking connection to close

### 1.2.1

* `[crypto]` Small improvements in hash generation
* `[csv]` Increased code coverage
* `[easing]` Increased code coverage
* `[fmtc]` Increased code coverage
* `[httputil]` Increased code coverage
* `[jsonutil]` Increased code coverage
* `[pid]` Increased code coverage
* `[req]` Increased code coverage
* `[req]` Increased default timeout to 10 seconds
* `[strutil]` Increased code coverage
* `[timeutil]` Increased code coverage

### 1.2.0

* `[log]` Now buffered I/O must be enabled manually
* `[log]` Auto flushing for bufio

### 1.1.1

* `[system]` Added JSON tags for User, Group and SessionInfo structs
* `[usage]` Info now can use os.Args`[0]` for info rendering
* `[version]` Added package for working with version in semver notation

### 1.1.0

* `[arg]` Changed default fail values (int -1 → 0, float -1.0 → 0.0)
* `[arg]` Increased code coverage
* `[arg]` Many minor fixes
* `[cron]` Fixed rare bug
* `[cron]` Increased code coverage
* `[crypto]` Increased code coverage
* `[easing]` Increased code coverage
* `[errutil]` Increased code coverage
* `[fmtc]` Increased code coverage
* `[fmtutil]` Increased code coverage
* `[jsonutil]` Increased code coverage
* `[knf]` Fixed bug in Reload method for global config 
* `[knf]` Improved Reload method
* `[knf]` Increased code coverage
* `[log]` Increased code coverage
* `[mathutil]` Increased code coverage
* `[pid]` Increased code coverage
* `[rand]` Increased code coverage
* `[req]` Fixed bug with Accept header
* `[req]` Increased code coverage
* `[sliceutil]` Increased code coverage
* `[sortutil]` Increased code coverage
* `[spellcheck]` Increased code coverage
* `[strutil]` Increased code coverage
* `[system]` Added method system.SetProcName for changing process name
* `[timeutil]` Fixed bug in PrettyDuration method
* `[timeutil]` Increased code coverage
* `[tmp]` Increased code coverage

### 1.0.1

* `[system]` Fixed bug in fs usage calculation
* `[usage]` Improved new Info struct creation

### 1.0.0

Initial public release
