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
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var symbols = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string with random chars
func String(length int) string {
	if length <= 0 {
		return ""
	}

	rnd := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	symbolsLength := len(symbols)
	result := make([]byte, length)

	for i := range length {
		result[i] = symbols[rnd.Intn(symbolsLength)]
	}

	return string(result)
}

// Slice returns slice with random chars
func Slice(length int) []string {
	if length == 0 {
		return []string{}
	}

	rnd := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	symbolsLength := len(symbols)
	result := make([]string, length)

	for i := range length {
		result[i] = string(symbols[rnd.Intn(symbolsLength)])
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //
