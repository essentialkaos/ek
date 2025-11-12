## Changelog

### [13.36.1](https://kaos.sh/ek/13.36.1)

- **`[spinner]`** Description now limited by the size of the line
- **`[terminal/tty]`** Fixed bug with closing tty file descriptor

### [13.36.0](https://kaos.sh/ek/13.36.0)

- **`[spinner]`** More efficient spinner rendering

### [13.35.4](https://kaos.sh/ek/13.35.4)

- **`[errors]`** Added method `Bundle.Addf`

### [13.35.3](https://kaos.sh/ek/13.35.3)

- **`[i18n]`** Added support for prefix/suffix data to methods `Text.With` and `Text.Write`

### [13.35.2](https://kaos.sh/ek/13.35.2)

- **`[i18n]`** Added method `Text.Start`
- **`[i18n]`** Added method `Text.End`
- **`[i18n]`** Added method `Text.Write`

### [13.35.1](https://kaos.sh/ek/13.35.1)

- **`[timeutil]`** Added method `Period.Intersects`

### [13.35.0](https://kaos.sh/ek/13.35.0)

- **`[cache/fs]`** Added keys/items iterators
- **`[cache/fs]`** Added `Invalidate` method
- **`[cache/memory]`** Added keys/items iterators
- **`[cache/memory]`** Added `Invalidate` method

### [13.34.1](https://kaos.sh/ek/13.34.1)

- **`[i18n]`** Added method `Data.Has`
- **`[i18n]`** Improved working with empty `Data` map

### [13.34.0](https://kaos.sh/ek/13.34.0)

- **`[i18n]`** Type `String` has been renamed to `Text`
- **`[i18n]`** Added method `Text.Format`

### [13.33.1](https://kaos.sh/ek/13.33.1)

- **`[fmtc]`** Code refactoring
- **`[log]`** Code refactoring

### [13.33.0](https://kaos.sh/ek/13.33.0)

- **`[sliceutil]`** Added method `ToAny`
- **`[sliceutil]`** Method `IntToInterface` marked as deprecated
- **`[sliceutil]`** Method `StringToInterface` marked as deprecated

### [13.32.0](https://kaos.sh/ek/13.32.0)

- **`[log]`** Added handler `PanicHandler` for `panic`
- **`[req]`** Fixed minor bug with parsing `Retry-After` header
- **`[knf]`** Code refactoring

### [13.31.3](https://kaos.sh/ek/13.31.3)

- **`[timeutil]`** Added usage example for `AddWorkdays`

### [13.31.2](https://kaos.sh/ek/13.31.2)

- **`[timeutil]`** Added method `AddWorkdays`

### [13.31.1](https://kaos.sh/ek/13.31.1)

- **`[passthru]`** Calculate speed even if total data size is not provided

### [13.31.0](https://kaos.sh/ek/13.31.0)

- **`[req]`** Added support of [`Retry-After`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Retry-After) header to `Retrier`
- **`[knf/validators]`** Added support of string values to `SizeGreater` validator
- **`[knf/validators]`** Added support of string values to `SizeLess` validator
- **`[knf/validators]`** Improved support of empty property values in `SizeLess`, `SizeGreater`, and `InRange` validators
- **`[knf/validators]`** Improved error messages

### [13.30.1](https://kaos.sh/ek/13.30.1)

- **`[hashutil]`** Added method `Hash.Equal`
- **`[hashutil]`** Added method `Hash.EqualString`
- **`[hashutil]`** Added method `Hash.IsEmpty`

### [13.30.0](https://kaos.sh/ek/13.30.0)

- **`[options]`** Added method `Map.SetIf`
- **`[options]`** Improved method `Map.Set`

### [13.29.1](https://kaos.sh/ek/13.29.1)

- **`[support]`** Fixed minor bug with printing custom binary name

### [13.29.0](https://kaos.sh/ek/13.29.0)

- **`[selfupdate]`** Added new package for application self-update
- **`[hashutil]`** Using `Hash` type instead of strings
- **`[timeutil]`** Added `time.Duration` wrapper for printing pretty duration
- **`[spinner]`** Added support of different duration formats
- **`[strutil]`** Added method `Wrap`
- **`[color]`** Code refactoring
- **`[cron]`** Code refactoring
- **`[csv]`** Code refactoring
- **`[events]`** Code refactoring
- **`[fmtutil]`** Code refactoring
- **`[fmtutil/filetree]`** Code refactoring
- **`[i18n]`** Code refactoring
- **`[initsystem]`** Code refactoring
- **`[knf]`** Code refactoring
- **`[log]`** Code refactoring
- **`[options]`** Code refactoring
- **`[progress]`** Code refactoring
- **`[req]`** Code refactoring
- **`[secstr]`** Code refactoring
- **`[spellcheck]`** Code refactoring
- **`[strutil]`** Code refactoring
- **`[support]`** Code refactoring
- **`[system]`** Code refactoring
- **`[timeutil]`** Code refactoring
- **`[usage]`** Code refactoring
- **`[usage/update]`** Code refactoring
- **`[mathutil]`** Methods `Min` and `Max` marked as deprecated
- **`[fmtc]`** Deprecated methods have been removed
- **`[fsutil]`** Deprecated methods have been removed
- **`[knf]`** Deprecated methods have been removed
- **`[knf/validators]`** Deprecated validators have been removed
- **`[log]`** Deprecated variables have been removed
- **`[sliceutil]`** Deprecated methods have been removed

### [13.28.2](https://kaos.sh/ek/13.28.2)

- **`[options]`** Improved support of `--` option (_POSIX standard convention that means "everything that follows is not an option"_)
- **`[options]`** Code refactoring

### [13.28.1](https://kaos.sh/ek/13.28.1)

- **`[sliceutil]`** Added method `Filter`

### [13.28.0](https://kaos.sh/ek/13.28.0)

- **`[hashutil]`** Added new package with helpers for calculating hashes of files, byte slices, and strings using any hasher compatible with `hash.Hash`
- **`[req]`** Added methods `Response.JSONWithHash` and `Response.SaveWithHash` for transparent data's hash calculation
- **`[req]`** Code refactoring
- **`[hash]`** Package deprecated
- **`[errutil]`** Package deleted

### [13.27.4](https://kaos.sh/ek/13.27.4)

- **`[timeutil]`** Added method `ParseWithAny`

### [13.27.3](https://kaos.sh/ek/13.27.3)

- **`[log]`** Added method `Fields.AddF`

### [13.27.2](https://kaos.sh/ek/13.27.2)

- **`[sliceutil]`** Added method `Shuffle`

### [13.27.1](https://kaos.sh/ek/13.27.1)

- **`[system]`** Collect and add a unique system ID to `SystemInfo`
- **`[support]`** Added unique system ID to system info

### [13.27.0](https://kaos.sh/ek/13.27.0)

- **`[system]`** Read sessions data from [`utmp`](https://www.man7.org/linux/man-pages/man5/utmp.5.html) file
- **`[system]`** Added `Host` info to `SessionInfo`
- **`[system]`** User info struct replaced by user name in `SessionInfo`

### [13.26.2](https://kaos.sh/ek/13.26.2)

- **`[support]`** Improved handling empty slices with data (_deps, apps, services…_)
- **`[support/apps]`** Added method `ExtractVersion` for extracting version info from command output

### [13.26.1](https://kaos.sh/ek/13.26.1)

- **`[usage/update]`** Use token from `GH_TOKEN` environment variable for requests to Github API
- **`[usage/update]`** Use token from `GL_TOKEN` environment variable for requests to Gitlab API
- **`[usage/update]`** Gitlab update checker refactoring

### [13.26.0](https://kaos.sh/ek/13.26.0)

- **`[usage]`** Added support of environment variables to `Info`

### [13.25.0](https://kaos.sh/ek/13.25.0)

- **`[log]`** Added `DiscardFields` option to discard all fields from log messages

### [13.24.2](https://kaos.sh/ek/13.24.2)

- **`[fmtutil/table]`** Fixed separator size when full screen mode is disabled

### [13.24.1](https://kaos.sh/ek/13.24.1)

- **`[log]`** Fixed a bug with printing divider when using JSON format

### [13.24.0](https://kaos.sh/ek/13.24.0)

- **`[knf/validators]`** Added more validators for time duration (`TypeDur`, `DurShorter`, `DurLonger`)
- **`[strutil]`** Added support for negative substring end parameter
- **`[log]`** Print messages with `AUX` level as `INFO` for JSON format
- **`[knf/validators]`** `LenLess` validator renamed to `LenShorter`
- **`[knf/validators]`** `LenGreater` validator renamed to `LenLonger`
- **`[emoji]`** Code refactoring
- **`[fmtc]`** Code refactoring
- **`[knf/value]`** Code refactoring
- **`[log]`** Code refactoring
- **`[system]`** Code refactoring

### [13.23.1](https://kaos.sh/ek/13.23.1)

- **`[support]`** Fixed separator title alignment for info sections

### [13.23.0](https://kaos.sh/ek/13.23.0)

- **`[fmtutil]`** Added title alignment for separator

### [13.22.1](https://kaos.sh/ek/13.22.1)

- **`[req]`** Added new helpers for working with `Headers`

### [13.22.0](https://kaos.sh/ek/13.22.0)

- **`[uuid/prefixed]`** Added package for encoding and decoding prefixed UUID's
- **`[reutil]`** Added package with helpers for working with regular expressions
- **`[req]`** Fixed bug with canceling context of request with timeout too early

### [13.21.1](https://kaos.sh/ek/13.21.1)

- **`[timeutil]`** Added method `LocalTimezone`
- **`[req]`** Added helper `Response.Save`

### [13.21.0](https://kaos.sh/ek/13.21.0)

- **`[knf/united]`** Added methods `Is` and `Has`
- **`[knf]`** Method `HasProp` renamed to `Has`
- **`[knf]`** Code refactoring

### [13.20.5](https://kaos.sh/ek/13.20.5)

- **`[system/container]`** Added [`containerd`](https://containerd.io) support
- **`[system/container]`** Added [`K8s`](https://kubernetes.io) support
- **`[support]`** Added `containerd` support

### [13.20.4](https://kaos.sh/ek/13.20.4)

- **`[req]`** Added new helpers for working with `Query`

### [13.20.3](https://kaos.sh/ek/13.20.3)

- **`[timeutil]`** Better duration calculation for `DurationAs` and `Period.DurationAs`
- **`[log]`** Improved output of empty fields in text format
- **`[req]`** Added new auth provider `AuthHeader`

### [13.20.2](https://kaos.sh/ek/13.20.2)

- **`[knf/validators/time]`** Added timezone validator
- **`[timeutil]`** Added method `DurationAs`
- **`[timeutil]`** Method `Period.DurationIn` renamed to `Period.DurationAs`
- **`[knf/validators]`** Code refactoring
- **`[timeutil]`** Code refactoring

### [13.20.1](https://kaos.sh/ek/13.20.1)

- **`[timeutil]`** Added method `Period.Stringf`
- **`[timeutil]`** `Period.String` returns period range instead of duration
- **`[log]`** Fixed bug with fields color for debug messages

### [13.20.0](https://kaos.sh/ek/13.20.0)

- **`[timeutil]`** Added new struct for period with the start and end date

### [13.19.2](https://kaos.sh/ek/13.19.2)

- **`[timeutil]`** Added helper `Since`
- **`[spellcheck]`** Code refactoring
- **`[system]`** Code refactoring

### [13.19.1](https://kaos.sh/ek/13.19.1)

- **`[ansi]`** Code refactoring
- **`[cron]`** Code refactoring
- **`[csv]`** Code refactoring
- **`[directio]`** Code refactoring
- **`[fmtc]`** Code refactoring
- **`[fmtutil]`** Code refactoring
- **`[i18n]`** Code refactoring
- **`[knf]`** Code refactoring
- **`[log]`** Code refactoring
- **`[passwd]`** Code refactoring
- **`[path]`** Code refactoring
- **`[progress]`** Code refactoring
- **`[rand]`** Code refactoring
- **`[req]`** Code refactoring
- **`[sliceutil]`** Code refactoring
- **`[spellcheck]`** Code refactoring
- **`[strutil]`** Code refactoring
- **`[support/network]`** Code refactoring
- **`[system]`** Code refactoring
- **`[timeutil]`** Code refactoring
- **`[uuid]`** Code refactoring

### [13.19.0](https://kaos.sh/ek/13.19.0)

- **`[timeutil]`** Methods `PrevDay`, `PrevMonth`, `PrevYear`, `NextDay`, `NextMonth`, `NextYear`, `PrevWorkday`, `PrevWeekend`, `NextWorkday`, and `NextWeekend` returns dates rounded to the start of the date
- **`[timeutil]`** Added helper `PrevWeek`
- **`[timeutil]`** Added helper `NextWeek`

### [13.18.0](https://kaos.sh/ek/13.18.0)

- **`[log]`** Added helper `Field.Mask`
- **`[log]`** Added helper `Field.Compact`
- **`[log]`** Added helper `Field.Head`
- **`[log]`** Added helper `Field.Tail`

### [13.17.0](https://kaos.sh/ek/13.17.0)

- **`[knf/validators/cron]`** Added new sub-package with KNF validator for cron expressions
- **`[knf/validators/time]`** Added new sub-package with KNF validator for time format

### [13.16.0](https://kaos.sh/ek/13.16.0)

- **`[i18n]`** Added method for bundle validation `ValidateBundle`
- **`[timeutil]`** Added helper `IsWeekend`
- **`[timeutil]`** Added helper `EndOfHour`
- **`[timeutil]`** Added helper `EndOfDay`
- **`[timeutil]`** Added helper `EndOfWeek`
- **`[timeutil]`** Added helper `EndOfMonth`
- **`[timeutil]`** Added helper `EndOfYear`
- **`[timeutil]`** Added helper `Until`
- **`[i18n]`** Code refactoring
- **`[timeutil]`** Code refactoring

### [13.15.12](https://kaos.sh/ek/13.15.12)

- **`[sliceutil]`** Added method `Diff`

### [13.15.11](https://kaos.sh/ek/13.15.11)

- **`[strutil]`** Added method `JoinFunc`

### [13.15.10](https://kaos.sh/ek/13.15.10)

- **`[terminal/input]`** Fixed bug with hiding password when using custom validator
- **`[terminal/input]`** Improved handling of input data passed through stdin

### [13.15.9](https://kaos.sh/ek/13.15.9)

- **`[support/pkgs]`** Fixed bug with collecting info about multilib versions of RPM packages

### [13.15.8](https://kaos.sh/ek/13.15.8)

- **`[terminal/input]`** Improved support for custom title formatting with `fmtc` tags

### [13.15.7](https://kaos.sh/ek/13.15.7)

- **`[strutil]`** Improved method `Mask`

### [13.15.6](https://kaos.sh/ek/13.15.6)

- **`[strutil]`** Added method `Mask`

### [13.15.5](https://kaos.sh/ek/13.15.5)

- **`[log]`** Added method `NewFields`

### [13.15.4](https://kaos.sh/ek/13.15.4)

- **`[log]`** Added `Levels` method which returns all supported log levels
- **`[log]`** `LogLevels` variable marked as deprecated
- **`[fmtutil/panel]`** Code refactoring

### [13.15.3](https://kaos.sh/ek/13.15.3)

- **`[knf/united]`** Return `errors.Errors` instead of `[]error` from `Validate`

### [13.15.2](https://kaos.sh/ek/13.15.2)

- **`[color]`** Improved named colors support

### [13.15.1](https://kaos.sh/ek/13.15.1)

- **`[options]`** Added method `Argument.Uint64`
- **`[options]`** Return type of method `Argument.Uint` set from `uint64` to `uint`

### [13.15.0](https://kaos.sh/ek/13.15.0)

- **`[color]`** Added `fmt.GoStringer` support for structs
- **`[req]`** Added more authentication methods (`OAuth`, `Digest`, `VAPID`, `AWS4`, `API Key`)
- **`[req]`** Added support of custom authentication methods
- **`[color]`** Code refactoring

### [13.14.2](https://kaos.sh/ek/13.14.2)

- **`[log]`** Added variable `LogLevels` with all supported log level names

### [13.14.1](https://kaos.sh/ek/13.14.1)

- **`[usage]`** Fixed bug with rendering examples (_introduced in `13.14.0`_)
- **`[path]`** Added Windows support

### [13.14.0](https://kaos.sh/ek/13.14.0)

- **`[fmtutil/filetree]`** Added experimental package for printing file tree
- **`[fmtc]`** Added methods `Printfn`, `Fprintfn`, `Sprintfn`, and `LPrintfn`
- **`[req]`** Added custom encoder for arrays in `Query`
- **`[fmtutil/table]`** Improved support of data with ANSI escape sequences
- **`[fmtc]`** `NameColor` marked as deprecated (_use method `AddColor` instead_)

### [13.13.1](https://kaos.sh/ek/13.13.1)

- **`[spellcheck]`** Distance calculation method now public

### [13.13.0](https://kaos.sh/ek/13.13.0)

- **`[req]`** Added support for different types of slices to `Query`
- **`[req]`** Added support for `fmt.Stringer` interface to `Query`
- **`[req]`** Added interface for custom struct encoding for `Query`
- **`[req]`** Improved `Query`encoding

### [13.12.0](https://kaos.sh/ek/13.12.0)

- **`[req]`** Added custom timeout per request
- **`[req]`** Added `Retrier`
- **`[req]`** Make `Limiter` public
- **`[log]`** Added `WithFullCallerPath` option to enable the output of the full caller path
- **`[strutil]`** Added support of escaped strings to `Fields`
- **`[strutil]`** Added fuzz tests for `Fields` method
- **`[knf]`** Fixed build of fuzz tests

### [13.11.0](https://kaos.sh/ek/13.11.0)

- **`[req]`** Added request limiter

### [13.10.1](https://kaos.sh/ek/13.10.1)

- **`[mathutil]`** Added shorthand helper `B`

### [13.10.0](https://kaos.sh/ek/13.10.0)

> [!IMPORTANT]
> This release contains breaking changes to the `input.Read`, `input.ReadPassword`, and `input.ReadPasswordSecure` methods. Prior to this release, all of these methods took a boolean argument to disallow empty input. Since we are adding input validators, you will need to use the `NotEmpty` validator for the same behaviour.

- **`[fmtutil/table]`** Added automatic breaks feature
- **`[terminal/input]`** Added input validation feature
- **`[terminal/input]`** Fixed bug with hiding the password when `HidePassword` is set to true and an empty input error is displayed
- **`[terminal/input]`** Fixed bug with printing new line after input field on error

### [13.9.2](https://kaos.sh/ek/13.9.2)

- **`[knf]`** Added helper `Q`
- **`[fmtc]`** Code refactoring
- **`[usage]`** Code refactoring

### [13.9.1](https://kaos.sh/ek/13.9.1)

- **`[errors]`** Fixed bug with extra newline character at the end of `Error` output

### [13.9.0](https://kaos.sh/ek/13.9.0)

- **`[errors]`** Added new package of utilities for working with errors
- **`[knf]`** Added type `Validators` for `Validator` slice
- **`[knf]`** Code refactoring
- **`[options]`** Code refactoring
- **`[errutil]`** Package deprecated

### [13.8.1](https://kaos.sh/ek/13.8.1)

- **`[req]`** `AutoDiscard` now doesn't affect responses with successful response status codes (`200`-`299`)

### [13.8.0](https://kaos.sh/ek/13.8.0)

- **`[knf/validators]`** Added support of `int64`, `uint`, and `uint64` to `Less` validator
- **`[knf/validators]`** Added support of `int64`, `uint`, and `uint64` to `Greater` validator
- **`[knf/validators]`** Added validator `SizeGreater`
- **`[knf/validators]`** Added validator `SizeLess`

### [13.7.0](https://kaos.sh/ek/13.7.0)

- **`[fmtutil]`** Added support of `kk`, `kkk`, `kib`, `mib`, `gib`, and `tib` in `ParseSize`
- **`[knf]`** Added getter `GetSZ` for reading value as the size
- **`[knf/united]`** Added getter `GetSZ` for reading value as the size
- **`[knf/validators]`** Added validator `TypeSize`
- **`[knf/validators]`** Added validator `InRange`
- **`[fmtutil]`** Code refactoring
- **`[knf/united]`** Improved usage examples

### [13.6.1](https://kaos.sh/ek/13.6.1)

- **`[req]`** Guess the value of the `Content-Type` header based on the request body type

### [13.6.0](https://kaos.sh/ek/13.6.0)

- **`[setup]`** Added package to install/uninstall application as a service
- **`[support/deps]`** Improved collecting and filtering dependencies info
- **`[support/kernel]`** Added simple globs support for parameter names

### [13.5.1](https://kaos.sh/ek/13.5.1)

- **`[mathutil]`** Added method `FromPerc`
- **`[mathutil]`** Code refactoring

### [13.5.0](https://kaos.sh/ek/13.5.0)

- **`[support/resources]`** Added package for collecting info about CPU and memory
- **`[support/kernel]`** Added package for collecting OS kernel parameters
- **`[system/sysctl]`** Added method `All` to get all kernel parameters
- **`[system]`** Added macOS support for `GetCPUInfo`
- **`[system]`** Added macOS support for `GetMemUsage`
- **`[support]`** Added locale to default OS info output
- **`[mathutil]`** Added methods `IsInt`, `IsFloat`, and `IsNumber`
- **`[support]`** Added info if CGO is used for build

### [13.4.0](https://kaos.sh/ek/13.4.0)

- **`[system/sysctl]`** Added new package for reading kernel parameters
- **`[strutil]`** Added method `SqueezeRepeats`

### [13.3.5](https://kaos.sh/ek/13.3.5)

- **`[timeutil]`** Improved formatting of days and seconds in `MiniDuration`

### [13.3.4](https://kaos.sh/ek/13.3.4)

- **`[timeutil]`** Added support of minutes, hours and days to `MiniDuration`
- **`[timeutil]`** Added separator customization to `MiniDuration`
- **`[timeutil]`** Added usage examples for `MiniDuration`

### [13.3.3](https://kaos.sh/ek/13.3.3)

- **`[support/apps]`** Fixed Docker version extraction
- **`[support]`** Improved output of large list of IPv4 and IPv6 addresses

### [13.3.2](https://kaos.sh/ek/13.3.2)

- **`[support/apps]`** Added support for Docker, Podman, and LXC
- **`[rand]`** Removed method `Int`
- **`[rand]`** Code refactoring

### [13.3.1](https://kaos.sh/ek/13.3.1)

- **`[cache]`** Added constants with duration
- **`[cache/fs]`** Using `cache.Duration` instead of `time.Duration`
- **`[cache/memory]`** Using `cache.Duration` instead of `time.Duration`
- **`[cache/fs]`** Added check for passed expiration duration in `Set`
- **`[cache/fs]`** Code refactoring
- **`[cache/fs]`** Fixed bug with closing item file after reading data
- **`[cache/fs]`** Fixed bug with removing expired items in `Get`

### [13.3.0](https://kaos.sh/ek/13.3.0)

- **`[cache/fs]`** Added cache with file system storage
- **`[cache]`** In-memory cache moved to `cache/memory`
- **`[sliceutil]`** Added method `Join`

### [13.2.1](https://kaos.sh/ek/13.2.1)

- **`[terminal/input]`** Added `NewLine` flag
- **`[sliceutil]`** Methods `Copy`, `IsEqual`, `Index`, `Contains`, and `Deduplicate` marked as deprecated
- **`[terminal/input]`** Improved TMUX support


### [13.2.0](https://kaos.sh/ek/13.2.0)

- **`[errutil]`** Added method `Wrap`
- **`[passthru]`** `Reader` now implements only `io.Reader` interface, not `io.ReadCloser`
- **`[passthru]`** `Writer` now implements only `io.Writer` interface, not `io.WriteCloser`

### [13.1.0](https://kaos.sh/ek/13.1.0)

- **`[env]`** Fixed compatibility with Windows
- **`[support]`** Add basic support of collecting info on Windows
- **`[support/apps]`** Added Windows support
- **`[support/network]`** Added Windows support

### [13.0.0](https://kaos.sh/ek/13.0.0)

> [!CAUTION]
> In this release, we have changed the logic behind some knf validators (`Less`, `Greater`, `LenLess`, `LenGreater`) from negative to positive check. This means that if you have any of these validators, you need to swap them (`Less` → `Greater`, `Greater` → `Less`, `LenLess` → `LenGreater`, `LenGreater` → `LenLess`) to keep the logic of the validation.

- **`[knf/validators]`** Changed logic from negative to positive check for `Less`, `Greater`, `LenLess`, and `LenGreater` validators.
- **`[knf/validators]`** Validator `LenNotEquals` renamed to `LenEquals`
- **`[knf/validators]`** Validator `NotPrefix` renamed to `NotPrefix`
- **`[knf/validators]`** Validator `HasSuffix` renamed to `HasSuffix`
- **`[knf/validators]`** Validator `Equals` renamed to `NotEquals`
- **`[log]`** Code refactoring

---

### [12.130.0](https://kaos.sh/ek/12.130.0)

- **`[knf/validators]`** Added validators `Set`, `SetToAny`, `SetToAnyIgnoreCase`, `LenLess`, `LenGreater`, and `LenNotEquals`
- **`[knf/validators]`** Validators `Empty`, `NotContains`, and `NotLen` marked as deprecated
- **`[knf/validators]`** Improved validators input validation
- **`[knf/validators/fs]`** Improved validators input validation
- **`[knf/validators/regexp]`** Improved validators input validation
- **`[knf/validators]`** Code refactoring
- **`[knf/validators/system]`** Code refactoring

### [12.129.0](https://kaos.sh/ek/12.129.0)

- **`[fsutil]`** Added method `GetModeOctal`

### [12.128.0](https://kaos.sh/ek/12.128.0)

- **`[pager]`** Disable `PAGER` environment variable usage by default
- **`[pager]`** Added `AllowEnv` option to allow the use of the `PAGER` environment variable

### [12.127.0](https://kaos.sh/ek/12.127.0)

- **`[options]`** Added method `Delete`
- **`[usage]`** Added support for optional arguments to usage info
- **`[uuid]`** Added [UUID7](https://uuid7.com) generator
- **`[options]`** Code refactoring

### [12.126.1](https://kaos.sh/ek/12.126.1)

- **`[usage]`** Added `Info.WrapLen` option for text wrapping configuration

### [12.126.0](https://kaos.sh/ek/12.126.0)

- **`[support]`** Added support of Docker with gVisor runtime
- **`[system/container]`** Added support of Docker with gVisor runtime
- **`[usage]`** Added automatic wrapping for example, command, and option descriptions
- **`[color]`** Documentation improvements
- **`[sliceutil]`** Code refactoring

### [12.125.1](https://kaos.sh/ek/12.125.1)

- **`[support/network]`** Added Cloudflare trace support for public IP resolution
- **`[terminal]`** Code refactoring

### [12.125.0](https://kaos.sh/ek/12.125.0)

- **`[i18n]`** Added new package for internationalization
- **`[log]`** Added field collection support
- **`[log]`** Ignore fields without key
- **`[log]`** Fixed formatting of JSON output when no message is passed
- **`[pluralize]`** Fixed bug with handling negative numbers
- **`[pluralize]`** Code refactoring

### [12.124.0](https://kaos.sh/ek/12.124.0)

- **`[env]`** Add `Variable` struct for lazy environment reading
- **`[fmtc]`** Add support of `FMTC_NO_BOLD`, `FMTC_NO_ITALIC`, and `FMTC_NO_BLINK` environment variables

### [12.123.2](https://kaos.sh/ek/12.123.2)

- **`[terminal]`** Fixed bug with output messages from `Error` and `Warn` to stdout instead of stderr

### [12.123.1](https://kaos.sh/ek/12.123.1)

- **`[support/network]`** Sort and deduplicate IPs

### [12.123.0](https://kaos.sh/ek/12.123.0)

- **`[csv]`** Added method `Reader.Line`
- **`[csv]`** Added more helpers for working with CSV row data

### [12.122.0](https://kaos.sh/ek/12.122.0)

- **`[csv]`** Added helpers for working with CSV row
- **`[csv]`** Added option to skip header
- **`[option]`** Removed `Required` flag from option struct

### [12.121.0](https://kaos.sh/ek/12.121.0)

- **`[initsystem/sdnotify]`** Added new package for sending messages to systemd
- **`[support/deps]`** Updated for compatibility with the latest version of [depsy](https://kaos.sh/depsy)
- **`[terminal/tty]`** Improved check for systemd

### [12.120.0](https://kaos.sh/ek/12.120.0)

- **`[knf]`** Added methods `Alias` and `Config.Alias`
- **`[knf]`** Added property name validation for all getters
- **`[sliceutil]`** Added method `IsEqual`
- **`[knf]`** Code refactoring
- **`[knf]`** Added more tests

### [12.119.0](https://kaos.sh/ek/12.119.0)

- **`[initsystem]`** Added [launchd](https://www.launchd.info) support
- **`[support/services]`** Added package for collecting services info

### [12.118.0](https://kaos.sh/ek/12.118.0)

- **`[terminal/input]`** Added method `SetHistoryCapacity`
- **`[options]`** Improved `Errors.String` output formatting
- **`[terminal/input]`** Added more usage examples
- **`[terminal/input]`** Code refactoring

### [12.117.0](https://kaos.sh/ek/12.117.0)

- **`[terminal/input]`** Added new package for working with user input
- **`[options]`** Added helper for working with errors
- **`[terminal]`** All methods for working with user input moved to separate package
- **`[terminal]`** Improved rendering messages with prefix
- **`[color]`** Tests refactoring
- **`[options]`** Code refactoring
- **`[terminal]`** Code refactoring

> [!IMPORTANT]
> - `[mathutil]` Removed all deprecated methods
> - `[passwd]` Removed all deprecated methods
> - `[sliceutil]` Removed all deprecated methods
> - `[terminal]` Removed all deprecated methods
> - `[usage]` Removed all deprecated methods
> - `[uuid]` Removed all deprecated methods
> - `[terminal/window]` Package removed due to deprecation of all methods

### [12.116.0](https://kaos.sh/ek/12.116.0)

- **`[support]`** Added Yandex Serverless support
- **`[system/container]`** Added Yandex Serverless support

### [12.115.0](https://kaos.sh/ek/12.115.0)

- **`[knf/united]`** Added method `GetMapping`

### [12.114.0](https://kaos.sh/ek/12.114.0)

- **`[knf/united]`** Added method `CombineSimple`

### [12.113.1](https://kaos.sh/ek/12.113.1)

- **`[support/pkgs]`** Added compatibility with macOS

### [12.113.0](https://kaos.sh/ek/12.113.0)

- **`[options]`** `Alias`, `Conflicts`, and `Bound` now supports string slices
- **`[options]`** Improved string representation format of `Map`, `V` and option name
- **`[terminal/tty]`** Windows stubs refactoring

### [12.112.1](https://kaos.sh/ek/12.112.1)

- **`[support]`** Added binary name info
- **`[terminal/tty]`** Added method `IsSystemd`
- **`[fsutil]`** Method `IsNonEmpty` marked as deprecated
- **`[fsutil]`** Fixed stubs for Windows
- **`[terminal/window]`** Fixed deprivation message

### [12.111.1](https://kaos.sh/ek/12.111.1)

- **`[knf/united]`** United configuration separated from global KNF configuration

### [12.111.0](https://kaos.sh/ek/12.111.0)

- **`[knf/united]`** Added helper method `AddOptions`
- **`[options]`** Added methods `V.String` and `Map.String`
- **`[support]`** Minor UI improvement
- **`[options]`** Code refactoring

### [12.110.1](https://kaos.sh/ek/12.110.1)

- **`[support/pkgs]`** Improved formatting
- **`[support/network]`** Local IPs removed from IPv4/IPv6 address lists

### [12.110.0](https://kaos.sh/ek/12.110.0)

- **`[support/pkgs]`** Added `tdnf` (_Photon Linux_) support
- **`[support/pkgs]`** Added `pacman` (_Arch Linux_) support
- **`[support/pkgs]`** Package name removed from version info
- **`[support]`** Fixed stubs for Windows

### [12.109.0](https://kaos.sh/ek/12.109.0)

- **`[support]`** Added custom checks support
- **`[knf/united]`** Improved handling of environment variables
- **`[support]`** Improved application info formatting
- **`[support]`** Code refactoring
- **`[support]`** Added more tests

### [12.108.1](https://kaos.sh/ek/12.108.1)

- **`[support]`** Fixed documentation formatting

### [12.108.0](https://kaos.sh/ek/12.108.0)

- **`[support]`** Added new package for collecting support information
- **`[options]`** Added method `Map.Set`
- **`[options]`** Added method `Map.Delete`
- **`[options]`** Added shortcut `F` for method `Format`
- **`[system/container]`** Added method `IsContainer`
- **`[system/container]`** Added engine info caching
- **`[pager]`** Fixed panic when pager stdin is not a file
- **`[usage]`** Fixed bug with changing color for certain command or option
- **`[lock]`** Fixed build tags
- **`[options]`** Code refactoring
- **`[usage]`** Code refactoring
- **`[options]`** Tests refactoring
- **`[system/container]`** Added usage examples

### [12.107.0](https://kaos.sh/ek/12.107.0)

- **`[knf/united]`** Added method `Simple`

### [12.106.0](https://kaos.sh/ek/12.106.0)

- **`[knf/united]`** Added method `ToOption` with shortcut `O`
- **`[knf/united]`** Added method `ToEnvVar` with shortcut `E`

### [12.105.0](https://kaos.sh/ek/12.105.0)

- **`[log]`** Added JSON output format
- **`[log]`** Added caller info to messages
- **`[strutil]`** Added method `IndexByteSkip`
- **`[timeutil]`** Added method `FromISOWeek`
- **`[log]`** Code refactoring

### [12.104.0](https://kaos.sh/ek/12.104.0)

- **`[knf/united]`** Added validation using `knf` validators
- **`[knf/united]`** Added usage examples

### [12.103.0](https://kaos.sh/ek/12.103.0)

- **`[knf/united]`** Added new package for working with united configuration (_knf + options + environment variables_)
- **`[knf]`** Added method `GetTD`
- **`[knf]`** Added method `GetTS`
- **`[knf]`** Added method `GetTZ`
- **`[knf]`** Added method `GetL`
- **`[options]`** Fixed panic when parsing unsupported option with value passed with equal sign (`=`)
- **`[options]`** Code refactoring
- **`[knf]`** Code refactoring

### [12.102.0](https://kaos.sh/ek/12.102.0)

- **`[knf/validators/network]`** Added `Mail` validator
- **`[log]`** Added `Divider` method
- **`[knf/validators/fs]`** Code refactoring
- **`[knf/validators/network]`** Code refactoring
- **`[knf/validators/regexp]`** Code refactoring
- **`[knf/validators/system]`** Code refactoring

### [12.101.0](https://kaos.sh/ek/12.101.0)

- **`[req]`** Added Bearer Token property to `Request` struct

### [12.100.0](https://kaos.sh/ek/12.100.0)

- **`[log]`** Added `NilLogger`
- **`[path]`** Handle both directions in `DirN`
- **`[strutil]`** Code refactoring

### [12.99.0](https://kaos.sh/ek/12.99.0)

- **`[knf]`** Add method `Parse`
- **`[knf]`** Add method `Config.Merge`
- **`[knf]`** Code refactoring
- **`[knf]`** Improved usage examples

### [12.98.0](https://kaos.sh/ek/12.98.0)

- **`[mathutil]`** Added method `Perc`
- **`[system]`** Added methods `MemUsage.MemUsedPerc` and `MemUsage.SwapUsedPerc`
- **`[fmtutil]`** Code refactoring
- **`[fmtutil/table]`** Improved tests
- **`[mathutil]`** Added usage examples

### [12.97.0](https://kaos.sh/ek/12.97.0)

- **`[passthru]`** Added package with pass-thru reader and writer
- **`[progress]`** Migrate to `passthru` package
- **`[fmtutil/table]`** Improved borders and separators rendering
- **`[usage]`** Improved environment info output
- **`[spinner]`** Improved message rendering

### [12.96.1](https://kaos.sh/ek/12.96.1)

- **`[terminal/tty]`** Fixed bug in checking for TTY when stdin is a character device
- **`[terminal/tty]`** Improved tests

### [12.96.0](https://kaos.sh/ek/12.96.0)

- **`[fmtutil/table]`** Added global flag `FullScreen` for full screen table mode
- **`[fmtutil/table]`** Full-screen table mode enabled by default

### [12.95.0](https://kaos.sh/ek/12.95.0)

- **`[fmtutil/table]`** Added table border symbol customization
- **`[fmtutil/table]`** Added table border symbol color customization
- **`[fmtutil/table]`** Added table separator symbol customization
- **`[fmtutil/table]`** Added table separator symbol color customization
- **`[fmtutil/table]`** Added table header color customization
- **`[fmtutil/table]`** Added table options for hiding top and bottom borders
- **`[fmtutil/table]`** Added data preprocessing feature using custom input processing function
- **`[fmtutil/table]`** Improved data rendering
- **`[fmtutil/table]`** Code refactoring

### [12.94.0](https://kaos.sh/ek/12.94.0)

- **`[options]`** Added method `Format`
- **`[options]`** Code refactoring

### [12.93.0](https://kaos.sh/ek/12.93.0)

- **`[fmtc]`** Added complex tags support for named colors
- **`[fmtc]`** Added named colors nesting

### [12.92.0](https://kaos.sh/ek/12.92.0)

- **`[color]`** Added method `Term2RGB`

### [12.91.0](https://kaos.sh/ek/12.91.0)

- **`[terminal/tty]`** Added method `IsTMUX`
- **`[terminal/tty]`** Added more usage examples
- **`[timeutil]`** `ParseDuration` now returns `time.Duration` instead of seconds

### [12.90.1](https://kaos.sh/ek/12.90.1)

- **`[pager]`** Use `more -f` by default
- **`[pager]`** Code refactoring

### [12.90.0](https://kaos.sh/ek/12.90.0)

- **`[color]`** Added alpha channel info to `HSL` and `HSV`
- **`[color]`** Added method `RGBA.WithAlpha`
- **`[color]`** Improved support for hex colors with alpha
- **`[color]`** Use web color representation for `Hex.String`
- **`[color]`** Added flag to `Hex.ToWeb` to disable shorthand generation
- **`[color]`** Fixed shorthand hex generation for `#FFF`
- **`[color]`** Fixed `RGBA` to `Hex` conversion

### [12.89.0](https://kaos.sh/ek/12.89.0)

- **`[usage]`** Added color customization for example description
- **`[usage]`** Changed default color for example description

### [12.88.1](https://kaos.sh/ek/12.88.1)

- **`[pager]`** Improved pager search
- **`[fmtutil]`** Code refactoring
- **`[fmtutil/table]`** Code refactoring
- **`[pager]`** Fixed build tags
- **`[terminal/tty]`** Fixed build tags

### [12.88.0](https://kaos.sh/ek/12.88.0)

- **`[pager]`** Added new package for pager (`less`/`more`) setup
- **`[terminal/tty]`** Added new package for working with TTY
- **`[fmtc]`** Added method `IsColorsSupported`

### [12.87.0](https://kaos.sh/ek/12.87.0)

- **`[fmtc]`** Added tag for italic text (`{&}`)
- **`[fmtc]`** Added tag for striked text (`{=}`)
- **`[fmtc]`** Added tag for hidden text (`{+}`)

### [12.86.0](https://kaos.sh/ek/12.86.0)

- **`[usage]`** `Info.AddCommand` now returns pointer to added command
- **`[usage]`** `Info.AddOption` now returns pointer to added option
- **`[usage]`** Added color customization for release and build info
- **`[usage]`** Added release separator customization
- **`[usage]`** Improved command group rendering
- **`[usage]`** Improved options rendering

### [12.85.0](https://kaos.sh/ek/12.85.0)

- **`[log]`** Added method `Is` and `Logger.Is`
- **`[log]`** Added more usage examples
- **`[terminal]`** Updated Windows stubs

### [12.84.0](https://kaos.sh/ek/12.84.0)

- **`[errutil]`** Added method `Errors.First`
- **`[errutil]`** Added method `Errors.Get`
- **`[fmtutil/table]`** Added short form of align flags
- **`[terminal]`** `Error`, `Warn` and `Info` now accept custom message objects
- **`[errutil]`** Added more usage examples

### [12.83.2](https://kaos.sh/ek/12.83.2)

- **`[fmtutil/panel]`** Improved panel rendering with disabled colors
- **`[fmtc]`** Fixed `IsTag` compatibility with sequence of tags (e.g. `{*}{_}{r}`)
- **`[fmtc]`** Fixed bug in `Clean` with writing reset escape sequence if there is no reset tag in the given string

### [12.83.1](https://kaos.sh/ek/12.83.1)

- **`[protip]`** Disabling tips using environment variable (`PROTIP=0`)

### [12.83.0](https://kaos.sh/ek/12.83.0)

- **`[protip]`** Added new package for showing usage tips
- **`[fmtc]`** Added method `IsTag` for color tag validation
- **`[fmtutil/panel]`** Added `TOP_LINE` option
- **`[fmtutil/panel]`** Added `DefaultOptions` variable to set default options for all panels
- **`[strutil]`** Added method `LenVisual`
- **`[log]`** Added color tag validation
- **`[progress]`** Added settings validation
- **`[spinner]`** Added color tag validation
- **`[usage]`** Added color tag validation
- **`[fmtutil/panel]`** Improved panel rendering when `BOTTOM_LINE` option is used
- **`[fmtutil/panel]`** Improved panel rendering when `label` is empty
- **`[fmtutil/panel]`** Added limit for minimal panel size (_256 symbols_)
- **`[fmtutil/panel]`** Added limit for maximum panel size (_256 symbols_)
- **`[fmtutil/panel]`** Added limit for maximum indent size (_24 symbols_)
- **`[fmtutil/panel]`** Fixed bug with panel rendering if `Indent` > 0
- **`[fmtutil/panel]`** Code refactoring
- **`[netutil]`** Fixed stubs for Windows

### [12.82.0](https://kaos.sh/ek/12.82.0)

- **`[netutil]`** Added method `GetAllIP`
- **`[netutil]`** Added method `GetAllIP6`

### [12.81.0](https://kaos.sh/ek/12.81.0)

- **`[knf/validators/network]`** Added `HasIP` validator
- **`[strutil]`** Added method `ReplaceIgnoreCase`
- **`[uuid]`** `GenUUID4` renamed to `UUID4`
- **`[uuid]`** `GenUUID5` renamed to `UUID5`
- **`[uuid]`** Code refactoring

### [12.80.0](https://kaos.sh/ek/12.80.0)

- **`[system]`** Added ANSI color info to `OSInfo`
- **`[system]`** Added methods `OSInfo.ColoredPrettyName` and `OSInfo.ColoredName`
- **`[strutil]`** Improved usage examples

### [12.79.0](https://kaos.sh/ek/12.79.0)

- **`[fmtutil/panel]`** Added indent customization

### [12.78.0](https://kaos.sh/ek/12.78.0)

- **`[barcode]`** New package with methods to generate colored representation of unique data

### [12.77.1](https://kaos.sh/ek/12.77.1)

- **`[options]`** Fixed bug with `Split` result for empty options

### [12.77.0](https://kaos.sh/ek/12.77.0)

- **`[options]`** Added merge symbol customization
- **`[options]`** Added method `Split` for splitting string value of mergeble option
- **`[options]`** Improve usage examples

### [12.76.1](https://kaos.sh/ek/12.76.1)

- **`[knf]`** Added dedicated type for duration modifiers

### [12.76.0](https://kaos.sh/ek/12.76.0)

- **`[knf]`** Added modificator support for `GetD`
- **`[spinner]`** Added initial spinner animation

### [12.75.1](https://kaos.sh/ek/12.75.1)

- **`[terminal]`** Improved `AlwaysYes` flag handling

### [12.75.0](https://kaos.sh/ek/12.75.0)

- **`[terminal]`** Added flag `AlwaysYes`, if set `ReadAnswer` will always return true without reading user input (_useful for working with option for forced actions_)

### [12.74.0](https://kaos.sh/ek/12.74.0)

- **`[timeutil]`** Added method `PrettyDurationSimple` for printing duration in simple format

### [12.73.2](https://kaos.sh/ek/12.73.2)

- **`[fmtutil]`** Fixed handling negative numbers in `PrettySize`
- **`[fsutil]`** Fixed handling empty paths in `ProperPath`

### [12.73.1](https://kaos.sh/ek/12.73.1)

- **`[fmtutil/panel]`** Panel rendering moved from `terminal` sub-package to it's own sub-package

### [12.72.0](https://kaos.sh/ek/12.72.0)

- **`[terminal]`** Added support of options for panels
- **`[mathutil]`** Sub-package migrated to generics
- **`[sliceutil]`** Sub-package migrated to generics
- **`[color]`** Code refactoring
- **`[spinner]`** Code refactoring

### [12.71.0](https://kaos.sh/ek/12.71.0)

- **`[terminal]`** Added panel size configuration feature
- **`[terminal]`** Improved panel rendering for messages with newlines
- **`[initsystem]`** Removed systemd statuses `activating` and `deactivating` from checking service state

### [12.70.0](https://kaos.sh/ek/12.70.0)

- **`[terminal]`** Added flag `HidePassword` for masking passwords while typing
- **`[terminal]`** Added method `Info` for showing informative messages
- **`[terminal]`** Added method `Panel` for showing panel with custom label, title, and message
- **`[terminal]`** Added method `ErrorPanel` for showing panel with error message
- **`[terminal]`** Added method `WarnPanel` for showing panel with warning message
- **`[terminal]`** Added method `InfoPanel` for showing panel with informative message
- **`[terminal]`** Method `PrintErrorMessage` marked as deprecated (_use method `Error` instead_)
- **`[terminal]`** Method `PrintWarnMessage` marked as deprecated (_use method `Warn` instead_)
- **`[fmtc]`** Code refactoring
- **`[initsystem]`** Code refactoring

### [12.69.0](https://kaos.sh/ek/12.69.0)

- **`[csv]`** Added method `WithComma` to CSV reader
- **`[spinner]`** Added symbols customization
- **`[spinner]`** Change default skip symbol to check mark
- **`[spinner]`** Change default skip symbol color to dark grey
- **`[spinner]`** Code refactoring

### [12.68.0](https://kaos.sh/ek/12.68.0)

- **`[options]`** Improved short options parsing logic
- **`[progress]`** Added window size configuration for passthru calculator
* Fixed typos

### [12.67.1](https://kaos.sh/ek/12.67.1)

- **`[options]`** Fixed bug with flattening empty arguments

### [12.67.0](https://kaos.sh/ek/12.67.0)

- **`[usage]`** Added support of raw version printing
- **`[path]`** Added method `JoinSecure` - more "secure" alternative to standard `Join`
- **`[usage]`** Code refactoring

### [12.66.0](https://kaos.sh/ek/12.66.0)

- **`[options]`** Added method `Arguments.Flatten` for converting arguments into a string
- **`[usage/update]`** Added update checker for GitLab

### [12.65.0](https://kaos.sh/ek/12.65.0)

- **`[fmtutil]`** Added method `Align` for aligning text with ANSI control sequences (_for example colors_)
- **`[usage]`** Added feature for adding and rendering environment info

### [12.64.1](https://kaos.sh/ek/12.64.1)

- **`[processes]`** `ProcessInfo.Childs` renamed to `ProcessInfo.Children`
* Fixed typos

### [12.64.0](https://kaos.sh/ek/12.64.0)

- **`[timeutil]`** Added method `MiniDuration` which returns formatted value for short durations (_e.g. s/ms/us/ns_)
- **`[terminal]`** Method `ReadUI` marked as deprecated (_use method `Read` instead_)

### [12.63.0](https://kaos.sh/ek/12.63.0)

- **`[knf]`** Fixed bug with using method `Is` for checking for empty values
- **`[terminal]`** Added prefix feature for error and warning messages

### [12.62.0](https://kaos.sh/ek/12.62.0)

- **`[system/containers]`** More precise method for checking container engine
- **`[system/containers]`** Added LXC support

### [12.61.0](https://kaos.sh/ek/12.61.0)

- **`[lscolors]`** Sub-package moved from `fmtc` to root of the package
- **`[lscolors]`** `GetColor` returns colors for types of objects (_like directory, block device, link…_)
- **`[lscolors]`** Added flag `DisableColors` for disabling all colors in output
- **`[usage]`** Methods `Render` marked as deprecated (_use `Print` methods instead_)
- **`[cron]`** Code refactoring
- **`[csv]`** Code refactoring
- **`[fmtutil/table]`** Code refactoring
- **`[knf]`** Code refactoring
- **`[options]`** Code refactoring
- **`[progress]`** Code refactoring
- **`[req]`** Code refactoring
- **`[spellcheck]`** Code refactoring
- **`[tmp]`** Code refactoring
- **`[usage]`** Code refactoring
- **`[cache]`** Better tests for panics
- **`[cron]`** Better tests for panics
- **`[csv]`** Better tests for panics
- **`[fmtutil/table]`** Better tests for panics
- **`[log]`** Better tests for panics
- **`[progress]`** Better tests for panics
- **`[req]`** Better tests for panics
- **`[spellcheck]`** Better tests for panics
- **`[tmp]`** Better tests for panics

### [12.60.1](https://kaos.sh/ek/12.60.1)

- **`[initsystem]`** Improved systemd support

### [12.60.0](https://kaos.sh/ek/12.60.0)

- **`[system/container]`** Added container sub-package with methods for checking container engine info
- **`[system]`** Added container engine info to `SystemInfo`
- **`[fmtutil/table]`** Improved separator rendering
* Code refactoring

### [12.59.0](https://kaos.sh/ek/12.59.0)

- **`[fmtutil/table]`** Improved separator rendering
* Code refactoring

### [12.58.0](https://kaos.sh/ek/12.58.0)

- **`[system]`** Added system arch name to `SystemInfo`
- **`[system]`** `Version` and `Distribution` info removed from `SystemInfo` (_use `OSInfo` instead_)
- **`[system]`** `GetOSInfo` now works on macOS

### [12.57.1](https://kaos.sh/ek/12.57.1)

- **`[progress]`** Fixed bug with updating progress settings

### [12.57.0](https://kaos.sh/ek/12.57.0)

- **`[log]`** Added interface for compatible loggers
- **`[httputil]`** Added more helpers for checking URL scheme
- **`[log]`** Color for critical errors set from magenta to bold red
- **`[tmp]`** Improved unique name generation
- **`[req]`** Code refactoring
- **`[tmp]`** Improved usage examples

### [12.56.0](https://kaos.sh/ek/12.56.0)

- **`[fmtc]`** Added named colors support
- **`[system]`** Fixed fuzz tests

### [12.55.2](https://kaos.sh/ek/12.55.2)

- **`[fmtc]`** Reverted changes made in 12.55.1

### [12.55.1](https://kaos.sh/ek/12.55.1)

- **`[fmtc]`** Fixed bug with printing useless carriage return symbol in `TPrint*` commands

### [12.55.0](https://kaos.sh/ek/12.55.0)

- **`[terminal]`** Added color customization for warning and error messages
- **`[fmtc]`** Method `NewLine` now can print more than one new line
- **`[progress]`** Fixed bug with handling finish stage
- **`[progress]`** Improved usage example
- **`[passwd]`** Code refactoring

### [12.54.0](https://kaos.sh/ek/12.54.0)

- **`[timeutil]`** Added more helpers for working with dates
- **`[options]`** Fixed panic in `GetS` when mixed or string option contains non-string value

### [12.53.0](https://kaos.sh/ek/12.53.0)

- **`[strutil]`** Added method `B` for choosing value by condition
- **`[system/process]`** Tests updated for compatibility with GitHub Actions

### [12.52.0](https://kaos.sh/ek/12.52.0)

- **`[fmtc]`** Added method `If` for conditional message printing

### [12.51.0](https://kaos.sh/ek/12.51.0)

- **`[lock]`** New package for working with lock files
- **`[fsutil]`** Better errors messages from `ValidatePerms`
- **`[pid]`** Code refactoring

### [12.50.1](https://kaos.sh/ek/12.50.1)

- **`[progress]`** Fixed bug with duplicating progress bar
- **`[progress]`** Fixed bug with duplicating percentage symbol
- **`[fmtc]`** Fixed simple mistakes in usage examples
- **`[fmtutil/table]`** Fixed strings formatting
- **`[req]`** Fixed strings formatting
- **`[spinner]`** Fixed strings formatting
- **`[usage]`** Fixed strings formatting

### [12.50.0](https://kaos.sh/ek/12.50.0)

- **`[fmtc]`** Added methods `TPrint`, `LPrint` and `TLPrint`
- **`[system]`** More fields support from `os-release` file
- **`[system]`** Added method `ParseOSInfo` for parsing `os-release` files
- **`[system]`** Improved tests
- **`[fmtc]`** Added more usage examples

### [12.49.0](https://kaos.sh/ek/12.49.0)

- **`[options]`** Added method `Is` for checking argument value
- **`[options]`** Fix bug with filtering unexpanded globs

### [12.48.0](https://kaos.sh/ek/12.48.0)

- **`[terminal]`** Added option HideLength for hiding password length
- **`[fmtutil/table]`** Fixed bug with calculating number of columns

### [12.47.0](https://kaos.sh/ek/12.47.0)

- **`[fsutil]`** Added bufio writer for copying files
- **`[knf]`** Added method `Is` for checking property value
- **`[options]`** Added method `Is` for checking option value
- **`[fsutil]`** Fixed bug with changing file mode for copied file if custom umask is used
- **`[options]`** Code refactoring
- **`[spinner]`** Code refactoring
- **`[knf]`** Added more usage examples

### [12.46.0](https://kaos.sh/ek/12.46.0)

- **`[fmtc/lscolors]`** Added method `GetColor` for getting control sequence for file
- **`[fmtc/lscolors]`** Added method `ColorizePath` for colorizing whole path
- **`[spinner]`** Added method `Skip` for marking actions as skipped
- **`[tmp]`** Use `os.TempDir` for default directory instead of hardcoded path
- **`[options]`** Fixed arguments filtering feature
- **`[system]`** Fixed stubs for Windows
- **`[terminal]`** Fixed stubs for Windows
- **`[log]`** Improved documentation

### [12.45.0](https://kaos.sh/ek/12.45.0)

- **`[knf]`** Added method `Config.File()` which returns path to configuration file
- **`[options]`** Added many helpful methods for working with arguments
- **`[spinner]`** Added `fmtc` color codes support
- **`[terminal]`** Title color customization support
- **`[fsutil]`** Added method `CopyAttr` for copying attributes from one object to another
- **`[fsutil]`** Code refactoring
- **`[options]`** Code refactoring
- **`[tmp]`** Code refactoring
- **`[options]`** Added more usage examples
- **`[fsutil]`** Improved tests

### [12.44.1](https://kaos.sh/ek/12.44.1)

- **`[ansi]`** Fixed examples
- **`[color]`** Tests refactoring
- **`[directio]`** Code refactoring
- **`[easing]`** Code refactoring
- **`[fmtc]`** Code refactoring
- **`[fsutil]`** Code refactoring
- **`[knf]`** Code refactoring
- **`[log]`** Code refactoring
- **`[log]`** Tests refactoring
- **`[options]`** Code refactoring
- **`[req]`** Code refactoring
- **`[signal]`** Use buffered channels for signals
- **`[system]`** Code refactoring
- **`[timeutil]`** Code refactoring
- **`[usage]`** Code refactoring

### [12.44.0](https://kaos.sh/ek/12.44.0)

- **`[path]`** Added method `Compact` for converting path to compact representation
- **`[fmtc]`** Added method `Render` for converting color tags to ANSI escape codes
- **`[knf/validators/fs]`** Fixed bug with formatting `FileMode` in error messages
- **`[cron]`** Improved parsing error messages
- **`[fsutil]`** Improved parsing error messages
- **`[cron]`** Improved tests
- **`[directio]`** Improved tests
- **`[fsutil]`** Improved tests
- **`[initsystem]`** Improved tests
- **`[jsonutil]`** Improved tests
- **`[knf]`** Improved tests
- **`[knf/validators/fs]`** Improved tests
- **`[knf/validators/network]`** Improved tests
- **`[knf/validators/regexp]`** Improved tests
- **`[knf/validators/system]`** Improved tests
- **`[knf/validators]`** Improved tests
- **`[log]`** Improved tests
- **`[options]`** Improved tests
- **`[req]`** Improved tests
- **`[system]`** Improved tests
- **`[system/process]`** Improved tests
- **`[system/sensors]`** Improved tests
- **`[tmp]`** Improved tests
- **`[version]`** Improved tests
- **`[spellcheck]`** Added usage examples
- **`[log]`** Fixed bug with rendering messages with colors when `fmtc.DisableColors` set to `true`
- **`[terminal]`** Fixed stubs for Windows
- **`[secstr]`** Fixed stubs for Windows

### [12.43.0](https://kaos.sh/ek/12.43.0)

- **`[terminal]`** Added method `ReadPasswordSecure` for reading password into a secure string
- **`[secstr]`** Fixed panic in `IsEmpty` and `Destroy` if string struct is nil

### [12.42.1](https://kaos.sh/ek/12.42.1)

* Module path set to `v12`

### [12.42.0](https://kaos.sh/ek/12.42.0)

- **`[fsutil]`** Improved helpers `CopyFile`, `MoveFile` and `CopyDir`
- **`[fsutil]`** Code refactoring
- **`[log]`** Code refactoring
* Removed `pkg.re` usage

### [12.41.0](https://kaos.sh/ek/12.41.0)

- **`[usage]`** Added helpers for generating a string representation of `Command` and `Option`
- **`[fsutil]`** Improved errors description for `ValidatePerms`

### [12.40.0](https://kaos.sh/ek/12.40.0)

- **`[ansi]`** Added methods for working with byte slices
- **`[secstr]`** Added method `IsEmpty` for checking empty strings
- **`[fmtutil]`** Method `Wrap` was rewritten from scratch
- **`[fmtutil]`** Method `Wrap` now ignores ANSI/VT100 control sequences
- **`[timeutil]`** Method `ParseDuration` was rewritten from scratch
- **`[timeutil]`** Method `ParseDuration` now allows to define default modificator
- **`[usage]`** Code refactoring

### [12.39.1](https://kaos.sh/ek/12.39.1)

- **`[knf]`** Fixed bug with naming tests source files

### [12.39.0](https://kaos.sh/ek/12.39.0)

- **`[fmtc]`** Added TrueColor support
- **`[system]`** Improved macOS support
- **`[netutil]`** Improved macOS support
- **`[system]`** Fixed bug with extracting macOS version
- **`[system]`** Fixed bug with extracting macOS arch info
- **`[fmtc]`** Code refactoring
- **`[system]`** Code refactoring
* Added more stubs for macOS
* Improved stubs for Windows

### [12.38.1](https://kaos.sh/ek/12.38.1)

* Fixed build tags for Go ≤ 1.16

### [12.38.0](https://kaos.sh/ek/12.38.0)

- **`[usage]`** Improved color customization for application name and version
- **`[usage]`** Added color customization for application name in usage info

### [12.37.0](https://kaos.sh/ek/12.37.0)

- **`[strutil]`** Added helper `Q` for working with default values
- **`[usage]`** Added color customization for application name and version

### [12.36.0](https://kaos.sh/ek/12.36.0)

- **`[system]`** Added CPU architecture bits info to `SystemInfo`

### [12.35.1](https://kaos.sh/ek/12.35.1)

- **`[passwd]`** Fixed typo in deprecation notice

### [12.35.0](https://kaos.sh/ek/12.35.0)

- **`[secstr]`** New package for working with protected (secure) strings
- **`[req]`** Method `Query.String()` renamed to `Query.Encode()`
- **`[passwd]`** Added method `GenPasswordVariations` and `GenPasswordBytesVariations` for generating password variations with possible typos fixes
- **`[passwd]`** Added methods `HashBytes`, `CheckBytes`, `GenPasswordBytes` and `GetPasswordBytesStrength`
- **`[passwd]`** Method `Encrypt` marked as deprecated (_use `Hash` method instead_)
- **`[passwd]`** Added more usage examples

### [12.34.0](https://kaos.sh/ek/12.34.0)

- **`[fsutil]`** Added method `TouchFile` for creating empty files
- **`[fsutil]`** Code refactoring
- **`[fmtc]`** Documentation refactoring
- **`[timeutil]`** Documentation refactoring
- **`[fsutil]`** Added usage examples

### [12.33.0](https://kaos.sh/ek/12.33.0)

- **`[errutil]`** Added support for `string`, `[]string`, and `errutil.Errors` types to method `Errors.Add`
- **`[fmtutil]`** Added method `PrettyBool` for formatting boolean values

### [12.32.0](https://kaos.sh/ek/12.32.0)

- **`[fmtutil]`** Added method `PrettyDiff` for formatting diff numbers
- **`[fmtutil]`** Fixed bug in `PrettyNum` with formatting negative numbers
- **`[fmtutil]`** Code refactoring

### [12.31.0](https://kaos.sh/ek/12.31.0)

- **`[errutil]`** Added method `Reset` for resetting `Errors` instance
- **`[sliceutil]`** Added methods `Copy`, `CopyInts` and `CopyFloats` for copying slices
- **`[csv]`** Code refactoring
- **`[sliceutil]`** Code refactoring

### [12.30.0](https://kaos.sh/ek/12.30.0)

- **`[ansi]`** New package for working with ANSI/VT100 control sequences
- **`[terminal]`** Added fmtc color codes support in input prompt
- **`[terminal]`** Fixed bug with masking password if prompt contains ANSI color codes
- **`[strutil]`** Code refactoring
- **`[options]`** Added more usage examples

### [12.29.0](https://kaos.sh/ek/12.29.0)

- **`[errutil]`** Added method `Cap` for getting max capacity
- **`[system/sensors]`** Added sorting by the name for slice with devices info
- **`[errutil]`** Added more usage examples

### [12.28.0](https://kaos.sh/ek/12.28.0)

- **`[fmtc]`** `NO_COLOR` support
- **`[fmtc]`** Code refactoring

### [12.27.0](https://kaos.sh/ek/12.27.0)

- **`[httputil]`** Added method `GetPortByScheme`
- **`[events]`** Improved unknown events handling
- **`[system/sensors]`** Code refactoring
- **`[system/sensors]`** Increased code coverage (0.0% → 100.0%)
- **`[events]`** Added usage examples
- **`[httputil]`** Added usage examples
* Error check moved at the beginning of every test

### [12.26.0](https://kaos.sh/ek/12.26.0)

- **`[events]`** New package for creating event-driven systems
- **`[system]`** Fixed bug with parsing CPU info data
- **`[fsutil]`** Added method `IsEmpty` for checking empty files
- **`[system/process]`** Fixed bug with searching info for creating process tree
- **`[knf]`** Code refactoring
- **`[usage]`** Added more usage examples

### [12.25.0](https://kaos.sh/ek/12.25.0)

- **`[color]`** Added method `Parse` for parsing colors (`#RGB`/`#RGBA`/`#RRGGBB`/`#RRGGBBAA`)
- **`[color]`** Fixed bug with formatting small values using `Hex.ToWeb`
- **`[color]`** Fixed bug with converting `Hex` to `RGBA`

### [12.24.2](https://kaos.sh/ek/12.24.2)

- **`[easing]`** Added links to examples for every function
- **`[easing]`** Added usage examples

### [12.24.1](https://kaos.sh/ek/12.24.1)

- **`[color]`** Added three-digit RGB notation generation to `Hex.ToWeb`

### [12.24.0](https://kaos.sh/ek/12.24.0)

- **`[color]`** Using structs for color models instead of bare numbers
- **`[color]`** Much simpler converting between color models
- **`[color]`** Added method `RGB2CMYK` for converting RGB colors to CMYK
- **`[color]`** Added method `CMYK2RGB` for converting CMYK colors to RGB
- **`[color]`** Added method `RGB2HSL` for converting RGB colors to HSL
- **`[color]`** Added method `HSL2RGB` for converting HSL colors to RGB
- **`[color]`** Added method `Luminance` for calculating relative luminance for RGB color
- **`[color]`** Added method `Contrast` for calculating contrast ratio of foreground and background colors
- **`[color]`** Method `RGB2HSB` rewritten from scratch and renamed to `RGB2HSV`
- **`[color]`** Method `HSB2RGB` rewritten from scratch and renamed to `HSV2RGB`
- **`[color]`** Added more usage examples


### [12.23.0](https://kaos.sh/ek/12.23.0)

- **`[cache]`** Renamed `Store` to `Cache`
- **`[cache]`** Added method `Size` for checking cache size
- **`[cache]`** Added method `Expired` for checking number of expired items

### [12.22.0](https://kaos.sh/ek/12.22.0)

- **`[timeutil]`** Added more new helpers
- **`[log]`** Code refactoring
- **`[log]`** Improved tests

### [12.21.0](https://kaos.sh/ek/12.21.0)

- **`[knf]`** Added new getter `GetD` which returns value as duration in seconds
- **`[system/process]`** Improved tests

### [12.20.3](https://kaos.sh/ek/12.20.3)

- `go-check` package replaced by [our fork](https://kaos.sh/check)
- **`[cron]`** Removed useless example
- **`[knf/validators/fs]`** Fixed bug with handling pattern matching error
- **`[path]`** Fixed usage examples

### [12.20.2](https://kaos.sh/ek/12.20.2)

- **`[terminal]`** Usage examples improvements

### [12.20.1](https://kaos.sh/ek/12.20.1)

- **`[system/process]`** Fixed stubs for Windows

### [12.20.0](https://kaos.sh/ek/12.20.0)

- **`[system/process]`** Added methods for setting and getting CPU scheduler priority
- **`[system/process]`** Added methods for setting and getting IO scheduler priority

### [12.19.0](https://kaos.sh/ek/12.19.0)

- **`[system]`** Added method `GetCPUCount` for getting info about number of CPU's
- **`[system/process]`** Added method `GetMountInfo` for getting info about process mounts
- **`[system]`** Code refactoring
- **`[system/process]`** Code refactoring
- **`[system]`** Increased code coverage (78.5% → 90.5%)
- **`[system/process]`** Increased code coverage (82.4% → 98.0%)

### [12.18.1](https://kaos.sh/ek/12.18.1)

- **`[sliceutil]`** Added usage examples

### [12.18.0](https://kaos.sh/ek/12.18.0)

- **`[strutil]`** Added methods `HasPrefixAny` and `HasSuffixAny` for checking multiple prefixes or suffixes at once

### [12.17.1](https://kaos.sh/ek/12.17.1)

- **`[path]`** Fixed stubs for Windows

### [12.17.0](https://kaos.sh/ek/12.17.0)

- **`[options]`** Method `Parse` now returns arguments as `Arguments` struct with additional methods for working with them
- **`[strutil]`** Added methods `Before` and `After` for extracting strings before and after some substring
- **`[progress]`** Fixed bug with rendering resulting progress bar in some situations
- **`[progress]`** Using integer instead of floats for progress if total size is less than 1000 and `IsSize` set to false

### [12.16.0 ](https://kaos.sh/ek/12.16.0 )

- **`[path]`** Added new method for checking Unix-type globs
- **`[fsutil]`** Fixed stubs for Windows
- **`[progress]`** Minor UI fix

### [12.15.1](https://kaos.sh/ek/12.15.1)

- **`[usage/completion/bash]`** Code refactoring

### [12.15.0](https://kaos.sh/ek/12.15.0)

- **`[spinner]`** Added new package for showing spinner animation for long-running tasks
- **`[timeutil]`** Added high precision mode for `ShortDuration`

### [12.14.1](https://kaos.sh/ek/12.14.1)

- **`[fmtc/lscolors]`** Improved environment variable parsing

### [12.14.0](https://kaos.sh/ek/12.14.0)

- **`[fmtc/lscolors]`** Added new package for colorizing file names with colors from dircolors

### [12.13.0](https://kaos.sh/ek/12.13.0)

- **`[usage/completion/bash]`** Improved completion generation
- **`[usage/completion/zsh]`** Improved completion generation
- **`[usage/completion/bash]`** Fixed bug with showing files with autocomplete
- **`[usage/completion/zsh]`** Fixed bug with showing files with autocomplete

### [12.12.0](https://kaos.sh/ek/12.12.0)

- **`[timeutil]`** Added method `PrettyDurationInDays` for rendering pretty duration in days
- **`[timeutil]`** Code refactoring

### [12.11.0](https://kaos.sh/ek/12.11.0)

- **`[timeutil]`** Added checking for parsing errors to `ParseDuration` method

### [12.10.1](https://kaos.sh/ek/12.10.1)

- **`[req]`** Code refactoring

### [12.10.0](https://kaos.sh/ek/12.10.0)

- **`[usage/man]`** Added package for generating man pages from usage info

### [12.9.0](https://kaos.sh/ek/12.9.0)

- **`[usage/update]`** Disabled update check from CI environments

### [12.8.1](https://kaos.sh/ek/12.8.1)

- **`[knf/validators/system]`** Fixed bug with source file naming

### [12.8.0](https://kaos.sh/ek/12.8.0)

- **`[log]`** Removed useless return value from `Aux`, `Debug`, `Info`, `Warn`, `Error`, `Crit` and `Print` methods

### [12.7.0](https://kaos.sh/ek/12.7.0)

- **`[knf/validators/regexp]`** Added new KNF validator for checking regular expression pattern matching
- **`[knf/validators/fs]`** Added new KNF validator for checking shell pattern matching
- **`[fsutil]`** Fixed bug with checking empty dirs on osx
- **`[initsystem]`** Disabled tests on osx
- **`[knf/validators/fs]`** Fixed tests on osx
- **`[knf/validators/system]`** Fixed compatibility with osx
- **`[log]`** Fixed tests on OSX
- **`[system]`** Fixed checking user or group existence on OSX
- **`[system]`** Fixed group lookup on osx
- **`[system]`** Improved user info fetching on OSX

### [12.6.1](https://kaos.sh/ek/12.6.1)

- **`[path]`** Code refactoring
- **`[path]`** Added more usage examples
- **`[timeutil]`** Added more usage examples

### [12.6.0](https://kaos.sh/ek/12.6.0)

- **`[timeutil]`** Added method `SecondsToDuration` for conversion `float64` to `time.Duration`
- **`[timeutil]`** `DurationToSeconds` now returns the result as a float64 number
- **`[hash]`** Code refactoring

### [12.5.2](https://kaos.sh/ek/12.5.2)

- **`[system]`** Fixed compatibility with Go ≥ 1.15

### [12.5.1](https://kaos.sh/ek/12.5.1)

* License changed from EKOL to Apache-2.0

### [12.5.0](https://kaos.sh/ek/12.5.0)

- **`[req]`** Added method `Bytes()` for reading response body as byte slice
- **`[env]`** Fixed tests

### [12.4.0](https://kaos.sh/ek/12.4.0)

- **`[timeutil]`** Added method `ShortDuration` for duration formatting
- **`[timeutil]`** Code refactoring

### [12.3.0](https://kaos.sh/ek/12.3.0)

- **`[progress]`** Calculate speed and remaining time using Exponentially Weighted Moving Average (EWMA)
- **`[progress]`** Added pass thru writer

### [12.2.0](https://kaos.sh/ek/12.2.0)

- **`[progress]`** Added package for creating terminal progress bar
- **`[usage/update]`** Increased dial and request timeouts to 3 seconds
- **`[fmtutil]`** Added possibility to define custom separators in `PrettySize` and `PrettyNum` methods
- **`[passwd]`** Increased code coverage (94.9% → 96.8%)
- **`[usage/update]`** Increased code coverage (92.1% → 100%)
- **`[req]`** Tests refactoring

### [12.1.0](https://kaos.sh/ek/12.1.0)

- **`[usage/update]`** Added update checker for custom storages

### [12.0.0](https://kaos.sh/ek/12.0.0)

- **`[path]`** Added method `DirN` for reading N elements from path
- **`[pluralize]`** Methods `Pluralize` and `PluralizeSpecial` now return only pluralized word (_incompatible changes_)
- **`[pluralize]`** Added methods `P` and `PS` for pluralization with custom formatting
- **`[pluralize]`** Added usage examples

---

### [11.6.3](https://kaos.sh/ek/11.6.3)

- **`[usage]`** Added more examples

### [11.6.2](https://kaos.sh/ek/11.6.2)

- **`[cron]`** Added usage examples

### [11.6.1](https://kaos.sh/ek/11.6.1)

- **`[system]`** Fixed bug with parsing group info in `id` output

### [11.6.0](https://kaos.sh/ek/11.6.0)

- **`[usage]`** Added support of raw examples (_without prefix with command name_)

### [11.5.2](https://kaos.sh/ek/11.5.2)

- **`[errutil]`** Fixed panic in `Add` if given Errors struct is nil

### [11.5.1](https://kaos.sh/ek/11.5.1)

- **`[color]`** Fixed compatibility with ARM
- **`[fmtutil]`** Fixed compatibility with ARM
- **`[system]`** Fixed compatibility with ARM

### [11.5.0](https://kaos.sh/ek/11.5.0)

- **`[signal]`** Added method `GetByName` for getting signal by its name
- **`[signal]`** Added method `GetByCode` for getting signal by its code

### [11.4.0](https://kaos.sh/ek/11.4.0)

- **`[fsutil]`** Added method `ValidatePerms` for permissions validation
- **`[system]`** Improved current user info caching mechanic
- **`[fsutil]`** Increased code coverage (98.0% → 98.8%)

### [11.3.1](https://kaos.sh/ek/11.3.1)

- **`[initsystem]`** Fixed stubs for Windows

### [11.3.0](https://kaos.sh/ek/11.3.0)

- **`[log]`** Logger is now more concurrency friendly

### [11.2.2](https://kaos.sh/ek/11.2.2)

- **`[log]`** Default color for debug messages set to light gray

### [11.2.1](https://kaos.sh/ek/11.2.1)

- **`[cache]`** Added data removal from cache with disabled janitor

### [11.2.0](https://kaos.sh/ek/11.2.0)

- **`[cache]`** Added method `Has` for checking item existence
- **`[cache]`** Janitor thread will not run if the cleaning interval is equal to 0

### [11.1.0](https://kaos.sh/ek/11.1.0)

- **`[pid]`** Added method `Read` for reading PID files without any configuration

### [11.0.1](https://kaos.sh/ek/11.0.1)

- **`[knf]`** Minor documentation fixes

### [11.0.0](https://kaos.sh/ek/11.0.0)

- **`[fsutil]`** `GetPerms` renamed to `GetMode`
- **`[fsutil]`** Added support of checking for character and block devices (`C` and `B`)
- **`[knf]`** Validators moved to sub-package (_incompatible changes_)
- **`[knf]`** Added more validators
- **`[knf]`** Removed useless dependencies
- **`[fsutil]`** Increased code coverage (97.4% → 98.0%)
- **`[kv]`** Package removed

---

### [10.18.1](https://kaos.sh/ek/10.18.1)

- **`[strutil]`** Fixed bug in `Substr` method for a situation when the index of start symbol is greater than the length of the string
- **`[strutil]`** Fixed bug in `Substring` method for a situation when the index of start symbol is greater than the length of the string

### [10.18.0](https://kaos.sh/ek/10.18.0)

- **`[knf]`** Added `no` as a valid boolean value for `GetB`
- **`[knf]`** Added new validators for property type validation
- **`[knf]`** Code refactoring

### [10.17.0](https://kaos.sh/ek/10.17.0)

- **`[cache]`** Added package which provides simple in-memory key:value store

### [10.16.0](https://kaos.sh/ek/10.16.0)

- **`[timeutil]`** Added support of short durations (_milliseconds, microseconds or nanoseconds_) to `PrettyDuration` method

### [10.15.0](https://kaos.sh/ek/10.15.0)

- **`[log]`** Added support of ANSI colors in the output
- **`[log]`** Using `uint8` for level codes instead of `int`

### [10.14.0](https://kaos.sh/ek/10.14.0)

- **`[version]`** Added method `IsZero` for checking empty version struct

### [10.13.1](https://kaos.sh/ek/10.13.1)

- **`[errutil]`** Method `Add` now allows adding slices with errors

### [10.13.0](https://kaos.sh/ek/10.13.0)

- **`[errutil]`** Added possibility to limit the number of errors to store
- **`[errutil]`** Method `Add` now allows adding errors from other Errors struct
- **`[sliceutil]`** Using in-place deduplication in `Deduplicate` method

### [10.12.2](https://kaos.sh/ek/10.12.2)

- **`[req]`** Changed default user-agent to `go-ek-req/10`

### [10.12.1](https://kaos.sh/ek/10.12.1)

- **`[usage]`** Fixed bug with formatting options without short name

### [10.12.0](https://kaos.sh/ek/10.12.0)

- **`[req]`** Added method `PostFile` for multipart file uploading

### [10.11.1](https://kaos.sh/ek/10.11.1)

- **`[fsutil]`** Fixed bug with filtering listing data

### [10.11.0](https://kaos.sh/ek/10.11.0)

- **`[pid]`** Added method `IsProcessWorks` for checking process state by PID
- **`[pid]`** Improved process state check
- **`[pid]`** Improved Mac OS X support

### [10.10.2](https://kaos.sh/ek/10.10.2)

- **`[terminal]`** Reading user input now is more stdin friendly (_you can pass the input through the stdin_)

### [10.10.1](https://kaos.sh/ek/10.10.1)

- **`[usage]`** Fixed bug with formatting options
- **`[fmtutil/table]`** More copy&paste friendly table rendering

### [10.10.0](https://kaos.sh/ek/10.10.0)

- **`[emoji]`** New package for working with emojis

### [10.9.1](https://kaos.sh/ek/10.9.1)

- **`[usage/completion/bash]`** Improved bash completion generation

### [10.9.0](https://kaos.sh/ek/10.9.0)

- **`[usage/completion/bash]`** Added bash completion generator
- **`[usage/completion/zsh]`** Added zsh completion generator
- **`[usage/completion/fish]`** Added fish completion generator
- **`[usage]`** Added method `info.BoundOptions` for linking command with options
- **`[csv]`** Added method `reader.ReadTo` for reading CSV data into slice
- **`[strutil]`** Fixed bug in `Fields` method
- **`[initsystem]`** Added caching for initsystem usage status
- **`[initsystem]`** Improved service state search for SysV scripts on systems with Systemd
- **`[usage]`** Code refactoring

### [10.8.0](https://kaos.sh/ek/10.8.0)

- **`[strutil]`** Added method `Exclude` as the faster replacement for `strings.ReplaceAll`

### [10.7.1](https://kaos.sh/ek/10.7.1)

- **`[fmtutil]`** Fixed bug with formatting small float numbers using `PrettySize` method

### [10.7.0](https://kaos.sh/ek/10.7.0)

- **`[jsonutil]`** Added `Write` as alias for `EncodeToFile`
- **`[jsonutil]`** Added `Read` as alias for `DecodeFile`
- **`[jsonutil]`** Added `WriteGz` for writing gzipped JSON data
- **`[jsonutil]`** Added `ReadGz` for reading gzipped JSON data

### [10.6.0](https://kaos.sh/ek/10.6.0)

- **`[strutil]`** Improved parsing logic for the `Fields` method
- **`[strutil]`** Added additional supported quotation marks types

### [10.5.1](https://kaos.sh/ek/10.5.1)

- **`[initsystem]`** Fixed bug with parsing systemd's `failed` ActiveState status
- **`[initsystem]`** Added tests for output parsers
- **`[initsystem]`** Code refactoring

### [10.5.0](https://kaos.sh/ek/10.5.0)

- **`[fmtc]`** Added new methods `LPrintf`, `LPrintln`, `TLPrintf` and `TLPrintln`
- **`[fmtc]`** Fixed bug with parsing reset and modification tags (_found by go-fuzz_)
- **`[fmtc]`** Code refactoring

### [10.4.0](https://kaos.sh/ek/10.4.0)

- **`[fmtc]`** Improved work with temporary output (`TPrintf`, `TPrintln`)

### [10.3.0](https://kaos.sh/ek/10.3.0)

- **`[fsutil]`** Added method `IsReadableByUser` for checking read permission for some user
- **`[fsutil]`** Added method `IsWritableByUser` for checking write permission for some user
- **`[fsutil]`** Added method `IsExecutableByUser` for checking execution permission for some user

### [10.2.0](https://kaos.sh/ek/10.2.0)

- **`[version]`** Added method `Simple()` which returns simple version
- **`[version]`** More usage examples added

### [10.1.0](https://kaos.sh/ek/10.1.0)

- **`[system]`** Improved OS version search
- **`[tmp]`** Package refactoring

### [10.0.0](https://kaos.sh/ek/10.0.0)
- **`[system]`** Added method `GetCPUInfo` for fetching info about CPUs from procfs
- **`[fmtutil/table]`** Added global variable `MaxWidth` for configuration of maximum table width
- **`[system]`** `FSInfo` now is `FSUsage` (_incompatible changes_)
- **`[system]`** `MemInfo` now is `MemUsage` (_incompatible changes_)
- **`[system]`** `CPUInfo` now is `CPUUsage` (_incompatible changes_)
- **`[system]`** `InterfaceInfo` now is `InterfaceStats` (_incompatible changes_)
- **`[system]`** `GetFSInfo()` now is `GetFSUsage()` (_incompatible changes_)
- **`[system]`** `GetMemInfo()` now is `GetMemUsage()` (_incompatible changes_)
- **`[system]`** `GetCPUInfo()` now is `GetCPUUsage()` (_incompatible changes_)
- **`[system]`** `GetInterfacesInfo()` now is `GetInterfacesStats()` (_incompatible changes_)
- **`[initsystem]`** `HasService()` now is `IsPresent()` (_incompatible changes_)
- **`[initsystem]`** `IsServiceWorks()` now is `IsWorks()` (_incompatible changes_)
- **`[system]`** Fixed bug with parsing CPU stats data (_found by go-fuzz_)
- **`[fmtc]`** Fixed bug with parsing reset and modification tags (_found by go-fuzz_)
- **`[initsystem]`** Fixed examples
- **`[fmtc]`** Fixed examples
- **`[system]`** Added fuzz testing
- **`[cron]`** Code refactoring
- **`[timeutil]`** Code refactoring
- **`[fmtutil]`** Increased code coverage (97.9% → 100.0%)
- **`[fmtutil/table]`** Increased code coverage (99.4% → 100.0%)
- **`[knf]`** Increased code coverage (99.6% → 100.0%)
- **`[req]`** Increased code coverage (97.1% → 100.0%)
- **`[pid]`** Increased code coverage (97.4% → 100.0%)
- **`[system]`** Increased code coverage (73.8% → 79.0%)

---

### [9.28.1](https://kaos.sh/ek/9.28.1)

- **`[initsystem]`** Improved application state checking in systemd
- **`[system]`** Fixed typo in json tag for `CPUInfo.Average`

### [9.28.0](https://kaos.sh/ek/9.28.0)

- **`[system]`** Improved memory usage calculation
- **`[system]`** Added `Shmem` and `SReclaimable` values to `MemInfo` struct
- **`[system]`** Fixed typo in json tag for `MemInfo.SwapCached`
- **`[system]`** Improved tests

### [9.27.0](https://kaos.sh/ek/9.27.0)

- **`[system/sensors]`** Added package for collecting sensors data
- **`[strutil]`** Added method `Substring` for safe substring extraction
- **`[strutil]`** Added method `Extract` for safe substring extraction
- **`[strutil]`** Fixed tests and example for `Substr` method
- **`[strutil]`** Improved tests
- **`[strutil]`** Code refactoring

### [9.26.3](https://kaos.sh/ek/9.26.3)

- **`[strutil]`** Optimization and improvements for `ReadField` method
- **`[easing]`** Code refactoring
- **`[fmtutil]`** Code refactoring
- **`[knf]`** Code refactoring
- **`[log]`** Code refactoring
- **`[options]`** Code refactoring
- **`[pid]`** Code refactoring
- **`[req]`** Code refactoring
- **`[sliceutil]`** Code refactoring
- **`[strutil]`** Code refactoring
- **`[system]`** Code refactoring
- **`[terminal]`** Code refactoring
- **`[timeutil]`** Code refactoring
- **`[uuid]`** Code refactoring

### [9.26.2](https://kaos.sh/ek/9.26.2)

- **`[fmtc]`** Fixed bug with parsing `{}` and `{-}` as valid color tags
- **`[fmtc]`** Added fuzz testing

### [9.26.1](https://kaos.sh/ek/9.26.1)

- **`[fmtutil/table]`** Fixed bug with rendering data using not-configured table

### [9.26.0](https://kaos.sh/ek/9.26.0)

- **`[sliceutil]`** Added method `Index` which return index of given item in slice

### [9.25.2](https://kaos.sh/ek/9.25.2)

- **`[fmtutil]`** Improved size parser

### [9.25.1](https://kaos.sh/ek/9.25.1)

- **`[fmtutil]`** Fixed various bugs with processing NaN values

### [9.25.0](https://kaos.sh/ek/9.25.0)

- **`[req]`** Added constants with status codes

### [9.24.0](https://kaos.sh/ek/9.24.0)
- **`[req]`** Added method `String` for `Query` struct for query encoding

### [9.23.0](https://kaos.sh/ek/9.23.0)

- **`[log]`** Added wrapper for compatibility with stdlib logger
- **`[log]`** Fixed race condition issue

### [9.22.3](https://kaos.sh/ek/9.22.3)

- **`[usage]`** Fixed bug with aligning option info with Unicode symbols
- **`[options]`** Guess option type by default value type
- **`[options]`** Added check for unsupported default value type

### [9.22.2](https://kaos.sh/ek/9.22.2)

- **`[system/process]`** Fixed windows stubs

### [9.22.1](https://kaos.sh/ek/9.22.1)

- **`[fsutil]`** Improved `CopyDir` method

### [9.22.0](https://kaos.sh/ek/9.22.0)

- **`[fsutil]`** Added method `CopyDir` for recursive directories copying
- **`[fsutil]`** Removed useless method `Current`
- **`[fsutil]`** Tests refactoring
- **`[fsutil]`** Code refactoring

### [9.21.0](https://kaos.sh/ek/9.21.0)

- **`[system/process]`** Added new type `ProcSample` as a lightweight analog of ProcInfo for CPU usage calculation
- **`[system/process]`** Code refactoring
- **`[system/process]`** Increased code coverage (75.5% → 82.4%)
- **`[system]`** Code refactoring

### [9.20.1](https://kaos.sh/ek/9.20.1)

- **`[fmtutil]`** Added method `PrettyPerc` for formatting values in percentages

### [9.20.0](https://kaos.sh/ek/9.20.0)

- **`[fmtc]`** Added methods `Print` and `Sprintln` for better compatibility with `fmt` package
- **`[fmtutil/table]`** Fixed minor bug with output formatting
- **`[options]`** Code refactoring

### [9.19.0](https://kaos.sh/ek/9.19.0)

- **`[directio]`** Added sub-package `directio` for writing/reading data with using direct IO
- **`[fmtc]`** 256 colors support with new tags (foreground: `{#000}`, background: `{%000}`)
- **`[fmtc]`** Added method `Is256ColorsSupported` for checking support of 256 color output
- **`[fmtc]`** Improved color tags syntax
- **`[fmtc]`** Added tags for resetting modificators (e.g. `{!*}`)
- **`[fmtc]`** Removed color tags overriding (i.e. now `{*}{r} == {r*}`)
- **`[color]`** Added method `RGB2Term` for converting RGB colors to terminal color codes

### [9.18.1](https://kaos.sh/ek/9.18.1)

- **`[system]`** Fixed bug with extra new line symbol in user `Shell` field

### [9.18.0](https://kaos.sh/ek/9.18.0)

- **`[fmtc]`** Temporary output feature moved from T struct to `TPrintf` and `TPrintln`

### [9.17.4](https://kaos.sh/ek/9.17.4)

* Dependencies now download with initial `go get` command

### [9.17.3](https://kaos.sh/ek/9.17.3)

- **`[options]`** Fixed bug with using `Bound` or `Conflict` fields for options (thx to @gongled)
- **`[netutil]`** Code refactoring
- **`[netutil]`** Increased code coverage (78.8% → 87.9%)

### [9.17.2](https://kaos.sh/ek/9.17.2)

- **`[netutil]`** Improved main IP search

### [9.17.1](https://kaos.sh/ek/9.17.1)

- **`[strutil]`** Added usage example for `Copy` method
- **`[system/procname]`** Added usage examples

### [9.17.0](https://kaos.sh/ek/9.17.0)

- **`[netutil]`** Ignore TUN/TAP interfaces while searching main IP address
- **`[initsystem]`** Added method `IsEnabled` which return info about service autostart
- **`[initsystem]`** Method `GetServiceState` renamed to `IsServiceWorks`
- **`[strutil]`** Added method `Copy` for forced copying of strings

### [9.16.0](https://kaos.sh/ek/9.16.0)

- **`[strutil]`** Improved `Fields` parsing
- **`[fmtutil/table]`** Added method `RenderHeaders` for forced headers rendering

### [9.15.0](https://kaos.sh/ek/9.15.0)

- **`[strutil]`** Added ellipsis suffix customization
- **`[strutil]`** Added support of custom separators for `ReadField`
- **`[req]`** Closing response body after parsing data
- **`[system]`** Fixed bug with parsing `id` command output with empty group names
- **`[system]`** Fixed bug with calculating transferred bytes on active interfaces
- **`[system]`** Improved `id` and `getent` commands output parsing
- **`[system]`** Code refactoring

### [9.14.5](https://kaos.sh/ek/9.14.5)

- **`[terminal]`** Fixed bug with empty title output

### [9.14.4](https://kaos.sh/ek/9.14.4)

- **`[system]`** Code refactoring

### [9.14.3](https://kaos.sh/ek/9.14.3)

- **`[initsystem]`** Fixed bug with checking service state in systemd

### [9.14.2](https://kaos.sh/ek/9.14.2)

- **`[system]`** Fixed windows stubs
- **`[system]`** Fixed bug with unclosed file descriptor

### [9.14.1](https://kaos.sh/ek/9.14.1)

- **`[initsystem]`** Fixed bug in SysV service state determination

### [9.14.0](https://kaos.sh/ek/9.14.0)

- **`[strutil]`** Added new method `ReadField` for reading space/tab separated fields from given data
- **`[system]`** Code refactoring

### [9.13.0](https://kaos.sh/ek/9.13.0)

- **`[system]`** Improved CPU usage calculation
- **`[system/process]`** Code refactoring
- **`[system]`** Code refactoring

### [9.12.0](https://kaos.sh/ek/9.12.0)

- **`[knf]`** Added new validators: `NotLen`, `NotPrefix` and `NotSuffix`
- **`[knf]`** Validators code refactoring

### [9.11.2](https://kaos.sh/ek/9.11.2)

- **`[system/process]`** Fixed bug with parsing CPU data
- **`[system/process]`** Increased code coverage (0.0% → 87.5%)
- **`[usage/update]`** Increased code coverage (0.0% → 80.0%)

### [9.11.1](https://kaos.sh/ek/9.11.1)

- **`[system/process]`** Improved error handling in `GetInfo`

### [9.11.0](https://kaos.sh/ek/9.11.0)

- **`[system]`** Improved IO utilization calculation
- **`[system]`** Improved network speed calculation

### [9.10.0](https://kaos.sh/ek/9.10.0)

- **`[system]`** Added method `GetCPUStats` which return basic CPU info from `/proc/stat`
- **`[system]`** Improved IO utilization calculation

### [9.9.2](https://kaos.sh/ek/9.9.2)

- **`[initsystem]`** Added stubs for windows

### [9.9.1](https://kaos.sh/ek/9.9.1)

* Code refactoring

### [9.9.0](https://kaos.sh/ek/9.9.0)

- **`[system]`** Improved disk usage calculation (now it similar to `df` command output)

### [9.8.0](https://kaos.sh/ek/9.8.0)

- **`[initsystem]`** New package for working with different init systems (sysv, upstart, systemd)

### [9.7.1](https://kaos.sh/ek/9.7.1)

- **`[fmtc]`** Improved utf8 support in temporary messages

### [9.7.0](https://kaos.sh/ek/9.7.0)

- **`[fmtc]`** Added method `NewT` which creates a new struct for working with the temporary output
- **`[fmtc]`** More docs about color tags
- **`[knf]`** Removing trailing spaces from property values

### [9.6.0](https://kaos.sh/ek/9.6.0)

- **`[system/procname]`** Added method `Replace` which replace just one argument in process command

### [9.5.0](https://kaos.sh/ek/9.5.0)

- **`[knf]`** Added new getters `GetU`, `GetU64` and `GetI64`
- **`[usage]`** Improved API for `NewInfo` method

### [9.4.0](https://kaos.sh/ek/9.4.0)

- **`[options]`** Added support of mixed options (string / bool)

### [9.3.0](https://kaos.sh/ek/9.3.0)

- **`[terminal]`** Improved title rendering for `ReadAnswer` method
- **`[terminal]`** Simplified API for `ReadAnswer` method

### [9.2.0](https://kaos.sh/ek/9.2.0)

- **`[fmtutil]`** Improved floating numbers formatting with `PrettyNum`

### [9.1.4](https://kaos.sh/ek/9.1.4)

- **`[fmtutil/table]`** Fixed bug with color tags in headers when colors is disabled

### [9.1.3](https://kaos.sh/ek/9.1.3)

- **`[timeutil]`** Fixed bug with formatting milliseconds
- **`[timeutil]`** Improved tests

### [9.1.2](https://kaos.sh/ek/9.1.2)

- **`[terminal]`** Fixed bug with masking password in tmux

### [9.1.1](https://kaos.sh/ek/9.1.1)

- **`[fmtutil/table]`** Fixed bug with rendering data with color tags

### [9.1.0](https://kaos.sh/ek/9.1.0)

- **`[version]`** Fixed bug with version comparison
- **`[version]`** Added method `Int()` which return version as integer

### [9.0.0](https://kaos.sh/ek/9.0.0)

* Package `args` renamed to `options` (_incompatible changes_)
- **`[fmtutil/table]`** Added new package for rendering data as a table
- **`[fmtutil]`** Added support of separator symbol configuration
- **`[usage]`** Improved output about a newer version
- **`[usage]`** Increased code coverage (0.0% → 100%)
- **`[usage]`** Code refactoring

---

### [8.0.3](https://kaos.sh/ek/8.0.3)

- **`[usage]`** Improved options and commands info rendering

### [8.0.2](https://kaos.sh/ek/8.0.2)

* Overall documentation improvements

### [8.0.1](https://kaos.sh/ek/8.0.1)

- **`[system/process]`** Fixed windows stubs
- **`[system]`** Package refactoring
- **`[fsutil]`** Fixed checking empty directory on FreeBSD
- **`[pid]`** Fixed checking process state on FreeBSD

### [8.0.0](https://kaos.sh/ek/8.0.0)

- **`[system/process]`** Added method `GetMemInfo` for obtaining information about memory consumption by process.
- **`[system/process]`** Added method `GetInfo` which return partial info from `/proc/[PID]/stat`.
- **`[system/process]`** Added method `CalculateCPUUsage` which can be used for process CPU usage calculation.
- **`[system]`** Methods for executing commands moved to `system/exec` package (_incompatible changes_)
- **`[system]`** Methods for changing process name moved to `system/procname` package (_incompatible changes_)
- **`[system]`** Minor improvements
- **`[system]`** Code refactoring
- **`[system]`** Increased code coverage (0.0% → 79.5%)

---

### [7.5.0](https://kaos.sh/ek/7.5.0)

- **`[errutil]`** Implemented error interface (_added method_ `Error() string`)
- **`[errutil]`** Minor improvements
- **`[system]`** Fixed windows stubs

### [7.4.0](https://kaos.sh/ek/7.4.0)

- **`[fmtutil]`** Added flag `SeparatorFullscreen` which enable full size separator
- **`[terminal/window]`** Window size detection code moved from `terminal` to `terminal/window` package
- **`[terminal/window]`** Fixed bug with unclosed TTY file descriptor
- **`[fsutil]`** Fixed bug with `fsutil.IsLink` (_method returns true for symlinks_)
- **`[fsutil]`** Fixed bug with `fsutil.GetSize` (_method returns 0 for non-existent files_)
- **`[fsutil]`** Improved input arguments checks in `fsutil.CopyFile`
- **`[fsutil]`** Added input arguments checks to `fsutil.MoveFile`
- **`[fsutil]`** Increased code coverage (49.8% → 97.9%)
- **`[knf]`** Increased code coverage (99.2% → 99.6%)
- **`[jsonutil]`** Increased code coverage (92.3% → 100%)

### [7.3.0](https://kaos.sh/ek/7.3.0)

- **`[sortutil]`** Added methods `NatualLess` and `StringsNatual` for natural ordering
- **`[jsonutil]`** Added optional argument to `EncodeToFile` method with file permissions (0644 by default)
- **`[jsonutil]`** Code refactoring
- **`[jsonutil]`** Improved tests
- **`[jsonutil]`** Added usage examples

### [7.2.0](https://kaos.sh/ek/7.2.0)

- **`[knf]`** Return default value for the property even if config struct is nil

### [7.1.0](https://kaos.sh/ek/7.1.0)

- **`[system]`** Added methods `CalculateNetworkSpeed` and `CalculateIOUtil` for metrics calculation without blocking main thread
- **`[system]`** Code and examples refactoring

### [7.0.3](https://kaos.sh/ek/7.0.3)

- **`[passwd]`** Fixed panic in `Check` for some rare cases
- **`[fsutil]`** Fixed typo
- **`[pid]`** Fixed typo
- **`[system]`** Fixed typo
- **`[tmp]`** Fixed typo
- **`[knf]`** Increased code coverage

### [7.0.2](https://kaos.sh/ek/7.0.2)

- **`[version]`** Fixed bug with version comparison
- **`[version]`** Improved version data storing model
- **`[usage]`** Fixed bug with new application version checking mechanics

### [7.0.1](https://kaos.sh/ek/7.0.1)

- **`[fsutil]`** Fixed windows stubs for compatibility with latest changes

### [7.0.0](https://kaos.sh/ek/7.0.0)

- **`[usage]`** Added interface for different ways to check application updates
- **`[usage]`** Added Github update checker
- **`[usage]`** Moved `CommandsColorTag`, `OptionsColorTag`, `Breadcrumbs` to `Info` struct (_incompatible changes_)
- **`[fsutil]`** Now `ListingFilter` must be passed as value instead of pointer (_incompatible changes_)
- **`[fsutil]`** Added support of filtering by size for `ListingFilter`
- **`[version]`** Now `Parse` return value instead of pointer
- **`[cron]`** Improved expressions parsing
- **`[version]`** Added fuzz testing
- **`[cron]`** Added fuzz testing
- **`[knf]`** Added fuzz testing

---

### [6.2.1](https://kaos.sh/ek/6.2.1)

- **`[usage]`** Improved working with GitHub API

### [6.2.0](https://kaos.sh/ek/6.2.0)

- **`[netutil]`** Now GetIP return primary IPv4 address
- **`[netutil]`** Added method `GetIP6` which return main IPv6 address
- **`[usage]`** Showing info about latest available release on GitHub

### [6.1.0](https://kaos.sh/ek/6.1.0)

- **`[knf]`** Added tabs support in indentation
- **`[timeutil]`** Added new sequences `%n` (_new line symbol_) and `%K` (_milliseconds_)
- **`[timeutil]`** Code refactoring

### [6.0.0](https://kaos.sh/ek/6.0.0)

- **`[passwd]`** Much secure hash generation (now with sha512, bcrypt, and AES)
- **`[system]`** Improved changing process and arguments names
- **`[system/process]`** Fixed windows stubs

---

### [5.7.1](https://kaos.sh/ek/5.7.1)

- **`[usage]`** Improved build info output
- **`[system]`** Improved OS version search process

### [5.7.0](https://kaos.sh/ek/5.7.0)

- **`[system/process]`** `GetTree` now can return tree for custom root process
- **`[system/process]`** Fixed threads marking
- **`[fmtutil]`** Added method `CountDigits` for counting the number of digits in integer
- **`[terminal]`** Now `PrintWarnMessage` and `PrintErrorMessage` prints messages to stderr
- **`[usage]`** Added support for optional arguments in commands

### [5.6.0](https://kaos.sh/ek/5.6.0)

- **`[system]`** Added `Distribution` and `Version` info to `SystemInfo` struct
- **`[arg]`** Added bound arguments support
- **`[arg]`** Added conflicts arguments support
- **`[arg]`** Added method `Q` for merging several arguments to string (useful for `Alias`, `Bound` and `Conflicts`)

### [5.5.0](https://kaos.sh/ek/5.5.0)

- **`[system]`** Added method `CurrentTTY` which return path to current tty
- **`[system]`** Code refactoring

### [5.4.1](https://kaos.sh/ek/5.4.1)

- **`[fmtc]`** Fixed bug with parsing tags

### [5.4.0](https://kaos.sh/ek/5.4.0)

- **`[usage]`** Changed color for arguments from dark gray to light gray
- **`[usage]`** Added breadcrumbs output for commands and options
- **`[fmtutil]`** Fixed special symbols colorization in `ColorizePassword`

### [5.3.0](https://kaos.sh/ek/5.3.0)

- **`[fmtutil]`** Added method `ColorizePassword` for password colorization
- **`[passwd]`** Improved password generation and strength check

### [5.2.1](https://kaos.sh/ek/5.2.1)

- **`[log]`** Code refactoring
- **`[tmp]`** Added permissions customization for each temp struct

### [5.2.0](https://kaos.sh/ek/5.2.0)

- **`[terminal]`** Added password mask symbol color customization
- **`[terminal]`** [go-linenoise](https://github.com/essentialkaos/go-linenoise) updated to v3

### [5.1.1](https://kaos.sh/ek/5.1.1)

- **`[req]`** Improved `Engine` initialization routine
- **`[terminal]`** Fixed bug in windows stub with error variable name

### [5.1.0](https://kaos.sh/ek/5.1.0)

- **`[req]`** Improved `SetUserAgent` method for appending subpackages versions

### [5.0.1](https://kaos.sh/ek/5.0.1)

- **`[usage]`** Fixed examples header

### [5.0.0](https://kaos.sh/ek/5.0.0)

- **`[req]`** Fixed major bug with setting method through helper methods
- **`[req]`** Multi-client feature (_use `req.Engine` instead `req.Request` struct methods_)
- **`[crypto]`** Package divided into multiple packages (`hash`, `passwd`, `uuid`)
- **`[uuid]`** Added UUID generation based on SHA-1 hash of namespace UUID and name (_version 5_)
- **`[req]`** Added different types support for `Query`
- **`[knf]`** Added `NotContains` validator which checks if given config property contains any value from given slice
- **`[kv]`** Using values instead pointers
- **`[system]`** Added custom duration support for `GetNetworkSpeed` and `GetIOUtil`
- **`[version]`** Improved version parsing
- **`[system]`** More logical `RunAsUser` arguments naming
- **`[terminal]`** Minor fixes in windows stubs
- **`[netutil]`** Added tests
- **`[system]`** Code refactoring
* Added usage examples

---

### [3.5.1](https://kaos.sh/ek/3.5.1)

- **`[usage]`** Using dark gray color for license and copyright
- **`[fmtutil]`** Added global variable `SeparatorColorTag` for separator color customization
- **`[fmtutil]`** Added global variable `SeparatorTitleColorTag` for separator title color customization

### [3.5.0](https://kaos.sh/ek/3.5.0)

- **`[terminal]`** Using forked [go.linenoise](https://github.com/essentialkaos/go-linenoise) package instead original
- **`[terminal]`** Added hints support from new version of `go.linenoise`
- **`[fmtc]`** Light colors tag (`-`) support
- **`[usage]`** Using dark gray color for option values and example description
- **`[tmp]`** Added `DefaultDirPerms` and `DefaultFilePerms` global variables for permissions customization
- **`[tmp]`** Improved error handling

### [3.4.2](https://kaos.sh/ek/3.4.2)

- **`[strutil]`** Fixed bug with overflowing int in `Tail` method

### [3.4.1](https://kaos.sh/ek/3.4.1)

- **`[terminal]`** Improved reading user input

### [3.4.0](https://kaos.sh/ek/3.4.0)

- **`[httputil]`** Added `GetRequestAddr`, `GetRemoteAddr`, `GetRemoteHost`, `GetRemotePort` methods

### [3.3.1](https://kaos.sh/ek/3.3.1)

- **`[usage]`** Fixed bug with rendering command groups
- **`[terminal]`** Small fixes in windows stubs

### [3.3.0](https://kaos.sh/ek/3.3.0)

- **`[system/process]`** Added new package for getting information about active system processes
- **`[terminal]`** Fixed bug with title formatting in `ReadAnswer` method

### [3.2.3](https://kaos.sh/ek/3.2.3)

- **`[terminal]`** Fixed bug with title formatting in `ReadUI` method

### [3.2.2](https://kaos.sh/ek/3.2.2)

- **`[req]`** Added content types constants

### [3.2.1](https://kaos.sh/ek/3.2.1)

- **`[knf]`** Fixed typo in tests
- **`[strutil]`** Removed unreachable code

### [3.2.0](https://kaos.sh/ek/3.2.0)

- **`[strutil]`** Added method `Len` which returns number of symbols in string
- **`[strutil]`** UTF-8 support for `Substr`, `Tail`, `Head` and `Ellipsis` methods
- **`[strutil]`** Added some benchmarks to tests
- **`[fsutil]`** Fixed `GetPerm` stub for Windows
- **`[fsutil]`** Fixed package description

### [3.1.3](https://kaos.sh/ek/3.1.3)

- **`[req]`** `RequestTimeout` set to 0 (_disabled_) by default

### [3.1.2](https://kaos.sh/ek/3.1.2)

- **`[terminal]`** Fixed bug with source name file conventions
- **`[system]`** Fixed bug with appending real user info on MacOS X

### [3.1.1](https://kaos.sh/ek/3.1.1)

- **`[req]`** Small fixes in Request struct fields types

### [3.1.0](https://kaos.sh/ek/3.1.0)

- **`[req]`** Lazy transport initialization
- **`[req]`** Added `DialTimeout` and `RequestTimeout` variables for timeouts control

### [3.0.3](https://kaos.sh/ek/3.0.3)

- **`[system]`** Removed debug output

### [3.0.2](https://kaos.sh/ek/3.0.2)

* Added makefile with some helpful commands (`fmt`, `deps`, `test`)
* Small fixes in docs

### [3.0.1](https://kaos.sh/ek/3.0.1)

- **`[sliceutil]`** Code refactoring
- **`[knf]`** Typo fixed
- **`[terminal]`** Typo fixed
* Some minor changes

### [3.0.0](https://kaos.sh/ek/3.0.0)

- **`[fmtutil]`** Pluralization moved from `fmtutil` to separate package `pluralize` (_incompatible changes_)
- **`[pluralize]`** Brand new pluralization package with more than 140 languages support
- **`[timeutil]`** Improved `PrettyDuration` output
- **`[system]`** Now `SessionInfo` contnains full user info (`Info` struct) instead username (_incompatible changes_)
- **`[timeutil]`** Code refactoring
- **`[system]`** Code refactoring
- **`[log]`** Code refactoring
- **`[arg]`** Code refactoring

---

### [2.0.2](https://kaos.sh/ek/2.0.2)

- **`[pid]`** Added method `IsWorks` which return true if process with PID from PID file is active
- **`[pid]`** Increased code coverage

### [2.0.1](https://kaos.sh/ek/2.0.1)

- **`[terminal]`** Fixed bugs with Windows stubs
- **`[signal]`** Fixed bugs with Windows stubs

### [2.0.0](https://kaos.sh/ek/2.0.0)

- **`[color]`** New package for working with colors
- **`[usage]`** Added color tags support for description
- **`[terminal]`** Improved reading y/n answers (_incompatible changes_)
- **`[strutil]`** Added method `Fields` for "smart" string splitting
- **`[system]`** Methods `GetUsername` and `GetGroupname` deprecated
- **`[system]`** Added method `GroupList` for user struct which returns slice with user groups names
- **`[jsonutil]`** Code refactoring
- **`[usage]`** Code refactoring

---

### [1.8.3](https://kaos.sh/ek/1.8.3)

- **`[signal]`** Added method `Send` for sending signal to process

### [1.8.2](https://kaos.sh/ek/1.8.2)

- **`[log]`** Fixed bug with logging empty strings

### [1.8.1](https://kaos.sh/ek/1.8.1)

- **`[sortutil]`** Added method `VersionCompare` which can be used for custom version sorting

### [1.8.0](https://kaos.sh/ek/1.8.0)

- **`[sortutil]`** Added case insensitive strings sorting
- **`[sliceutil]`** Added `Deduplicate` method
- **`[strutil]`** Added `ReplaceAll` method
- **`[terminal]`** method `fmtutil.GetTermSize` moved to `terminal.GetSize`
- **`[timeutil]`** Added method `ParseDuration` which parses duration in `1w2d3h5m6s` format

### [1.7.8](https://kaos.sh/ek/1.7.8)

- **`[terminal]`** Custom prompt support
- **`[terminal]`** Custom masking symbols support
- **`[terminal]`** Code refactoring

### [1.7.7](https://kaos.sh/ek/1.7.7)

- **`[fsutil]`** Fixed bug in `List` method with filtering output
- **`[fsutil]`** Fixed bug with `NotPerms` filtering

### [1.7.6](https://kaos.sh/ek/1.7.6)

- **`[env]`** Added methods for getting env vars as string, int, and float

### [1.7.5](https://kaos.sh/ek/1.7.5)

- **`[usage]`** Added docs for exported fields in About struct

### [1.7.4](https://kaos.sh/ek/1.7.4)

- **`[fsutils]`** Added fs walker (bash `pushd`/`popd` analog)

### [1.7.3](https://kaos.sh/ek/1.7.3)

- **`[fsutil]`** Method `ListAbsolute` ranamed to `ListToAbsolute`

### [1.7.2](https://kaos.sh/ek/1.7.2)

- **`[errutil]`** Added method Chain

### [1.7.1](https://kaos.sh/ek/1.7.1)

- **`[log]`** Improved min level changing

### [1.7.0](https://kaos.sh/ek/1.7.0)

- **`[fsutil]`** Fixed major bug with closing file descriptor after directory listing
- **`[fsutil]`** Fixed major bug with closing file descriptor after counting lines in file
- **`[fsutil]`** Fixed major bug with closing file descriptor after checking number of files in directory

### [1.6.5](https://kaos.sh/ek/1.6.5)

- **`[fsutil]`** Improved docs
- **`[fsutil]`** Added method (wrapper) for moving files

### [1.6.4](https://kaos.sh/ek/1.6.4)

- **`[path]`** Added method IsDotfile for checking dotfile names

### [1.6.3](https://kaos.sh/ek/1.6.3)

- **`[strutil]`** Added methods PrefixSize and SuffixSize

### [1.6.2](https://kaos.sh/ek/1.6.2)

- **`[fsutil]`** Improved working with paths
- **`[fsutil]`** Added method ProperPath to windows stub

### [1.6.1](https://kaos.sh/ek/1.6.1)

- **`[path]`** Fixed windows stub

### [1.6.0](https://kaos.sh/ek/1.6.0)

- **`[path]`** Added package for working with paths

### [1.5.1](https://kaos.sh/ek/1.5.1)

- **`[knf]`** Fixed bug in HasProp method which returns true for unset properties

### [1.5.0](https://kaos.sh/ek/1.5.0)

- **`[tmp]`** Improved error handling
- **`[tmp]`** Changed name pattern of temporary files and directories

### [1.4.5](https://kaos.sh/ek/1.4.5)

- **`[pid]`** Fixed bug with PID file creation
- **`[pid]`** Increased coverage

### [1.4.4](https://kaos.sh/ek/1.4.4)

- **`[errutil]`** Added method Num which returns number of errors

### [1.4.3](https://kaos.sh/ek/1.4.3)

- **`[errutil]`** Improved Add method

### [1.4.2](https://kaos.sh/ek/1.4.2)

- **`[fsutil]`** Added method `ProperPath` which return first proper path from given slice

### [1.4.1](https://kaos.sh/ek/1.4.1)

- **`[fsutil]`** Added partial FreeBSD support
- **`[system]`** Added partial FreeBSD support
- **`[log]`** Some minor fixes in tests

### [1.4.0](https://kaos.sh/ek/1.4.0)

- **`[kv]`** Added package with simple key-value structs

### [1.3.3](https://kaos.sh/ek/1.3.3)

- **`[strutil]`** Fixed bug in Tail method

### [1.3.2](https://kaos.sh/ek/1.3.2)

- **`[strutil]`** Added method Head for subtraction first symbols from the string
- **`[strutil]`** Added method Tail for subtraction last symbols from the string

### [1.3.1](https://kaos.sh/ek/1.3.1)

* Improved TravisCI build script for support pkg.re
* Added pkg.re usage

### [1.3.0](https://kaos.sh/ek/1.3.0)

- **`[system]`** Fixed major bug with OS X compatibility
- **`[fmtutil]`** Fixed tests for OS X

### [1.2.2](https://kaos.sh/ek/1.2.2)

- **`[req]`** Added flag for marking connection to close

### [1.2.1](https://kaos.sh/ek/1.2.1)

- **`[crypto]`** Small improvements in hash generation
- **`[csv]`** Increased code coverage
- **`[easing]`** Increased code coverage
- **`[fmtc]`** Increased code coverage
- **`[httputil]`** Increased code coverage
- **`[jsonutil]`** Increased code coverage
- **`[pid]`** Increased code coverage
- **`[req]`** Increased code coverage
- **`[req]`** Increased default timeout to 10 seconds
- **`[strutil]`** Increased code coverage
- **`[timeutil]`** Increased code coverage

### [1.2.0](https://kaos.sh/ek/1.2.0)

- **`[log]`** Now buffered I/O must be enabled manually
- **`[log]`** Auto flushing for bufio

### [1.1.1](https://kaos.sh/ek/1.1.1)

- **`[system]`** Added JSON tags for User, Group and SessionInfo structs
- **`[usage]`** Info now can use os.Args`[0]` for info rendering
- **`[version]`** Added package for working with version in semver notation

### [1.1.0](https://kaos.sh/ek/1.1.0)

- **`[arg]`** Changed default fail values (int -1 → 0, float -1.0 → 0.0)
- **`[arg]`** Increased code coverage
- **`[arg]`** Many minor fixes
- **`[cron]`** Fixed rare bug
- **`[cron]`** Increased code coverage
- **`[crypto]`** Increased code coverage
- **`[easing]`** Increased code coverage
- **`[errutil]`** Increased code coverage
- **`[fmtc]`** Increased code coverage
- **`[fmtutil]`** Increased code coverage
- **`[jsonutil]`** Increased code coverage
- **`[knf]`** Fixed bug in Reload method for global config 
- **`[knf]`** Improved Reload method
- **`[knf]`** Increased code coverage
- **`[log]`** Increased code coverage
- **`[mathutil]`** Increased code coverage
- **`[pid]`** Increased code coverage
- **`[rand]`** Increased code coverage
- **`[req]`** Fixed bug with Accept header
- **`[req]`** Increased code coverage
- **`[sliceutil]`** Increased code coverage
- **`[sortutil]`** Increased code coverage
- **`[spellcheck]`** Increased code coverage
- **`[strutil]`** Increased code coverage
- **`[system]`** Added method system.SetProcName for changing process name
- **`[timeutil]`** Fixed bug in PrettyDuration method
- **`[timeutil]`** Increased code coverage
- **`[tmp]`** Increased code coverage

### [1.0.1](https://kaos.sh/ek/1.0.1)

- **`[system]`** Fixed bug in fs usage calculation
- **`[usage]`** Improved new Info struct creation

### [1.0.0](https://kaos.sh/ek/1.0.0)

_Initial public release_
