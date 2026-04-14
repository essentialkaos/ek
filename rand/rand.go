// Package rand provides methods for generating random data
package rand

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math/rand"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// alphabet is the set of characters used for random string generation
var alphabet = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns a randomly generated string of the given length.
// Returns an empty string if length is less than or equal to zero.
func String(length int) string {
	if length <= 0 {
		return ""
	}

	rnd := rand.New(rand.NewSource(rand.Int63()))
	alphabetLen := len(alphabet)
	result := make([]byte, length)

	for i := range length {
		result[i] = alphabet[rnd.Intn(alphabetLen)]
	}

	return string(result)
}

// Slice returns a slice of randomly selected single characters of the given length.
// Returns an empty slice if length is less than or equal to zero.
func Slice(length int) []string {
	if length <= 0 {
		return []string{}
	}

	rnd := rand.New(rand.NewSource(rand.Int63()))
	alphabetLen := len(alphabet)
	result := make([]string, length)

	for i := range length {
		result[i] = string(alphabet[rnd.Intn(alphabetLen)])
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //
