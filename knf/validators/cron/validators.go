// Package regexp provides KNF validators for cron expressions
package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v13/cron"
	"github.com/essentialkaos/ek/v13/knf"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Expression returns an error if configuration property contains invalid cron
	// expression
	Expression = validateCronExpression
)

// ////////////////////////////////////////////////////////////////////////////////// //

// validateCronExpression checks if the given property contains a valid cron expression
func validateCronExpression(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	_, err := cron.Parse(v)

	if err != nil {
		return fmt.Errorf("Property %s contains invalid cron expression: %w", prop, err)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
