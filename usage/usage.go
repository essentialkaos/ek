// Package usage provides methods and structs for rendering info for command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v10/fmtc"
	"pkg.re/essentialkaos/ek.v10/mathutil"
	"pkg.re/essentialkaos/ek.v10/strutil"
	"pkg.re/essentialkaos/ek.v10/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SPACES = "                                                                "
	_DOTS   = "................................................................"
)

const _BREADCRUMBS_MIN_SIZE = 8

// ////////////////////////////////////////////////////////////////////////////////// //

// About contains info about application
type About struct {
	App     string // App is application name
	Version string // Version is current application version in semver notation
	Release string // Release is current application release
	Build   string // Build is current application build
	Desc    string // Desc is short info about application
	Year    int    // Year is year when owner company was founded
	License string // License is name of license
	Owner   string // Owner is name of owner (company/developer)

	// Function for checking application updates
	UpdateChecker UpdateChecker
}

// Info contains info about commands, options, and examples
type Info struct {
	CommandsColorTag string // CommandsColor contains default commands color
	OptionsColorTag  string // OptionsColor contains default options color
	Breadcrumbs      bool   // Use bread crumbs for commands and options output

	Name    string
	Args    []string
	Spoiler string

	Commands []*Command
	Options  []*Option
	Examples []*Example

	curGroup string
}

// UpdateChecker is a base for all update checkers
type UpdateChecker struct {
	Data      string
	CheckFunc func(app, version, data string) (string, time.Time, bool)
}

// ////////////////////////////////////////////////////////////////////////////////// //

type Command struct {
	Name  string
	Desc  string
	Args  []string
	Group string
}

type Option struct {
	Name string
	Desc string
	Args []string
}

type Example struct {
	Cmd  string
	Desc string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewInfo create new info struct
func NewInfo(args ...string) *Info {
	var name string

	if len(args) != 0 {
		name = args[0]
		args = args[1:]
	}

	if name == "" {
		name = filepath.Base(os.Args[0])
	}

	info := &Info{
		Name: name,
		Args: args,

		CommandsColorTag: "{y}",
		OptionsColorTag:  "{g}",
		Breadcrumbs:      true,
	}

	return info
}

// AddGroup add new command group
func (info *Info) AddGroup(group string) {
	info.curGroup = group
}

// AddCommand add command (name, description, args)
func (info *Info) AddCommand(a ...string) {
	group := "Commands"

	if info.curGroup != "" {
		group = info.curGroup
	}

	if len(a) < 2 {
		return
	}

	info.Commands = append(
		info.Commands,
		&Command{
			Name:  a[0],
			Desc:  a[1],
			Args:  a[2:],
			Group: group,
		},
	)
}

// AddOption add option (name, description, args)
func (info *Info) AddOption(a ...string) {
	if len(a) < 2 {
		return
	}

	info.Options = append(
		info.Options,
		&Option{
			Name: a[0],
			Desc: a[1],
			Args: a[2:],
		},
	)
}

// AddExample add example for some command (command, description)
func (info *Info) AddExample(a ...string) {
	if len(a) == 0 {
		return
	}

	a = append(a, "")

	info.Examples = append(info.Examples, &Example{a[0], a[1]})
}

// AddSpoiler add spoiler
func (info *Info) AddSpoiler(spoiler string) {
	info.Spoiler = spoiler
}

// Render print usage info to console
func (info *Info) Render() {
	usageMessage := "\n{*}Usage:{!} " + info.Name

	if len(info.Options) != 0 {
		usageMessage += " " + info.OptionsColorTag + "{options}{!}"
	}

	if len(info.Commands) != 0 {
		usageMessage += " " + info.CommandsColorTag + "{command}{!}"
	}

	if len(info.Args) != 0 {
		usageMessage += " " + strings.Join(info.Args, " ")
	}

	fmtc.Println(usageMessage)

	if info.Spoiler != "" {
		fmtc.NewLine()
		fmtc.Println(info.Spoiler)
	}

	if len(info.Commands) != 0 {
		renderCommands(info)
	}

	if len(info.Options) != 0 {
		renderOptions(info)
	}

	if len(info.Examples) != 0 {
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
			fmtc.Printf(
				"{s-}Copyright (C) %d %s{!}\n",
				time.Now().Year(), about.Owner,
			)
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

	if about.UpdateChecker.CheckFunc != nil && about.UpdateChecker.Data != "" {
		newVersion, releaseDate, hasUpdate := about.UpdateChecker.CheckFunc(
			about.App,
			about.Version,
			about.UpdateChecker.Data,
		)

		if hasUpdate && isNewerVersion(about.Version, newVersion) {
			printNewVersionInfo(about.Version, newVersion, releaseDate)
		}
	}

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func renderCommands(info *Info) {
	maxSize := getMaxCommandSize(info.Commands)

	var curGroup string

	for _, command := range info.Commands {
		if curGroup != command.Group {
			printGroupHeader(command.Group)
			curGroup = command.Group
		}

		fmtc.Printf("  "+info.CommandsColorTag+"%s{!}", command.Name)

		if len(command.Args) != 0 {
			fmtc.Printf(" " + renderArgs(command.Args))
		}

		size := getItemSize(command.Name, command.Args)

		fmtc.Printf(getSeparator(size, maxSize, info.Breadcrumbs))
		fmtc.Printf(command.Desc)

		fmtc.NewLine()
	}
}

func renderOptions(info *Info) {
	maxSize := getMaxOptionSize(info.Options)

	printGroupHeader("Options")

	for _, option := range info.Options {
		fmtc.Printf("  "+info.OptionsColorTag+"%s{!}", formatOptionName(option.Name))

		if len(option.Args) != 0 {
			fmtc.Printf(" " + renderArgs(option.Args))
		}

		size := getItemSize(option.Name, option.Args)

		fmtc.Printf(getSeparator(size, maxSize, info.Breadcrumbs))
		fmtc.Printf(option.Desc)

		fmtc.NewLine()
	}
}

// renderExamples render examples
func renderExamples(info *Info) {
	printGroupHeader("Examples")

	total := len(info.Examples)

	for index, example := range info.Examples {
		fmtc.Printf("  %s %s\n", info.Name, example.Cmd)

		if example.Desc != "" {
			fmtc.Printf("  {s-}%s{!}\n", example.Desc)
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

// formatOptionName format option name
func formatOptionName(name string) string {
	if strings.Contains(name, ":") {
		return "--" + strutil.ReadField(name, 1, false, ":") +
			", -" + strutil.ReadField(name, 0, false, ":")
	}

	return "--" + name
}

// getSeparator return bread crumbs (or spaces if colors are disabled) for
// item name aligning
func getSeparator(size, maxSize int, breadcrumbs bool) string {
	if breadcrumbs && !fmtc.DisableColors && maxSize > _BREADCRUMBS_MIN_SIZE {
		return " {s-}" + _DOTS[:maxSize-size] + "{!} "
	}

	return " " + _SPACES[:maxSize-size] + " "
}

// getMaxCommandSize returns the biggest command size
func getMaxCommandSize(commands []*Command) int {
	var size int

	for _, command := range commands {
		size = mathutil.Max(size, getItemSize(command.Name, command.Args)+2)
	}

	return size
}

// getMaxOptionSize returns the biggest option size
func getMaxOptionSize(options []*Option) int {
	var size int

	for _, option := range options {
		size = mathutil.Max(size, getItemSize(option.Name, option.Args)+2)
	}

	return size
}

// getItemSize calculate rendered item size
func getItemSize(name string, args []string) int {
	var size int

	if strings.Contains(name, ":") {
		size += strutil.Len(name) + 4
	} else {
		size += strutil.Len(name) + 2
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "?") {
			size += strutil.Len(arg)
		} else {
			size += strutil.Len(arg) + 1
		}
	}

	return size
}

// printGroupHeader print category header
func printGroupHeader(name string) {
	fmtc.Printf("\n{*}%s{!}\n\n", name)
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

// printNewVersionInfo print info about latest release
func printNewVersionInfo(curVersion, newVersion string, releaseDate time.Time) {
	cv, err := version.Parse(curVersion)

	if err != nil {
		return
	}

	nv, err := version.Parse(newVersion)

	if err != nil {
		return
	}

	days := int(time.Since(releaseDate) / (time.Hour * 24))

	colorTag := "{s}"

	switch {
	case cv.Major() != nv.Major():
		colorTag = "{r}"
	case cv.Minor() != nv.Minor():
		colorTag = "{y}"
	}

	fmtc.NewLine()
	fmtc.Printf(colorTag+"Latest version is %s{!} ", newVersion)

	switch days {
	case 0:
		fmtc.Println("{s-}(released today){!}")
	case 1:
		fmtc.Println("{s-}(released 1 day ago){!}")
	default:
		fmtc.Printf("{s-}(released %d days ago){!}\n", days)
	}
}
