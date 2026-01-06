// Package regexp provides KNF validators with regular expressions
package regexp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"regexp"

	"github.com/essentialkaos/ek/v13/knf"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Regexp returns an error if configuration property does not match given regexp
	Regexp = validateRegexp
)

// ////////////////////////////////////////////////////////////////////////////////// //

// validateRegexp checks if the value of the property matches the given regexp pattern
func validateRegexp(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("Regexp", prop)
		}

		re, err := regexp.Compile(t)

		if err != nil {
			return fmt.Errorf("Invalid input for regexp.Regexp validator: %w", err)
		}

		if !re.MatchString(v) {
			return fmt.Errorf("Property %s must match regexp pattern %q", prop, t)
		}

	default:
		return getValidatorInputError("Regexp", prop, value)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getValidatorInputError returns an error for unsupported input type in regexp
// validator
func getValidatorInputError(validator, prop string, value any) error {
	return fmt.Errorf(
		"Validator regexp.%s doesn't support input with type <%T> for checking %s property",
		validator, value, prop,
	)
}

// getValidatorEmptyInputError returns an error for empty input in regexp validator
func getValidatorEmptyInputError(validator, prop string) error {
	return fmt.Errorf(
		"Validator regexp.%s requires non-empty input for checking %s property",
		validator, prop,
	)
}
