// Package bash provides methods for generating bash completion
package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"pkg.re/essentialkaos/ek.v12/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// _BASH_TEMPLATE is template used for completion generation
const _BASH_TEMPLATE = `# Completion for {{COMPNAME}}
# This completion is automatically generated

_{{COMPNAME_SAFE}}() {
  local cur prev cmds opts show_files

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="{{COMMANDS}}"
  opts="{{GLOBAL_OPTIONS}}"
  show_files="{{SHOW_FILES}}"

{{COMMANDS_HANDLERS}}

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  if [[ -z "$cmds" && -n "$show_files" ]] ; then
    _filedir && return 0
  fi

  COMPREPLY=($(compgen -W '$(_{{COMPNAME_SAFE}}_filter "$cmds" "$opts" "$show_files")' -- "$cur"))
}

_{{COMPNAME_SAFE}}_filter() {
  local cmds="$1"
  local opts="$2"
  local show_files="$3"

  local cmd1 cmd2

  for cmd1 in $1 ; do
    for cmd2 in ${COMP_WORDS[*]} ; do
      if [[ "$cmd1" == "$cmd2" ]] ; then
        echo "$2" && return 0
      fi
    done
  done

  if [[ -z "$show_files" ]] ; then
    echo "$opts" && return 0
  fi

  compgen -f -- "${COMP_WORDS[COMP_CWORD]}"
}

complete -F _{{COMPNAME_SAFE}} {{COMPNAME}} {{COMP_OPTS}}
`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates Bash completion code
func Generate(info *usage.Info, name string) string {
	result := _BASH_TEMPLATE

	result = strings.Replace(result, "{{COMMANDS}}", genCommandsList(info), -1)
	result = strings.Replace(result, "{{GLOBAL_OPTIONS}}", genGlobalOptionsList(info), -1)
	result = strings.Replace(result, "{{COMMANDS_HANDLERS}}", genCommandsHandlers(info), -1)
	result = strings.Replace(result, "{{COMPNAME}}", name, -1)

	if len(info.Args) != 0 {
		result = strings.Replace(result, "{{SHOW_FILES}}", "true", -1)
		result = strings.Replace(result, "{{COMP_OPTS}}", "-o filenames", -1)
	} else {
		result = strings.Replace(result, "{{SHOW_FILES}}", "", -1)
		result = strings.Replace(result, "{{COMP_OPTS}}", "", -1)
	}

	nameSafe := strings.Replace(name, "-", "_", -1)
	result = strings.Replace(result, "{{COMPNAME_SAFE}}", nameSafe, -1)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genGlobalOptionsList generates list with global options
func genGlobalOptionsList(info *usage.Info) string {
	var result []string

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

		result = append(result, "--"+opt.Long)
	}

	return strings.Join(result, " ")
}

// genCommandsHandlers generates command handler
func genCommandsHandlers(info *usage.Info) string {
	if !isCommandHandlersRequired(info) {
		return ""
	}

	result := "  case $prev in\n"

	for _, cmd := range info.Commands {
		if len(cmd.BoundOptions) != 0 {
			result += genCommandHandler(cmd, info)
		}
	}

	result += "  esac"

	return result
}

// genCommandHandler generates handler for given command
func genCommandHandler(cmd *usage.Command, info *usage.Info) string {
	result := fmt.Sprintf("    %s)\n", cmd.Name)

	var options []string

	for _, optName := range cmd.BoundOptions {
		opt := info.GetOption(optName)

		if opt == nil {
			continue
		}

		options = append(options, "--"+opt.Long)
	}

	result += fmt.Sprintf("      opts=\"%s\"\n", strings.Join(options, " "))
	result += "      COMPREPLY=($(compgen -W \"$opts\" -- \"$cur\"))\n"
	result += "      return 0\n"
	result += "      ;;\n\n"

	return result
}

// getCommandsList returns slice with available commands
func genCommandsList(info *usage.Info) string {
	var result []string

	for _, command := range info.Commands {
		result = append(result, command.Name)
	}

	return strings.Join(result, " ")
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
