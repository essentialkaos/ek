// Package provides methods for rendering info for command-line tools
package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
	"time"

	"github.com/essentialkaos/ek/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SPACES = "                                                                "

// ////////////////////////////////////////////////////////////////////////////////// //

// About contains info about application
type About struct {
	App     string
	Version string
	Release string
	Build   string
	Desc    string
	Year    int
	License string
	Owner   string
}

// Info contains info about commands, options and examples
type Info struct {
	Name     string
	SubArgs  string
	spoiler  string
	commands []*option
	options  []*option
	examples []*example
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

// NewInfo create new info struct
func NewInfo(args ...string) *Info {
	args = append(args, "", "")
	return &Info{Name: args[0], SubArgs: args[1]}
}

// AddGroup add new command group
func (info *Info) AddGroup(name string) {
	info.curGroup = name
}

// AddCommand add command (name, desc, args)
func (info *Info) AddCommand(args ...string) {
	if len(args) < 2 {
		return
	}

	oargs := ""

	if len(args) >= 3 {
		oargs = strings.Join(args[2:], " ")
	}

	if info.curGroup != "" {
		info.commands = append(info.commands, &option{args[0], args[1], oargs, info.curGroup})
	} else {
		info.commands = append(info.commands, &option{args[0], args[1], oargs, "Commands"})
	}
}

// AddOption add option (name, desc, args)
func (info *Info) AddOption(args ...string) {
	if len(args) < 2 {
		return
	}

	opt := parseOption(args[0])
	oargs := ""

	if len(args) >= 3 {
		oargs = strings.Join(args[2:], " ")
	}

	info.options = append(info.options, &option{opt, args[1], oargs, ""})
}

// AddExample add example for some command (command, desc)
func (info *Info) AddExample(args ...string) {
	if len(args) == 0 {
		return
	}

	args = append(args, "")

	info.examples = append(info.examples, &example{args[0], args[1]})
}

// AddSpoiler add spoiler
func (info *Info) AddSpoiler(spoiler string) {
	info.spoiler = spoiler
}

// Render print info to console
func (info *Info) Render() {
	usageMessage := fmt.Sprintf("\n{*}Usage:{!} %s", info.Name)

	if len(info.commands) != 0 {
		usageMessage += " {y}<command>{!}"
	}

	if len(info.options) != 0 {
		usageMessage += " {g}<options>{!}"
	}

	if info.SubArgs != "" {
		usageMessage += " " + info.SubArgs
	}

	fmtc.Println(usageMessage)

	if info.spoiler != "" {
		fmt.Println("")
		fmtc.Println(info.spoiler)
	}

	if len(info.commands) != 0 {
		renderCommands(info)
	}

	if len(info.options) != 0 {
		renderOptions(info)
	}

	if len(info.examples) != 0 {
		renderExamples(info)
	}

	fmt.Println("")
}

// Render print about info to console
func (about *About) Render() {
	switch {
	case about.Build != "":
		fmtc.Printf("\n{*c}%s {c}%s{!}{s}%s (%s){!} - %s\n\n", about.App, about.Version, about.Release, about.Build, about.Desc)
	default:
		fmtc.Printf("\n{*c}%s {c}%s{!}{s}%s{!} - %s\n\n", about.App, about.Version, about.Release, about.Desc)
	}

	if about.Owner != "" {
		if about.Year == 0 {
			fmtc.Printf("{s}Copyright (C) %s{!}\n", about.Owner)
		} else {
			fmtc.Printf("{s}Copyright (C) %d-%d %s{!}\n", about.Year, time.Now().Year(), about.Owner)
		}
	}

	if about.License != "" {
		fmtc.Printf("{s}%s{!}\n", about.License)
	}

	fmt.Println()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func parseOption(opt string) string {
	if opt[0:1] == "-" {
		return opt
	}

	if strings.Contains(opt, ":") {
		opts := strings.Split(opt, ":")
		return "--" + opts[1] + ", -" + opts[0]
	}

	return "--" + opt
}

func renderOptions(info *Info) {
	fmtc.Println("\n{*}Options:{!}\n")

	var (
		opt     *option
		ln, dln int
		maxlen  int
	)

	for _, opt = range info.options {
		ln = len(opt.name) + len(opt.args) + 2

		if ln > maxlen {
			maxlen = ln
		}
	}

	for _, opt = range info.options {
		ln = len(opt.name) + len(opt.args)
		dln = maxlen - ln

		if len(opt.args) != 0 {
			fmtc.Printf("  {g}%s{!} {s}%s{!} %s %s\n", opt.name, opt.args, _SPACES[0:dln], opt.desc)
		} else {
			fmtc.Printf("  {g}%s{!}  %s %s\n", opt.name, _SPACES[0:dln], opt.desc)
		}
	}
}

func renderCommands(info *Info) {
	var (
		cmd      *option
		ln, dln  int
		maxlen   int
		curGroup string
	)

	for _, cmd = range info.commands {
		ln = len(cmd.name) + len(cmd.args) + 2

		if ln > maxlen {
			maxlen = ln
		}
	}

	if info.commands[0].group == "" {
		fmtc.Println("\n{*}Commands:{!}\n")
	}

	for _, cmd = range info.commands {
		ln = len(cmd.name) + len(cmd.args)
		dln = maxlen - ln

		if curGroup != cmd.group {
			curGroup = cmd.group
			fmtc.Printf("\n{*}%s:{!}\n\n", curGroup)
		}

		if len(cmd.args) != 0 {
			fmtc.Printf("  {y}%s{!} {s}%s{!} %s %s\n", cmd.name, cmd.args, _SPACES[0:dln], cmd.desc)
		} else {
			fmtc.Printf("  {y}%s{!}  %s %s\n", cmd.name, _SPACES[0:dln], cmd.desc)
		}
	}
}

func renderExamples(info *Info) {
	fmtc.Println("\n{*}Examples:{!}\n")

	total := len(info.examples)

	for index, example := range info.examples {
		fmt.Printf("  %s %s\n", info.Name, example.cmd)

		if example.desc != "" {
			fmtc.Printf("  {s}%s{!}\n", example.desc)
		}

		if index < total-1 {
			fmtc.NewLine()
		}
	}
}
