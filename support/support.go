/*
Package support provides methods for collecting and printing support information
about system.

By default, it collects information about the application and environment:
  - App name
  - App version
  - Go version used for build
  - Binary SHA checksum
  - Git commit SHA checksum

There are also some sub-packages to collect/parse additional information:
  - apps: Extracting apps versions info
  - deps: Extracting dependency information from gomod data
  - pkgs: Collecting information about installed packages
  - services: Collecting information about services
  - fs: Collecting information about the file system
  - network: Collecting information about the network
  - resources: Collecting information about CPU and memory
  - kernel: Collecting information from OS kernel

Example of collecting maximum information about the application and system:

	support.Collect("TestApp", "12.3.4").
	  WithRevision("fc8d81e").
	  WithDeps(deps.Extract(gomodData)).
	  WithApps(apps.Golang(), apps.GCC()).
	  WithPackages(pkgs.Collect("rpm", "go,golang", "java,jre,jdk", "nano")).
	  WithServices(services.Collect("firewalld", "nginx")).
	  WithChecks(myAppAvailabilityCheck()).
	  WithEnvVars("LANG", "PAGER", "SSH_CLIENT").
	  WithNetwork(network.Collect("https://cloudflare.com/cdn-cgi/trace")).
	  WithFS(fs.Collect()).
	  WithResources(resources.Collect()).
	  WithKernel(kernel.Collect("vm.nr_hugepages*", "vm.swappiness")).
	  Print()

Also, you can't encode data to JSON/GOB and send it to your server instead of printing
it to the console.

	info := support.Collect("TestApp", "12.3.4").
	  WithRevision("fc8d81e").
	  WithDeps(deps.Extract(gomodData)).
	  WithApps(apps.Golang(), apps.GCC()).
	  WithPackages(pkgs.Collect("rpm", "go,golang", "java,jre,jdk", "nano")).
	  WithServices(services.Collect("firewalld", "nginx")).
	  WithChecks(myAppAvailabilityCheck()).
	  WithEnvVars("LANG", "PAGER", "SSH_CLIENT").
	  WithNetwork(network.Collect("https://cloudflare.com/cdn-cgi/trace")).
	  WithFS(fs.Collect()).
	  WithResources(resources.Collect()).
	  WithKernel(kernel.Collect("vm.nr_hugepages*", "vm.swappiness"))

	b, _ := json.Marshal(info)

	fmt.Println(string(b))
*/
package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/hash"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type CheckStatus string

const (
	CHECK_OK    CheckStatus = "ok"
	CHECK_ERROR CheckStatus = "error"
	CHECK_WARN  CheckStatus = "warn"
	CHECK_SKIP  CheckStatus = "skip"
)

type ServiceStatus string

const (
	STATUS_WORKS   ServiceStatus = "works"
	STATUS_STOPPED ServiceStatus = "stopped"
	STATUS_UNKNOWN ServiceStatus = "unknown"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Info contains all support information (can be encoded in JSON/GOB)
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Binary  string `json:"binary"`

	Build     *BuildInfo     `json:"build,omitempty"`
	OS        *OSInfo        `json:"os,omitempty"`
	System    *SystemInfo    `json:"system,omitempty"`
	Network   *NetworkInfo   `json:"network,omitempty"`
	Resources *ResourcesInfo `json:"resources,omitempty"`
	Kernel    []KernelParam  `json:"kernel,omitempty"`
	FS        []FSInfo       `json:"fs,omitempty"`
	Pkgs      []Pkg          `json:"pkgs,omitempty"`
	Services  []Service      `json:"services,omitempty"`
	Deps      []Dep          `json:"deps,omitempty"`
	Apps      []App          `json:"apps,omitempty"`
	Checks    []Check        `json:"checks,omitempty"`
	Env       []EnvVar       `json:"env,omitempty"`
}

// BuildInfo contains information about binary
type BuildInfo struct {
	GoVersion string `json:"go_version"`
	GoArch    string `json:"go_arch"`
	GoOS      string `json:"go_os"`
	CGO       bool   `json:"cgo"`

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

// ResourcesInfo contains information about system resources
type ResourcesInfo struct {
	CPU       []CPUInfo `json:"cpu"`
	MemTotal  uint64    `json:"mem_total,omitempty"`
	MemFree   uint64    `json:"mem_free,omitempty"`
	MemUsed   uint64    `json:"mem_used,omitempty"`
	SwapTotal uint64    `json:"swap_total,omitempty"`
	SwapFree  uint64    `json:"swap_free,omitempty"`
	SwapUsed  uint64    `json:"swap_used,omitempty"`
}

// CPUInfo contains info about CPU
type CPUInfo struct {
	Model   string `json:"model"`
	Threads int    `json:"threads"`
	Cores   int    `json:"cores"`
}

// Service contains basic info about service
type Service struct {
	Name      string        `json:"name"`
	Status    ServiceStatus `json:"status"`
	IsPresent bool          `json:"is_present"`
	IsEnabled bool          `json:"is_enabled"`
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

// KernelParam contains info about kernel parameter
type KernelParam = EnvVar

// Check contains info about custom check
type Check struct {
	Status  CheckStatus `json:"status"`
	Title   string      `json:"title"`
	Message string      `json:"message,omitempty"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var buildInfoProvider = debug.ReadBuildInfo

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects basic info about system
func Collect(app, ver string) *Info {
	bin, _ := os.Executable()

	if bin != "" {
		bin = filepath.Base(bin)
	}

	info := &Info{
		Name:    app,
		Version: ver,
		Binary:  bin,
	}

	info.appendBuildInfo()
	info.appendOSInfo()
	info.appendSystemInfo()

	return info
}

// ////////////////////////////////////////////////////////////////////////////////// //

// WithDeps adds information about dependencies
func (i *Info) WithDeps(deps []Dep) *Info {
	if i == nil {
		return nil
	}

	if len(deps) > 0 {
		i.Deps = deps
	}

	return i
}

// WithRevision adds git revision
func (i *Info) WithRevision(rev string) *Info {
	if i == nil {
		return nil
	}

	if rev != "" {
		i.Build.GitSHA = rev
		return i
	}

	return i
}

// WithPackages adds information about packages
func (i *Info) WithPackages(pkgs []Pkg) *Info {
	if i == nil {
		return nil
	}

	i.Pkgs = append(i.Pkgs, pkgs...)

	return i
}

// WithServices adds information about services
func (i *Info) WithServices(services []Service) *Info {
	if i == nil {
		return nil
	}

	i.Services = append(i.Services, services...)

	return i
}

// WithPackages adds information about system apps
func (i *Info) WithApps(apps ...App) *Info {
	if i == nil {
		return nil
	}

	i.Apps = append(i.Apps, apps...)

	return i
}

// WithChecks adds information custom checks
func (i *Info) WithChecks(check ...Check) *Info {
	if i == nil {
		return nil
	}

	i.Checks = append(i.Checks, check...)

	return i
}

// WithEnvVars adds information with environment variables
func (i *Info) WithEnvVars(vars ...string) *Info {
	if i == nil {
		return nil
	}

	for _, k := range vars {
		if k == "" {
			continue
		}

		v := os.Getenv(k)

		if v != "" {
			i.Env = append(i.Env, EnvVar{k, v})
		}
	}

	return i
}

// WithNetwork adds information about the network
func (i *Info) WithNetwork(info *NetworkInfo) *Info {
	if i == nil {
		return nil
	}

	i.Network = info

	return i
}

// WithFS adds file system information
func (i *Info) WithFS(info []FSInfo) *Info {
	if i == nil {
		return nil
	}

	i.FS = info

	return i
}

// WithResources adds system resources information
func (i *Info) WithResources(info *ResourcesInfo) *Info {
	if i == nil {
		return nil
	}

	i.Resources = info

	return i
}

// WithKernel adds kernel parameters info
func (i *Info) WithKernel(info []KernelParam) *Info {
	if i == nil {
		return nil
	}

	i.Kernel = info

	return i
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Print prints support info
func (i *Info) Print() {
	if i == nil {
		return
	}

	fmtutil.SeparatorTitleColorTag = "{s-}"
	fmtutil.SeparatorFullscreen = false
	fmtutil.SeparatorColorTag = "{s-}"
	fmtutil.SeparatorSize = 80

	i.printAppInfo()
	i.printOSInfo()
	i.printResourcesInfo()
	i.printKernelInfo()
	i.printNetworkInfo()
	i.printFSInfo()
	i.printEnvVars()
	i.printPackagesInfo()
	i.printServicesInfo()
	i.printAppsInfo()
	i.printChecksInfo()
	i.printDependencies()

	fmtutil.Separator(false)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// appendBuildInfo appends build info
func (i *Info) appendBuildInfo() {
	i.Build = &BuildInfo{
		GoVersion: strings.TrimPrefix(runtime.Version(), "go"),
		GoArch:    runtime.GOARCH,
		GoOS:      runtime.GOOS,
	}

	bin, _ := os.Executable()
	binSHA := hash.FileHash(bin)
	info, ok := buildInfoProvider()

	if ok {
		for _, s := range info.Settings {
			switch s.Key {
			case "CGO_ENABLED":
				if s.Value == "1" {
					i.Build.CGO = true
				}
			case "vcs.revision":
				if len(s.Value) > 7 {
					i.Build.GitSHA = s.Value[:7]
				}
			}
		}
	}

	i.Build.BinSHA = strutil.Head(binSHA, 7)
}

// printAppInfo prints info about app
func (i *Info) printAppInfo() {
	fmtutil.Separator(false, "APPLICATION")

	name := i.Name

	if strings.ToLower(i.Name) != strings.ToLower(i.Binary) {
		name += fmtc.Sprintf(" {s-}(%s){!}", i.Binary)
	}

	format(7, true,
		"Name", name,
		"Version", i.Version,
	)

	if i.Build == nil {
		return
	}

	goInfo := fmtc.Sprintf("%s {s}(%s/%s){!}", i.Build.GoVersion, i.Build.GoOS, i.Build.GoArch)

	if i.Build.CGO {
		goInfo += fmtc.Sprint(" {s}+CGO{!}")
	}

	format(7, false,
		"Go", goInfo,
		"Git SHA", i.Build.GitSHA+getHashColorBullet(i.Build.GitSHA),
		"Bin SHA", i.Build.BinSHA+getHashColorBullet(i.Build.BinSHA),
	)
}

// printOSInfo prints info about OS and system
func (i *Info) printOSInfo() {
	if i.OS != nil {
		fmtutil.Separator(false, "OS")

		format(12, true,
			"Name", strutil.Q(i.OS.coloredName, i.OS.Name),
			"Pretty Name", strutil.Q(i.OS.coloredPrettyName, i.OS.PrettyName),
			"Version", i.OS.Version,
			"ID", i.OS.ID,
			"ID Like", i.OS.IDLike,
			"Version ID", i.OS.VersionID,
			"Version Code", i.OS.VersionCode,
			"Platform ID", i.OS.PlatformID,
			"CPE", i.OS.CPE,
		)
	} else if i.System != nil {
		fmtutil.Separator(false, "SYSTEM INFO")

		format(12, true,
			"Name", i.System.Name,
			"Arch", i.System.Arch,
			"Kernel", i.System.Kernel,
		)
	}

	if i.System != nil {
		format(12, true,
			"Arch", i.System.Arch,
			"Kernel", i.System.Kernel,
		)
	}

	format(12, true, "Locale", os.Getenv("LANG"))

	if i.System.ContainerEngine != "" {
		fmtc.NewLine()
		switch i.System.ContainerEngine {
		case "docker":
			format(12, true, "Container", "Yes (Docker)")
		case "docker+runsc":
			format(12, true, "Container", "Yes (Docker+gVisor)")
		case "podman":
			format(12, true, "Container", "Yes (Podman)")
		case "lxc":
			format(12, true, "Container", "Yes (LXC)")
		case "yandex":
			format(12, true, "Container", "Yes (Yandex Serverless)")
		}
	}
}

// printResourcesInfo prints resources info
func (i *Info) printResourcesInfo() {
	if i.Resources == nil {
		return
	}

	fmtutil.Separator(false, "RESOURCES")

	if len(i.Resources.CPU) > 0 {
		var procs, cores int

		for i, p := range i.Resources.CPU {
			fmtc.Printf(
				"  {s-}%d.{!} %s {s}[ %dC × %dT → %d ]{!}\n",
				i+1, p.Model, p.Cores, p.Threads, p.Cores*p.Threads,
			)
			procs++
			cores += p.Cores * p.Threads
		}

		fmtc.NewLine()

		format(10, true,
			"Processors", fmtutil.PrettyNum(procs),
			"Cores", fmtutil.PrettyNum(cores),
		)

		fmtc.NewLine()
	}

	if i.Resources.MemTotal > 0 {
		perc := (float64(i.Resources.MemUsed) / float64(i.Resources.MemTotal)) * 100.0

		format(6, true, "Memory", fmtc.Sprintf(
			"%s {s}/{!} %s {s-}(%s){!}",
			fmtutil.PrettySize(i.Resources.MemUsed, " "),
			fmtutil.PrettySize(i.Resources.MemTotal, " "),
			fmtutil.PrettyPerc(perc),
		))
	}

	if i.Resources.SwapTotal > 0 {
		perc := (float64(i.Resources.SwapUsed) / float64(i.Resources.SwapTotal)) * 100.0

		format(6, true, "Swap", fmtc.Sprintf(
			"%s {s}/{!} %s {s-}(%s){!}",
			fmtutil.PrettySize(i.Resources.SwapUsed, " "),
			fmtutil.PrettySize(i.Resources.SwapTotal, " "),
			fmtutil.PrettyPerc(perc),
		))
	} else {
		format(6, true, "Swap", "")
	}
}

// printKernelInfo prints kernel parameters
func (i *Info) printKernelInfo() {
	if i.Kernel == nil {
		return
	}

	fmtutil.Separator(false, "KERNEL")

	keySize := getMaxKeySize(i.Kernel)

	for _, p := range i.Kernel {
		if mathutil.IsInt(p.Value) {
			vi, _ := strconv.Atoi(p.Value)
			format(keySize, true, p.Key, fmtutil.PrettyNum(vi))
		} else {
			format(keySize, true, p.Key, p.Value)
		}
	}
}

// printEnvVars prints environment variables
func (i *Info) printEnvVars() {
	if len(i.Env) == 0 {
		return
	}

	fmtutil.Separator(false, "ENVIRONMENT VARIABLES")

	size := getMaxKeySize(i.Env)

	for _, ev := range i.Env {
		format(size, true, ev.Key, ev.Value)
	}
}

// printPackagesInfo prints info about packages
func (i *Info) printPackagesInfo() {
	if len(i.Pkgs) == 0 {
		return
	}

	fmtutil.Separator(false, "PACKAGES")

	size := getMaxAppNameSize(i.Pkgs)

	for _, p := range i.Pkgs {
		if p.Name == "" {
			continue
		}

		format(size, true, p.Name, p.Version)
	}
}

// printServicesInfo prints services info
func (i *Info) printServicesInfo() {
	if len(i.Services) == 0 {
		return
	}

	fmtutil.Separator(false, "SERVICES")

	size := getMaxServiceNameSize(i.Services)

	for _, s := range i.Services {
		var status string

		switch s.Status {
		case STATUS_WORKS:
			status = "{g}works{!}"
		case STATUS_STOPPED:
			status = "{s}stopped{!}"
		}

		if s.IsEnabled {
			status += " {s-}(enabled){!}"
		}

		format(size, true, s.Name, fmtc.Sprint(status))
	}
}

// printAppsInfo prints info about applications
func (i *Info) printAppsInfo() {
	if len(i.Apps) == 0 {
		return
	}

	fmtutil.Separator(false, "APPLICATIONS")

	size := getMaxAppNameSize(i.Apps)

	for _, a := range i.Apps {
		if a.Name == "" {
			continue
		}

		v := a.Version

		v = strings.ReplaceAll(v, "(", "{s}(")
		v = strings.ReplaceAll(v, ")", "){!}")

		format(size, true, a.Name, fmtc.Sprint(v))
	}
}

// printChecksInfo prints checks info
func (i *Info) printChecksInfo() {
	if len(i.Checks) == 0 {
		return
	}

	fmtutil.Separator(false, "CHECKS")

	for _, c := range i.Checks {
		if c.Title == "" {
			continue
		}

		switch c.Status {
		case CHECK_OK:
			fmtc.Print("  {g}✔ {!}")
		case CHECK_SKIP:
			fmtc.Print("  {s-}✔ {!}")
		case CHECK_WARN:
			fmtc.Print("  {y}✖ {!}")
		case CHECK_ERROR:
			fmtc.Print("  {r}✖ {!}")
		}

		fmtc.Printf(" {*}%s{!}", c.Title)

		if c.Message == "" {
			fmtc.NewLine()
			continue
		}

		switch c.Status {
		case CHECK_OK, CHECK_SKIP:
			fmtc.Printfn(" {s}— {&}%s{!}", c.Message)
		case CHECK_WARN:
			fmtc.Printfn(" {s}— {y}{&}%s{!}", c.Message)
		case CHECK_ERROR:
			fmtc.Printfn(" {s}— {r}{&}%s{!}", c.Message)
		}
	}
}

// printNetworkInfo prints network info
func (i *Info) printNetworkInfo() {
	if i.Network == nil {
		return
	}

	fmtutil.Separator(false, "NETWORK")

	format(0, false,
		"Hostname", i.Network.Hostname,
		"Public IP", i.Network.PublicIP,
		"IP v4", strings.TrimLeft(fmtutil.Wrap(
			strings.Join(i.Network.IPv4, " "),
			strings.Repeat(" ", 13), 80,
		), " "),
		"IP v6", strings.TrimLeft(fmtutil.Wrap(
			strings.Join(i.Network.IPv6, " "),
			strings.Repeat(" ", 13), 80,
		), " "),
	)
}

// printFSInfo prints filesystem info
func (i *Info) printFSInfo() {
	if len(i.FS) == 0 {
		return
	}

	fmtutil.Separator(false, "FILESYSTEM")

	size := getMaxDeviceNameSize(i.FS)

	for _, m := range i.FS {
		if m.Path == "" || m.Device == "" {
			continue
		}

		format(size, false, m.Device, fmtc.Sprintf(
			"%s {s}(%s){!} %s{s}/{!}%s {s-}(%s){!}",
			m.Path, m.Type, fmtutil.PrettySize(m.Used),
			fmtutil.PrettySize(m.Used+m.Free),
			fmtutil.PrettyPerc(mathutil.Perc(m.Used, m.Used+m.Free)),
		))
	}
}

// printDependencies prints used dependencies
func (i *Info) printDependencies() {
	if len(i.Deps) == 0 {
		return
	}

	fmtutil.Separator(false, "DEPENDENCIES")

	for _, dep := range i.Deps {
		switch dep.Extra {
		case "":
			fmtc.Printfn(" {s}%8s{!}  %s", dep.Version, dep.Path)
		default:
			fmtc.Printfn(" {s}%8s{!}  %s {s-}(%s){!}", dep.Version, dep.Path, dep.Extra)
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// format formats and prints records
func format(size int, printEmpty bool, records ...string) {
	if size <= 0 {
		for i := 0; i < len(records); i += 2 {
			if records[i+1] == "" && !printEmpty {
				continue
			}

			size = mathutil.Max(size, len(records[i]))
		}
	}

	size++

	for i := 0; i < len(records); i += 2 {
		name, value := records[i]+":", records[i+1]

		if value == "" && printEmpty {
			fm := fmt.Sprintf("  {*}%%-%ds{!}  {s-}—{!}\n", size)
			fmtc.Printf(fm, name)
		} else if value != "" {
			fm := fmt.Sprintf("  {*}%%-%ds{!}  %%s\n", size)
			fmtc.Printf(fm, name, value)
		}
	}
}

// getHashColorBullet return bullet with color from hash
func getHashColorBullet(v string) string {
	if v == "" || fmtc.DisableColors || !fmtc.IsTrueColorSupported() {
		return ""
	}

	return fmtc.Sprintf(" {#" + strutil.Head(v, 6) + "}● {!}")
}

// getMaxKeySize returns max key size
func getMaxKeySize(vars []EnvVar) int {
	var size int

	for _, ev := range vars {
		size = mathutil.Max(size, len(ev.Key))
	}

	return size
}

// getMaxAppNameSize returns max package name size
func getMaxAppNameSize(apps []App) int {
	var size int

	for _, p := range apps {
		size = mathutil.Max(size, len(p.Name))
	}

	return size
}

// getMaxServiceNameSize returns max package name size
func getMaxServiceNameSize(apps []Service) int {
	var size int

	for _, s := range apps {
		size = mathutil.Max(size, len(s.Name))
	}

	return size
}

// getMaxDeviceNameSize returns max device name size
func getMaxDeviceNameSize(mounts []FSInfo) int {
	var size int

	for _, m := range mounts {
		size = mathutil.Max(size, len(m.Device))
	}

	return size
}
