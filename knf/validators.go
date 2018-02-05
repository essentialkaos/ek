package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// PropertyValidator is default type of property validation function
type PropertyValidator func(config *Config, prop string, value interface{}) error

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Empty return error if config property is empty
	Empty = validatorEmpty

	// NotContains return error if config property doesn't contains value from given slice
	NotContains = validatorNotContains

	// Less return error if config property is less than given integer
	Less = validatorLess

	// Greater return error if config property is greater than given integer
	Greater = validatorGreater

	// Equals return error if config property is equals to given string
	Equals = validatorEquals

	// NotLen return error if config property have wrong size
	NotLen = validatorNotLen

	// NotPrefix return error if config property doesn't have given prefix
	NotPrefix = validatorNotPrefix

	// NotPrefix return error if config property doesn't have given suffix
	NotSuffix = validatorNotSuffix
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validatorEmpty(config *Config, prop string, value interface{}) error {
	if config.GetS(prop) == "" {
		return fmt.Errorf("Property %s can't be empty", prop)
	}

	return nil
}

func validatorNotContains(config *Config, prop string, value interface{}) error {
	switch value.(type) {
	case []string:
		currentValue := config.GetS(prop)

		for _, v := range value.([]string) {
			if v == currentValue {
				return nil
			}
		}

		return fmt.Errorf("Property %s doesn't contains any valid value", prop)
	}

	return getWrongValidatorError(prop)
}

func validatorLess(config *Config, prop string, value interface{}) error {
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

func validatorGreater(config *Config, prop string, value interface{}) error {
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

func validatorEquals(config *Config, prop string, value interface{}) error {
	switch value.(type) {
	case int, int32, int64, uint, uint32, uint64:
		if config.GetI(prop) == value.(int) {
			return fmt.Errorf("Property %s can't be equal %d", prop, value.(int))
		}

	case float32, float64:
		if config.GetF(prop) == value.(float64) {
			return fmt.Errorf("Property %s can't be equal %f", prop, value.(float64))
		}

	case bool:
		if config.GetB(prop) == value.(bool) {
			return fmt.Errorf("Property %s can't be equal %t", prop, value.(bool))
		}

	case string:
		if config.GetS(prop) == value.(string) {
			return fmt.Errorf("Property %s can't be equal %s", prop, value.(string))
		}

	default:
		return getWrongValidatorError(prop)
	}

	return nil
}

func validatorNotLen(config *Config, prop string, value interface{}) error {
	if strutil.Len(config.GetS(prop)) != value.(int) {
		return fmt.Errorf("Property %s must be %d symbols long", prop, value.(int))
	}

	return nil
}

func validatorNotPrefix(config *Config, prop string, value interface{}) error {
	if !strings.HasPrefix(config.GetS(prop), value.(string)) {
		return fmt.Errorf("Property %s must have prefix \"%s\"", prop, value.(string))
	}

	return nil
}

func validatorNotSuffix(config *Config, prop string, value interface{}) error {
	if !strings.HasSuffix(config.GetS(prop), value.(string)) {
		return fmt.Errorf("Property %s must have suffix \"%s\"", prop, value.(string))
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getWrongValidatorError(prop string) error {
	return fmt.Errorf("Wrong validator for property %s", prop)
}
