// Package zsh provides methods for generating zsh completion
package zsh

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"pkg.re/essentialkaos/ek.v11/fmtc"
	"pkg.re/essentialkaos/ek.v11/options"
	"pkg.re/essentialkaos/ek.v11/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// _ZSH_TEMPLATE is zsh completion template
const _ZSH_TEMPLATE = `#compdef {{COMPNAME}}

# This completion is automatically generated

typeset -A opt_args

_{{COMPNAME}}() {
  _arguments \
{{GLOBAL_ARGS}}

{{COMMANDS_HANDLERS}}
}

{{COMMANDS_FUNC}}
_{{COMPNAME}} "$@"
`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates zsh completion code
func Generate(info *usage.Info, opts options.Map, name string) string {
	result := _ZSH_TEMPLATE

	result = strings.Replace(result, "{{GLOBAL_ARGS}}", genGlobalOptionList(info, opts), -1)
	result = strings.Replace(result, "{{COMMANDS_HANDLERS}}", genCommandsHandlers(info, opts), -1)
	result = strings.Replace(result, "{{COMMANDS_FUNC}}", genCommandsFunc(info), -1)
	result = strings.Replace(result, "{{COMPNAME}}", name, -1)

	nameSafe := strings.Replace(name, "-", "_", -1)

	result = strings.Replace(result, "{{COMPNAME_SAFE}}", nameSafe, -1)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genGlobalOptionList generates list with global options
func genGlobalOptionList(info *usage.Info, opts options.Map) string {
	var result string

	nonGlobalOptions := make(map[string]bool)

	for _, cmd := range info.Commands {
		for _, opt := range cmd.BoundOptions {
			nonGlobalOptions[opt] = true
		}
	}

	for _, opt := range info.Options {
		if nonGlobalOptions[opt.Long] {
			continue
		}

		result += genOptionDesc(opt, opts, 4)
	}

	if len(info.Commands) != 0 {
		result += "    '1: :_{{COMPNAME_SAFE}}_cmds' \\\n"
	}

	result += "    '*:: :->cmd_args' && ret=0"

	return result
}

// genOptionDesc generates description for some option
func genOptionDesc(opt *usage.Option, opts options.Map, prefixSize int) string {
	result := strings.Repeat(" ", prefixSize)

	var isBool, isMergeble bool
	var optV *options.V

	optLong := opt.Long

	if opts != nil {
		optV = opts[getOptionFullName(opt)]
	}

	if optV != nil {
		isBool = optV.Type == options.BOOL
		isMergeble = optV.Mergeble
	}

	if !isBool {
		optLong += "="
	}

	if !isMergeble {
		var exclusion string

		if opt.Short != "" {
			exclusion = fmt.Sprintf("-%s --%s", opt.Short, optLong)
		} else {
			exclusion = fmt.Sprintf("--%s", optLong)
		}

		if optV != nil && optV.Conflicts != "" {
			exclusion += genConflictsExclusion(optV.Conflicts)
		}

		result += "'(" + exclusion + ")'"
	}

	if opt.Short != "" {
		result += fmt.Sprintf("{-%s,--%s}", opt.Short, optLong)
	} else {
		result += fmt.Sprintf("--%s", opt.Long)
	}

	result += fmt.Sprintf("'[%s]'", fmtc.Clean(opt.Desc))
	result += " \\\n"

	return result
}

// genCommandsHandlers generates handlers for commands
func genCommandsHandlers(info *usage.Info, opts options.Map) string {
	if !isCommandHandlersRequired(info) {
		return ""
	}

	result := "  case $state in\n"
	result += "    cmd_args)\n"
	result += "      case $words[1] in\n"

	for _, cmd := range info.Commands {
		if len(cmd.BoundOptions) != 0 {
			result += genCommandHandler(cmd, info, opts)
		}
	}

	result += "      esac\n"
	result += "    ;;\n"
	result += "  esac\n"

	return result
}

// genCommandHandler generates handler for given command
func genCommandHandler(cmd *usage.Command, info *usage.Info, opts options.Map) string {
	result := fmt.Sprintf("        %s)\n", cmd.Name)
	result += "          _arguments \\\n"

	for _, optName := range cmd.BoundOptions {
		opt := info.GetOption(optName)

		if opt == nil {
			continue
		}

		result += genOptionDesc(opt, opts, 12)
	}

	result += "        ;;\n"

	return result
}

// genCommandsFunc generates function which generates list with all supported commands
func genCommandsFunc(info *usage.Info) string {
	if len(info.Commands) == 0 {
		return ""
	}

	result := "_{{COMPNAME_SAFE}}_cmds() {\n"
	result += "  local -a commands\n"
	result += "  commands=(\n"

	for _, cmd := range info.Commands {
		result += fmt.Sprintf("    '%s:%s'\n", cmd.Name, fmtc.Clean(cmd.Desc))
	}

	result += "  )\n"
	result += "  _describe 'command' commands\n"
	result += "}\n"

	return result
}

// genConflictsExclusion generates list with conflicts exlusions
func genConflictsExclusion(opts string) string {
	var result []string

	for _, opt := range strings.Split(opts, " ") {
		long, short := options.ParseOptionName(opt)

		if short != "" {
			result = append(result, "-"+short, "--"+long)
		} else {
			result = append(result, "--"+long)
		}
	}

	return " " + strings.Join(result, " ")
}

// isCommandHandlersRequired returns true if commands have bound options
func isCommandHandlersRequired(info *usage.Info) bool {
	for _, cmd := range info.Commands {
		if len(cmd.BoundOptions) != 0 {
			return true
		}
	}

	return false
}

// getOptionFullName generates combined option name
func getOptionFullName(opt *usage.Option) string {
	if opt.Short != "" {
		return opt.Short + ":" + opt.Long
	}

	return opt.Long
}
