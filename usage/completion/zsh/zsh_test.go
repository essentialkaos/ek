package zsh

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"github.com/essentialkaos/ek/v13/options"
	"github.com/essentialkaos/ek/v13/usage"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _RESULT = `#compdef test

# This completion is automatically generated

typeset -A opt_args

_test() {
  _arguments \
    '(-a --option-aaa=)'{-a,--option-aaa=}'[Test option A]' \
    '(-b --option-bbb)'{-b,--option-bbb}'[Test option B]' \
    '(--option-ccc=)'--option-ccc'[Test option C]' \
    '1: :_test_cmds' \
    '*:: :->cmd_args' && ret=0

  case $state in
    cmd_args)
      case $words[1] in
        clean)
          _arguments \
            '(-d --option-ddd)'{-d,--option-ddd}'[Test option D]' \
            '(-e --option-eee= -d --option-ddd --option-ccc)'{-e,--option-eee=}'[Test option E]' \
        ;;
      esac
    ;;
  esac

  _files -g "*.txt"
}

_test_cmds() {
  local -a commands
  commands=(
    'print:Print command'
    'clean:Clean command'
  )
  _describe 'command' commands
}

_test "$@"
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ZSHSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ZSHSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ZSHSuite) TestGenerator(c *C) {
	data := Generate(genTestUsageInfo(), genTestOptionsWithString(), "test", "*.txt")
	c.Assert(data, Equals, _RESULT)

	data = Generate(genTestUsageInfo(), genTestOptionsWithSlice(), "test", "*.txt")
	c.Assert(data, Equals, _RESULT)
}

func (s *ZSHSuite) TestAuxi(c *C) {
	info := usage.NewInfo("")

	c.Assert(genCommandsHandlers(info, nil), Equals, "")
	c.Assert(genCommandsFunc(info), Equals, "")

	c.Assert(genFilesHandler(info, nil), Equals, "")

	info = usage.NewInfo("", "files")

	c.Assert(genFilesHandler(info, []string{"*.txt"}), Equals, "  _files -g \"*.txt\"")
	c.Assert(genFilesHandler(info, nil), Equals, "  _files")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func genTestOptionsWithString() options.Map {
	return options.Map{
		"a:option-aaa": {},
		"b:option-bbb": {Type: options.BOOL},
		"option-ccc":   {},
		"e:option-eee": {Conflicts: "d:option-ddd option-ccc"},
		"d:option-ddd": {Type: options.BOOL},
	}
}

func genTestOptionsWithSlice() options.Map {
	return options.Map{
		"a:option-aaa": {},
		"b:option-bbb": {Type: options.BOOL},
		"option-ccc":   {},
		"e:option-eee": {Conflicts: []string{"d:option-ddd", "option-ccc"}},
		"d:option-ddd": {Type: options.BOOL},
	}
}

func genTestUsageInfo() *usage.Info {
	info := usage.NewInfo("", "file…")

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
