## Changelog

### v7.5.0

* `[errutil]` Implemented error interface (_added method_ `Error() string`)
* `[errutil]` Minor improvements
* `[system]` Fixed windows stubs

### v7.4.0

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

### v7.3.0

* `[sortutil]` Added methods `NatualLess` and `StringsNatual` for natural ordering
* `[jsonutil]` Added optional argument to `EncodeToFile` method with file permissions (0644 by default)
* `[jsonutil]` Code refactoring
* `[jsonutil]` Improved tests
* `[jsonutil]` Added usage examples

### v7.2.0

* `[knf]` Return default value for the property even if config struct is nil

### v7.1.0

* `[system]` Added methods `CalculateNetworkSpeed` and `CalculateIOUtil` for metrics calculation without blocking main thread
* `[system]` Code and examples refactoring

### v7.0.3

* `[passwd]` Fixed panic in `Check` for some rare cases
* `[fsutil]` Fixed typo
* `[pid]` Fixed typo
* `[system]` Fixed typo
* `[tmp]` Fixed typo
* `[knf]` Increased code coverage

### v7.0.2

* `[version]` Fixed bug with version comparison
* `[version]` Improved version data storing model
* `[usage]` Fixed bug with new application version checking mechanics

### v7.0.1

* `[fsutil]` Fixed windows stubs for compatibility with latest changes

### v7.0.0

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

#### v6.2.1

* `[usage]` Improved working with GitHub API

#### v6.2.0

* `[netutil]` Now GetIP return primary IPv4 address
* `[netutil]` Added method `GetIP6` which return main IPv6 address
* `[usage]` Showing info about latest available release on GitHub

#### v6.1.0

* `[knf]` Added tabs support in indentation
* `[timeutil]` Added new sequences `%n` (_new line symbol_) and `%K` (_milliseconds_)
* `[timeutil]` Code refactoring

#### v6.0.0

* `[passwd]` Much secure hash generation (now with sha512, bcrypt, and AES)
* `[system]` Improved changing process and arguments names
* `[system/process]` Fixed windows stubs

---

#### v5.7.1

* `[usage]` Improved build info output
* `[system]` Improved OS version search process

#### v5.7.0

* `[system/process]` `GetTree` now can return tree for custom root process
* `[system/process]` Fixed threads marking
* `[fmtutil]` Added method `CountDigits` for counting the number of digits in integer
* `[terminal]` Now `PrintWarnMessage` and `PrintErrorMessage` prints messages to stderr
* `[usage]` Added support for optional arguments in commands

#### v5.6.0

* `[system]` Added `Distribution` and `Version` info to `SystemInfo` struct
* `[arg]` Added bound arguments support
* `[arg]` Added conflicts arguments support
* `[arg]` Added method `Q` for merging several arguments to string (useful for `Alias`, `Bound` and `Conflicts`)

#### v5.5.0

* `[system]` Added method `CurrentTTY` which return path to current tty
* `[system]` Code refactoring

#### v5.4.1

* `[fmtc]` Fixed bug with parsing tags

#### v5.4.0

* `[usage]` Changed color for arguments from dark gray to light gray
* `[usage]` Added breadcrumbs output for commands and options
* `[fmtutil]` Fixed special symbols colorization in `ColorizePassword`

#### v5.3.0

* `[fmtutil]` Added method `ColorizePassword` for password colorization
* `[passwd]` Improved password generation and strength check

#### v5.2.1

* `[log]` Code refactoring
* `[tmp]` Added permissions customization for each temp struct

#### v5.2.0

* `[terminal]` Added password mask symbol color customization
* `[terminal]` [go-linenoise](https://github.com/essentialkaos/go-linenoise) updated to v3

#### v5.1.1

* `[req]` Improved `Engine` initialization routine
* `[terminal]` Fixed bug in windows stub with error variable name

#### v5.1.0

* `[req]` Improved `SetUserAgent` method for appending subpackages versions

#### v5.0.1

* `[usage]` Fixed examples header

#### v5.0.0

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

#### v3.5.1

* `[usage]` Using dark gray color for license and copyright
* `[fmtutil]` Added global variable `SeparatorColorTag` for separator color customization
* `[fmtutil]` Added global variable `SeparatorTitleColorTag` for separator title color customization

#### v3.5.0

* `[terminal]` Using forked [go.linenoise](https://github.com/essentialkaos/go-linenoise) package instead original
* `[terminal]` Added hints support from new version of `go.linenoise`
* `[fmtc]` Light colors tag (`-`) support
* `[usage]` Using dark gray color for option values and example description
* `[tmp]` Added `DefaultDirPerms` and `DefaultFilePerms` global variables for permissions customization
* `[tmp]` Improved error handling

#### v3.4.2

* `[strutil]` Fixed bug with overflowing int in `Tail` method

#### v3.4.1

* `[terminal]` Improved reading user input

#### v3.4.0

* `[httputil]` Added `GetRequestAddr`, `GetRemoteAddr`, `GetRemoteHost`, `GetRemotePort` methods

#### v3.3.1

* `[usage]` Fixed bug with rendering command groups
* `[terminal]` Small fixes in windows stubs

#### v3.3.0

* `[system/process]` Added new package for getting information about active system processes
* `[terminal]` Fixed bug with title formating in `ReadAnswer` method

#### v3.2.3

* `[terminal]` Fixed bug with title formating in `ReadUI` method

#### v3.2.2

* `[req]` Added content types constants

#### v3.2.1

* `[knf]` Fixed typo in tests
* `[strutil]` Removed unreachable code

#### v3.2.0

* `[strutil]` Added method `Len` which returns number of symbols in string
* `[strutil]` UTF-8 support for `Substr`, `Tail`, `Head` and `Ellipsis` methods
* `[strutil]` Added some benchmarks to tests
* `[fsutil]` Fixed `GetPerm` stub for Windows
* `[fsutil]` Fixed package description

#### v3.1.3

* `[req]` `RequestTimeout` set to 0 (_disabled_) by default

#### v3.1.2

* `[terminal]` Fixed bug with source name file conventions
* `[system]` Fixed bug with appending real user info on MacOS X

#### v3.1.1

* `[req]` Small fixes in Request struct fields types

#### v3.1.0

* `[req]` Lazy transport initialization
* `[req]` Added `DialTimeout` and `RequestTimeout` variables for timeouts control

#### v3.0.3

* `[system]` Removed debug output

#### v3.0.2

* Added makefile with some helpful commands (`fmt`, `deps`, `test`)
* Small fixes in docs

#### v3.0.1

* `[sliceutil]` Code refactoring
* `[knf]` Typo fixed
* `[terminal]` Typo fixed
* Some minor changes

#### v3.0.0

* `[fmtutil]` Pluralization moved from `fmtutil` to separate package `pluralize` (_incompatible changes_)
* `[pluralize]` Brand new pluralization package with more than 140 languages support
* `[timeutil]` Improved `PrettyDuration` output
* `[system]` Now `SessionInfo` contnains full user info (`Info` struct) instead username (_incompatible changes_)
* `[timeutil]` Code refactoring
* `[system]` Code refactoring
* `[log]` Code refactoring
* `[arg]` Code refactoring

---

#### v2.0.2

* `[pid]` Added method `IsWorks` which return true if process with PID from PID file is active
* `[pid]` Increased code coverage

#### v2.0.1

* `[terminal]` Fixed bugs with Windows stubs
* `[signal]` Fixed bugs with Windows stubs

#### v2.0.0

* `[color]` New package for working with colors
* `[usage]` Added color tags support for description
* `[terminal]` Improved reading y/n answers (_incompatible changes_)
* `[strutil]` Added method `Fields` for "smart" string splitting
* `[system]` Methods `GetUsername` and `GetGroupname` deprecated
* `[system]` Added method `GroupList` for user struct which returns slice with user groups names
* `[jsonutil]` Code refactoring
* `[usage]` Code refactoring

---

#### v1.8.3

* `[signal]` Added method `Send` for sending signal to process

#### v1.8.2

* `[log]` Fixed bug with logging empty strings

#### v1.8.1

* `[sortutil]` Added method `VersionCompare` which can be used for custom version sorting

#### v1.8.0

* `[sortutil]` Added case insensitive strings sorting
* `[sliceutil]` Added `Deduplicate` method
* `[strutil]` Added `ReplaceAll` method
* `[terminal]` method `fmtutil.GetTermSize` moved to `terminal.GetSize`
* `[timeutil]` Added method `ParseDuration` which parses duration in `1w2d3h5m6s` format

#### v1.7.8

* `[terminal]` Custom prompt support
* `[terminal]` Custom masking symbols support
* `[terminal]` Code refactoring

#### v1.7.7

* `[fsutil]` Fixed bug in `List` method with filtering output
* `[fsutil]` Fixed bug with `NotPerms` filtering

#### v1.7.6

* `[env]` Added methods for getting env vars as string, int, and float

#### v1.7.5

* `[usage]` Added docs for exported fields in About struct

#### v1.7.4

* `[fsutils]` Added fs walker (bash `pushd`/`popd` analog)

#### v1.7.3

* `[fsutil]` Method `ListAbsolute` ranamed to `ListToAbsolute`

#### v1.7.2

* `[errutil]` Added method Chain

#### v1.7.1

* `[log]` Improved min level changing

#### v1.7.0

* `[fsutil]` Fixed major bug with closing file descriptor after directory listing
* `[fsutil]` Fixed major bug with closing file descriptor after counting lines in file
* `[fsutil]` Fixed major bug with closing file descriptor after checking number of files in directory

#### v1.6.5

* `[fsutil]` Improved docs
* `[fsutil]` Added method (wrapper) for moving files

#### v1.6.4

* `[path]` Added method IsDotfile for checking dotfile names

#### v1.6.3

* `[strutil]` Added methods PrefixSize and SuffixSize

#### v1.6.2

* `[fsutil]` Improved working with paths
* `[fsutil]` Added method ProperPath to windows stub

#### v1.6.1

* `[path]` Fixed windows stub

#### v1.6.0

* `[path]` Added package for working with paths

#### v1.5.1

* `[knf]` Fixed bug in HasProp method which returns true for unset properties

#### v1.5.0

* `[tmp]` Improved error handling
* `[tmp]` Changed name pattern of temporary files and directories

#### v1.4.5

* `[pid]` Fixed bug with PID file creation
* `[pid]` Increased coverage

#### v1.4.4

* `[errutil]` Added method Num which returns number of errors

#### v1.4.3

* `[errutil]` Improved Add method

#### v1.4.2

* `[fsutil]` Added method `ProperPath` which return first proper path from given slice

#### v1.4.1

* `[fsutil]` Added partial FreeBSD support
* `[system]` Added partial FreeBSD support
* `[log]` Some minor fixes in tests

#### v1.4.0

* `[kv]` Added package with simple key-value structs

#### v1.3.3

* `[strutil]` Fixed bug in Tail method

#### v1.3.2

* `[strutil]` Added method Head for subtraction first symbols from the string
* `[strutil]` Added method Tail for subtraction last symbols from the string

#### v1.3.1

* Improved TravisCI build script for support pkg.re
* Added pkg.re usage

#### v1.3

* `[system]` Fixed major bug with OS X compatibility
* `[fmtutil]` Fixed tests for OS X

#### v1.2.2

* `[req]` Added flag for marking connection to close

#### v1.2.1

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

#### v1.2

* `[log]` Now buffered I/O must be enabled manually
* `[log]` Auto flushing for bufio

#### v1.1.1

* `[system]` Added JSON tags for User, Group and SessionInfo structs
* `[usage]` Info now can use os.Args`[0]` for info rendering
* `[version]` Added package for working with version in semver notation

#### v1.1

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

#### v1.0.1

* `[system]` Fixed bug in fs usage calculation
* `[usage]` Improved new Info struct creation

#### v1

Initial public release
