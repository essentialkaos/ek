// Package zsh provides methods for generating zsh completion
package zsh

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/options"
	"github.com/essentialkaos/ek/v13/usage"
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
{{FILES_HANDLER}}
}

{{COMMANDS_FUNC}}
_{{COMPNAME}} "$@"
`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates zsh completion code
func Generate(info *usage.Info, opts options.Map, name string, fileGlob ...string) string {
	result := _ZSH_TEMPLATE

	result = strings.ReplaceAll(result, "{{GLOBAL_ARGS}}", genGlobalOptionList(info, opts))
	result = strings.ReplaceAll(result, "{{COMMANDS_HANDLERS}}", genCommandsHandlers(info, opts))
	result = strings.ReplaceAll(result, "{{COMMANDS_FUNC}}", genCommandsFunc(info))
	result = strings.ReplaceAll(result, "{{FILES_HANDLER}}", genFilesHandler(info, fileGlob))
	result = strings.ReplaceAll(result, "{{COMPNAME}}", name)

	nameSafe := strings.ReplaceAll(name, "-", "_")

	result = strings.ReplaceAll(result, "{{COMPNAME_SAFE}}", nameSafe)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genFilesHandler generates handler for showing files
func genFilesHandler(info *usage.Info, fileGlob []string) string {
	if len(info.Args) == 0 && len(fileGlob) == 0 {
		return ""
	}

	if len(fileGlob) != 0 {
		return "  _files -g \"" + fileGlob[0] + "\""
	}

	return "  _files"
}

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

// genConflictsExclusion generates list with conflicts exclusions
func genConflictsExclusion(opts any) string {
	var result []string
	var optSlice []string

	switch t := opts.(type) {
	case string:
		optSlice = strings.Split(t, options.MergeSymbol)
	case []string:
		optSlice = t
	default:
		return ""
	}

	for _, opt := range optSlice {
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
