// Package arg provides methods for working with command-line arguments
package arg

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v3/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

/*
	STRING argument type is string
	INT argument type is integer
	BOOL argument type is boolean
	FLOAT argument type is floating number
*/
const (
	STRING = 0
	INT    = 1
	BOOL   = 2
	FLOAT  = 3
)

// Error codes
const (
	ERROR_UNSUPPORTED         = 0
	ERROR_NO_NAME             = 1
	ERROR_DUPLICATE_LONGNAME  = 2
	ERROR_DUPLICATE_SHORTNAME = 3
	ERROR_ARG_IS_NIL          = 4
	ERROR_EMPTY_VALUE         = 5
	ERROR_REQUIRED_NOT_SET    = 6
	ERROR_WRONG_FORMAT        = 7
)

// ////////////////////////////////////////////////////////////////////////////////// //

// V basic argument struct
type V struct {
	Type     int     // argument type
	Max      float64 // maximum integer argument value
	Min      float64 // minimum integer argument value
	Alias    string  // list of aliases
	Mergeble bool    // argument supports arguments value merging
	Required bool    // argument is required

	set bool // Non exported field

	Value interface{} // default value
}

// Map is map with list of argumens
type Map map[string]*V

// Arguments arguments struct
type Arguments struct {
	full        Map
	short       map[string]string
	initialized bool
	hasRequired bool
}

// ArgumentError argument parsing error
type ArgumentError struct {
	Arg  string
	Type int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var global *Arguments

// ////////////////////////////////////////////////////////////////////////////////// //

// Add add new supported argument
func (args *Arguments) Add(name string, arg *V) error {
	if !args.initialized {
		initArgs(args)
	}

	longName, shortName := parseName(name)

	switch {
	case arg == nil:
		return ArgumentError{"--" + longName, ERROR_ARG_IS_NIL}
	case longName == "":
		return ArgumentError{"", ERROR_NO_NAME}
	case args.full[longName] != nil:
		return ArgumentError{"--" + longName, ERROR_DUPLICATE_LONGNAME}
	case shortName != "" && args.short[shortName] != "":
		return ArgumentError{"-" + shortName, ERROR_DUPLICATE_SHORTNAME}
	}

	if arg.Required == true {
		args.hasRequired = true
	}

	args.full[longName] = arg

	if shortName != "" {
		args.short[shortName] = longName
	}

	if arg.Alias != "" {
		aliases := strings.Split(arg.Alias, " ")

		for _, v := range aliases {
			alLongName, alShortName := parseName(v)

			args.full[alLongName] = arg

			if alShortName != "" {
				args.short[alShortName] = longName
			}
		}
	}

	return nil
}

// AddMap add supported arguments as map
func (args *Arguments) AddMap(argsMap Map) []error {
	var errs []error

	for name, arg := range argsMap {
		err := args.Add(name, arg)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// GetS get argument value as string
func (args *Arguments) GetS(name string) string {
	longName, _ := parseName(name)
	arg, ok := args.full[longName]

	switch {
	case !ok:
		return ""
	case args.full[longName].Value == nil:
		return ""
	case arg.Type == INT:
		return strconv.Itoa(arg.Value.(int))
	case arg.Type == FLOAT:
		return strconv.FormatFloat(arg.Value.(float64), 'f', -1, 64)
	case arg.Type == BOOL:
		return strconv.FormatBool(arg.Value.(bool))
	default:
		return arg.Value.(string)
	}
}

// GetI get argument value as integer
func (args *Arguments) GetI(name string) int {
	longName, _ := parseName(name)
	arg, ok := args.full[longName]

	switch {
	case !ok:
		return 0

	case args.full[longName].Value == nil:
		return 0

	case arg.Type == STRING:
		result, err := strconv.Atoi(arg.Value.(string))
		if err == nil {
			return result
		}
		return 0

	case arg.Type == FLOAT:
		return int(arg.Value.(float64))

	case arg.Type == BOOL:
		if arg.Value.(bool) {
			return 1
		}
		return 0

	default:
		return arg.Value.(int)
	}
}

// GetB get argument value as boolean
func (args *Arguments) GetB(name string) bool {
	longName, _ := parseName(name)
	arg, ok := args.full[longName]

	switch {
	case !ok:
		return false

	case args.full[longName].Value == nil:
		return false

	case arg.Type == STRING:
		if arg.Value.(string) == "" {
			return false
		}
		return true

	case arg.Type == FLOAT:
		if arg.Value.(float64) > 0 {
			return true
		}
		return false

	case arg.Type == INT:
		if arg.Value.(int) > 0 {
			return true
		}
		return false

	default:
		return arg.Value.(bool)
	}
}

// GetF get argument value as floating number
func (args *Arguments) GetF(name string) float64 {
	longName, _ := parseName(name)
	arg, ok := args.full[longName]

	switch {
	case !ok:
		return 0.0

	case args.full[longName].Value == nil:
		return 0.0

	case arg.Type == STRING:
		result, err := strconv.ParseFloat(arg.Value.(string), 64)
		if err == nil {
			return result
		}
		return 0.0

	case arg.Type == INT:
		return float64(arg.Value.(int))

	case arg.Type == BOOL:
		if arg.Value.(bool) {
			return 1.0
		}
		return 0.0

	default:
		return arg.Value.(float64)
	}
}

// Has check that argument exists and set
func (args *Arguments) Has(name string) bool {
	longName, _ := parseName(name)
	arg, ok := args.full[longName]

	if !ok {
		return false
	}

	if !arg.set {
		return false
	}

	return true
}

// Parse parse arguments
func (args *Arguments) Parse(rawArgs []string, argsMap ...Map) ([]string, []error) {
	var errs []error

	if len(argsMap) != 0 {
		for _, amap := range argsMap {
			errs = append(errs, args.AddMap(amap)...)
		}
	}

	if len(errs) != 0 {
		return []string{}, errs
	}

	return args.parseArgs(rawArgs)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewArguments create new arguments struct
func NewArguments() *Arguments {
	return &Arguments{
		full:        make(Map),
		short:       make(map[string]string),
		initialized: true,
	}
}

// Add add new supported argument
func Add(name string, arg *V) error {
	if global == nil || global.initialized == false {
		global = NewArguments()
	}

	return global.Add(name, arg)
}

// AddMap add supported arguments as map
func AddMap(argsMap Map) []error {
	if global == nil || global.initialized == false {
		global = NewArguments()
	}

	return global.AddMap(argsMap)
}

// GetS get argument value as string
func GetS(name string) string {
	if global == nil || global.initialized == false {
		return ""
	}

	return global.GetS(name)
}

// GetI get argument value as integer
func GetI(name string) int {
	if global == nil || global.initialized == false {
		return 0
	}

	return global.GetI(name)
}

// GetB get argument value as boolean
func GetB(name string) bool {
	if global == nil || global.initialized == false {
		return false
	}

	return global.GetB(name)
}

// GetF get argument value as floating number
func GetF(name string) float64 {
	if global == nil || global.initialized == false {
		return 0.0
	}

	return global.GetF(name)
}

// Has check that argument exists and set
func Has(name string) bool {
	if global == nil || global.initialized == false {
		return false
	}

	return global.Has(name)
}

// Parse parse arguments
func Parse(argsMap ...Map) ([]string, []error) {
	if global == nil || global.initialized == false {
		global = NewArguments()
	}

	return global.Parse(os.Args[1:], argsMap...)
}

// ParseArgName parse combined name and return long and short arguments
func ParseArgName(arg string) (string, string) {
	return parseName(arg)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (args *Arguments) parseArgs(rawArgs []string) ([]string, []error) {
	if len(rawArgs) == 0 {
		return nil, args.getErrorsForRequiredArgs()
	}

	var (
		argName   string
		argList   []string
		errorList []error
	)

	for _, curArg := range rawArgs {
		if argName == "" {
			var (
				curArgName  string
				curArgValue string
				err         error
			)

			var curArgLen = len(curArg)

			switch {
			case strings.TrimRight(curArg, "-") == "":
				argList = append(argList, curArg)
				continue

			case curArgLen > 2 && curArg[0:2] == "--":
				curArgName, curArgValue, err = args.parseLongArgument(curArg[2:curArgLen])

			case curArgLen > 1 && curArg[0:1] == "-":
				curArgName, curArgValue, err = args.parseShortArgument(curArg[1:curArgLen])

			default:
				argList = append(argList, curArg)
				continue
			}

			if err != nil {
				errorList = append(errorList, err)
				continue
			}

			if curArgValue != "" {
				errorList = appendError(
					errorList,
					updateArgument(args.full[curArgName], curArgName, curArgValue),
				)
			} else {
				if args.full[curArgName] != nil && args.full[curArgName].Type == BOOL {
					errorList = appendError(
						errorList,
						updateArgument(args.full[curArgName], curArgName, ""),
					)
				} else {
					argName = curArgName
				}
			}
		} else {
			errorList = appendError(
				errorList,
				updateArgument(args.full[argName], argName, curArg),
			)

			argName = ""
		}
	}

	errorList = append(errorList, args.getErrorsForRequiredArgs()...)

	if argName != "" {
		errorList = append(errorList, ArgumentError{"--" + argName, ERROR_EMPTY_VALUE})
	}

	return argList, errorList
}

func (args *Arguments) parseLongArgument(arg string) (string, string, error) {
	if strings.Contains(arg, "=") {
		argSlice := strings.Split(arg, "=")

		if len(argSlice) <= 1 || argSlice[1] == "" {
			return "", "", ArgumentError{"--" + argSlice[0], ERROR_WRONG_FORMAT}
		}

		return argSlice[0], strings.Join(argSlice[1:len(argSlice)], "="), nil
	}

	if args.full[arg] != nil {
		return arg, "", nil
	}

	return "", "", ArgumentError{"--" + arg, ERROR_UNSUPPORTED}
}

func (args *Arguments) parseShortArgument(arg string) (string, string, error) {
	if strings.Contains(arg, "=") {
		argSlice := strings.Split(arg, "=")

		if len(argSlice) <= 1 || argSlice[1] == "" {
			return "", "", ArgumentError{"-" + argSlice[0], ERROR_WRONG_FORMAT}
		}

		argName := argSlice[0]

		if args.short[argName] == "" {
			return "", "", ArgumentError{"-" + argName, ERROR_UNSUPPORTED}
		}

		return args.short[argName], strings.Join(argSlice[1:len(argSlice)], "="), nil
	}

	if args.short[arg] == "" {
		return "", "", ArgumentError{"-" + arg, ERROR_UNSUPPORTED}
	}

	return args.short[arg], "", nil
}

func (args *Arguments) getErrorsForRequiredArgs() []error {
	if args.hasRequired == false {
		return nil
	}

	var errorList []error

	for n, v := range args.full {
		if v.Required == true && v.Value == nil {
			errorList = append(errorList, ArgumentError{n, ERROR_REQUIRED_NOT_SET})
		}
	}

	return errorList
}

// ////////////////////////////////////////////////////////////////////////////////// //

func initArgs(args *Arguments) {
	args.full = make(Map)
	args.short = make(map[string]string)
	args.initialized = true
}

func parseName(name string) (string, string) {
	na := strings.Split(name, ":")

	if len(na) == 1 {
		return na[0], ""
	}

	return na[1], na[0]
}

func updateArgument(arg *V, name string, value string) error {
	switch arg.Type {
	case STRING:
		return updateStringArgument(arg, value)

	case BOOL:
		return updateBooleanArgument(arg)

	case FLOAT:
		return updateFloatArgument(name, arg, value)

	case INT:
		return updateIntArgument(name, arg, value)
	}

	return fmt.Errorf("Unsuported argument type %d", arg.Type)
}

func updateStringArgument(arg *V, value string) error {
	if arg.set && arg.Mergeble {
		arg.Value = arg.Value.(string) + " " + value
	} else {
		arg.Value = value
		arg.set = true
	}

	return nil
}

func updateBooleanArgument(arg *V) error {
	arg.Value = true
	arg.set = true

	return nil
}

func updateFloatArgument(name string, arg *V, value string) error {
	floatValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return ArgumentError{"--" + name, ERROR_WRONG_FORMAT}
	}

	var resultFloat float64

	if arg.Min != arg.Max {
		resultFloat = mathutil.BetweenF(floatValue, arg.Min, arg.Max)
	} else {
		resultFloat = floatValue
	}

	if arg.set && arg.Mergeble {
		arg.Value = arg.Value.(float64) + resultFloat
	} else {
		arg.Value = resultFloat
		arg.set = true
	}

	return nil
}

func updateIntArgument(name string, arg *V, value string) error {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		return ArgumentError{"--" + name, ERROR_WRONG_FORMAT}
	}

	var resultInt int

	if arg.Min != arg.Max {
		resultInt = mathutil.Between(intValue, int(arg.Min), int(arg.Max))
	} else {
		resultInt = intValue
	}

	if arg.set && arg.Mergeble {
		arg.Value = arg.Value.(int) + resultInt
	} else {
		arg.Value = resultInt
		arg.set = true
	}

	return nil
}

func appendError(errList []error, err error) []error {
	if err == nil {
		return errList
	}

	return append(errList, err)
}

func (e ArgumentError) Error() string {
	switch e.Type {
	default:
		return fmt.Sprintf("Argument %s is not supported", e.Arg)
	case ERROR_EMPTY_VALUE:
		return fmt.Sprintf("Non-boolean argument %s is empty", e.Arg)
	case ERROR_REQUIRED_NOT_SET:
		return fmt.Sprintf("Required argument %s is not set", e.Arg)
	case ERROR_WRONG_FORMAT:
		return fmt.Sprintf("Argument %s has wrong format", e.Arg)
	case ERROR_ARG_IS_NIL:
		return fmt.Sprintf("Struct for argument %s is nil", e.Arg)
	case ERROR_DUPLICATE_LONGNAME, ERROR_DUPLICATE_SHORTNAME:
		return fmt.Sprintf("Argument %s defined 2 or more times", e.Arg)
	case ERROR_NO_NAME:
		return "Some argument does not have a name"
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //
