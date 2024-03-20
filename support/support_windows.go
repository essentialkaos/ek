//go:build !windows
// +build !windows

/*
Package support provides methods for collecting and printing support information
about system.

By default, it collects information about the application and environment:

- Name
- Version
- Go version used
- Binary SHA
- Git commit SHA
- Environment variables
- Applications

There are also some sub-packages to collect/parse additional information:

- apps: Package for extracting apps versions info
- deps: Package for extracting dependency information from gomod data
- pkgs: Package for collecting information about installed packages
- fs: Package for collecting information about the file system
- network: Package to collect information about the network

Example of collecting maximum information about the application and system:

	support.Collect("TestApp", "12.3.4").
	  WithRevision("fc8d81e").
	  WithDeps(deps.Extract(gomodData)).
	  WithApps(apps.Golang(), apps.GCC()).
	  WithPackages(pkgs.Collect("rpm", "go,golang", "java,jre,jdk", "nano")).
	  WithEnvVars("LANG", "PAGER", "SSH_CLIENT").
	  WithNetwork(network.Collect("https://domain.com/ip-echo")).
	  WithFS(fs.Collect()).
	  Print()

Also, you can't encode data to JSON/GOB and send it to your server instead of printing
it to the console.

	info := support.Collect("TestApp", "12.3.4").
	  WithRevision("fc8d81e").
	  WithDeps(deps.Extract(gomodData)).
	  WithApps(apps.Golang(), apps.GCC()).
	  WithPackages(pkgs.Collect("rpm", "go,golang", "java,jre,jdk", "nano")).
	  WithEnvVars("LANG", "PAGER", "SSH_CLIENT").
	  WithNetwork(network.Collect("https://domain.com/ip-echo")).
	  WithFS(fs.Collect())

	b, _ := json.Marshal(info)

	fmt.Println(string(b))
*/
package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Info contains all support information (can be encoded in JSON/GOB)
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`

	Build   *BuildInfo   `json:"build,omitempty"`
	OS      *OSInfo      `json:"os,omitempty"`
	System  *SystemInfo  `json:"system,omitempty"`
	Network *NetworkInfo `json:"network,omitempty"`
	FS      []FSInfo     `json:"fs,omitempty"`
	Pkgs    []Pkg        `json:"pkgs,omitempty"`
	Deps    []Dep        `json:"deps,omitempty"`
	Apps    []App        `json:"apps,omitempty"`
	Env     []EnvVar     `json:"env,omitempty"`
}

// BuildInfo contains information about binary
type BuildInfo struct {
	GoVersion string `json:"go_version"`
	GoArch    string `json:"go_arch"`
	GoOS      string `json:"go_os"`

	GitSHA string `json:"git_sha,omitempty"`
	BinSHA string `json:"bin_sha,omitempty"`
}

// OSInfo contains extended information about OS
type OSInfo struct {
	Name        string `json:"name,omitempty"`
	PrettyName  string `json:"pretty_name,omitempty"`
	Version     string `json:"version,omitempty"`
	Build       string `json:"build,omitempty"`
	ID          string `json:"id,omitempty"`
	IDLike      string `json:"id_like,omitempty"`
	VersionID   string `json:"version_id,omitempty"`
	VersionCode string `json:"version_code,omitempty"`
	PlatformID  string `json:"platform_id,omitempty"`
	CPE         string `json:"cpe,omitempty"`

	coloredName       string
	coloredPrettyName string
}

// SystemInfo contains basic information about system
type SystemInfo struct {
	Name            string `json:"name"`
	Arch            string `json:"arch"`
	Kernel          string `json:"kernel"`
	ContainerEngine string `json:"container_engine,omitempty"`
}

// NetworkInfo contains basic information about network
type NetworkInfo struct {
	Hostname string   `json:"hostname"`
	PublicIP string   `json:"public_ip,omitempty"`
	IPv4     []string `json:"ipv4"`
	IPv6     []string `json:"ipv6,omitempty"`
}

// FSInfo contains basic information about file system mount
type FSInfo struct {
	Path   string `json:"path,omitempty"`
	Device string `json:"device,omitempty"`
	Type   string `json:"type,omitempty"`
	Used   uint64 `json:"used,omitempty"`
	Free   uint64 `json:"free,omitempty"`
}

// App contains basic information about app
type App struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Pkg contains basic information about package
type Pkg = App

// Dep contains dependency information
type Dep struct {
	Path    string `json:"path"`
	Version string `json:"version"`
	Extra   string `json:"extra"`
}

// EnvVar contains information about environment variable
type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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
