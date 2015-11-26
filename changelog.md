### Changelog

#### r6

* [crypto] Small improvements in hash generation
* [csv] Increased code coverage
* [req] Increased code coverage
* [req] Increased default timeout to 10 seconds

#### r5

* [log] Now buffered I/O must be enabled manually
* [log] Auto flushing for bufio

#### r4

* [system] Added json tags for User, Group and SessionInfo structs
* [usage] Info now can use os.Args[0] for info rendering
* [version] Added package for working with version in simver notation

#### r3

* [arg] Changed default fail values (int -1 → 0, float -1.0 → 0.0)
* [arg] Increased code coverage
* [arg] Many minor fixes
* [cron] Fixed rare bug
* [cron] Increased code coverage
* [crypto] Increased code coverage
* [easing] Increased code coverage
* [errutil] Increased code coverage
* [fmtc] Increased code coverage
* [fmtutil] Increased code coverage
* [jsonutil] Increased code coverage
* [knf] Fixed bug in Reload method for global config 
* [knf] Improved Reload method
* [knf] Increased code coverage
* [log] Increased code coverage
* [mathutil] Increased code coverage
* [pid] Increased code coverage
* [rand] Increased code coverage
* [req] Fixed bug with Accept header
* [req] Increased code coverage
* [sliceutil] Increased code coverage
* [sortutil] Increased code coverage
* [spellcheck] Increased code coverage
* [strutil] Increased code coverage
* [system] Added method system.SetProcName for changing process name
* [timeutil] Fixed bug in PrettyDuration method
* [timeutil] Increased code coverage
* [tmp] Increased code coverage

#### r2

* [system] Fixed bug in fs usage calculation
* [usage] Improved new Info struct creation

#### r1

Initial public release
