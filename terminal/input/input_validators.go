package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Validator is input validator type
type Validator interface {
	Validate(input string) (string, error)
}

// ////////////////////////////////////////////////////////////////////////////////// //

type notEmptyValidator struct{}
type isNumberValidator struct{}
type isFloatValidator struct{}
type isEmailValidator struct{}
type isURLValidator struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// NotEmpty returns an error if input is empty
	NotEmpty = notEmptyValidator{}

	// IsNumber returns an error if the input is not a valid number
	IsNumber = isNumberValidator{}

	// IsFloat returns an error if the input is not a valid floating number
	IsFloat = isFloatValidator{}

	// IsEmail returns an error if the input is not a valid email
	IsEmail = isEmailValidator{}

	// IsURL returns an error if the input is not a valid URL
	IsURL = isURLValidator{}
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrIsEmpty is error for empty input
	ErrIsEmpty = errors.New("You must enter non-empty value")

	// ErrInvalidNumber is error for invalid number
	ErrInvalidNumber = errors.New("Entered value is not a valid number")

	// ErrInvalidFloat is error for invalid floating number
	ErrInvalidFloat = errors.New("Entered value is not a valid floating number")

	// ErrInvalidEmail is error for invalid email
	ErrInvalidEmail = errors.New("Entered value is not a valid e-mail")

	// ErrInvalidURL is error for invalid URL
	ErrInvalidURL = errors.New("Entered value is not a valid URL")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate validates for empty input
func (v notEmptyValidator) Validate(input string) (string, error) {
	if strings.TrimSpace(input) != "" {
		return input, nil
	}

	return input, ErrIsEmpty
}

// Validate checks if the given input is a number
func (v isNumberValidator) Validate(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return input, nil // Empty input is okay
	}

	_, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		return input, ErrInvalidNumber
	}

	return input, nil
}

// Validate checks if the given input is a floating number
func (v isFloatValidator) Validate(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return input, nil // Empty input is okay
	}

	_, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return input, ErrInvalidFloat
	}

	return input, nil
}

// Validate checks if the given input is an email
func (v isEmailValidator) Validate(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return input, nil // Empty input is okay
	}

	name, domain, ok := strings.Cut(input, "@")

	if !ok || strings.TrimSpace(name) == "" ||
		strings.TrimSpace(domain) == "" || !strings.Contains(domain, ".") {
		return input, ErrInvalidEmail
	}

	return input, nil
}

// Validate checks if the given input is a URL
func (v isURLValidator) Validate(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return input, nil // Empty input is okay
	}

	switch {
	case strings.HasPrefix(input, "http://"),
		strings.HasPrefix(input, "https://"),
		strings.HasPrefix(input, "ftp://"):
		// OK
	default:
		return input, ErrInvalidURL
	}

	if !strings.Contains(input, ".") {
		return input, ErrInvalidURL
	}

	return input, nil
}
