// Package rand provides methods for generating random data
package rand

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

	symbolsLength := len(symbols)
	result := make([]byte, length)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < length; i++ {
		result[i] = symbols[rand.Intn(symbolsLength)]
	}

	return string(result)
}

// Int returns random int
func Int(n int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(n)
}

// Slice returns slice with random chars
func Slice(length int) []string {
	if length == 0 {
		return []string{}
	}

	symbolsLength := len(symbols)
	result := make([]string, length)

	for i := 0; i < length; i++ {
		result[i] = string(symbols[rand.Intn(symbolsLength)])
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //
