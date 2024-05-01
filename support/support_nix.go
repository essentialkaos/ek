//go:build !windows
// +build !windows

package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/hash"
	"github.com/essentialkaos/ek/v12/mathutil"
	"github.com/essentialkaos/ek/v12/path"
	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects basic info about system
func Collect(app, ver string) *Info {
	bin, _ := os.Executable()

	if bin != "" {
		bin = path.Base(bin)
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

	i.Build.GitSHA = extractGitRevFromBuildInfo()

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

	i.Build.BinSHA = strutil.Head(binSHA, 7)
}

// printAppInfo prints info about app
func (i *Info) printAppInfo() {
	fmtutil.Separator(false, "APPLICATION INFO")

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

	format(7, false,
		"Go", fmtc.Sprintf("%s {s}(%s/%s){!}", i.Build.GoVersion, i.Build.GoOS, i.Build.GoArch),
		"Git SHA", i.Build.GitSHA+getHashColorBullet(i.Build.GitSHA),
		"Bin SHA", i.Build.BinSHA+getHashColorBullet(i.Build.BinSHA),
	)
}

// printOSInfo prints info about OS and system
func (i *Info) printOSInfo() {
	if i.OS != nil {
		fmtutil.Separator(false, "OS INFO")

		format(12, true,
			"Name", i.OS.coloredName,
			"Pretty Name", i.OS.coloredPrettyName,
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

	if i.System.ContainerEngine != "" {
		fmtc.NewLine()
		switch i.System.ContainerEngine {
		case "docker":
			format(12, true, "Container", "Yes (Docker)")
		case "podman":
			format(12, true, "Container", "Yes (Podman)")
		case "lxc":
			format(12, true, "Container", "Yes (LXC)")
		case "yandex":
			format(12, true, "Container", "Yes (Yandex Serverless)")
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
			fmtc.Printf(" {s}— {&}%s{!}\n", c.Message)
		case CHECK_WARN:
			fmtc.Printf(" {s}— {y}{&}%s{!}\n", c.Message)
		case CHECK_ERROR:
			fmtc.Printf(" {s}— {r}{&}%s{!}\n", c.Message)
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
		"IP v4", strings.Join(i.Network.IPv4, " "),
		"IP v6", strings.Join(i.Network.IPv6, " "),
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
			fmtc.Printf(" {s}%8s{!}  %s\n", dep.Version, dep.Path)
		default:
			fmtc.Printf(" {s}%8s{!}  %s {s-}(%s){!}\n", dep.Version, dep.Path, dep.Extra)
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

// extractGitRevFromBuildInfo extracts git SHA from embedded build info
func extractGitRevFromBuildInfo() string {
	info, ok := debug.ReadBuildInfo()

	if !ok {
		return ""
	}

	for _, s := range info.Settings {
		if s.Key == "vcs.revision" && len(s.Value) > 7 {
			return s.Value[:7]
		}
	}

	return ""
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
