package fish

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"pkg.re/essentialkaos/ek.v12/usage"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _RESULT1 = `# Completion for test
# This completion is automatically generated

function __fish_test_no_command
  set cmd (commandline -opc)
  if [ (count $cmd) -eq 1 ]
    return 0
  end
  return 1
end

function __fish_test_using_command
  set cmd (commandline -opc)
  if [ (count $cmd) -gt 1 ]
    if [ $argv[1] = $cmd[2] ]
      return 0
    end
  end
  return 1
end

complete -f -n '__fish_test_no_command' -c test -l option-aaa -s a -d 'Test option A'
complete -f -n '__fish_test_no_command' -c test -l option-bbb -s b -d 'Test option B'
complete -f -n '__fish_test_no_command' -c test -l option-ccc -d 'Test option B'

complete -f -n '__fish_test_no_command' -c test -a 'print' -d 'Print command'

complete -f -n '__fish_test_no_command' -c test -a 'clean' -d 'Clean command'
complete -f -n '__fish_test_using_command clean' -c test -l option-ddd -s e -d 'Test option D'
complete -f -n '__fish_test_using_command clean' -c test -l option-eee -s d -d 'Test option E'


`

const _RESULT2 = `# Completion for test
# This completion is automatically generated

function __fish_test_no_command
  set cmd (commandline -opc)
  if [ (count $cmd) -eq 1 ]
    return 0
  end
  return 1
end

function __fish_test_using_command
  set cmd (commandline -opc)
  if [ (count $cmd) -gt 1 ]
    if [ $argv[1] = $cmd[2] ]
      return 0
    end
  end
  return 1
end


`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type FishSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&FishSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FishSuite) TestGenerator(c *C) {
	data := Generate(genTestUsageInfo(), "test")

	c.Assert(data, Equals, _RESULT1)

	data = Generate(usage.NewInfo(""), "test")

	c.Assert(data, Equals, _RESULT2)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func genTestUsageInfo() *usage.Info {
	info := usage.NewInfo("")

	info.AddCommand("print", "Print command")
	info.AddCommand("clean", "Clean command")

	info.AddOption("a:option-aaa", "Test option A")
	info.AddOption("b:option-bbb", "Test option B")
	info.AddOption("option-ccc", "Test option B")

	info.AddOption("e:option-ddd", "Test option D")
	info.AddOption("d:option-eee", "Test option E")

	info.BoundOptions("clean", "e:option-ddd", "d:option-eee", "option-hhh")

	return info
}
