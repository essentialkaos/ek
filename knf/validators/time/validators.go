// Package time provides KNF validators for time-related elements (layouts, formats…)
package time

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/timeutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Format returns an error if the configuration property contains an invalid time
	// conversion format
	Format = validateFormat

	// Timezone returns an error if the configuration property contains an invalid time
	// zone name
	Timezone = validateTimezone
)

// ////////////////////////////////////////////////////////////////////////////////// //

// validateFormat checks if the given property contains a valid time conversion format
func validateFormat(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	str := timeutil.Format(time.Now(), v)

	if !strings.ContainsRune(str, '%') {
		return nil
	}

	seq := strutil.Substr(str, strings.IndexRune(str, '%'), 2)

	return fmt.Errorf(
		"Property %s contains invalid time format: Invalid control sequence %q",
		prop, seq,
	)
}

// validateTimezone checks if the given property contains a valid time zone name
func validateTimezone(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	_, err := time.LoadLocation(v)

	if err == nil {
		return nil
	}

	return fmt.Errorf("Property %s contains invalid time zone name: %v", prop, err)
}
