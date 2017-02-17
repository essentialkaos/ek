// Package usage provides methods and structs for rendering info for command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v6/fmtc"
	"pkg.re/essentialkaos/ek.v6/req"
	"pkg.re/essentialkaos/ek.v6/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SPACES = "                                                                "
	_DOTS   = "................................................................"
)

const _BREADCRUMBS_MIN_SIZE = 16

// ////////////////////////////////////////////////////////////////////////////////// //

// About contains info about application
type About struct {
	App        string // App is application name
	Version    string // Version is current application version in semver notation
	Release    string // Release is current application release
	Build      string // Build is current application build
	Desc       string // Desc is short info about application
	Year       int    // Year is year when owner company was founded
	License    string // License is name of license
	Owner      string // Owner is name of owner (company/developer)
	Repository string // GitHub repository (owner/name)
}

// Info contains info about commands, options and examples
type Info struct {
	name     string
	args     string
	spoiler  string
	commands []option
	options  []option
	examples []example
	curGroup string
}

// ////////////////////////////////////////////////////////////////////////////////// //

type option struct {
	name  string
	desc  string
	args  []string
	group string
}

type example struct {
	cmd  string
	desc string
}

type release struct {
	Tag       string    `json:"tag_name"`
	Published time.Time `json:"published_at"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// CommandsColor contains default commands color
	CommandsColorTag = "{y}"

	// OptionsColor contains default options color
	OptionsColorTag = "{g}"

	// Use bread crumbs for commands and options output
	Breadcrumbs = false
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewInfo create new info struct
func NewInfo(name string, args ...string) *Info {
	if name == "" {
		return &Info{
			name:     filepath.Base(os.Args[0]),
			args:     strings.Join(args, " "),
			commands: make([]option, 0),
			options:  make([]option, 0),
			examples: make([]example, 0),
		}
	}

	return &Info{
		name:     name,
		args:     strings.Join(args, " "),
		commands: make([]option, 0),
		options:  make([]option, 0),
		examples: make([]example, 0),
	}
}

// AddGroup add new command group
func (info *Info) AddGroup(group string) {
	info.curGroup = group
}

// AddCommand add command (name, desc, args)
func (info *Info) AddCommand(a ...string) {
	group := "Commands"

	if info.curGroup != "" {
		group = info.curGroup
	}

	appendOption(a, &info.commands, group)
}

// AddOption add option (name, desc, args)
func (info *Info) AddOption(a ...string) {
	appendOption(a, &info.options, "Options")
}

// AddExample add example for some command (command, desc)
func (info *Info) AddExample(a ...string) {
	if len(a) == 0 {
		return
	}

	a = append(a, "")

	info.examples = append(info.examples,
		example{
			cmd:  a[0],
			desc: a[1],
		},
	)
}

// AddSpoiler add spoiler
func (info *Info) AddSpoiler(spoiler string) {
	info.spoiler = spoiler
}

// Render print usage info to console
func (info *Info) Render() {
	usageMessage := "\n{*}Usage:{!} " + info.name

	if len(info.commands) != 0 {
		usageMessage += " " + CommandsColorTag + "{command}{!}"
	}

	if len(info.options) != 0 {
		usageMessage += " " + OptionsColorTag + "{options}{!}"
	}

	if info.args != "" {
		usageMessage += " " + info.args
	}

	fmtc.Println(usageMessage)

	if info.spoiler != "" {
		fmtc.NewLine()
		fmtc.Println(info.spoiler)
	}

	if len(info.commands) != 0 {
		renderOptions(info.commands, CommandsColorTag)
	}

	if len(info.options) != 0 {
		renderOptions(info.options, OptionsColorTag)
	}

	if len(info.examples) != 0 {
		renderExamples(info)
	}

	fmtc.NewLine()
}

// Render print version info to console
func (about *About) Render() {
	switch {
	case about.Build != "":
		fmtc.Printf(
			"\n{*c}%s {c}%s{!}{s}%s{!} {s-}(%s){!} - %s\n\n",
			about.App, about.Version,
			about.Release, about.Build, about.Desc,
		)
	default:
		fmtc.Printf(
			"\n{*c}%s {c}%s{!}{s}%s{!} - %s\n\n",
			about.App, about.Version,
			about.Release, about.Desc,
		)
	}

	if about.Owner != "" {
		if about.Year == 0 {
			fmtc.Printf("{s-}Copyright (C) %s{!}\n", about.Owner)
		} else {
			fmtc.Printf(
				"{s-}Copyright (C) %d-%d %s{!}\n",
				about.Year, time.Now().Year(), about.Owner,
			)
		}
	}

	if about.License != "" {
		fmtc.Printf("{s-}%s{!}\n", about.License)
	}

	if about.Repository != "" {
		printLatestReleaseInfo(about.App, about.Version, about.Repository)
	}

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// appendOption append new option to options slice
func appendOption(data []string, options *[]option, group string) {
	if len(data) < 2 {
		return
	}

	var optionArgs []string

	if len(data) >= 3 {
		optionArgs = data[2:]
	}

	var name = data[0]

	if group == "Options" {
		name = parseOption(name)
	}

	*options = append(*options,
		option{
			name:  name,
			desc:  data[1],
			args:  optionArgs,
			group: group,
		},
	)
}

// parseOption parse option in format used by ek.arg
func parseOption(option string) string {
	if strings.HasPrefix(option, "-") {
		return option
	}

	if strings.Contains(option, ":") {
		optionSlice := strings.Split(option, ":")
		return "--" + optionSlice[1] + ", -" + optionSlice[0]
	}

	return "--" + option
}

// renderOptions render options
func renderOptions(options []option, colorTag string) {
	var (
		curGroup string
		opt      option
		maxSize  int
	)

	maxSize = getMaxOptionSize(options)

	for _, opt = range options {
		if curGroup != opt.group {
			printGroupHeader(opt.group)
			curGroup = opt.group
		}

		fmtc.Printf("  "+colorTag+"%s{!}", opt.name)

		if len(opt.args) != 0 {
			fmtc.Printf(" " + renderArgs(opt.args))
		}

		fmtc.Printf(getBreadcrumbs(opt, maxSize))
		fmtc.Printf(opt.desc)

		fmtc.NewLine()
	}
}

// renderExamples render examples
func renderExamples(info *Info) {
	printGroupHeader("Examples")

	total := len(info.examples)

	for index, example := range info.examples {
		fmtc.Printf("  %s %s\n", info.name, example.cmd)

		if example.desc != "" {
			fmtc.Printf("  {s-}%s{!}\n", example.desc)
		}

		if index < total-1 {
			fmtc.NewLine()
		}
	}
}

// renderArgs render args with colors
func renderArgs(args []string) string {
	var result string

	for _, a := range args {
		if strings.HasPrefix(a, "?") {
			result += "{s-}" + a[1:] + "{!} "
		} else {
			result += "{s}" + a + "{!} "
		}
	}

	return fmtc.Sprintf(strings.TrimRight(result, " "))
}

// getRenderedArgsSize return size of string with rendered arguments
func getRenderedArgsSize(args []string) int {
	var result int

	for _, a := range args {
		if strings.HasPrefix(a, "?") {
			result += len(a)
		} else {
			result += len(a) + 1
		}
	}

	return result
}

// getBreadCrumbs return bread crumbs (or spaces if colors are disabled) for
// option name aligning
func getBreadcrumbs(opt option, maxSize int) string {
	optLen := len(opt.name) + getRenderedArgsSize(opt.args)

	if Breadcrumbs && !fmtc.DisableColors && maxSize > _BREADCRUMBS_MIN_SIZE {
		return " {s-}" + _DOTS[:maxSize-optLen] + "{!} "
	}

	return " " + _SPACES[:maxSize-optLen] + " "
}

// getMaxOptionSize return longest option name size
func getMaxOptionSize(options []option) int {
	var result = 0

	for _, opt := range options {
		argsLen := getRenderedArgsSize(opt.args)
		optLen := len(opt.name) + argsLen + 2

		if optLen > result {
			result = optLen
		}
	}

	return result
}

// printGroupHeader print category header
func printGroupHeader(name string) {
	fmtc.Printf("\n{*}%s{!}\n\n", name)
}

// printLatestReleaseInfo print info about latest release on GitHub
func printLatestReleaseInfo(app, currentVersion, repository string) {
	latestRelease := getLatestRelease(app, currentVersion, repository)

	if latestRelease == nil || len(latestRelease.Tag) < 2 {
		return
	}

	latestVersion := latestRelease.Tag[1:]

	if !isNewerVersion(currentVersion, latestVersion) {
		return
	}

	fmtc.NewLine()

	days := int(time.Since(latestRelease.Published) / (time.Hour * 24))

	var colorTag string

	switch {
	case days < 14:
		colorTag = "{s}"
	case days < 60:
		colorTag = "{y}"
	default:
		colorTag = "{r}"
	}

	if days < 2 {
		fmtc.Printf(
			colorTag+"Latest version is %s (released 1 day ago){!}\n",
			latestVersion,
		)
	} else {
		fmtc.Printf(
			colorTag+"Latest version is %s (released %d days ago){!}\n",
			latestVersion, days,
		)
	}
}

// getLatestRelease fetch latest release from GitHub
func getLatestRelease(app, version, repository string) *release {
	engine := req.Engine{}

	engine.SetDialTimeout(2)
	engine.SetRequestTimeout(2)
	engine.SetUserAgent(app, version, "go.ek/6")

	response, err := engine.Get(req.Request{
		URL:         "https://api.github.com/repos/" + repository + "/releases/latest",
		AutoDiscard: true,
	})

	if err != nil || response.StatusCode != 200 {
		return nil
	}

	if response.Header.Get("X-RateLimit-Remaining") == "0" {
		return nil
	}

	var rel = &release{}

	err = response.JSON(rel)

	if err != nil {
		return nil
	}

	return rel
}

// isNewerVersion return true if latest version is greater than current
func isNewerVersion(current, latest string) bool {
	v1, err := version.Parse(current)

	if err != nil {
		return false
	}

	v2, err := version.Parse(latest)

	if err != nil {
		return false
	}

	return v2.Greater(v1)
}
