package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"github.com/essentialkaos/ek/v13/usage"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _RESULT_FILES = `# Completion for test
# This completion is automatically generated

_test() {
  local cur prev cmds opts show_files

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="print clean"
  opts="--option-aaa --option-bbb --option-ccc"
  show_files="true"
  file_glob="[sr]pm"

  case $prev in
    clean)
      opts="--option-ddd --option-eee"
      COMPREPLY=($(compgen -W "$opts" -- "$cur"))
      return 0
      ;;

  esac

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  _test_filter "$cmds" "$opts" "$show_files" "$file_glob"
}

_test_filter() {
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

complete -F _test test -o filenames
`

const _RESULT_NO_FILES = `# Completion for test
# This completion is automatically generated

_test() {
  local cur prev cmds opts show_files

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="print clean"
  opts="--option-aaa --option-bbb --option-ccc"
  show_files=""
  file_glob=""

  case $prev in
    clean)
      opts="--option-ddd --option-eee"
      COMPREPLY=($(compgen -W "$opts" -- "$cur"))
      return 0
      ;;

  esac

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  _test_filter "$cmds" "$opts" "$show_files" "$file_glob"
}

_test_filter() {
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

complete -F _test test 
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BashSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BashSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BashSuite) TestGenerator(c *C) {
	completion := Generate(genTestUsageInfo(true), "test", "[sr]pm")
	c.Assert(completion, Equals, _RESULT_FILES)

	completion = Generate(genTestUsageInfo(false), "test")
	c.Assert(completion, Equals, _RESULT_NO_FILES)
}

func (s *BashSuite) TestAuxi(c *C) {
	info := usage.NewInfo("")

	c.Assert(isCommandHandlersRequired(info), Equals, false)
	c.Assert(genCommandsHandlers(info), Equals, "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func genTestUsageInfo(withFiles bool) *usage.Info {
	var info *usage.Info

	if withFiles {
		info = usage.NewInfo("", "spec-file")
	} else {
		info = usage.NewInfo("")
	}

	info.AddCommand("print", "Print command")
	info.AddCommand("clean", "Clean command")

	info.AddOption("a:option-aaa", "Test option A")
	info.AddOption("b:option-bbb", "Test option B", "?bbb")
	info.AddOption("option-ccc", "Test option C")

	info.AddOption("d:option-ddd", "Test option D")
	info.AddOption("e:option-eee", "Test option E")

	info.BoundOptions("clean", "d:option-ddd", "e:option-eee", "unknown")

	return info
}
