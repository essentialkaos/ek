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

// Validator is the interface implemented by all input validators. Validators are
// applied in order; the first failure stops the chain and re-prompts the user.
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
	// NotEmpty rejects an empty or whitespace-only input
	NotEmpty = notEmptyValidator{}

	// IsNumber rejects input that cannot be parsed as a 64-bit integer
	IsNumber = isNumberValidator{}

	// IsFloat rejects input that cannot be parsed as a 64-bit floating-point number
	IsFloat = isFloatValidator{}

	// IsEmail rejects input that does not resemble a valid email address
	IsEmail = isEmailValidator{}

	// IsURL rejects input that does not start with http://, https://, or ftp://
	IsURL = isURLValidator{}
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrIsEmpty is returned when input is empty or contains only whitespace
	ErrIsEmpty = errors.New("value must be non-empty")

	// ErrInvalidNumber is returned when input cannot be parsed as an integer
	ErrInvalidNumber = errors.New("value is not a valid number")

	// ErrInvalidFloat is returned when input cannot be parsed as a floating-point number
	ErrInvalidFloat = errors.New("value is not a valid floating number")

	// ErrInvalidEmail is returned when input does not match a basic email format
	ErrInvalidEmail = errors.New("value is not a valid e-mail")

	// ErrInvalidURL is returned when input does not begin with a recognised URL scheme
	ErrInvalidURL = errors.New("value is not a valid URL")
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
