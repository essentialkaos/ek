package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Collect collects basic info about system
func Collect(app, ver string) *Info {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ WithDeps adds information about dependencies
func (i *Info) WithDeps(deps []Dep) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithRevision adds git revision
func (i *Info) WithRevision(rev string) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithPackages adds information about packages
func (i *Info) WithPackages(pkgs []Pkg) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithPackages adds information about system apps
func (i *Info) WithApps(apps ...App) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithChecks adds information custom checks
func (i *Info) WithChecks(check ...Check) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithEnvVars adds information with environment variables
func (i *Info) WithEnvVars(vars ...string) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithNetwork adds information about the network
func (i *Info) WithNetwork(info *NetworkInfo) *Info {
	panic("UNSUPPORTED")
}

// ❗ WithFS adds file system information
func (i *Info) WithFS(info []FSInfo) *Info {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Print prints support info
func (i *Info) Print() {
	panic("UNSUPPORTED")
}
