// Package usage provides methods and structs for generating usage info for
// command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v14/fmtc"
	"github.com/essentialkaos/ek/v14/strutil"
	"github.com/essentialkaos/ek/v14/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SPACES = "                                                                "
	_DOTS   = "................................................................"
)

const _BREADCRUMBS_MIN_SIZE = 8

const _DEFAULT_WRAP_LEN = 88

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	// DEFAULT_COMMANDS_COLOR_TAG is the default fmtc color tag for command names
	DEFAULT_COMMANDS_COLOR_TAG = "{y}"

	// DEFAULT_OPTIONS_COLOR_TAG is the default fmtc color tag for option names
	DEFAULT_OPTIONS_COLOR_TAG = "{g}"

	// DEFAULT_ENV_VAR_COLOR_TAG is the default fmtc color tag for environment variable
	// names
	DEFAULT_ENV_VAR_COLOR_TAG = "{m}"

	// DEFAULT_EXAMPLE_DESC_COLOR_TAG is the default fmtc color tag for example
	// descriptions
	DEFAULT_EXAMPLE_DESC_COLOR_TAG = "{&}{s-}"

	// DEFAULT_APP_NAME_COLOR_TAG is the default fmtc color tag for the application name
	DEFAULT_APP_NAME_COLOR_TAG = "{c*}"

	// DEFAULT_APP_VER_COLOR_TAG is the default fmtc color tag for the application
	// version
	DEFAULT_APP_VER_COLOR_TAG = "{c}"

	// DEFAULT_APP_REL_COLOR_TAG is the default fmtc color tag for the application
	// release
	DEFAULT_APP_REL_COLOR_TAG = "{s}"

	// DEFAULT_APP_BUILD_COLOR_TAG is the default fmtc color tag for the application
	// build
	DEFAULT_APP_BUILD_COLOR_TAG = "{s-}"
)

const (
	// VERSION_FULL returns the full version string including release and build
	VERSION_FULL = "full"

	// VERSION_SIMPLE returns only the version number (e.g. "1.2.3")
	VERSION_SIMPLE = "simple"

	// VERSION_MAJOR returns only the major version number
	VERSION_MAJOR = "major"

	// VERSION_MINOR returns only the minor version number
	VERSION_MINOR = "minor"

	// VERSION_PATCH returns only the patch version number
	VERSION_PATCH = "patch"

	// VERSION_RELEASE returns only the release label
	VERSION_RELEASE = "release"

	// VERSION_BUILD returns only the build metadata
	VERSION_BUILD = "build"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Environment is a list of environment variables shown in the application about section
type Environment []EnvironmentInfo

// EnvironmentInfo holds the name and version of a runtime environment dependency
type EnvironmentInfo struct {
	Name    string
	Version string
}

// About holds application metadata printed by the version/about command
type About struct {
	App        string // Application name
	Version    string // Application version in semver notation
	Release    string // Release label (e.g. "β4")
	Build      string // Build metadata (e.g. a commit hash)
	Desc       string // Short one-line description of the application
	Year       int    // Year the owner company or project was founded
	License    string // License name (e.g. "Apache 2.0")
	Owner      string // Name of the owner company or developer
	BugTracker string // URL of the issue tracker

	AppNameColorTag string // fmtc color tag for the app name; defaults to [DEFAULT_APP_NAME_COLOR_TAG]
	VersionColorTag string // fmtc color tag for the version; defaults to [DEFAULT_APP_VER_COLOR_TAG]
	ReleaseColorTag string // fmtc color tag for the release label; defaults to [DEFAULT_APP_REL_COLOR_TAG]
	BuildColorTag   string // fmtc color tag for the build metadata; defaults to [DEFAULT_APP_BUILD_COLOR_TAG]

	Copyright string // Custom copyright prefix; defaults to "Copyright (C)"

	ReleaseSeparator string // Separator between version and release (default: "-"
	DescSeparator    string // Separator between version and description (default: "-")

	Environment Environment // Runtime environment dependencies shown below the version line

	UpdateChecker UpdateChecker // Optional checker for announcing newer releases
}

// Info holds the full usage information for a command-line application,
// including its commands, options, environment variables, and examples
type Info struct {
	AppNameColorTag     string // fmtc color tag for the app name; defaults to [DEFAULT_APP_NAME_COLOR_TAG]
	CommandsColorTag    string // fmtc color tag for command names; defaults to DEFAULT_COMMANDS_COLOR_TAG
	CommandsHeader      string // Custom header label for the commands section (default: "Commands")
	OptionsColorTag     string // fmtc color tag for option names; defaults to [DEFAULT_OPTIONS_COLOR_TAG]
	OptionsHeader       string // Custom header label for the options section (default: "Options")
	EnvVarsColorTag     string // fmtc color tag for env var names; defaults to [DEFAULT_ENV_VAR_COLOR_TAG]
	EnvVarsHeader       string // Custom header label for the env vars section (default: "Environment variables")
	ExamplesHeader      string // Custom header label for the examples section (default: "Examples")
	ExampleDescColorTag string // fmtc color tag for example descriptions; defaults to [DEFAULT_EXAMPLE_DESC_COLOR_TAG]
	UsageHeader         string // Custom header label for the usage line (default: "Usage")

	CommandPlaceholder string // Placeholder text for the command token in the usage line (default: "{command}")
	OptionsPlaceholder string // Placeholder text for the options token in the usage line (default: "{options}")

	Breadcrumbs bool // When true, dots are used to align descriptions; enabled by default
	WrapLen     int  // Column at which long descriptions are wrapped; defaults to 88

	Name    string   // Application name shown in the usage line
	Args    []string // Positional argument placeholders shown after options and commands
	Spoiler string   // Optional free-form text printed below the usage line

	Commands []*Command // Registered commands
	Options  []*Option  // Registered options
	EnvVars  []*Env     // Registered environment variables
	Examples []*Example // Registered usage examples

	curGroup string
}

// UpdateChecker carries the data and logic needed to check for a newer release
type UpdateChecker struct {
	// Arbitrary data passed to CheckFunc (e.g. a repo slug or API key)
	Payload string

	// Returns the latest version, its release date, and whether an update is available
	CheckFunc func(app, version, data string) (string, time.Time, bool)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Command describes a single CLI command with its arguments and related options
type Command struct {
	Name         string   // Command name as typed by the user
	Desc         string   // Short description shown in the commands list
	Group        string   // Section header this command belongs to
	Args         []string // Argument placeholders; prefix with "?" to mark as optional
	BoundOptions []string // Long option names shown as related when this command is printed

	ColorTag string // fmtc color tag override; falls back to Info.CommandsColorTag

	info *Info
}

// Option describes a single CLI option with its short and long forms
type Option struct {
	Short string // Short form without the leading dash (e.g. "v" for -v)
	Long  string // Long form without the leading dashes (e.g. "verbose" for --verbose)
	Desc  string // Short description shown in the options list
	Arg   string // Argument placeholder; prefix with "?" to mark as optional

	ColorTag string // fmtc color tag override; falls back to Info.OptionsColorTag

	info *Info
}

// Env describes a single environment variable recognised by the application
type Env struct {
	Name string // Environment variable name (e.g. "GOMAXPROCS")
	Desc string // Short description shown in the env vars list

	ColorTag string // fmtc color tag override; falls back to Info.EnvVarsColorTag

	info *Info
}

// Example holds a single usage example shown in the examples section
type Example struct {
	Cmd   string // Command string appended to the app name (or the full line when Raw is true)
	Desc  string // Optional description printed below the example command
	IsRaw bool   // Raw disables automatic prepending of the application name to Cmd

	info *Info
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewInfo creates an Info struct for the current application. The first optional
// argument is the app name; remaining arguments are positional argument placeholders
func NewInfo(args ...any) *Info {
	var name string

	if len(args) != 0 {
		name = evalString(args[0])
		args = args[1:]
	}

	name = strutil.Q(name, filepath.Base(os.Args[0]))

	info := &Info{
		Name:        name,
		Args:        evalStrings(args),
		Breadcrumbs: true,
	}

	return info
}

// AddGroup sets the group label applied to all subsequently added commands
func (i *Info) AddGroup(group any) {
	groupName := evalString(group)

	if i == nil || groupName == "" {
		return
	}

	i.curGroup = evalString(groupName)
}

// AddCommand registers a new command under the current group. Optional args are
// positional argument placeholders; prefix with "?" to mark as optional
func (i *Info) AddCommand(name, desc any, args ...any) *Command {
	if i == nil || name == "" || desc == "" {
		return nil
	}

	group := strutil.Q(i.CommandsHeader, "Commands")

	if i.curGroup != "" {
		group = i.curGroup
	}

	cmd := &Command{
		Name:  evalString(name),
		Desc:  evalString(desc),
		Args:  evalStrings(args),
		Group: group,
		info:  i,
	}

	i.Commands = append(i.Commands, cmd)

	return cmd
}

// AddOption registers a new option. Name must be in "short:long" or "long" format.
// Optional args define the option's argument placeholder; prefix with "?" to mark
// as optional.
func (i *Info) AddOption(name, desc any, args ...any) *Option {
	if i == nil || name == "" || desc == "" {
		return nil
	}

	long, short := parseOptionName(evalString(name))

	opt := &Option{
		Long:  long,
		Short: short,
		Desc:  evalString(desc),
		Arg:   strings.Join(evalStrings(args), " "),
		info:  i,
	}

	i.Options = append(i.Options, opt)

	return opt
}

// AddEnv registers an environment variable with its description
func (i *Info) AddEnv(name, desc any) *Env {
	if i == nil || name == "" || desc == "" {
		return nil
	}

	env := &Env{
		Name: evalString(name),
		Desc: evalString(desc),
		info: i,
	}

	i.EnvVars = append(i.EnvVars, env)

	return env
}

// AddExample registers a usage example, automatically prepending the application name
func (i *Info) AddExample(cmd any, desc ...any) {
	if i == nil || cmd == "" {
		return
	}

	var cmdDesc string

	if len(desc) != 0 {
		cmdDesc = evalString(desc[0])
	}

	i.Examples = append(i.Examples, &Example{evalString(cmd), cmdDesc, false, i})
}

// AddRawExample registers a usage example printed exactly as given, without the
// application name prefix
func (i *Info) AddRawExample(cmd any, desc ...any) {
	if i == nil || cmd == "" {
		return
	}

	var cmdDesc string

	if len(desc) != 0 {
		cmdDesc = evalString(desc[0])
	}

	i.Examples = append(i.Examples, &Example{evalString(cmd), cmdDesc, true, i})
}

// AddSpoiler sets the free-form text printed below the usage line
func (i *Info) AddSpoiler(spoiler any) {
	if i == nil {
		return
	}

	i.Spoiler = evalString(spoiler)
}

// BoundOptions associates the named options with a command so they are shown as
// related. Option names follow the same "short:long" or "long" format used in
// AddOption.
func (i *Info) BoundOptions(cmd any, options ...any) {
	if i == nil || cmd == "" {
		return
	}

	for _, command := range i.Commands {
		if command.Name == cmd {
			for _, opt := range options {
				longOption, _ := parseOptionName(evalString(opt))
				command.BoundOptions = append(command.BoundOptions, longOption)
			}

			return
		}
	}
}

// GetCommand returns the registered command with the given name, or nil if not found
func (i *Info) GetCommand(name any) *Command {
	if i == nil {
		return nil
	}

	cmdName := evalString(name)

	for _, command := range i.Commands {
		if command.Name == cmdName {
			return command
		}
	}

	return nil
}

// GetOption returns the registered option with the given long name, or nil if not found
func (i *Info) GetOption(name any) *Option {
	if i == nil {
		return nil
	}

	optName, _ := parseOptionName(evalString(name))

	for _, option := range i.Options {
		if option.Long == optName {
			return option
		}
	}

	return nil
}

// Print renders the full usage information to stdout
func (i *Info) Print() {
	if i == nil {
		return
	}

	appNameColorTag := strutil.B(
		i.AppNameColorTag != "" && fmtc.IsTag(i.AppNameColorTag),
		i.AppNameColorTag, DEFAULT_APP_NAME_COLOR_TAG,
	)

	optionsColorTag := strutil.B(
		i.OptionsColorTag != "" && fmtc.IsTag(i.OptionsColorTag),
		i.OptionsColorTag, DEFAULT_OPTIONS_COLOR_TAG,
	)

	commandsColorTag := strutil.B(
		i.CommandsColorTag != "" && fmtc.IsTag(i.CommandsColorTag),
		i.CommandsColorTag, DEFAULT_COMMANDS_COLOR_TAG,
	)

	usageMessage := "\n{*}" + strutil.Q(i.UsageHeader, "Usage") + ":{!} "
	usageMessage += appNameColorTag + i.Name + "{!}"

	if len(i.Options) != 0 {
		usageMessage += " " + optionsColorTag + strutil.Q(i.OptionsPlaceholder, "{options}") + "{!}"
	}

	if len(i.Commands) != 0 {
		usageMessage += " " + commandsColorTag + strutil.Q(i.CommandPlaceholder, "{command}") + "{!}"
	}

	if len(i.Args) != 0 {
		usageMessage += " " + printArgs(i.Args...)
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

	if len(i.EnvVars) != 0 {
		printEnvVars(i)
	}

	if len(i.Examples) != 0 {
		printExamples(i)
	}

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns the command name, satisfying fmt.Stringer
func (c *Command) String() string {
	if c == nil {
		return ""
	}

	return c.Name
}

// String returns the option in "--long" form, satisfying fmt.Stringer
func (o *Option) String() string {
	if o == nil {
		return ""
	}

	return "--" + o.Long
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Print renders this command as a single formatted line aligned with its group peers
func (c *Command) Print() {
	if c == nil {
		return
	}

	size := getCommandSize(c)
	useBreadcrumbs := true
	maxSize := size
	wrapLen := _DEFAULT_WRAP_LEN

	colorTag := strutil.Q(
		strutil.B(c.ColorTag != "" && fmtc.IsTag(c.ColorTag), c.ColorTag, ""),
		DEFAULT_COMMANDS_COLOR_TAG,
	)

	if c.info != nil {
		colorTag = strutil.Q(c.info.CommandsColorTag, colorTag)
		maxSize = getMaxCommandSize(c.info.Commands, c.Group)
		useBreadcrumbs = c.info.Breadcrumbs
		wrapLen = c.info.WrapLen
	}

	fmtc.Printf("  "+colorTag+"%s{!}", c.Name)

	if len(c.Args) != 0 {
		fmtc.Print(" " + printArgs(c.Args...))
	}

	fmtc.Print(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Print(wrapText(c.Desc, maxSize+2, wrapLen))

	fmtc.NewLine()
}

// Print renders this option as a single formatted line aligned with its section peers
func (o *Option) Print() {
	if o == nil {
		return
	}

	size := getOptionSize(o)
	useBreadcrumbs := true
	maxSize := size
	wrapLen := _DEFAULT_WRAP_LEN

	colorTag := strutil.Q(
		strutil.B(o.ColorTag != "" && fmtc.IsTag(o.ColorTag), o.ColorTag, ""),
		DEFAULT_OPTIONS_COLOR_TAG,
	)

	if o.info != nil {
		colorTag = strutil.Q(o.info.OptionsColorTag, colorTag)
		maxSize = getMaxOptionSize(o.info.Options)
		useBreadcrumbs = o.info.Breadcrumbs
		wrapLen = o.info.WrapLen
	}

	fmtc.Printf("  " + formatOptionName(o, colorTag))

	if o.Arg != "" {
		fmtc.Print(" " + printArgs(o.Arg))
	}

	fmtc.Print(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Print(wrapText(o.Desc, maxSize+5, wrapLen))

	fmtc.NewLine()
}

// Print renders this environment variable as a single formatted line. The name is
// shown in bold when the variable is currently set in the process environment.
func (e *Env) Print() {
	if e == nil {
		return
	}

	size := strutil.Len(e.Name)
	useBreadcrumbs := true
	maxSize := size
	wrapLen := _DEFAULT_WRAP_LEN

	colorTag := strutil.Q(
		strutil.B(e.ColorTag != "" && fmtc.IsTag(e.ColorTag), e.ColorTag, ""),
		DEFAULT_ENV_VAR_COLOR_TAG,
	)

	if e.info != nil {
		colorTag = strutil.Q(e.info.EnvVarsColorTag, colorTag)
		maxSize = getMaxEnvVarSize(e.info.EnvVars)
		useBreadcrumbs = e.info.Breadcrumbs
		wrapLen = e.info.WrapLen
	}

	if os.Getenv(e.Name) != "" {
		fmtc.Printf("  "+colorTag+"{*}%s{!}", e.Name)
	} else {
		fmtc.Printf("  "+colorTag+"%s{!}", e.Name)
	}

	fmtc.Print(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Print(wrapText(e.Desc, maxSize+2, wrapLen))

	fmtc.NewLine()
}

// Print renders this example, optionally prefixed with the application name.
// If Desc is set, it is printed on a second line below the command.
func (e *Example) Print() {
	if e == nil {
		return
	}

	appName := os.Args[0]
	wrapLen := _DEFAULT_WRAP_LEN

	if e.info != nil {
		appName = e.info.Name
		wrapLen = e.info.WrapLen
	}

	if e.IsRaw {
		fmtc.Printfn("  %s", e.Cmd)
	} else {
		fmtc.Printfn("  %s %s", appName, e.Cmd)
	}

	if e.Desc != "" {
		descColor := DEFAULT_EXAMPLE_DESC_COLOR_TAG

		if e.info != nil {
			descColor = strutil.B(
				e.info.ExampleDescColorTag != "" && fmtc.IsTag(e.info.ExampleDescColorTag),
				e.info.ExampleDescColorTag, "",
			)
		}

		fmtc.Printfn("  "+descColor+"%s{!}", wrapText(e.Desc, 2, wrapLen))
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Print renders the application about/version block to stdout. An optional infoType
// argument (VERSION_FULL, VERSION_MAJOR, etc.) prints only that version component.
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

	nameColor := strutil.Q(
		strutil.B(a.AppNameColorTag != "" && fmtc.IsTag(a.AppNameColorTag), a.AppNameColorTag, ""),
		DEFAULT_APP_NAME_COLOR_TAG,
	)

	versionColor := strutil.Q(
		strutil.B(a.VersionColorTag != "" && fmtc.IsTag(a.VersionColorTag), a.VersionColorTag, ""),
		DEFAULT_APP_VER_COLOR_TAG,
	)

	releaseColor := strutil.Q(
		strutil.B(a.ReleaseColorTag != "" && fmtc.IsTag(a.ReleaseColorTag), a.ReleaseColorTag, ""),
		DEFAULT_APP_REL_COLOR_TAG,
	)

	buildColor := strutil.Q(
		strutil.B(a.BuildColorTag != "" && fmtc.IsTag(a.BuildColorTag), a.BuildColorTag, ""),
		DEFAULT_APP_BUILD_COLOR_TAG,
	)

	relSeparator := strutil.Q(a.ReleaseSeparator, "-")
	descSeparator := strutil.Q(a.DescSeparator, "-")

	fmtc.Printf("\n"+nameColor+"%s{!} "+versionColor+"%s{!}", a.App, a.Version)

	fmtc.If(a.Release != "").Printf(releaseColor+relSeparator+"%s{!}", a.Release)
	fmtc.If(a.Build != "").Printf(buildColor+" (%s){!}", a.Build)

	fmtc.Printfn(" "+descSeparator+" %s", a.Desc)

	if len(a.Environment) > 0 {
		fmtc.Printfn("{s-}│{!}")

		for i, env := range a.Environment {
			if len(a.Environment) != i+1 {
				fmtc.Printfn("{s-}├ %s: %s{!}", env.Name, strutil.Q(env.Version, "—"))
			} else {
				fmtc.Printfn("{s-}└ %s: %s{!}", env.Name, strutil.Q(env.Version, "—"))
			}
		}
	}

	fmtc.NewLine()

	if a.Owner != "" {
		fmtc.If(a.Year == 0).Printf(
			"{s-}%s %d %s{!}\n",
			strutil.Q(a.Copyright, "Copyright (C)"),
			time.Now().Year(), a.Owner,
		)

		fmtc.If(a.Year != 0).Printf(
			"{s-}%s %d-%d %s{!}\n",
			strutil.Q(a.Copyright, "Copyright (C)"),
			a.Year, time.Now().Year(), a.Owner,
		)
	}

	fmtc.If(a.License != "").Printfn("{s-}%s{!}", a.License)

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
	printGroupHeader(strutil.Q(info.OptionsHeader, "Options"))

	for _, option := range info.Options {
		option.Print()
	}
}

// printEnvVars prints all supported environment variables
func printEnvVars(info *Info) {
	printGroupHeader(strutil.Q(info.EnvVarsHeader, "Environment variables"))

	for _, env := range info.EnvVars {
		env.Print()
	}
}

// printExamples prints all usage examples
func printExamples(info *Info) {
	printGroupHeader(strutil.Q(info.ExamplesHeader, "Examples"))

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
	var result strings.Builder

	for _, a := range args {
		if strings.HasPrefix(a, "?") {
			result.WriteString("{s-}" + a[1:] + "{!} ")
		} else {
			result.WriteString("{s}" + a + "{!} ")
		}
	}

	return fmtc.Sprint(strings.TrimRight(result.String(), " "))
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
		return strutil.ReadField(name, 1, false, ':'),
			strutil.ReadField(name, 0, false, ':')
	}

	return name, ""
}

// getSeparator return bread crumbs (or spaces if colors are disabled) for
// item name aligning
func getSeparator(size, maxSize int, breadcrumbs bool) string {
	if size >= maxSize {
		return "  "
	}

	if breadcrumbs && !fmtc.DisableColors && maxSize > _BREADCRUMBS_MIN_SIZE {
		return " {s-}" + _DOTS[:maxSize-size] + "{!} "
	}

	return " " + _SPACES[:maxSize-size] + " "
}

// getMaxCommandSize returns the size of the longest command
func getMaxCommandSize(commands []*Command, group string) int {
	var size int

	for _, command := range commands {
		if group != "" && command.Group != group {
			continue
		}

		size = max(size, getCommandSize(command)+2)
	}

	return size
}

// getMaxOptionSize returns the size of the longest option
func getMaxOptionSize(options []*Option) int {
	var size int

	for _, option := range options {
		size = max(size, getOptionSize(option)+2)
	}

	return size
}

// getMaxEnvVarSize returns the size of the longest environment variable
func getMaxEnvVarSize(envVars []*Env) int {
	var size int

	for _, env := range envVars {
		size = max(size, strutil.Len(env.Name)+2)
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
	fmtc.Printfn("\n{*}%s{!}\n", name)
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
		fmtc.Printfn("{s-}(released %d days ago){!}", days)
	}
}

// wrapText wraps text to the specified length with indentation
func wrapText(text string, indent, maxLen int) string {
	size := max(_DEFAULT_WRAP_LEN, maxLen) - indent

	if strutil.LenVisual(fmtc.Clean(text)) <= size {
		return text
	}

	var buf bytes.Buffer

	for {
		if len(text) < size {
			buf.WriteString(text)
			break
		}

		wi := strings.LastIndex(text[:size], " ")

		if wi == -1 {
			buf.WriteString(text)
			break
		}

		buf.WriteString(text[:wi])
		buf.WriteRune('\n')
		buf.WriteString(strings.Repeat(" ", indent))
		text = text[wi+1:]
	}

	return buf.String()
}

// evalStrings converts slice with any into slice with strings
func evalStrings(data []any) []string {
	var result []string

	for _, v := range data {
		result = append(result, evalString(v))
	}

	return result
}

// evalString converts any into string
func evalString(v any) string {
	switch vv := v.(type) {
	case string:
		return vv
	case fmt.Stringer:
		return vv.String()
	case nil:
		return ""
	}

	return fmt.Sprintf("%v", v)
}
