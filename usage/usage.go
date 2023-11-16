// Package usage provides methods and structs for generating usage info for
// command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/mathutil"
	"github.com/essentialkaos/ek/v12/strutil"
	"github.com/essentialkaos/ek/v12/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SPACES = "                                                                "
	_DOTS   = "................................................................"
)

const _BREADCRUMBS_MIN_SIZE = 8

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DEFAULT_COMMANDS_COLOR_TAG  = "{y}"
	DEFAULT_OPTIONS_COLOR_TAG   = "{g}"
	DEFAULT_APP_NAME_COLOR_TAG  = "{c*}"
	DEFAULT_APP_VER_COLOR_TAG   = "{c}"
	DEFAULT_APP_REL_COLOR_TAG   = "{s}"
	DEFAULT_APP_BUILD_COLOR_TAG = "{s-}"
)

const (
	VERSION_FULL    = "full"
	VERSION_SIMPLE  = "simple"
	VERSION_MAJOR   = "major"
	VERSION_MINOR   = "minor"
	VERSION_PATCH   = "patch"
	VERSION_RELEASE = "release"
	VERSION_BUILD   = "build"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Environment []EnvironmentInfo

type EnvironmentInfo struct {
	Name    string
	Version string
}

// About contains info about app
type About struct {
	App     string // App is app name
	Version string // Version is current app version in semver notation
	Release string // Release is current app release
	Build   string // Build is current app build
	Desc    string // Desc is short info about app
	Year    int    // Year is year when owner company was founded
	License string // License is name of license
	Owner   string // Owner is name of owner (company/developer)

	AppNameColorTag string // AppNameColorTag contains default app name color tag
	VersionColorTag string // VersionColorTag contains default app version color tag
	ReleaseColorTag string // ReleaseColorTag contains default app release color tag
	BuildColorTag   string // BuildColorTag contains default app build color tag

	ReleaseSeparator string // ReleaseSeparator contains symbol for version and release separation (default: -)
	DescSeparator    string // DescSeparator contains symbol for version and description separation (default: -)

	Environment Environment // Environment contains info about environment

	// Function for checking app updates
	UpdateChecker UpdateChecker
}

// Info contains info about commands, options, and examples
type Info struct {
	AppNameColorTag  string // AppNameColorTag contains default app name color tag
	CommandsColorTag string // CommandsColorTag contains default commands color tag
	OptionsColorTag  string // OptionsColorTag contains default options color tag
	Breadcrumbs      bool   // Breadcrumbs is flag for using bread crumbs for commands and options output

	Name    string   // Name is app name
	Args    []string // Args is slice with app arguments
	Spoiler string   // Spoiler contains additional info

	Commands []*Command // Commands is list of supported commands
	Options  []*Option  // Options is list of supported options
	Examples []*Example // Examples is list of usage examples

	curGroup string
}

// UpdateChecker is a base for all update checkers
type UpdateChecker struct {
	Payload   string
	CheckFunc func(app, version, data string) (string, time.Time, bool)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Command contains info about supported command
type Command struct {
	Name         string   // Name is command name
	Desc         string   // Desc is command description
	Group        string   // Group is command group name
	Args         []string // Args is slice with arguments
	BoundOptions []string // BoundOptions is slice with long names of related options

	ColorTag string // ColorTag contains default color tag

	info *Info
}

// Option contains info about supported option
type Option struct {
	Short string // Short is short option name (with one minus prefix)
	Long  string // Long is long option name (with two minuses prefix)
	Desc  string // Desc is option description
	Arg   string // Arg is option argument

	ColorTag string // ColorTag contains default color tag

	info *Info
}

// Example contains usage example
type Example struct {
	Cmd  string // Cmd is command usage example
	Desc string // Desc is usage description
	Raw  bool   // Raw is raw example flag (without automatic binary name appending)

	info *Info
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewInfo creates new info struct
func NewInfo(args ...string) *Info {
	var name string

	if len(args) != 0 {
		name = args[0]
		args = args[1:]
	}

	name = strutil.Q(name, filepath.Base(os.Args[0]))

	info := &Info{
		Name: name,
		Args: args,

		CommandsColorTag: DEFAULT_COMMANDS_COLOR_TAG,
		OptionsColorTag:  DEFAULT_OPTIONS_COLOR_TAG,
		Breadcrumbs:      true,
	}

	return info
}

// AddGroup adds new command group
func (i *Info) AddGroup(group string) {
	if i == nil {
		return
	}

	i.curGroup = group
}

// AddCommand adds command (name, description, args)
func (i *Info) AddCommand(a ...string) *Command {
	if i == nil || len(a) < 2 {
		return nil
	}

	group := "Commands"

	if i.curGroup != "" {
		group = i.curGroup
	}

	cmd := &Command{
		Name:  a[0],
		Desc:  a[1],
		Args:  a[2:],
		Group: group,
		info:  i,
	}

	i.Commands = append(i.Commands, cmd)

	return cmd
}

// AddOption adds option (name, description, args)
func (i *Info) AddOption(a ...string) *Option {
	if i == nil || len(a) < 2 {
		return nil
	}

	long, short := parseOptionName(a[0])
	opt := &Option{
		Long:  long,
		Short: short,
		Desc:  a[1],
		Arg:   strings.Join(a[2:], " "),
		info:  i,
	}

	i.Options = append(i.Options, opt)

	return opt
}

// AddExample adds example of application usage
func (i *Info) AddExample(a ...string) {
	if i == nil || len(a) == 0 {
		return
	}

	a = append(a, "")

	i.Examples = append(i.Examples, &Example{a[0], a[1], false, i})
}

// AddRawExample adds example of application usage without command prefix
func (i *Info) AddRawExample(a ...string) {
	if i == nil || len(a) == 0 {
		return
	}

	a = append(a, "")

	i.Examples = append(i.Examples, &Example{a[0], a[1], true, i})
}

// AddSpoiler adds spoiler
func (i *Info) AddSpoiler(spoiler string) {
	if i == nil {
		return
	}

	i.Spoiler = spoiler
}

// BoundOptions bounds command with options
func (i *Info) BoundOptions(cmd string, options ...string) {
	if i == nil || cmd == "" {
		return
	}

	for _, command := range i.Commands {
		if command.Name == cmd {
			for _, opt := range options {
				longOption, _ := parseOptionName(opt)
				command.BoundOptions = append(command.BoundOptions, longOption)
			}

			return
		}
	}
}

// GetCommand tries to find command with given name
func (i *Info) GetCommand(name string) *Command {
	if i == nil {
		return nil
	}

	for _, command := range i.Commands {
		if command.Name == name {
			return command
		}
	}

	return nil
}

// GetOption tries to find option with given name
func (i *Info) GetOption(name string) *Option {
	if i == nil {
		return nil
	}

	name, _ = parseOptionName(name)

	for _, option := range i.Options {
		if option.Long == name {
			return option
		}
	}

	return nil
}

// Render prints usage info
//
// Deprecated: Use method Print instead
func (i *Info) Render() {
	i.Print()
}

// Print prints usage info
func (i *Info) Print() {
	if i == nil {
		return
	}

	appNameColorTag := strutil.B(fmtc.IsTag(i.AppNameColorTag), i.AppNameColorTag, DEFAULT_APP_NAME_COLOR_TAG)
	optionsColorTag := strutil.B(fmtc.IsTag(i.OptionsColorTag), i.OptionsColorTag, DEFAULT_OPTIONS_COLOR_TAG)
	commandsColorTag := strutil.B(fmtc.IsTag(i.CommandsColorTag), i.CommandsColorTag, DEFAULT_COMMANDS_COLOR_TAG)

	usageMessage := "\n{*}Usage:{!} " + appNameColorTag + i.Name + "{!}"

	if len(i.Options) != 0 {
		usageMessage += " " + optionsColorTag + "{options}{!}"
	}

	if len(i.Commands) != 0 {
		usageMessage += " " + commandsColorTag + "{command}{!}"
	}

	if len(i.Args) != 0 {
		usageMessage += " " + strings.Join(i.Args, " ")
	}

	fmtc.Println(usageMessage)

	if i.Spoiler != "" {
		fmtc.NewLine()
		fmtc.Println(i.Spoiler)
	}

	if len(i.Commands) != 0 {
		printCommands(i)
	}

	if len(i.Options) != 0 {
		printOptions(i)
	}

	if len(i.Examples) != 0 {
		printExamples(i)
	}

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns a string representation of the command
func (c *Command) String() string {
	if c == nil {
		return ""
	}

	return c.Name
}

// String returns a string representation of the option
func (o *Option) String() string {
	if o == nil {
		return ""
	}

	return "--" + o.Long
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Render prints info about command
//
// Deprecated: Use method Print instead
func (c *Command) Render() {
	c.Print()
}

// Render prints info about option
//
// Deprecated: Use method Print instead
func (o *Option) Render() {
	o.Print()
}

// Render prints usage example
//
// Deprecated: Use method Print instead
func (e *Example) Render() {
	e.Print()
}

// Print prints info about command
func (c *Command) Print() {
	if c == nil {
		return
	}

	size := getCommandSize(c)
	useBreadcrumbs := true
	maxSize := size

	colorTag := strutil.Q(strutil.B(fmtc.IsTag(c.ColorTag), c.ColorTag, DEFAULT_COMMANDS_COLOR_TAG), DEFAULT_COMMANDS_COLOR_TAG)

	if c.info != nil {
		colorTag = c.info.CommandsColorTag
		maxSize = getMaxCommandSize(c.info.Commands, c.Group)
		useBreadcrumbs = c.info.Breadcrumbs
	}

	fmtc.Printf("  "+colorTag+"%s{!}", c.Name)

	if len(c.Args) != 0 {
		fmtc.Print(" " + printArgs(c.Args...))
	}

	fmtc.Print(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Print(c.Desc)

	fmtc.NewLine()
}

// Print prints info about option
func (o *Option) Print() {
	if o == nil {
		return
	}

	size := getOptionSize(o)
	useBreadcrumbs := true
	maxSize := size

	colorTag := strutil.Q(strutil.B(fmtc.IsTag(o.ColorTag), o.ColorTag, DEFAULT_OPTIONS_COLOR_TAG), DEFAULT_OPTIONS_COLOR_TAG)

	if o.info != nil {
		colorTag = o.info.OptionsColorTag
		maxSize = getMaxOptionSize(o.info.Options)
		useBreadcrumbs = o.info.Breadcrumbs
	}

	fmtc.Printf("  " + formatOptionName(o, colorTag))

	if o.Arg != "" {
		fmtc.Print(" " + printArgs(o.Arg))
	}

	fmtc.Print(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Print(o.Desc)

	fmtc.NewLine()
}

// Print prints usage example
func (e *Example) Print() {
	if e == nil {
		return
	}

	appName := os.Args[0]

	if e.info != nil {
		appName = e.info.Name
	}

	if e.Raw {
		fmtc.Printf("  %s\n", e.Cmd)
	} else {
		fmtc.Printf("  %s %s\n", appName, e.Cmd)
	}

	if e.Desc != "" {
		fmtc.Printf("  {s-}%s{!}\n", e.Desc)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Render prints version info
//
// Deprecated: Use method Print instead
func (a *About) Render() {
	a.Print()
}

// Print prints version info
func (a *About) Print(infoType ...string) {
	if a == nil {
		return
	}

	if len(infoType) != 0 {
		ver := getRawVersion(a, infoType[0])

		if ver != "" {
			fmtc.Println(ver)
			return
		}
	}

	nameColor := strutil.Q(strutil.B(fmtc.IsTag(a.AppNameColorTag), a.AppNameColorTag, DEFAULT_APP_NAME_COLOR_TAG), DEFAULT_APP_NAME_COLOR_TAG)
	versionColor := strutil.Q(strutil.B(fmtc.IsTag(a.VersionColorTag), a.VersionColorTag, DEFAULT_APP_VER_COLOR_TAG), DEFAULT_APP_VER_COLOR_TAG)
	releaseColor := strutil.Q(strutil.B(fmtc.IsTag(a.ReleaseColorTag), a.ReleaseColorTag, DEFAULT_APP_REL_COLOR_TAG), DEFAULT_APP_REL_COLOR_TAG)
	buildColor := strutil.Q(strutil.B(fmtc.IsTag(a.BuildColorTag), a.BuildColorTag, DEFAULT_APP_BUILD_COLOR_TAG), DEFAULT_APP_BUILD_COLOR_TAG)
	relSeparator := strutil.Q(a.ReleaseSeparator, "-")
	descSeparator := strutil.Q(a.DescSeparator, "-")

	fmtc.Printf("\n"+nameColor+"%s{!} "+versionColor+"%s{!}", a.App, a.Version)

	fmtc.If(a.Release != "").Printf(releaseColor+relSeparator+"%s{!}", a.Release)
	fmtc.If(a.Build != "").Printf(buildColor+" (%s){!}", a.Build)

	fmtc.Printf(" "+descSeparator+" %s\n", a.Desc)

	if len(a.Environment) > 0 {
		fmtc.Printf("{s-}│{!}\n")

		for i, env := range a.Environment {
			if len(a.Environment) != i+1 {
				fmtc.Printf("{s-}├ %s: %s{!}\n", env.Name, env.Version)
			} else {
				fmtc.Printf("{s-}└ %s: %s{!}\n", env.Name, env.Version)
			}
		}
	}

	fmtc.NewLine()

	if a.Owner != "" {
		fmtc.If(a.Year == 0).Printf(
			"{s-}Copyright (C) %d %s{!}\n",
			time.Now().Year(), a.Owner,
		)

		fmtc.If(a.Year != 0).Printf(
			"{s-}Copyright (C) %d-%d %s{!}\n",
			a.Year, time.Now().Year(), a.Owner,
		)
	}

	fmtc.If(a.License != "").Printf("{s-}%s{!}\n", a.License)

	if a.UpdateChecker.CheckFunc != nil && a.UpdateChecker.Payload != "" {
		newVersion, releaseDate, hasUpdate := a.UpdateChecker.CheckFunc(
			a.App, a.Version, a.UpdateChecker.Payload,
		)

		if hasUpdate && isNewerVersion(a.Version, newVersion) {
			printNewVersionInfo(a.Version, newVersion, releaseDate)
		}
	}

	fmtc.If(a.Owner != "" || a.License != "").NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// printCommands prints all supported commands
func printCommands(info *Info) {
	var curGroup string

	for _, command := range info.Commands {
		if curGroup != command.Group {
			printGroupHeader(command.Group)
			curGroup = command.Group
		}

		command.Print()
	}
}

// printOptions prints all supported options
func printOptions(info *Info) {
	printGroupHeader("Options")

	for _, option := range info.Options {
		option.Print()
	}
}

// printExamples prints all usage examples
func printExamples(info *Info) {
	printGroupHeader("Examples")

	total := len(info.Examples)

	for index, example := range info.Examples {
		example.Print()

		if index < total-1 {
			fmtc.NewLine()
		}
	}
}

// printArgs prints arguments with colors
func printArgs(args ...string) string {
	var result string

	for _, a := range args {
		if strings.HasPrefix(a, "?") {
			result += "{s-}" + a[1:] + "{!} "
		} else {
			result += "{s}" + a + "{!} "
		}
	}

	return fmtc.Sprint(strings.TrimRight(result, " "))
}

// getRawVersion prints raw (just numbers) version info
func getRawVersion(about *About, infoType string) string {
	switch infoType {
	case VERSION_FULL:
		return about.Version +
			fmtc.If(about.Release != "").Sprint("-"+about.Release) +
			fmtc.If(about.Build != "").Sprint("+"+about.Build)
	case VERSION_SIMPLE:
		return about.Version
	case VERSION_RELEASE:
		return about.Release
	case VERSION_BUILD:
		return about.Build
	}

	ver, _ := version.Parse(about.Version)

	if ver.IsZero() {
		return ""
	}

	switch infoType {
	case VERSION_MAJOR:
		return fmtc.Sprintf("%d", ver.Major())
	case VERSION_MINOR:
		return fmtc.Sprintf("%d", ver.Minor())
	case VERSION_PATCH:
		return fmtc.Sprintf("%d", ver.Patch())
	}

	return ""
}

// formatOptionName formats option name
func formatOptionName(opt *Option, colorTag string) string {
	if opt.Short != "" {
		return fmt.Sprintf(
			"%s--%s{s}, %s-%s{!}",
			colorTag, opt.Long, colorTag, opt.Short,
		)
	}

	return colorTag + "--" + opt.Long + "{!}"
}

// parseOptionName parses option name
func parseOptionName(name string) (string, string) {
	if strings.Contains(name, ":") {
		return strutil.ReadField(name, 1, false, ":"),
			strutil.ReadField(name, 0, false, ":")
	}

	return name, ""
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
func getMaxCommandSize(commands []*Command, group string) int {
	var size int

	for _, command := range commands {
		if group != "" && command.Group != group {
			continue
		}

		size = mathutil.Max(size, getCommandSize(command)+2)
	}

	return size
}

// getMaxOptionSize returns the biggest option size
func getMaxOptionSize(options []*Option) int {
	var size int

	for _, option := range options {
		size = mathutil.Max(size, getOptionSize(option)+2)
	}

	return size
}

// getOptionSize calculates final command size
func getCommandSize(cmd *Command) int {
	size := strutil.Len(cmd.Name) + 2

	for _, arg := range cmd.Args {
		if strings.HasPrefix(arg, "?") {
			size += strutil.Len(arg)
		} else {
			size += strutil.Len(arg) + 1
		}
	}

	return size
}

// getOptionSize calculates final option size
func getOptionSize(opt *Option) int {
	var size int

	if opt.Short != "" {
		size += strutil.Len(opt.Long) + strutil.Len(opt.Short) + 4
	} else {
		size += strutil.Len(opt.Long) + 1
	}

	if opt.Arg != "" {
		size += strutil.Len(opt.Arg)

		if !strings.HasPrefix(opt.Arg, "?") {
			size++
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
