// Package usage provides methods and structs for rendering info for command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v4/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SPACES = "                                                                "

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

type option struct {
	name  string
	desc  string
	args  string
	group string
}

type example struct {
	cmd  string
	desc string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// CommandsColor contains default commands color
	CommandsColor = "y"

	// OptionsColor contains default options color
	OptionsColor = "g"
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
		usageMessage += " {" + CommandsColor + "}{command}{!}"
	}

	if len(info.options) != 0 {
		usageMessage += " {" + OptionsColor + "}{options}{!}"
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
		renderOptions(info.commands, CommandsColor)
	}

	if len(info.options) != 0 {
		renderOptions(info.options, OptionsColor)
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
			"\n{*c}%s {c}%s{!}{s}%s (%s){!} - %s\n\n",
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

	fmtc.NewLine()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// appendOption append new option to options slice
func appendOption(data []string, options *[]option, group string) {
	if len(data) < 2 {
		return
	}

	var optionArgs string

	if len(data) >= 3 {
		optionArgs = strings.Join(data[2:], " ")
	}

	var name = data[0]

	if group == "Options" {
		name = parseOption(data[0])
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
func renderOptions(options []option, color string) {
	var (
		curGroup string
		opt      option
		maxSize  int
	)

	maxSize = getMaxOptionSize(options)

	for _, opt = range options {
		if curGroup != opt.group {
			fmtc.Printf("\n{*}%s{!}\n\n", opt.group)
			curGroup = opt.group
		}

		if len(opt.args) != 0 {
			fmtc.Printf(
				"  {"+color+"}%s{!} {s-}%s{!} %s %s\n",
				opt.name,
				opt.args,
				getOptionSpaces(opt, maxSize),
				fmtc.Sprintf(opt.desc),
			)
		} else {
			fmtc.Printf(
				"  {"+color+"}%s{!}  %s %s\n",
				opt.name,
				getOptionSpaces(opt, maxSize),
				fmtc.Sprintf(opt.desc),
			)
		}
	}
}

// renderExamples render examples
func renderExamples(info *Info) {
	fmtc.Println("\n{*}Examples:{!}\n")

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

// getOptionSpaces return spaces for option name aligning
func getOptionSpaces(opt option, maxSize int) string {
	optLen := len(opt.name) + len(opt.args)
	spaces := maxSize - optLen

	return _SPACES[:spaces]
}

// getMaxOptionSize return longest option name size
func getMaxOptionSize(options []option) int {
	var result = 0

	for _, opt := range options {
		optLen := len(opt.name) + len(opt.args) + 2

		if optLen > result {
			result = optLen
		}
	}

	return result
}
