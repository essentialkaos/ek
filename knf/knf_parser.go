package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_COMMENT_SYMBOL       = "#"
	_SECTION_START_SYMBOL = "["
	_SECTION_END_SYMBOL   = "["
	_PROP_DELIMITER       = ":"
	_MACRO_START_SYMBOL   = "{"
	_MACRO_END_SYMBOL     = "}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// macroRE is a regexp for extracting macroses
var macroRE = regexp.MustCompile(`\{([\w\-]+):([\w\-]+)\}`)

// ////////////////////////////////////////////////////////////////////////////////// //

// readKNFFile reads KNF file
func readKNFFile(file string) (*Config, error) {
	fd, err := os.OpenFile(path.Clean(file), os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	config, err := readKNFData(fd)

	if err != nil {
		return nil, err
	}

	config.file = file

	return config, nil
}

// readKNFData reads data from given reader
func readKNFData(r io.Reader) (*Config, error) {
	reader := bufio.NewReader(r)
	scanner := bufio.NewScanner(reader)

	config := &Config{
		data: make(map[string]string),
		mx:   &sync.RWMutex{},
	}

	var isDataRead bool
	var section string
	var lineNum int

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t")
		lineNum++

		if line == "" || strings.HasPrefix(line, _COMMENT_SYMBOL) {
			continue
		}

		isDataRead = true

		if strings.HasPrefix(line, _SECTION_START_SYMBOL) &&
			strings.HasPrefix(line, _SECTION_END_SYMBOL) {
			section = strings.Trim(line, "[]")
			config.data[strings.ToLower(section)] = "!"
			config.sections = append(config.sections, section)
			continue
		}

		if section == "" {
			return nil, fmt.Errorf("Error at line %d: Data defined before section", lineNum)
		}

		propName, propValue, err := parseKNFProperty(line, config)

		if err != nil {
			return nil, fmt.Errorf("Error at line %d: %w", lineNum, err)
		}

		fullPropName := genPropName(section, propName)

		if config.HasProp(fullPropName) {
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

// parseKNFProperty parses line with property name and value
func parseKNFProperty(line string, config *Config) (string, string, error) {
	di := strings.Index(line, _PROP_DELIMITER)

	if di == -1 {
		return "", "", fmt.Errorf("Property must have \":\" as a delimiter")
	}

	name, value := line[:di], line[di+1:]

	name = strings.Trim(name, " \t")
	value = strings.Trim(value, " \t")

	if !strings.Contains(value, _MACRO_START_SYMBOL) &&
		!strings.Contains(value, _MACRO_END_SYMBOL) {
		return name, value, nil
	}

	var err error

	value, err = evalMacroses(value, config)

	if err != nil {
		return "", "", err
	}

	return name, value, nil
}

// evalMacroses evaluates all macroses in given string
func evalMacroses(value string, config *Config) (string, error) {
	macroses := macroRE.FindAllStringSubmatch(value, -1)

	for _, macros := range macroses {
		full, section, prop := macros[0], macros[1], macros[2]

		if !config.HasProp(genPropName(section, prop)) {
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
	return section + _PROP_DELIMITER + prop
}
