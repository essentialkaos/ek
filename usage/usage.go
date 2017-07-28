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

	"pkg.re/essentialkaos/ek.v9/fmtc"
	"pkg.re/essentialkaos/ek.v9/version"
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

	name     string
	args     string
	spoiler  string
	commands []*entity
	options  []*entity
	examples []*example
	curGroup string
}

// UpdateChecker is a base for all update checkers
type UpdateChecker struct {
	Data      string
	CheckFunc func(app, version, data string) (string, time.Time, bool)
}

// ////////////////////////////////////////////////////////////////////////////////// //

type entity struct {
	name  string
	desc  string
	args  []string
	group string
}

type example struct {
	cmd  string
	desc string
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
		name:     name,
		args:     strings.Join(args, " "),
		commands: make([]*entity, 0),
		options:  make([]*entity, 0),
		examples: make([]*example, 0),

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

	appendEntity(a, &info.commands, group)
}

// AddOption add option (name, description, args)
func (info *Info) AddOption(a ...string) {
	appendEntity(a, &info.options, "Options")
}

// AddExample add example for some command (command, description)
func (info *Info) AddExample(a ...string) {
	if len(a) == 0 {
		return
	}

	a = append(a, "")

	info.examples = append(info.examples, &example{cmd: a[0], desc: a[1]})
}

// AddSpoiler add spoiler
func (info *Info) AddSpoiler(spoiler string) {
	info.spoiler = spoiler
}

// Render print usage info to console
func (info *Info) Render() {
	usageMessage := "\n{*}Usage:{!} " + info.name

	if len(info.options) != 0 {
		usageMessage += " " + info.OptionsColorTag + "{options}{!}"
	}

	if len(info.commands) != 0 {
		usageMessage += " " + info.CommandsColorTag + "{command}{!}"
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
		renderEntities(info.commands, info.CommandsColorTag, info.Breadcrumbs)
	}

	if len(info.options) != 0 {
		renderEntities(info.options, info.OptionsColorTag, info.Breadcrumbs)
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
			fmtc.Printf(
				"{s-}Copyright (C) %d %s{!}\n",
				time.Now().Year(), about.Owner)
		} else {
			fmtc.Printf(
				"{s-}Copyright (C) %d-%d %s{!}\n",
				about.Year, time.Now().Year(), about.Owner)
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

// appendEntity append new entity to given slice
func appendEntity(data []string, entities *[]*entity, group string) {
	if len(data) < 2 {
		return
	}

	var args []string

	if len(data) >= 3 {
		args = data[2:]
	}

	var name = data[0]

	*entities = append(*entities,
		&entity{
			name:  name,
			desc:  data[1],
			args:  args,
			group: group,
		},
	)
}

// formatOption format entity name
func formatOption(entity *entity) string {
	if strings.Contains(entity.name, ":") {
		optionSlice := strings.Split(entity.name, ":")
		return "--" + optionSlice[1] + ", -" + optionSlice[0]
	}

	return "--" + entity.name
}

// renderEntities render entities
func renderEntities(entities []*entity, colorTag string, breadcrumbs bool) {
	var curGroup string
	var maxSize int

	maxSize = getMaxEntitySize(entities)

	for _, entity := range entities {
		if curGroup != entity.group {
			printGroupHeader(entity.group)
			curGroup = entity.group
		}

		if entity.group == "Options" {
			fmtc.Printf("  "+colorTag+"%s{!}", formatOption(entity))
		} else {
			fmtc.Printf("  "+colorTag+"%s{!}", entity.name)
		}

		if len(entity.args) != 0 {
			fmtc.Printf(" " + renderArgs(entity.args))
		}

		fmtc.Printf(getEntitySeparator(entity, maxSize, breadcrumbs))
		fmtc.Printf(entity.desc)

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

// getEntitySeparator return bread crumbs (or spaces if colors are disabled) for
// entity name aligning
func getEntitySeparator(entity *entity, maxSize int, breadcrumbs bool) string {
	entLen := getEntitySize(entity)

	if breadcrumbs && !fmtc.DisableColors && maxSize > _BREADCRUMBS_MIN_SIZE {
		return " {s-}" + _DOTS[:maxSize-entLen] + "{!} "
	}

	return " " + _SPACES[:maxSize-entLen] + " "
}

// getMaxEntitySize return longest entity name size
func getMaxEntitySize(entities []*entity) int {
	var result int

	for _, entity := range entities {
		entLen := getEntitySize(entity) + 2

		if entLen > result {
			result = entLen
		}
	}

	return result
}

// getEntitySize calculate rendered entity size
func getEntitySize(entity *entity) int {
	var size int

	if strings.Contains(entity.name, ":") {
		size += len(entity.name) + 4
	} else {
		size += len(entity.name) + 2
	}

	for _, arg := range entity.args {
		if strings.HasPrefix(arg, "?") {
			size += len(arg)
		} else {
			size += len(arg) + 1
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
