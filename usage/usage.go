// Package usage provides methods and structs for generating usage info for
// command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v12/fmtc"
	"pkg.re/essentialkaos/ek.v12/mathutil"
	"pkg.re/essentialkaos/ek.v12/strutil"
	"pkg.re/essentialkaos/ek.v12/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SPACES = "                                                                "
	_DOTS   = "................................................................"
)

const _BREADCRUMBS_MIN_SIZE = 8

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DEFAULT_COMMANDS_COLOR_TAG = "{y}"
	DEFAULT_OPTIONS_COLOR_TAG  = "{g}"
	DEFAULT_APP_NAME_COLOR_TAG = "{c*}"
	DEFAULT_APP_VER_COLOR_TAG  = "{c}"
)

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
	BugTracker string // BugTracker is URL of bug tracker

	AppNameColorTag string // AppNameColorTag contains default app name color tag
	VersionColorTag string // VersionColorTag contains default app version color tag

	// Function for checking application updates
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
	Group        string   // Group is group name
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
	i.curGroup = group
}

// AddCommand adds command (name, description, args)
func (i *Info) AddCommand(a ...string) {
	group := "Commands"

	if i.curGroup != "" {
		group = i.curGroup
	}

	if len(a) < 2 {
		return
	}

	i.Commands = append(
		i.Commands,
		&Command{
			Name:  a[0],
			Desc:  a[1],
			Args:  a[2:],
			Group: group,
			info:  i,
		},
	)
}

// AddOption adds option (name, description, args)
func (i *Info) AddOption(a ...string) {
	if len(a) < 2 {
		return
	}

	long, short := parseOptionName(a[0])

	i.Options = append(
		i.Options,
		&Option{
			Long:  long,
			Short: short,
			Desc:  a[1],
			Arg:   strings.Join(a[2:], " "),
			info:  i,
		},
	)
}

// AddExample adds example of application usage
func (i *Info) AddExample(a ...string) {
	if len(a) == 0 {
		return
	}

	a = append(a, "")

	i.Examples = append(i.Examples, &Example{a[0], a[1], false, i})
}

// AddRawExample adds example of application usage without command prefix
func (i *Info) AddRawExample(a ...string) {
	if len(a) == 0 {
		return
	}

	a = append(a, "")

	i.Examples = append(i.Examples, &Example{a[0], a[1], true, i})
}

// AddSpoiler adds spoiler
func (i *Info) AddSpoiler(spoiler string) {
	i.Spoiler = spoiler
}

// BoundOptions bounds command with options
func (i *Info) BoundOptions(cmd string, options ...string) {
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
	for _, command := range i.Commands {
		if command.Name == name {
			return command
		}
	}

	return nil
}

// GetOption tries to find option with given name
func (i *Info) GetOption(name string) *Option {
	name, _ = parseOptionName(name)

	for _, option := range i.Options {
		if option.Long == name {
			return option
		}
	}

	return nil
}

// Render prints usage info to console
func (i *Info) Render() {
	usageMessage := "\n{*}Usage:{!} " + i.AppNameColorTag + i.Name + "{!}"

	if len(i.Options) != 0 {
		usageMessage += " " + i.OptionsColorTag + "{options}{!}"
	}

	if len(i.Commands) != 0 {
		usageMessage += " " + i.CommandsColorTag + "{command}{!}"
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
		renderCommands(i)
	}

	if len(i.Options) != 0 {
		renderOptions(i)
	}

	if len(i.Examples) != 0 {
		renderExamples(i)
	}

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns a string representation of the command
func (c *Command) String() string {
	return c.Name
}

// String returns a string representation of the option
func (o *Option) String() string {
	return "--" + o.Long
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Render renders info about command
func (c *Command) Render() {
	colorTag := strutil.Q(DEFAULT_COMMANDS_COLOR_TAG, c.ColorTag)
	size := getCommandSize(c)
	useBreadcrumbs := true
	maxSize := size

	if c.info != nil {
		colorTag = c.info.CommandsColorTag
		maxSize = getMaxCommandSize(c.info.Commands)
		useBreadcrumbs = c.info.Breadcrumbs
	}

	fmtc.Printf("  "+colorTag+"%s{!}", c.Name)

	if len(c.Args) != 0 {
		fmtc.Printf(" " + renderArgs(c.Args...))
	}

	fmtc.Printf(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Printf(c.Desc)

	fmtc.NewLine()
}

// Render renders info about option
func (o *Option) Render() {
	colorTag := strutil.Q(DEFAULT_OPTIONS_COLOR_TAG, o.ColorTag)
	size := getOptionSize(o)
	useBreadcrumbs := true
	maxSize := size

	if o.info != nil {
		colorTag = o.info.OptionsColorTag
		maxSize = getMaxOptionSize(o.info.Options)
		useBreadcrumbs = o.info.Breadcrumbs
	}

	fmtc.Printf("  "+colorTag+"%s{!}", formatOptionName(o))

	if o.Arg != "" {
		fmtc.Printf(" " + renderArgs(o.Arg))
	}

	fmtc.Printf(getSeparator(size, maxSize, useBreadcrumbs))
	fmtc.Printf(o.Desc)

	fmtc.NewLine()
}

// Render renders usage example
func (e *Example) Render() {
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

// Render prints version info to console
func (a *About) Render() {
	nc := strutil.Q(a.AppNameColorTag, DEFAULT_APP_NAME_COLOR_TAG)
	vc := strutil.Q(a.VersionColorTag, DEFAULT_APP_VER_COLOR_TAG)

	switch {
	case a.Build != "":
		fmtc.Printf(
			"\n"+nc+"%s{!} "+vc+"%s{!}{s}%s{!} {s-}(%s){!} - %s\n\n",
			a.App, a.Version,
			a.Release, a.Build, a.Desc,
		)
	default:
		fmtc.Printf(
			"\n"+nc+"%s{!} "+vc+"%s{!}{s}%s{!} - %s\n\n",
			a.App, a.Version,
			a.Release, a.Desc,
		)
	}

	if a.Owner != "" {
		if a.Year == 0 {
			fmtc.Printf(
				"{s-}Copyright (C) %d %s{!}\n",
				time.Now().Year(), a.Owner,
			)
		} else {
			fmtc.Printf(
				"{s-}Copyright (C) %d-%d %s{!}\n",
				a.Year, time.Now().Year(), a.Owner,
			)
		}
	}

	if a.License != "" {
		fmtc.Printf("{s-}%s{!}\n", a.License)
	}

	if a.UpdateChecker.CheckFunc != nil && a.UpdateChecker.Payload != "" {
		newVersion, releaseDate, hasUpdate := a.UpdateChecker.CheckFunc(
			a.App,
			a.Version,
			a.UpdateChecker.Payload,
		)

		if hasUpdate && isNewerVersion(a.Version, newVersion) {
			printNewVersionInfo(a.Version, newVersion, releaseDate)
		}
	}

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// renderCommands renders all supported commands
func renderCommands(info *Info) {
	var curGroup string

	for _, command := range info.Commands {
		if curGroup != command.Group {
			printGroupHeader(command.Group)
			curGroup = command.Group
		}

		command.Render()
	}
}

// renderOptions renders all supported options
func renderOptions(info *Info) {
	printGroupHeader("Options")

	for _, option := range info.Options {
		option.Render()
	}
}

// renderExamples renders all usage examples
func renderExamples(info *Info) {
	printGroupHeader("Examples")

	total := len(info.Examples)

	for index, example := range info.Examples {
		example.Render()

		if index < total-1 {
			fmtc.NewLine()
		}
	}
}

// renderArgs renders args with colors
func renderArgs(args ...string) string {
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

// formatOptionName formats option name
func formatOptionName(opt *Option) string {
	if opt.Short != "" {
		return "--" + opt.Long + ", -" + opt.Short
	}

	return "--" + opt.Long
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
func getMaxCommandSize(commands []*Command) int {
	var size int

	for _, command := range commands {
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

// getOptionSize calculate rendered command size
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

// getOptionSize calculate rendered option size
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
