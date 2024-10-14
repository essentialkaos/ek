// Package validators provides basic KNF validators
package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Empty returns error if property is not set
	Set = validatorSet

	// SetToAny returns error if property doesn't contain any value from given slice
	SetToAny = validatorSetToAny

	// SetToAnyIgnoreCase returns error if property doesn't contain value from given
	// slice in any letter case
	SetToAnyIgnoreCase = validatorSetToAnyIgnoreCase

	// Less returns an error if the property value is smaller than the given number
	Less = validatorLess

	// Greater returns an error if the property value is greater than the given number
	Greater = validatorGreater

	// InRange returns an error if the property value is not in the given range
	InRange = validatorInRange

	// NotEquals returns an error if the property value is equal to the given string
	NotEquals = validatorNotEquals

	// LenLess returns an error if the length of the property value is greater than
	// given number
	LenLess = validatorLenLess

	// LenGreater returns an error if the length of the property value is less than
	// given number
	LenGreater = validatorLenGreater

	// LenNotEquals an error if the length of the property value is not equal to the
	// given number
	LenEquals = validatorLenEquals

	// HasPrefix returns error if property doesn't have given prefix
	HasPrefix = validatorHasPrefix

	// HasSuffix returns error if property doesn't have given suffix
	HasSuffix = validatorHasSuffix

	// TypeBool returns error if property contains non-boolean value
	TypeBool = validatorTypeBool

	// TypeNum returns error if property contains non-numeric (int/uint) value
	TypeNum = validatorTypeNum

	// TypeNum returns error if property contains non-float value
	TypeFloat = validatorTypeFloat

	// TypeSize returns error if property contains non-size value
	TypeSize = validatorTypeSize
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Range is numeric range
type Range struct {
	From any
	To   any
}

// ////////////////////////////////////////////////////////////////////////////////// //

func validatorSet(config knf.IConfig, prop string, value any) error {
	if config.GetS(prop) == "" {
		return fmt.Errorf("Property %s must be set", prop)
	}

	return nil
}

func validatorTypeBool(config knf.IConfig, prop string, value any) error {
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

func validatorTypeNum(config knf.IConfig, prop string, value any) error {
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

func validatorTypeFloat(config knf.IConfig, prop string, value any) error {
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

func validatorTypeSize(config knf.IConfig, prop string, value any) error {
	propValue := config.GetS(prop)

	if propValue == "" {
		return nil
	}

	size := config.GetSZ(prop)
	propValueNorm := strings.TrimRight(propValue, " bB")
	_, err := strconv.ParseFloat(propValueNorm, 64)

	if size == 0 && err != nil {
		return fmt.Errorf(
			"Property %s contains unsupported size value (%s)",
			prop, propValue,
		)
	}

	return nil
}

func validatorSetToAny(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case []string:
		if !isSliceContainsValue(t, config.GetS(prop), false) {
			return fmt.Errorf("Property %s doesn't contains any valid value", prop)
		}

		return nil
	}

	return getValidatorInputError("SetToAny", prop, value)
}

func validatorSetToAnyIgnoreCase(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case []string:
		if !isSliceContainsValue(t, config.GetS(prop), true) {
			return fmt.Errorf("Property %s doesn't contains any valid value", prop)
		}

		return nil
	}

	return getValidatorInputError("SetToAnyIgnoreCase", prop, value)
}

func validatorLess(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case int:
		if config.GetI(prop) > t {
			return fmt.Errorf("Property %s can't be greater than %d", prop, t)
		}

	case float64:
		if config.GetF(prop) > t {
			return fmt.Errorf("Property %s can't be greater than %g", prop, t)
		}

	default:
		return getValidatorInputError("Less", prop, value)
	}

	return nil
}

func validatorGreater(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case int:
		if config.GetI(prop) < t {
			return fmt.Errorf("Property %s can't be less than %d", prop, t)
		}

	case float64:
		if config.GetF(prop) < t {
			return fmt.Errorf("Property %s can't be less than %g", prop, t)
		}

	default:
		return getValidatorInputError("Greater", prop, value)
	}

	return nil
}

func validatorInRange(config knf.IConfig, prop string, value any) error {
	rng, ok := value.(Range)

	if !ok {
		return getValidatorInputError("InRange", prop, value)
	}

	var from, to float64

	switch u := rng.From.(type) {
	case int:
		from = float64(u)
	case uint:
		from = float64(u)
	case float64:
		from = u
	default:
		return getValidatorRangeError("From", rng.From)
	}

	switch u := rng.To.(type) {
	case int:
		to = float64(u)
	case uint:
		to = float64(u)
	case float64:
		to = u
	default:
		return getValidatorRangeError("To", rng.To)
	}

	v := config.GetF(prop)

	if v < from || v > to {
		return fmt.Errorf("Property %s must be in range %g-%g", prop, from, to)
	}

	return nil
}

func validatorNotEquals(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case int:
		if config.GetI(prop) == t {
			return fmt.Errorf("Property %s can't be equal %d", prop, t)
		}

	case float64:
		if config.GetF(prop) == t {
			return fmt.Errorf("Property %s can't be equal %f", prop, t)
		}

	case bool:
		if config.GetB(prop) == t {
			return fmt.Errorf("Property %s can't be equal %t", prop, t)
		}

	case string:
		if config.GetS(prop) == t {
			return fmt.Errorf("Property %s can't be equal %q", prop, t)
		}

	default:
		return getValidatorInputError("NotEquals", prop, value)
	}

	return nil
}

func validatorLenLess(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case int:
		if strutil.Len(config.GetS(prop)) > t {
			return fmt.Errorf("Property %s value can't be longer than %d symbols", prop, t)
		}

	default:
		return getValidatorInputError("LenLess", prop, value)
	}

	return nil
}

func validatorLenGreater(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case int:
		if strutil.Len(config.GetS(prop)) < t {
			return fmt.Errorf("Property %s value can't be shorter than %d symbols", prop, t)
		}

	default:
		return getValidatorInputError("LenGreater", prop, value)
	}

	return nil
}

func validatorLenEquals(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case int:
		if strutil.Len(config.GetS(prop)) != t {
			return fmt.Errorf("Property %s must be %d symbols long", prop, t)
		}

	default:
		return getValidatorInputError("LenEquals", prop, value)
	}

	return nil
}

func validatorHasPrefix(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("HasPrefix", prop)
		}

		if !strings.HasPrefix(config.GetS(prop), t) {
			return fmt.Errorf("Property %s must have prefix %q", prop, t)
		}

	default:
		return getValidatorInputError("HasPrefix", prop, value)
	}

	return nil
}

func validatorHasSuffix(config knf.IConfig, prop string, value any) error {
	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("HasSuffix", prop)
		}

		if !strings.HasSuffix(config.GetS(prop), t) {
			return fmt.Errorf("Property %s must have suffix %q", prop, t)
		}

	default:
		return getValidatorInputError("HasSuffix", prop, value)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isSliceContainsValue(s []string, value string, ignoreCase bool) bool {
	value = strings.ToLower(value)

	for _, v := range s {
		switch ignoreCase {
		case true:
			if strings.ToLower(v) == value {
				return true
			}
		default:
			if v == value {
				return true
			}
		}
	}

	return false
}

func getValidatorInputError(validator, prop string, value any) error {
	return fmt.Errorf(
		"Validator knf.%s doesn't support input with type <%T> for checking %s property",
		validator, value, prop,
	)
}

func getValidatorEmptyInputError(validator, prop string) error {
	return fmt.Errorf(
		"Validator knf.%s requires non-empty input for checking %s property",
		validator, prop,
	)
}

func getValidatorRangeError(prop string, value any) error {
	return fmt.Errorf(
		"Validator knf.InRange doesn't support type <%T> for 'Range.%s' value",
		value, prop,
	)
}
