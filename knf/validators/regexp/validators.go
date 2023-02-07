// Package regexp provides KNF validators with regular expressions
package regexp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"regexp"

	"github.com/essentialkaos/ek/v12/knf"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Regexp returns an error if config property does not match given regexp
	Regexp = validateRegexp
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateRegexp(config *knf.Config, prop string, value any) error {
	pattern := value.(string)
	confVal := config.GetS(prop)

	if confVal == "" || pattern == "" {
		return nil
	}

	isMatch, err := regexp.MatchString(pattern, confVal)

	if err != nil {
		return fmt.Errorf("Can't use given regexp pattern: %v", err)
	}

	if !isMatch {
		return fmt.Errorf("Property %s must match regexp pattern %s", prop, pattern)
	}

	return nil
}
