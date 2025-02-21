// Package regexp provides KNF validators for cron expressions
package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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
	// Expression returns an error if config property contains invalid cron expression
	Expression = validateCronExpression
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateCronExpression(config knf.IConfig, prop string, value any) error {
	confVal := config.GetS(prop)

	if confVal == "" {
		return nil
	}

	_, err := cron.Parse(confVal)

	if err != nil {
		return fmt.Errorf("Invalid cron expression: %w", err)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
