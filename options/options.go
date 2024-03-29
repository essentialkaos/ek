// Package options provides methods for working with command-line options
package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Options types
const (
	STRING uint8 = iota // String option
	INT                 // Int/Uint option
	BOOL                // Boolean option
	FLOAT               // Floating number option
	MIXED               // String or boolean option
)

// Error codes
const (
	ERROR_UNSUPPORTED = iota
	ERROR_DUPLICATE_LONGNAME
	ERROR_DUPLICATE_SHORTNAME
	ERROR_OPTION_IS_NIL
	ERROR_EMPTY_VALUE
	ERROR_REQUIRED_NOT_SET
	ERROR_WRONG_FORMAT
	ERROR_CONFLICT
	ERROR_BOUND_NOT_SET
	ERROR_UNSUPPORTED_VALUE
)

// ////////////////////////////////////////////////////////////////////////////////// //

// V is basic option struct
type V struct {
	Type      uint8   // option type
	Max       float64 // maximum integer option value
	Min       float64 // minimum integer option value
	Alias     string  // list of aliases
	Conflicts string  // list of conflicts options
	Bound     string  // list of bound options
	Mergeble  bool    // option supports options value merging
	Required  bool    // option is required

	set bool // non-exported field

	Value any // default value
}

// Map is map with list of options
type Map map[string]*V

// Options is options struct
type Options struct {
	short       map[string]string
	initialized bool
	full        Map
}

// OptionError is argument parsing error
type OptionError struct {
	Option      string
	BoundOption string
	Type        int
}

// ////////////////////////////////////////////////////////////////////////////////// //

type optionName struct {
	Long  string
	Short string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilOptions returns if options struct is nil
	ErrNilOptions = fmt.Errorf("Options struct is nil")

	// ErrNilMap returns if options map is nil
	ErrNilMap = fmt.Errorf("Options map is nil")

	// ErrEmptyName returns if option have no name
	ErrEmptyName = fmt.Errorf("One or more options do not have a name")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MergeSymbol used for joining several mergeble options with string value
var MergeSymbol = " "

// ////////////////////////////////////////////////////////////////////////////////// //

// global is global options
var global *Options

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds a new option
func (opts *Options) Add(name string, option *V) error {
	if opts == nil {
		return ErrNilOptions
	}

	if !opts.initialized {
		initOptions(opts)
	}

	optName := parseName(name)

	switch {
	case optName.Long == "":
		return ErrEmptyName
	case option == nil:
		return OptionError{"--" + optName.Long, "", ERROR_OPTION_IS_NIL}
	case opts.full[optName.Long] != nil:
		return OptionError{"--" + optName.Long, "", ERROR_DUPLICATE_LONGNAME}
	case optName.Short != "" && opts.short[optName.Short] != "":
		return OptionError{"-" + optName.Short, "", ERROR_DUPLICATE_SHORTNAME}
	}

	opts.full[optName.Long] = option

	if optName.Short != "" {
		opts.short[optName.Short] = optName.Long
	}

	if option.Alias != "" {
		aliases := parseOptionsList(option.Alias)

		for _, l := range aliases {
			opts.full[l.Long] = option

			if l.Short != "" {
				opts.short[l.Short] = optName.Long
			}
		}
	}

	return nil
}

// AddMap adds map with supported options
func (opts *Options) AddMap(optMap Map) []error {
	if optMap == nil {
		return []error{ErrNilMap}
	}

	var errs []error

	for name, opt := range optMap {
		err := opts.Add(name, opt)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// GetS returns option value as string
func (opts *Options) GetS(name string) string {
	if opts == nil {
		return ""
	}

	optName := parseName(name)
	opt, ok := opts.full[optName.Long]

	switch {
	case !ok:
		return ""
	case opts.full[optName.Long].Value == nil:
		return ""
	case opt.Type == INT:
		return strconv.Itoa(opt.Value.(int))
	case opt.Type == FLOAT:
		return strconv.FormatFloat(opt.Value.(float64), 'f', -1, 64)
	case opt.Type == BOOL:
		return strconv.FormatBool(opt.Value.(bool))
	default:
		return fmt.Sprintf("%s", opt.Value)
	}
}

// GetI returns option value as integer
func (opts *Options) GetI(name string) int {
	if opts == nil {
		return 0
	}

	optName := parseName(name)
	opt, ok := opts.full[optName.Long]

	switch {
	case !ok:
		return 0

	case opts.full[optName.Long].Value == nil:
		return 0

	case opt.Type == STRING, opt.Type == MIXED:
		result, err := strconv.Atoi(opt.Value.(string))
		if err == nil {
			return result
		}
		return 0

	case opt.Type == FLOAT:
		return int(opt.Value.(float64))

	case opt.Type == BOOL:
		if opt.Value.(bool) {
			return 1
		}
		return 0

	default:
		return opt.Value.(int)
	}
}

// GetB returns option value as boolean
func (opts *Options) GetB(name string) bool {
	if opts == nil {
		return false
	}

	optName := parseName(name)
	opt, ok := opts.full[optName.Long]

	switch {
	case !ok:
		return false

	case opts.full[optName.Long].Value == nil:
		return false

	case opt.Type == STRING, opt.Type == MIXED:
		if opt.Value.(string) == "" {
			return false
		}
		return true

	case opt.Type == FLOAT:
		if opt.Value.(float64) > 0 {
			return true
		}
		return false

	case opt.Type == INT:
		if opt.Value.(int) > 0 {
			return true
		}
		return false

	default:
		return opt.Value.(bool)
	}
}

// GetF returns option value as floating number
func (opts *Options) GetF(name string) float64 {
	if opts == nil {
		return 0.0
	}

	optName := parseName(name)
	opt, ok := opts.full[optName.Long]

	switch {
	case !ok:
		return 0.0

	case opts.full[optName.Long].Value == nil:
		return 0.0

	case opt.Type == STRING, opt.Type == MIXED:
		result, err := strconv.ParseFloat(opt.Value.(string), 64)
		if err == nil {
			return result
		}
		return 0.0

	case opt.Type == INT:
		return float64(opt.Value.(int))

	case opt.Type == BOOL:
		if opt.Value.(bool) {
			return 1.0
		}
		return 0.0

	default:
		return opt.Value.(float64)
	}
}

// Split splits mergeble option to it's values
func (opts *Options) Split(name string) []string {
	value := opts.GetS(name)

	if value == "" {
		return nil
	}

	return strings.Split(value, MergeSymbol)
}

// Is checks if option with given name has given value
func (opts *Options) Is(name string, value any) bool {
	if opts == nil {
		return false
	}

	if !opts.Has(name) {
		return false
	}

	switch t := value.(type) {
	case string:
		return opts.GetS(name) == t
	case int:
		return opts.GetI(name) == t
	case float64:
		return opts.GetF(name) == t
	case bool:
		return opts.GetB(name) == t
	}

	return false
}

// Has checks if option with given name exists and set
func (opts *Options) Has(name string) bool {
	if opts == nil {
		return false
	}

	opt, ok := opts.full[parseName(name).Long]

	if !ok {
		return false
	}

	if !opt.set {
		return false
	}

	return true
}

// Parse parses slice with raw options
func (opts *Options) Parse(rawOpts []string, optMap ...Map) (Arguments, []error) {
	if opts == nil {
		return nil, []error{ErrNilOptions}
	}

	var errs []error

	if len(optMap) != 0 {
		for _, m := range optMap {
			errs = append(errs, opts.AddMap(m)...)
		}
	}

	if len(errs) != 0 {
		return Arguments{}, errs
	}

	return opts.parseOptions(rawOpts)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Set set option in map
// Note that if the option is already set, it will be replaced
func (m Map) Set(name string, opt *V) error {
	optName := parseName(name)

	switch {
	case m == nil:
		return ErrNilMap
	case optName.Long == "":
		return ErrEmptyName
	case opt == nil:
		return OptionError{"--" + optName.Long, "", ERROR_OPTION_IS_NIL}
	}

	m[name] = opt

	return nil
}

// Delete removes option from map
func (m Map) Delete(name string) bool {
	if m == nil || m[name] == nil {
		return false
	}

	delete(m, name)

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewOptions creates new options struct
func NewOptions() *Options {
	return &Options{
		full:        make(Map),
		short:       make(map[string]string),
		initialized: true,
	}
}

// Add adds new supported option
func Add(name string, opt *V) error {
	if global == nil || !global.initialized {
		global = NewOptions()
	}

	return global.Add(name, opt)
}

// AddMap adds map with supported options
func AddMap(optMap Map) []error {
	if global == nil || !global.initialized {
		global = NewOptions()
	}

	return global.AddMap(optMap)
}

// GetS returns option value as string
func GetS(name string) string {
	if global == nil || !global.initialized {
		return ""
	}

	return global.GetS(name)
}

// GetI returns option value as integer
func GetI(name string) int {
	if global == nil || !global.initialized {
		return 0
	}

	return global.GetI(name)
}

// GetB returns option value as boolean
func GetB(name string) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.GetB(name)
}

// GetF returns option value as floating number
func GetF(name string) float64 {
	if global == nil || !global.initialized {
		return 0.0
	}

	return global.GetF(name)
}

// Split splits mergeble option to it's values
func Split(name string) []string {
	if global == nil || !global.initialized {
		return nil
	}

	return global.Split(name)
}

// Has checks if option with given name exists and set
func Has(name string) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.Has(name)
}

// Is checks if option with given name has given value
func Is(name string, value any) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.Is(name, value)
}

// Parse parses slice with raw options
func Parse(optMap ...Map) (Arguments, []error) {
	if global == nil || !global.initialized {
		global = NewOptions()
	}

	return global.Parse(os.Args[1:], optMap...)
}

// ParseOptionName parses combined name and returns long and short options
func ParseOptionName(opt string) (string, string) {
	a := parseName(opt)
	return a.Long, a.Short
}

// Format formats option name
func Format(opt string) string {
	a := parseName(opt)

	switch {
	case a.Long == "":
		return ""
	case a.Short == "":
		return "--" + a.Long
	default:
		return fmt.Sprintf("-%s/--%s", a.Short, a.Long)
	}
}

// Merge merges several options into string
func Merge(opts ...string) string {
	return strings.Join(opts, " ")
}

// Q merges several options into string (shortcut for Merge)
func Q(opts ...string) string {
	return Merge(opts...)
}

// F formats option name (shortcut for Format)
func F(opt string) string {
	return Format(opt)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string representation of options map
func (m Map) String() string {
	if m == nil {
		return "options.Map[Nil]"
	}

	if len(m) == 0 {
		return "options.Map[]"
	}

	result := "options.Map["

	for n, v := range m {
		result += parseName(n).Long + ":" + v.String() + " "
	}

	return strings.TrimRight(result, " ") + "]"
}

// String returns string representation of option
func (v *V) String() string {
	if v == nil {
		return "Nil{}"
	}

	var result string

	switch v.Type {
	case STRING:
		result = "String{"
	case INT:
		result = "Int{"
	case BOOL:
		result = "Bool{"
	case FLOAT:
		result = "Float{"
	case MIXED:
		result = "Mixed{"
	default:
		result = "Unknown{"
	}

	if v.Value != nil {
		result += fmt.Sprintf("Value:%v ", v.Value)
	}

	if v.Min != 0 {
		result += fmt.Sprintf("Min:%g ", v.Min)
	}

	if v.Max != 0 {
		result += fmt.Sprintf("Max:%g ", v.Max)
	}

	if v.Alias != "" {
		result += fmt.Sprintf("Alias:%s ", v.Alias)
	}

	if v.Conflicts != "" {
		result += fmt.Sprintf("Conflicts:%s ", v.Conflicts)
	}

	if v.Bound != "" {
		result += fmt.Sprintf("Bound:%s ", v.Bound)
	}

	if v.Mergeble {
		result += "Mergeble:Yes "
	}

	if v.Required {
		result += "Required:Yes "
	}

	return strings.TrimRight(result, " ") + "}"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// I think it is okay to have such a long and complicated method for parsing data
// because it has a lot of logic which can't be separated into different methods
// without losing code readability
// codebeat:disable[LOC,BLOCK_NESTING,CYCLO]

func (opts *Options) parseOptions(rawOpts []string) (Arguments, []error) {
	opts.prepare()

	if len(rawOpts) == 0 {
		return nil, opts.validate()
	}

	var optName string
	var mixedOpt bool
	var arguments Arguments
	var errorList []error

	for _, curOpt := range rawOpts {
		if optName == "" || mixedOpt {
			var (
				curOptName  string
				curOptValue string
				err         error
			)

			var curOptLen = len(curOpt)

			switch {
			case strings.TrimRight(curOpt, "-") == "":
				arguments = append(arguments, Argument(curOpt))
				continue

			case curOptLen > 2 && curOpt[:2] == "--":
				curOptName, curOptValue, err = opts.parseLongOption(curOpt[2:curOptLen])

			case curOptLen > 1 && curOpt[:1] == "-":
				curOptName, curOptValue, err = opts.parseShortOption(curOpt[1:curOptLen])

			case mixedOpt:
				errorList = appendError(
					errorList,
					updateOption(opts.full[optName], optName, curOpt),
				)

				optName, mixedOpt = "", false

			default:
				arguments = append(arguments, Argument(curOpt))
				continue
			}

			if err != nil {
				errorList = append(errorList, err)
				continue
			}

			if curOptName == "" && curOptValue == "" {
				arguments = append(arguments, Argument(curOpt))
				continue
			}

			if curOptName != "" && mixedOpt {
				errorList = appendError(
					errorList,
					updateOption(opts.full[optName], optName, "true"),
				)

				mixedOpt = false
			}

			if curOptValue != "" {
				errorList = appendError(
					errorList,
					updateOption(opts.full[curOptName], curOptName, curOptValue),
				)
			} else {
				switch {
				case opts.full[curOptName] != nil && opts.full[curOptName].Type == BOOL:
					errorList = appendError(
						errorList,
						updateOption(opts.full[curOptName], curOptName, ""),
					)

				case opts.full[curOptName] != nil && opts.full[curOptName].Type == MIXED:
					optName = curOptName
					mixedOpt = true

				default:
					optName = curOptName
				}
			}
		} else {
			errorList = appendError(
				errorList,
				updateOption(opts.full[optName], optName, curOpt),
			)

			optName = ""
		}
	}

	errorList = append(errorList, opts.validate()...)

	if optName != "" {
		if opts.full[optName].Type == MIXED {
			errorList = appendError(
				errorList,
				updateOption(opts.full[optName], optName, "true"),
			)
		} else {
			errorList = append(errorList, OptionError{"--" + optName, "", ERROR_EMPTY_VALUE})
		}
	}

	return arguments, errorList
}

// codebeat:enable[LOC,BLOCK_NESTING,CYCLO]

func (opts *Options) parseLongOption(opt string) (string, string, error) {
	if strings.Contains(opt, "=") {
		optName, optValue, ok := strings.Cut(opt, "=")

		if ok && optValue == "" {
			return "", "", OptionError{"--" + optName, "", ERROR_WRONG_FORMAT}
		}

		if opts.full[optName] == nil {
			return "", "", OptionError{"--" + optName, "", ERROR_UNSUPPORTED}
		}

		return optName, optValue, nil
	}

	if opts.full[opt] != nil {
		return opt, "", nil
	}

	return "", "", OptionError{"--" + opt, "", ERROR_UNSUPPORTED}
}

func (opts *Options) parseShortOption(opt string) (string, string, error) {
	if strings.Contains(opt, "=") {
		optName, optValue, ok := strings.Cut(opt, "=")

		if ok && optValue == "" {
			return "", "", OptionError{"-" + optName, "", ERROR_WRONG_FORMAT}
		}

		if opts.short[optName] == "" {
			return "", "", OptionError{"-" + optName, "", ERROR_UNSUPPORTED}
		}

		return opts.short[optName], optValue, nil
	}

	if len(opt) > 2 {
		return "", "", nil
	}

	if opts.short[opt] == "" {
		return "", "", OptionError{"-" + opt, "", ERROR_UNSUPPORTED}
	}

	return opts.short[opt], "", nil
}

func (opts *Options) prepare() {
	for _, v := range opts.full {
		// String is default type
		if v.Type == STRING && v.Value != nil {
			v.Type = guessType(v.Value)
		}
	}
}

func (opts *Options) validate() []error {
	var errorList []error

	for n, v := range opts.full {
		if !isSupportedType(v.Value) {
			errorList = append(errorList, OptionError{n, "", ERROR_UNSUPPORTED_VALUE})
		}

		if v.Required && v.Value == nil {
			errorList = append(errorList, OptionError{n, "", ERROR_REQUIRED_NOT_SET})
		}

		if v.Conflicts != "" {
			conflicts := parseOptionsList(v.Conflicts)

			for _, c := range conflicts {
				if opts.Has(c.Long) && opts.Has(n) {
					errorList = append(errorList, OptionError{n, c.Long, ERROR_CONFLICT})
				}
			}
		}

		if v.Bound != "" {
			bound := parseOptionsList(v.Bound)

			for _, b := range bound {
				if !opts.Has(b.Long) && opts.Has(n) {
					errorList = append(errorList, OptionError{n, b.Long, ERROR_BOUND_NOT_SET})
				}
			}
		}
	}

	return errorList
}

// ////////////////////////////////////////////////////////////////////////////////// //

func initOptions(opts *Options) {
	opts.full = make(Map)
	opts.short = make(map[string]string)
	opts.initialized = true
}

func parseName(name string) optionName {
	short, long, ok := strings.Cut(name, ":")

	if !ok {
		return optionName{short, ""}
	}

	return optionName{long, short}
}

func parseOptionsList(list string) []optionName {
	var result []optionName

	for _, a := range strings.Split(list, " ") {
		result = append(result, parseName(a))
	}

	return result
}

func updateOption(opt *V, name, value string) error {
	switch opt.Type {
	case STRING, MIXED:
		return updateStringOption(opt, value)

	case BOOL:
		return updateBooleanOption(opt)

	case FLOAT:
		return updateFloatOption(name, opt, value)

	case INT:
		return updateIntOption(name, opt, value)
	}

	return fmt.Errorf("Option %s has unsupported type", Format(name))
}

func updateStringOption(opt *V, value string) error {
	if opt.set && opt.Mergeble {
		opt.Value = opt.Value.(string) + MergeSymbol + value
	} else {
		opt.Value = value
		opt.set = true
	}

	return nil
}

func updateBooleanOption(opt *V) error {
	opt.Value = true
	opt.set = true

	return nil
}

func updateFloatOption(name string, opt *V, value string) error {
	floatValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return OptionError{"--" + name, "", ERROR_WRONG_FORMAT}
	}

	var resultFloat float64

	if opt.Min != opt.Max {
		resultFloat = betweenFloat(floatValue, opt.Min, opt.Max)
	} else {
		resultFloat = floatValue
	}

	if opt.set && opt.Mergeble {
		opt.Value = opt.Value.(float64) + resultFloat
	} else {
		opt.Value = resultFloat
		opt.set = true
	}

	return nil
}

func updateIntOption(name string, opt *V, value string) error {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		return OptionError{"--" + name, "", ERROR_WRONG_FORMAT}
	}

	var resultInt int

	if opt.Min != opt.Max {
		resultInt = betweenInt(intValue, int(opt.Min), int(opt.Max))
	} else {
		resultInt = intValue
	}

	if opt.set && opt.Mergeble {
		opt.Value = opt.Value.(int) + resultInt
	} else {
		opt.Value = resultInt
		opt.set = true
	}

	return nil
}

func appendError(errList []error, err error) []error {
	if err == nil {
		return errList
	}

	return append(errList, err)
}

func betweenInt(val, min, max int) int {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

func betweenFloat(val, min, max float64) float64 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

func isSupportedType(v any) bool {
	switch v.(type) {
	case nil, string, bool, int, float64:
		return true
	}

	return false
}

func guessType(v any) uint8 {
	switch v.(type) {
	case string:
		return STRING
	case bool:
		return BOOL
	case int:
		return INT
	case float64:
		return FLOAT
	}

	return STRING
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (e OptionError) Error() string {
	switch e.Type {
	default:
		return fmt.Sprintf("Option %s is not supported", e.Option)
	case ERROR_EMPTY_VALUE:
		return fmt.Sprintf("Non-boolean option %s is empty", e.Option)
	case ERROR_REQUIRED_NOT_SET:
		return fmt.Sprintf("Required option %s is not set", e.Option)
	case ERROR_WRONG_FORMAT:
		return fmt.Sprintf("Option %s has wrong format", e.Option)
	case ERROR_OPTION_IS_NIL:
		return fmt.Sprintf("Struct for option %s is nil", e.Option)
	case ERROR_DUPLICATE_LONGNAME, ERROR_DUPLICATE_SHORTNAME:
		return fmt.Sprintf("Option %s defined 2 or more times", e.Option)
	case ERROR_CONFLICT:
		return fmt.Sprintf("Option %s conflicts with option %s", e.Option, e.BoundOption)
	case ERROR_BOUND_NOT_SET:
		return fmt.Sprintf("Option %s must be defined with option %s", e.BoundOption, e.Option)
	case ERROR_UNSUPPORTED_VALUE:
		return fmt.Sprintf("Option %s contains unsupported default value", e.Option)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //
