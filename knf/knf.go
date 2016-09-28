// Package knf provides methods for working with configs in KNF format
package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v4/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_COMMENT_SYMBOL   = "#"
	_SECTION_SYMBOL   = "["
	_SEPARATOR_SYMBOL = ":"
	_MACRO_SYMBOL     = "{"
	_DELIMITER        = ":"
)

const _MACRO_REGEXP = "{([a-zA-Z0-9_-]{2,}):([a-zA-Z0-9_-]{2,})}"

// ////////////////////////////////////////////////////////////////////////////////// //

// Config is basic config struct
type Config struct {
	sections []string
	props    []string
	data     map[string]string
	file     string
}

// Validator is config propery validator struct
type Validator struct {
	Property string            // Property name
	Func     PropertyValidator // Validation func
	Value    interface{}       // Expected value
}

// ////////////////////////////////////////////////////////////////////////////////// //

// RegExp strcut for searching and parsing macroses
var macroRE = regexp.MustCompile(_MACRO_REGEXP)

// Global config struct
var global *Config

// ////////////////////////////////////////////////////////////////////////////////// //

// Global read and parse config file
// global config will be accessible globally from any part of code
func Global(file string) error {
	config, err := Read(file)

	if err != nil {
		return err
	}

	global = config

	return nil
}

// Read read and parse config file
func Read(file string) (*Config, error) {
	switch {
	case fsutil.IsExist(file) == false:
		return nil, errors.New("File " + file + " is not exist")
	case fsutil.IsNonEmpty(file) == false:
		return nil, errors.New("File " + file + " is empty")
	case fsutil.IsReadable(file) == false:
		return nil, errors.New("File " + file + " is not readable")
	}

	config := &Config{
		data:     make(map[string]string),
		sections: make([]string, 0),
		props:    make([]string, 0),
		file:     file,
	}

	fd, err := os.Open(path.Clean(file))

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	var sectionName = ""

	reader := bufio.NewReader(fd)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.Replace(line, " ", "", -1) == "" {
			continue
		}

		if strings.HasPrefix(strings.Trim(line, " "), _COMMENT_SYMBOL) {
			continue
		}

		if strings.HasPrefix(strings.Trim(line, " "), _SECTION_SYMBOL) {
			sectionName = strings.Trim(line, "[ ]")
			config.data[sectionName+_DELIMITER] = "true"
			config.sections = append(config.sections, sectionName)
			continue
		}

		if sectionName == "" {
			return nil, errors.New("Configuration file " + file + " is malformed")
		}

		propName, propValue := parseRecord(line, config)
		fullPropName := sectionName + _DELIMITER + propName

		config.props = append(config.props, fullPropName)
		config.data[fullPropName] = propValue
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

// Reload read and parse global config file
func Reload() (map[string]bool, error) {
	if global == nil {
		return nil, errors.New("Global config is not loaded")
	}

	return global.Reload()
}

// GetS return global config value as string
func GetS(name string, defvals ...string) string {
	if global == nil {
		return ""
	}

	return global.GetS(name, defvals...)
}

// GetI return global config value as string
func GetI(name string, defvals ...int) int {
	if global == nil {
		return 0
	}

	return global.GetI(name, defvals...)
}

// GetF return global config value as floating number
func GetF(name string, defvals ...float64) float64 {
	if global == nil {
		return 0.0
	}

	return global.GetF(name, defvals...)
}

// GetB return global config value as boolean
func GetB(name string, defvals ...bool) bool {
	if global == nil {
		return false
	}

	return global.GetB(name, defvals...)
}

// GetM return global config value as file mode
func GetM(name string, defvals ...os.FileMode) os.FileMode {
	if global == nil {
		return os.FileMode(0)
	}

	return global.GetM(name, defvals...)
}

// HasSection check if section exist
func HasSection(section string) bool {
	if global == nil {
		return false
	}

	return global.HasSection(section)
}

// HasProp check if property exist
func HasProp(name string) bool {
	if global == nil {
		return false
	}

	return global.HasProp(name)
}

// Sections return slice with section names
func Sections() []string {
	if global == nil {
		return []string{}
	}

	return global.Sections()
}

// Props return slice with properties names in some section
func Props(section string) []string {
	if global == nil {
		return []string{}
	}

	return global.Props(section)
}

// Validate require slice with pointers to validators and
// return slice with validation errors
func Validate(validators []*Validator) []error {
	if global == nil {
		return []error{errors.New("Global config struct is nil")}
	}

	return global.Validate(validators)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reload read and parse config file
func (c *Config) Reload() (map[string]bool, error) {
	if c == nil {
		return nil, errors.New("Config is nil")
	}

	if c.file == "" {
		return nil, errors.New("Path to config file is empty (non initialized struct?)")
	}

	nc, err := Read(c.file)

	if err != nil {
		return nil, err
	}

	changes := make(map[string]bool)

	for prop, value := range c.data {
		changes[prop] = value != nc.data[prop]
	}

	c.data, c.sections = nc.data, nc.sections

	return changes, nil
}

// GetS return config value as string
func (c *Config) GetS(name string, defvals ...string) string {
	if c == nil {
		return ""
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	return val
}

// GetI return config value as string
func (c *Config) GetI(name string, defvals ...int) int {
	if c == nil {
		return 0
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	// HEX Parsing
	if len(val) >= 3 && val[0:2] == "0x" {
		valHex, err := strconv.ParseInt(val[2:], 16, 0)

		if err != nil {
			return 0
		}

		return int(valHex)
	}

	valInt, err := strconv.Atoi(val)

	if err != nil {
		return 0
	}

	return valInt
}

// GetF return config value as floating number
func (c *Config) GetF(name string, defvals ...float64) float64 {
	if c == nil {
		return 0.0
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	valFl, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return 0.0
	}

	return valFl
}

// GetB return config value as boolean
func (c *Config) GetB(name string, defvals ...bool) bool {
	if c == nil {
		return false
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	switch {
	case val == "", val == "0", val == "false":
		return false
	default:
		return true
	}
}

// GetM return config value as file mode
func (c *Config) GetM(name string, defvals ...os.FileMode) os.FileMode {
	if c == nil {
		return 0
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	valM, err := strconv.ParseUint(val, 8, 32)

	if err != nil {
		return 0
	}

	return os.FileMode(valM)
}

// HasSection check if section exist
func (c *Config) HasSection(section string) bool {
	if c == nil {
		return false
	}

	return c.data[section+_DELIMITER] == "true"
}

// HasProp check if property exist
func (c *Config) HasProp(name string) bool {
	if c == nil {
		return false
	}

	return c.data[name] != ""
}

// Sections return slice with section names
func (c *Config) Sections() []string {
	if c == nil {
		return []string{}
	}

	return c.sections
}

// Props return slice with properties names in some section
func (c *Config) Props(section string) []string {
	if c == nil || !c.HasSection(section) {
		return []string{}
	}

	var result []string

	// Section name + delimiter
	snLength := len(section) + 1

	for _, prop := range c.props {
		if len(prop) <= snLength {
			continue
		}

		if prop[0:snLength] == section+_DELIMITER {
			result = append(result, prop[snLength:])
		}
	}

	return result
}

// Validate require slice with pointers to validators and
// return slice with validation errors
func (c *Config) Validate(validators []*Validator) []error {
	if c == nil {
		return []error{errors.New("Config is nil")}
	}

	var result []error

	for _, v := range validators {
		err := v.Func(c, v.Property, v.Value)

		if err != nil {
			result = append(result, err)
		}
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

func parseRecord(data string, config *Config) (string, string) {
	va := strings.Split(data, _SEPARATOR_SYMBOL)

	propName := va[0]
	propValue := strings.Join(va[1:], _SEPARATOR_SYMBOL)

	propName = strings.TrimLeft(propName, " ")
	propValue = strings.TrimLeft(propValue, " ")

	if strings.Contains(propValue, _MACRO_SYMBOL) {
		macroses := macroRE.FindAllStringSubmatch(propValue, -1)

		for _, macros := range macroses {
			macroFull := macros[0]
			macroSect := macros[1]
			macroProp := macros[2]

			macroVal := config.GetS(macroSect + _DELIMITER + macroProp)

			propValue = strings.Replace(propValue, macroFull, macroVal, -1)
		}
	}

	return propName, propValue
}

// ////////////////////////////////////////////////////////////////////////////////// //
