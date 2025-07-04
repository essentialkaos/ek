// Package bash provides methods for generating bash completion
package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"github.com/essentialkaos/ek/v13/usage"
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
  file_glob="{{FILE_GLOB}}"

{{COMMANDS_HANDLERS}}

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  _{{COMPNAME_SAFE}}_filter "$cmds" "$opts" "$show_files" "$file_glob"
}

_{{COMPNAME_SAFE}}_filter() {
  local cmds="$1"
  local opts="$2"
  local show_files="$3"
  local file_glob="$4"

  local cmd1 cmd2

  for cmd1 in $cmds ; do
    for cmd2 in ${COMP_WORDS[*]} ; do
      if [[ "$cmd1" == "$cmd2" ]] ; then
        if [[ -z "$show_files" ]] ; then
          COMPREPLY=($(compgen -W "$opts" -- "$cur"))
        else
          _filedir "$file_glob"
        fi

        return 0
      fi
    done
  done

  if [[ -z "$show_files" ]] ; then
    COMPREPLY=($(compgen -W "$cmds" -- "$cur"))
    return 0
  fi

  _filedir "$file_glob"
}

complete -F _{{COMPNAME_SAFE}} {{COMPNAME}} {{COMP_OPTS}}
`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates Bash completion code
func Generate(info *usage.Info, name string, fileExt ...string) string {
	result := _BASH_TEMPLATE

	result = strings.ReplaceAll(result, "{{COMMANDS}}", genCommandsList(info))
	result = strings.ReplaceAll(result, "{{GLOBAL_OPTIONS}}", genGlobalOptionsList(info))
	result = strings.ReplaceAll(result, "{{COMMANDS_HANDLERS}}", genCommandsHandlers(info))
	result = strings.ReplaceAll(result, "{{COMPNAME}}", name)

	if len(info.Args) != 0 {
		result = strings.ReplaceAll(result, "{{SHOW_FILES}}", "true")
		result = strings.ReplaceAll(result, "{{COMP_OPTS}}", "-o filenames")
		if len(fileExt) != 0 {
			result = strings.ReplaceAll(result, "{{FILE_GLOB}}", fileExt[0])
		}
	} else {
		result = strings.ReplaceAll(result, "{{SHOW_FILES}}", "")
		result = strings.ReplaceAll(result, "{{COMP_OPTS}}", "")
		result = strings.ReplaceAll(result, "{{FILE_GLOB}}", "")
	}

	nameSafe := strings.ReplaceAll(name, "-", "_")
	result = strings.ReplaceAll(result, "{{COMPNAME_SAFE}}", nameSafe)

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
