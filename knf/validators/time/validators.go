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
	// Format returns an error if the config property contains an invalid time conversion
	// format
	Format = validateFormat
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateFormat(config knf.IConfig, prop string, value any) error {
	confVal := config.GetS(prop)

	if confVal == "" {
		return nil
	}

	str := timeutil.Format(time.Now(), confVal)

	if !strings.ContainsRune(str, '%') {
		return nil
	}

	seq := strutil.Substr(str, strings.IndexRune(str, '%'), 2)

	return fmt.Errorf(
		"Property %s contains invalid time format: Invalid control sequence %q",
		prop, seq,
	)
}
