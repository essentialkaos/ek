// Package options provides methods for working with command-line options
package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/errors"
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
	ERROR_WRONG_FORMAT
	ERROR_CONFLICT
	ERROR_BOUND_NOT_SET
	ERROR_UNSUPPORTED_VALUE
	ERROR_UNSUPPORTED_ALIAS_LIST_FORMAT
	ERROR_UNSUPPORTED_CONFLICT_LIST_FORMAT
	ERROR_UNSUPPORTED_BOUND_LIST_FORMAT
)

// ////////////////////////////////////////////////////////////////////////////////// //

// V is basic option struct
type V struct {
	Type      uint8   // option type
	Max       float64 // maximum integer option value
	Min       float64 // minimum integer option value
	Alias     any     // string or slice of strings with aliases
	Conflicts any     // string or slice of strings with conflicts options
	Bound     any     // string or slice of strings with bound options
	Mergeble  bool    // option supports options value merging

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
	// ErrNilOptions is returned if options struct is nil
	ErrNilOptions = fmt.Errorf("Options struct is nil")

	// ErrNilMap is returned if options map is nil
	ErrNilMap = fmt.Errorf("Options map is nil")

	// ErrEmptyName is returned if option have no name
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
func (o *Options) Add(name string, option *V) error {
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

// AddMap adds map with supported options
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

// GetS returns option value as string
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

// GetI returns option value as integer
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

// GetB returns option value as boolean
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

// Split splits mergeble option to it's values
func (o *Options) Split(name string) []string {
	value := o.GetS(name)

	if value == "" {
		return nil
	}

	return strings.Split(value, MergeSymbol)
}

// Is checks if option with given name has given value
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

// Has checks if option with given name exists and set
func (o *Options) Has(name string) bool {
	if o == nil || len(o.full) == 0 {
		return false
	}

	opt, ok := o.full[parseName(name).Long]

	if !ok {
		return false
	}

	if !opt.set {
		return false
	}

	return true
}

// Delete deletes option with given name
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

// Parse parses slice with raw options
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
func AddMap(optMap Map) errors.Errors {
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

// Delete deletes option with given name
//
// You can use this method to remove options with private data such as passwords,
// tokens, etc.
func Delete(name string) bool {
	if global == nil || !global.initialized {
		return false
	}

	return global.Delete(name)
}

// Parse parses slice with raw options
func Parse(optMap ...Map) (Arguments, errors.Errors) {
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
	return parseName(opt).String()
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

func (o *Options) parseOptions(data []string) (Arguments, errors.Errors) {
	o.prepare()

	if len(data) == 0 {
		return nil, o.validate()
	}

	var optName string
	var mixedOpt bool
	var arguments Arguments
	var errs errors.Errors

	for _, curOpt := range data {
		if optName == "" || mixedOpt {
			var err error
			var curOptName, curOptValue string

			curOptLen := len(curOpt)

			switch {
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
		if o.full[optName].Type == MIXED {
			errs = appendError(errs,
				updateOption(o.full[optName], optName, "true"),
			)
		} else {
			errs = append(errs, OptionError{"--" + optName, "", ERROR_EMPTY_VALUE})
		}
	}

	return arguments, errs
}

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

	return fmt.Errorf("Option %q has unsupported type", Format(name))
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

func appendError(errs errors.Errors, err error) errors.Errors {
	if err == nil {
		return errs
	}

	return append(errs, err)
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
		return fmt.Sprintf("Option %q is not supported", e.Option)
	case ERROR_EMPTY_VALUE:
		return fmt.Sprintf("Non-boolean option %q is empty", e.Option)
	case ERROR_WRONG_FORMAT:
		return fmt.Sprintf("Option %q has wrong format", e.Option)
	case ERROR_OPTION_IS_NIL:
		return fmt.Sprintf("Struct for option %q is nil", e.Option)
	case ERROR_DUPLICATE_LONGNAME, ERROR_DUPLICATE_SHORTNAME:
		return fmt.Sprintf("Option %q defined 2 or more times", e.Option)
	case ERROR_CONFLICT:
		return fmt.Sprintf("Option %q conflicts with option %q", e.Option, e.BoundOption)
	case ERROR_BOUND_NOT_SET:
		return fmt.Sprintf("Option %q must be defined with option %q", e.BoundOption, e.Option)
	case ERROR_UNSUPPORTED_VALUE:
		return fmt.Sprintf("Option %q contains unsupported default value", e.Option)
	case ERROR_UNSUPPORTED_ALIAS_LIST_FORMAT:
		return fmt.Sprintf("Option %q contains unsupported list format of aliases", e.Option)
	case ERROR_UNSUPPORTED_CONFLICT_LIST_FORMAT:
		return fmt.Sprintf("Option %q contains unsupported list format of conflicting options", e.Option)
	case ERROR_UNSUPPORTED_BOUND_LIST_FORMAT:
		return fmt.Sprintf("Option %q contains unsupported list format of bound options", e.Option)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //
