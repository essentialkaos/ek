package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SYMBOL_COMMENT       = "#"
	_SYMBOL_SECTION_START = "["
	_SYMBOL_SECTION_END   = "["
	_SYMBOL_DELIMITER     = ":"
	_SYMBOL_MACRO_START   = "{"
	_SYMBOL_MACRO_END     = "}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// macroRE is a regexp for extracting macro
var macroRE = regexp.MustCompile(`\{([\w\-]+):([\w\-]+)\}`)

// ////////////////////////////////////////////////////////////////////////////////// //

// readData reads data from given reader
func readData(r io.Reader) (*Config, error) {
	config := &Config{
		data: make(map[string]string),
		mx:   &sync.RWMutex{},
	}

	var isDataRead bool
	var section string
	var lineNum int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t")
		lineNum++

		if line == "" || strings.HasPrefix(line, _SYMBOL_COMMENT) {
			continue
		}

		isDataRead = true

		if strings.HasPrefix(line, _SYMBOL_SECTION_START) &&
			strings.HasPrefix(line, _SYMBOL_SECTION_END) {
			section = strings.Trim(line, "[]")
			config.data[strings.ToLower(section)] = "!"
			config.sections = append(config.sections, section)
			continue
		}

		if section == "" {
			return nil, fmt.Errorf("Error at line %d: Data defined before section", lineNum)
		}

		propName, propValue, err := parseProperty(line, config)

		if err != nil {
			return nil, fmt.Errorf("Error at line %d: %w", lineNum, err)
		}

		fullPropName := genPropName(section, propName)

		if config.Has(fullPropName) {
			return nil, fmt.Errorf("Error at line %d: Property %q defined more than once", lineNum, propName)
		}

		config.props = append(config.props, fullPropName)
		config.data[strings.ToLower(fullPropName)] = propValue
	}

	if !isDataRead {
		return nil, fmt.Errorf("Configuration file doesn't contain any valid data")
	}

	return config, scanner.Err()
}

// parseProperty parses line with property name and value
func parseProperty(line string, config *Config) (string, string, error) {
	di := strings.Index(line, _SYMBOL_DELIMITER)

	if di == -1 {
		return "", "", fmt.Errorf(`Property must have ":" as a delimiter`)
	}

	name, value := line[:di], line[di+1:]

	name = strings.Trim(name, " \t")
	value = strings.Trim(value, " \t")

	if !strings.Contains(value, _SYMBOL_MACRO_START) &&
		!strings.Contains(value, _SYMBOL_MACRO_END) {
		return name, value, nil
	}

	var err error

	value, err = evalMacros(value, config)

	if err != nil {
		return "", "", err
	}

	return name, value, nil
}

// evalMacros evaluates all macros in given string
func evalMacros(value string, config *Config) (string, error) {
	macros := macroRE.FindAllStringSubmatch(value, -1)

	for _, macro := range macros {
		full, section, prop := macro[0], macro[1], macro[2]

		if !config.Has(genPropName(section, prop)) {
			return "", fmt.Errorf("Unknown property %s", full)
		}

		propVal := config.GetS(genPropName(section, prop))
		value = strings.ReplaceAll(value, full, propVal)
	}

	return value, nil
}

// genPropName generates "full property name" which contains section and
// property name
func genPropName(section, prop string) string {
	return section + _SYMBOL_DELIMITER + prop
}
