## Changelog

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

* `[knf]` Fixed bug in HasProp method which return true for unset properties

#### v1.5.0

* `[tmp]` Improved error handling
* `[tmp]` Changed name pattern of temporary files and directories

#### v1.4.5

* `[pid]` Fixed bug with pid file creation
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

* `[system]` Added json tags for User, Group and SessionInfo structs
* `[usage]` Info now can use os.Args`[0]` for info rendering
* `[version]` Added package for working with version in simver notation

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
