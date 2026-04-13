// Package options provides methods for working with command-line options
package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v14/errors"
	"github.com/essentialkaos/ek/v14/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Option type constants define the expected value kind for a parsed option
const (
	STRING uint8 = iota // String option
	INT                 // Int/Uint option
	BOOL                // Boolean option
	FLOAT               // Floating number option
	MIXED               // String or boolean option
)

// Error code constants identify the specific reason an option parsing error occurred
const (
	ERROR_UNSUPPORTED                      = iota // Option is not registered
	ERROR_DUPLICATE_LONGNAME                      // Long name is already defined
	ERROR_DUPLICATE_SHORTNAME                     // Short name is already defined
	ERROR_OPTION_IS_NIL                           // Option struct pointer is nil
	ERROR_EMPTY_VALUE                             // Non-boolean option has no value
	ERROR_WRONG_FORMAT                            // Option value has invalid format
	ERROR_CONFLICT                                // Option conflicts with another option
	ERROR_BOUND_NOT_SET                           // Required bound option is missing
	ERROR_UNSUPPORTED_VALUE                       // Default value type is not supported
	ERROR_UNSUPPORTED_ALIAS_LIST_FORMAT           // Alias list has unsupported format
	ERROR_UNSUPPORTED_CONFLICT_LIST_FORMAT        // Conflict list has unsupported format
	ERROR_UNSUPPORTED_BOUND_LIST_FORMAT           // Bound list has unsupported format
)

// ////////////////////////////////////////////////////////////////////////////////// //

// V is a shorthand alias for [Option]
type V = Option

// Option holds the definition and current value of a single command-line option
type Option struct {
	Type      uint8   // Option type ([STRING], [INT], [BOOL], [FLOAT], or [MIXED])
	Max       float64 // Maximum allowed value for numeric options
	Min       float64 // Minimum allowed value for numeric options
	Alias     any     // Additional name(s) for the option; string or []string
	Conflicts any     // Option name(s) that cannot be used together with this one; string or []string
	Bound     any     // Option name(s) that must be set together with this one; string or []string
	Mergeble  bool    // If true, repeated occurrences of the option are merged into one value

	set bool // indicates whether the option was explicitly set during parsing

	Value any // Default value; overwritten by the parsed value
}

// Map is a mapping of option name strings to their Option definitions
type Map map[string]*Option

// Options holds all registered options and their parsed state
type Options struct {
	short       map[string]string
	initialized bool
	full        Map
}

// OptionError describes an error encountered while registering or parsing an option
type OptionError struct {
	Option      string // Name of the option that caused the error
	BoundOption string // Name of the related bound or conflicting option, if applicable
	Type        int    // Error code (one of the ERROR_* constants)
}

// ////////////////////////////////////////////////////////////////////////////////// //

type optionName struct {
	Long  string
	Short string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilOptions is returned when a method is called on a nil [Options] pointer
	ErrNilOptions = errors.New("options struct is nil")

	// ErrNilMap is returned when a nil [Map] is passed to [AddMap]
	ErrNilMap = errors.New("options map is nil")

	// ErrEmptyName is returned when an option is registered without a name
	ErrEmptyName = errors.New("one or more options do not have a name")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MergeSymbol is the separator used to join values of mergeable string options
var MergeSymbol = " "

// ////////////////////////////////////////////////////////////////////////////////// //

// global is global options
var global *Options

// ////////////////////////////////////////////////////////////////////////////////// //

// Add registers a new option under the given name in the options set.
// The name may be in "short:long" format to register both a short and long form.
func (o *Options) Add(name string, option *Option) error {
	if o == nil {
		return ErrNilOptions
	}

	if !o.initialized {
		initOptions(o)
	}

	optName := parseName(name)

	switch {
	case optName.Long == "":
		return ErrEmptyName
	case option == nil:
		return OptionError{"--" + optName.Long, "", ERROR_OPTION_IS_NIL}
	case o.full[optName.Long] != nil:
		return OptionError{"--" + optName.Long, "", ERROR_DUPLICATE_LONGNAME}
	case optName.Short != "" && o.short[optName.Short] != "":
		return OptionError{"-" + optName.Short, "", ERROR_DUPLICATE_SHORTNAME}
	}

	o.full[optName.Long] = option

	if optName.Short != "" {
		o.short[optName.Short] = optName.Long
	}

	if option.Alias != "" {
		aliases, ok := parseOptionsList(option.Alias)

		if !ok {
			return OptionError{"--" + optName.Long, "", ERROR_UNSUPPORTED_ALIAS_LIST_FORMAT}
		}

		for _, l := range aliases {
			o.full[l.Long] = option

			if l.Short != "" {
				o.short[l.Short] = optName.Long
			}
		}
	}

	return nil
}

// AddMap registers all options from the provided Map, collecting any errors encountered
func (o *Options) AddMap(optMap Map) errors.Errors {
	if optMap == nil {
		return errors.Errors{ErrNilMap}
	}

	var errs errors.Errors

	for name, opt := range optMap {
		err := o.Add(name, opt)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// GetS returns the value of the named option as a string.
// Returns an empty string if the option is unset or not found.
func (o *Options) GetS(name string) string {
	if o == nil || len(o.full) == 0 {
		return ""
	}

	optName := parseName(name)
	opt, ok := o.full[optName.Long]

	switch {
	case !ok:
		return ""
	case o.full[optName.Long].Value == nil:
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

// GetI returns the value of the named option as an integer.
// Returns 0 if the option is unset, not found, or cannot be converted.
func (o *Options) GetI(name string) int {
	if o == nil || len(o.full) == 0 {
		return 0
	}

	optName := parseName(name)
	opt, ok := o.full[optName.Long]

	switch {
	case !ok:
		return 0

	case o.full[optName.Long].Value == nil:
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

// GetB returns the value of the named option as a boolean.
// Returns false if the option is unset, not found, or has an empty/zero value.
func (o *Options) GetB(name string) bool {
	if o == nil || len(o.full) == 0 {
		return false
	}

	optName := parseName(name)
	opt, ok := o.full[optName.Long]

	switch {
	case !ok:
		return false

	case o.full[optName.Long].Value == nil:
		return false

	case opt.Type == STRING, opt.Type == MIXED:
		return opt.Value.(string) != ""

	case opt.Type == FLOAT:
		return opt.Value.(float64) > 0

	case opt.Type == INT:
		return opt.Value.(int) > 0

	default:
		return opt.Value.(bool)
	}
}

// GetF returns the value of the named option as a float64.
// Returns 0.0 if the option is unset, not found, or cannot be converted.
func (o *Options) GetF(name string) float64 {
	if o == nil || len(o.full) == 0 {
		return 0.0
	}

	optName := parseName(name)
	opt, ok := o.full[optName.Long]

	switch {
	case !ok:
		return 0.0

	case o.full[optName.Long].Value == nil:
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

// Split splits the value of a mergeable option into its individual components.
// Returns nil if the option is unset or empty.
func (o *Options) Split(name string) []string {
	value := o.GetS(name)

	if value == "" {
		return nil
	}

	return strings.Split(value, MergeSymbol)
}

// Is reports whether the named option's current value equals the given value
func (o *Options) Is(name string, value any) bool {
	if o == nil || len(o.full) == 0 {
		return false
	}

	if !o.Has(name) {
		return false
	}

	switch t := value.(type) {
	case string:
		return o.GetS(name) == t
	case int:
		return o.GetI(name) == t
	case float64:
		return o.GetF(name) == t
	case bool:
		return o.GetB(name) == t
	}

	return false
}

// Has reports whether the named option was explicitly set during parsing
func (o *Options) Has(name string) bool {
	if o == nil || len(o.full) == 0 {
		return false
	}

	opt, ok := o.full[parseName(name).Long]

	return ok && opt.set
}

// Delete removes the named option from the options set and returns true on success.
//
// You can use this method to remove options with private data such as passwords,
// tokens, etc.
func (o *Options) Delete(name string) bool {
	if o == nil || len(o.full) == 0 {
		return false
	}

	optName := parseName(name)
	_, ok := o.full[optName.Long]

	if !ok {
		return false
	}

	delete(o.full, optName.Long)

	return true
}

// Parse parses the given raw argument slice, optionally registering additional
// option maps first. Returns the non-option arguments and any errors encountered.
func (o *Options) Parse(data []string, optMap ...Map) (Arguments, errors.Errors) {
	if o == nil {
		return nil, errors.Errors{ErrNilOptions}
	}

	var errs errors.Errors

	if len(optMap) != 0 {
		for _, m := range optMap {
			errs = append(errs, o.AddMap(m)...)
		}
	}

	if len(errs) != 0 {
		return Arguments{}, errs
	}

	return o.parseOptions(data)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Set stores the given option under the specified name, replacing any existing entry.
// Panics if the map, name, or option is nil/empty.
func (m Map) Set(name string, opt *Option) Map {
	optName := parseName(name)

	switch {
	case m == nil:
		panic(ErrNilMap.Error())
	case optName.Long == "":
		panic(ErrEmptyName.Error())
	case opt == nil:
		panic(OptionError{"--" + optName.Long, "", ERROR_OPTION_IS_NIL}.Error())
	}

	m[name] = opt

	return m
}

// SetIf stores the option only when cond is true; it is a no-op otherwise
func (m Map) SetIf(cond bool, name string, opt *Option) Map {
	if cond {
		m.Set(name, opt)
	}

	return m
}

// Delete removes the named option from the map and returns true if it was present
func (m Map) Delete(name string) bool {
	if m == nil || m[name] == nil {
		return false
	}

	delete(m, name)

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewOptions creates and returns a new, initialized [Options] instance
func NewOptions() *Options {
	return &Options{
		full:        make(Map),
		short:       make(map[string]string),
		initialized: true,
	}
}

// Add registers a new option in the global options set
func Add(name string, opt *Option) error {
	if global == nil || !global.initialized {
		global = NewOptions()
	}

	return global.Add(name, opt)
}

// AddMap registers all options from the provided Map in the global options set
func AddMap(optMap Map) errors.Errors {
	if global == nil || !global.initialized {
		global = NewOptions()
	}

	return global.AddMap(optMap)
}

// GetS returns the value of the named global option as a string
func GetS(name string) string {
	if global == nil || !global.initialized {
		return ""
	}

	return global.GetS(name)
}

// GetI returns the value of the named global option as an integer
func GetI(name string) int {
	if global == nil || !global.initialized {
		return 0
	}

	return global.GetI(name)
}

// GetB returns the value of the named global option as a boolean
func GetB(name string) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.GetB(name)
}

// GetF returns the value of the named global option as a float64
func GetF(name string) float64 {
	if global == nil || !global.initialized {
		return 0.0
	}

	return global.GetF(name)
}

// Split splits the value of a mergeable global option into its individual
// components
func Split(name string) []string {
	if global == nil || !global.initialized {
		return nil
	}

	return global.Split(name)
}

// Has reports whether the named global option was explicitly set during parsing
func Has(name string) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.Has(name)
}

// Is reports whether the named global option's current value equals the given value
func Is(name string, value any) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.Is(name, value)
}

// Delete removes the named option from the global options set and returns true
// on success.
//
// You can use this method to remove options with private data such as passwords,
// tokens, etc.
func Delete(name string) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.Delete(name)
}

// Parse parses os.Args[1:] using the global options set, registering any
// provided maps first. Returns the non-option arguments and any errors encountered.
func Parse(optMap ...Map) (Arguments, errors.Errors) {
	if global == nil || !global.initialized {
		global = NewOptions()
	}

	return global.Parse(os.Args[1:], optMap...)
}

// ParseOptionName parses a combined "short:long" option name and returns the long
// and short parts
func ParseOptionName(opt string) (string, string) {
	a := parseName(opt)
	return a.Long, a.Short
}

// Format returns the canonical CLI representation of an option name (e.g.
// "--verbose" or "-v/--verbose")
func Format(opt string) string {
	return parseName(opt).String()
}

// Merge joins multiple option values into a single space-separated string
func Merge(opts ...string) string {
	return strings.Join(opts, " ")
}

// Q joins multiple option values into a single string; shortcut for [Merge]
func Q(opts ...string) string {
	return Merge(opts...)
}

// F returns the canonical CLI representation of an option name; shortcut for [Format]
func F(opt string) string {
	return Format(opt)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns a human-readable representation of the options map
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

// String returns a human-readable representation of the option including its type
// and value
func (v *Option) String() string {
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

	if v.Alias != nil {
		result += fmt.Sprintf("Alias:%v ", formatOptionsList(v.Alias))
	}

	if v.Conflicts != nil {
		result += fmt.Sprintf("Conflicts:%v ", formatOptionsList(v.Conflicts))
	}

	if v.Bound != nil {
		result += fmt.Sprintf("Bound:%v ", formatOptionsList(v.Bound))
	}

	if v.Mergeble {
		result += "Mergeble:Yes "
	}

	return strings.TrimRight(result, " ") + "}"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string representation of optionName
func (o optionName) String() string {
	switch {
	case o.Long == "":
		return ""
	case o.Short == "":
		return "--" + o.Long
	}

	return fmt.Sprintf("-%s/--%s", o.Short, o.Long)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseOptions parses slice with raw options data
func (o *Options) parseOptions(data []string) (Arguments, errors.Errors) {
	o.prepare()

	if len(data) == 0 {
		return nil, o.validate()
	}

	var optName string
	var mixedOpt bool
	var arguments Arguments
	var errs errors.Errors

LOOP:
	for index, curOpt := range data {
		if optName == "" || mixedOpt {
			var err error
			var curOptName, curOptValue string

			curOptLen := len(curOpt)

			switch {
			case curOpt == "--":
				arguments = append(arguments, NewArguments(data[index:]...)...)
				break LOOP

			case strings.TrimRight(curOpt, "-") == "":
				arguments = append(arguments, Argument(curOpt))
				continue

			case curOptLen > 2 && curOpt[:2] == "--":
				curOptName, curOptValue, err = o.parseLongOption(curOpt[2:curOptLen])

			case curOptLen > 1 && curOpt[:1] == "-":
				curOptName, curOptValue, err = o.parseShortOption(curOpt[1:curOptLen])

			case mixedOpt:
				errs = appendError(errs,
					updateOption(o.full[optName], optName, curOpt),
				)

				optName, mixedOpt = "", false

			default:
				arguments = append(arguments, Argument(curOpt))
				continue
			}

			if err != nil {
				errs = append(errs, err)
				continue
			}

			if curOptName == "" && curOptValue == "" {
				arguments = append(arguments, Argument(curOpt))
				continue
			}

			if curOptName != "" && mixedOpt {
				errs = appendError(errs,
					updateOption(o.full[optName], optName, "true"),
				)

				mixedOpt = false
			}

			if curOptValue != "" {
				errs = appendError(errs,
					updateOption(o.full[curOptName], curOptName, curOptValue),
				)
			} else {
				switch {
				case o.full[curOptName] != nil && o.full[curOptName].Type == BOOL:
					errs = appendError(errs,
						updateOption(o.full[curOptName], curOptName, ""),
					)

				case o.full[curOptName] != nil && o.full[curOptName].Type == MIXED:
					optName = curOptName
					mixedOpt = true

				default:
					optName = curOptName
				}
			}
		} else {
			errs = appendError(errs,
				updateOption(o.full[optName], optName, curOpt),
			)

			optName = ""
		}
	}

	errs = append(errs, o.validate()...)

	if optName != "" {
		opt, ok := o.full[optName]
		if ok && opt.Type == MIXED {
			errs = appendError(errs,
				updateOption(o.full[optName], optName, "true"),
			)
		} else {
			errs = append(errs, OptionError{"--" + optName, "", ERROR_EMPTY_VALUE})
		}
	}

	return arguments, errs
}

// parseLongOption parses long option (started with --) and returns option name
// and value
func (o *Options) parseLongOption(opt string) (string, string, error) {
	if strings.Contains(opt, "=") {
		optName, optValue, ok := strings.Cut(opt, "=")

		if ok && optValue == "" {
			return "", "", OptionError{"--" + optName, "", ERROR_WRONG_FORMAT}
		}

		if o.full[optName] == nil {
			return "", "", OptionError{"--" + optName, "", ERROR_UNSUPPORTED}
		}

		return optName, optValue, nil
	}

	if o.full[opt] != nil {
		return opt, "", nil
	}

	return "", "", OptionError{"--" + opt, "", ERROR_UNSUPPORTED}
}

// parseShortOption parses short option (started with -) and returns option name
// and value
func (o *Options) parseShortOption(opt string) (string, string, error) {
	if strings.Contains(opt, "=") {
		optName, optValue, ok := strings.Cut(opt, "=")

		if ok && optValue == "" {
			return "", "", OptionError{"-" + optName, "", ERROR_WRONG_FORMAT}
		}

		if o.short[optName] == "" {
			return "", "", OptionError{"-" + optName, "", ERROR_UNSUPPORTED}
		}

		return o.short[optName], optValue, nil
	}

	if len(opt) > 2 {
		return "", "", nil
	}

	if o.short[opt] == "" {
		return "", "", OptionError{"-" + opt, "", ERROR_UNSUPPORTED}
	}

	return o.short[opt], "", nil
}

func (o *Options) prepare() {
	for _, v := range o.full {
		// String is default type
		if v.Type == STRING && v.Value != nil {
			v.Type = guessType(v.Value)
		}
	}
}

// validate checks options for errors
func (o *Options) validate() errors.Errors {
	var errs errors.Errors

	for n, v := range o.full {
		if !isSupportedType(v.Value) {
			errs = append(errs, OptionError{F(n), "", ERROR_UNSUPPORTED_VALUE})
		}

		if v.Conflicts != "" {
			conflicts, ok := parseOptionsList(v.Conflicts)

			if !ok {
				errs = append(errs, OptionError{F(n), "", ERROR_UNSUPPORTED_CONFLICT_LIST_FORMAT})
			} else {
				for _, c := range conflicts {
					if o.Has(c.Long) && o.Has(n) {
						errs = append(errs, OptionError{F(n), F(c.Long), ERROR_CONFLICT})
					}
				}
			}
		}

		if v.Bound != "" {
			bound, ok := parseOptionsList(v.Bound)

			if !ok {
				errs = append(errs, OptionError{F(n), "", ERROR_UNSUPPORTED_BOUND_LIST_FORMAT})
			} else {
				for _, b := range bound {
					if !o.Has(b.Long) && o.Has(n) {
						errs = append(errs, OptionError{F(n), F(b.Long), ERROR_BOUND_NOT_SET})
					}
				}
			}
		}
	}

	return errs
}

// ////////////////////////////////////////////////////////////////////////////////// //

// initOptions initializes options struct
func initOptions(opts *Options) {
	opts.full = make(Map)
	opts.short = make(map[string]string)
	opts.initialized = true
}

// parseName parses option name and returns long and short names
func parseName(name string) optionName {
	short, long, ok := strings.Cut(name, ":")

	if !ok {
		return optionName{short, ""}
	}

	return optionName{long, short}
}

// parseOptionsList parses options list and returns slice with option names
func parseOptionsList(list any) ([]optionName, bool) {
	var result []optionName

	switch t := list.(type) {
	case nil:
		return nil, true

	case string:
		for _, a := range strings.Split(t, MergeSymbol) {
			result = append(result, parseName(a))
		}

	case []string:
		for _, a := range t {
			result = append(result, parseName(a))
		}

	default:
		return nil, false
	}

	return result, true
}

// formatOptionsList formats options list to string
func formatOptionsList(list any) string {
	opts, ok := parseOptionsList(list)

	if !ok {
		return "{InvalidList}"
	}

	switch len(opts) {
	case 0:
		return "{Empty}"
	case 1:
		return opts[0].String()
	}

	return fmt.Sprintf("%v", opts)
}

// updateOption updates option value in options map
func updateOption(opt *Option, name, value string) error {
	if opt == nil {
		return OptionError{"--" + name, "", ERROR_UNSUPPORTED}
	}

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

	return fmt.Errorf("option %q has unsupported type", Format(name))
}

// updateStringOption updates string option value
func updateStringOption(opt *Option, value string) error {
	if opt.set && opt.Mergeble {
		opt.Value = opt.Value.(string) + MergeSymbol + value
	} else {
		opt.Value = value
		opt.set = true
	}

	return nil
}

// updateBooleanOption updates boolean option value
func updateBooleanOption(opt *Option) error {
	opt.Value = true
	opt.set = true

	return nil
}

// updateFloatOption updates float option value
func updateFloatOption(name string, opt *Option, value string) error {
	floatValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return OptionError{"--" + name, "", ERROR_WRONG_FORMAT}
	}

	var resultFloat float64

	if opt.Min != opt.Max {
		resultFloat = mathutil.Between(floatValue, opt.Min, opt.Max)
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

// updateIntOption updates integer option value
func updateIntOption(name string, opt *Option, value string) error {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		return OptionError{"--" + name, "", ERROR_WRONG_FORMAT}
	}

	var resultInt int

	if opt.Min != opt.Max {
		resultInt = mathutil.Between(intValue, int(opt.Min), int(opt.Max))
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

// appendError appends error to errors slice
func appendError(errs errors.Errors, err error) errors.Errors {
	if err == nil {
		return errs
	}

	return append(errs, err)
}

// isSupportedType checks if value has supported type for options
func isSupportedType(v any) bool {
	switch v.(type) {
	case nil, string, bool, int, float64:
		return true
	}

	return false
}

// guessType guesses type of value based on its type
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

// Error returns string representation of OptionError
func (e OptionError) Error() string {
	switch e.Type {
	default:
		return fmt.Sprintf("option %q is not supported", e.Option)
	case ERROR_EMPTY_VALUE:
		return fmt.Sprintf("non-boolean option %q is empty", e.Option)
	case ERROR_WRONG_FORMAT:
		return fmt.Sprintf("option %q has wrong format", e.Option)
	case ERROR_OPTION_IS_NIL:
		return fmt.Sprintf("struct for option %q is nil", e.Option)
	case ERROR_DUPLICATE_LONGNAME, ERROR_DUPLICATE_SHORTNAME:
		return fmt.Sprintf("option %q defined 2 or more times", e.Option)
	case ERROR_CONFLICT:
		return fmt.Sprintf("option %q conflicts with option %q", e.Option, e.BoundOption)
	case ERROR_BOUND_NOT_SET:
		return fmt.Sprintf("option %q must be defined with option %q", e.BoundOption, e.Option)
	case ERROR_UNSUPPORTED_VALUE:
		return fmt.Sprintf("option %q contains unsupported default value", e.Option)
	case ERROR_UNSUPPORTED_ALIAS_LIST_FORMAT:
		return fmt.Sprintf("option %q contains unsupported list format of aliases", e.Option)
	case ERROR_UNSUPPORTED_CONFLICT_LIST_FORMAT:
		return fmt.Sprintf("option %q contains unsupported list format of conflicting options", e.Option)
	case ERROR_UNSUPPORTED_BOUND_LIST_FORMAT:
		return fmt.Sprintf("option %q contains unsupported list format of bound options", e.Option)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //
