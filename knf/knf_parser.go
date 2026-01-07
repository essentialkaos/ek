package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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

		fullPropName := Q(section, propName)

		if config.Has(fullPropName) {
			return nil, fmt.Errorf("Error at line %d: Property %q defined more than once", lineNum, propName)
		}

		config.props = append(config.props, fullPropName)
		config.data[fullPropName] = propValue
	}

	if !isDataRead {
		return nil, fmt.Errorf("Configuration file doesn't contain any valid data")
	}

	return config, scanner.Err()
}

// parseProperty parses line with property name and value
func parseProperty(line string, config *Config) (string, string, error) {
	name, value, ok := strings.Cut(line, _SYMBOL_DELIMITER)

	if !ok {
		return "", "", fmt.Errorf("Property must have %q as a delimiter", _SYMBOL_DELIMITER)
	}

	name = strings.Trim(name, " \t")
	value = strings.Trim(value, " \t")

	if !strings.ContainsAny(value, _SYMBOL_MACRO_START+_SYMBOL_MACRO_END) {
		return name, value, nil
	}

	var err error

	value, err = evalMacros(value, config)

	if err != nil {
		return "", "", err
	}

	return name, value, nil
}

// evalMacros evaluates all macro in given string
func evalMacros(value string, config *Config) (string, error) {
	macros := macroRE.FindAllStringSubmatch(value, -1)

	for _, macro := range macros {
		full, section, prop := macro[0], macro[1], macro[2]

		if !config.Has(Q(section, prop)) {
			return "", fmt.Errorf("Unknown property %s", full)
		}

		propVal := config.GetS(Q(section, prop))
		value = strings.ReplaceAll(value, full, propVal)
	}

	return value, nil
}
