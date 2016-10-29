// Package passwd contains methods for working with passwords
package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// AuthData contains password, salt and hash
type AuthData struct {
	Password string // Password
	Salt     string // Salt
	Hash     string // Salted hash
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Password strength
const (
	STRENGTH_WEAK = iota
	STRENGTH_MEDIUM
	STRENGTH_STRONG
)

const (
	_SYMBOLS_WEAK   = "abcdefghijklmnopqrstuvwxyz"
	_SYMBOLS_MEDIUM = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_SYMBOLS_STRONG = "!\";%:?*()_+=-~/\\<>,.[]{}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GenPassword generate random password
func GenPassword(length, strength int) string {
	return getRandomPassword(length, between(strength, 0, 2))
}

// GenAuth return struct with generated password, hashed password and salt
func GenAuth(length, strength int) AuthData {
	password := getRandomPassword(length, between(strength, 0, 2))
	return CreateAuth(password)
}

// CreateAuth return struct with given password, hashed password and salt
func CreateAuth(password string) AuthData {
	salt := getRandomPassword(16, STRENGTH_MEDIUM)
	hash := GenHash(password, salt)

	return AuthData{password, salt, hash}
}

// GenHash generate SHA-256 hash for given password and salt
func GenHash(password, salt string) string {
	hasher := sha256.New()

	hasher.Write([]byte(password + salt))

	prehash := fmt.Sprintf("%064x", hasher.Sum(nil))
	hasher2 := sha256.New()

	hasher2.Write([]byte(salt + prehash))

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// GetPasswordStrength return password strength
func GetPasswordStrength(password string) int {
	if password == "" {
		return STRENGTH_WEAK
	}

	var conditions int

	if strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
		strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		conditions++
	}

	if strings.ContainsAny(password, "1234567890") {
		conditions++
	}

	if strings.ContainsAny(password, _SYMBOLS_STRONG) {
		conditions++
	}

	if len(password) < 6 {
		conditions = 1
	} else {
		conditions++
	}

	switch conditions {
	case 4:
		return STRENGTH_STRONG

	case 3:
		return STRENGTH_MEDIUM

	default:
		return STRENGTH_WEAK
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func between(val, min, max int) int {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

func getRandomPassword(length, strength int) string {
	if length == 0 {
		return ""
	}

	if strength == STRENGTH_STRONG && length < 6 {
		length = 6
	}

	var symbols = _SYMBOLS_WEAK

	switch strength {
	case STRENGTH_MEDIUM:
		symbols += _SYMBOLS_MEDIUM
	case STRENGTH_STRONG:
		symbols += _SYMBOLS_MEDIUM + _SYMBOLS_STRONG
	}

	for {
		ls := len(symbols)
		r := make([]byte, length)

		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < length; i++ {
			r[i] = symbols[rand.Intn(ls)]
		}

		if GetPasswordStrength(string(r)) == strength {
			return string(r)
		}
	}
}
