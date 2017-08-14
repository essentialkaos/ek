package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// PropertyValidator is default type of property validation function
type PropertyValidator func(config *Config, prop string, value interface{}) error

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Empty check if given config property is empty or not
	Empty = validatorEmpty

	// NotContains check if given config property contains any value from given slice
	NotContains = validatorNotContains

	// Less check if given config property is less than defined value or not
	Less = validatorLess

	// Greater check if given config property is greater than defined value or not
	Greater = validatorGreater

	// Equals check if given config property equals to defined value or not
	Equals = validatorEquals
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

// ////////////////////////////////////////////////////////////////////////////////// //

func getWrongValidatorError(prop string) error {
	return fmt.Errorf("Wrong validator for property %s", prop)
}
