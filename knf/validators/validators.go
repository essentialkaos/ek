// Package validators provides basic KNF validators
package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Empty returns error if config property is empty
	Empty = validatorEmpty

	// NotContains returns error if config property doesn't contains value from given slice
	NotContains = validatorNotContains

	// Less returns error if config property is less than given integer
	Less = validatorLess

	// Greater returns error if config property is greater than given integer
	Greater = validatorGreater

	// Equals returns error if config property is equals to given string
	Equals = validatorEquals

	// NotLen returns error if config property have wrong size
	NotLen = validatorNotLen

	// NotPrefix returns error if config property doesn't have given prefix
	NotPrefix = validatorNotPrefix

	// NotPrefix returns error if config property doesn't have given suffix
	NotSuffix = validatorNotSuffix

	// TypeBool returns error if config property contains non-boolean value
	TypeBool = validatorTypeBool

	// TypeNum returns error if config property contains non-numeric (int/uint) value
	TypeNum = validatorTypeNum

	// TypeNum returns error if config property contains non-float value
	TypeFloat = validatorTypeFloat
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validatorEmpty(config *knf.Config, prop string, value any) error {
	if config.GetS(prop) == "" {
		return fmt.Errorf("Property %s can't be empty", prop)
	}

	return nil
}

func validatorTypeBool(config *knf.Config, prop string, value any) error {
	propValue := config.GetS(prop)

	switch strings.ToLower(propValue) {
	case "", "0", "1", "true", "false", "yes", "no":
		return nil
	default:
		return fmt.Errorf(
			"Property %s contains unsupported boolean value (%s)",
			prop, propValue,
		)
	}
}

func validatorTypeNum(config *knf.Config, prop string, value any) error {
	propValue := config.GetS(prop)

	if propValue == "" {
		return nil
	}

	_, err := strconv.Atoi(propValue)

	if err != nil {
		return fmt.Errorf(
			"Property %s contains unsupported numeric value (%s)",
			prop, propValue,
		)
	}

	return nil
}

func validatorTypeFloat(config *knf.Config, prop string, value any) error {
	propValue := config.GetS(prop)

	if propValue == "" {
		return nil
	}

	_, err := strconv.ParseFloat(propValue, 64)

	if err != nil {
		return fmt.Errorf(
			"Property %s contains unsupported float value (%s)",
			prop, propValue,
		)
	}

	return nil
}

func validatorNotContains(config *knf.Config, prop string, value any) error {
	switch u := value.(type) {
	case []string:
		currentValue := config.GetS(prop)

		for _, v := range u {
			if v == currentValue {
				return nil
			}
		}

		return fmt.Errorf("Property %s doesn't contains any valid value", prop)
	}

	return getWrongValidatorError(prop)
}

func validatorLess(config *knf.Config, prop string, value any) error {
	switch value.(type) {
	case int, int32, int64, uint, uint32, uint64:
		if config.GetI(prop) < value.(int) {
			return fmt.Errorf("Property %s can't be less than %d", prop, value.(int))
		}
	case float32, float64:
		if config.GetF(prop) < value.(float64) {
			return fmt.Errorf("Property %s can't be less than %g", prop, value.(float64))
		}
	default:
		return getWrongValidatorError(prop)
	}

	return nil
}

func validatorGreater(config *knf.Config, prop string, value any) error {
	switch value.(type) {
	case int, int32, int64, uint, uint32, uint64:
		if config.GetI(prop) > value.(int) {
			return fmt.Errorf("Property %s can't be greater than %d", prop, value.(int))
		}

	case float32, float64:
		if config.GetF(prop) > value.(float64) {
			return fmt.Errorf("Property %s can't be greater than %g", prop, value.(float64))
		}

	default:
		return getWrongValidatorError(prop)
	}

	return nil
}

func validatorEquals(config *knf.Config, prop string, value any) error {
	switch u := value.(type) {
	case int, int32, int64, uint, uint32, uint64:
		if config.GetI(prop) == value.(int) {
			return fmt.Errorf("Property %s can't be equal %d", prop, value.(int))
		}

	case float32, float64:
		if config.GetF(prop) == value.(float64) {
			return fmt.Errorf("Property %s can't be equal %f", prop, value.(float64))
		}

	case bool:
		if config.GetB(prop) == u {
			return fmt.Errorf("Property %s can't be equal %t", prop, value.(bool))
		}

	case string:
		if config.GetS(prop) == u {
			return fmt.Errorf("Property %s can't be equal %q", prop, value.(string))
		}

	default:
		return getWrongValidatorError(prop)
	}

	return nil
}

func validatorNotLen(config *knf.Config, prop string, value any) error {
	if strutil.Len(config.GetS(prop)) != value.(int) {
		return fmt.Errorf("Property %s must be %d symbols long", prop, value.(int))
	}

	return nil
}

func validatorNotPrefix(config *knf.Config, prop string, value any) error {
	if !strings.HasPrefix(config.GetS(prop), value.(string)) {
		return fmt.Errorf("Property %s must have prefix %q", prop, value.(string))
	}

	return nil
}

func validatorNotSuffix(config *knf.Config, prop string, value any) error {
	if !strings.HasSuffix(config.GetS(prop), value.(string)) {
		return fmt.Errorf("Property %s must have suffix %q", prop, value.(string))
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getWrongValidatorError(prop string) error {
	return fmt.Errorf("Wrong validator for property %s", prop)
}
