// Package passwd contains methods for working with passwords
package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"io"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	STRENGTH_WEAK   = iota // Only lowercase English alphabet characters
	STRENGTH_MEDIUM        // Lowercase and uppercase English alphabet characters, digits
	STRENGTH_STRONG        // Lowercase and uppercase English alphabet characters, digits, special symbols
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SYMBOLS_WEAK   = "abcdefghijklmnopqrstuvwxyz"
	_SYMBOLS_MEDIUM = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_SYMBOLS_STRONG = "!\";%:?*()_+=-~/\\<>,.[]{}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Encrypt creates hash and encrypts password with salt and pepper
func Encrypt(password, pepper string) (string, error) {
	switch {
	case password == "":
		return "", errors.New("Password can't be empty")
	case pepper == "":
		return "", errors.New("Pepper can't be empty")
	}

	if !isValidPepper(pepper) {
		return "", errors.New("Pepper have invalid size")
	}

	hasher := sha512.New()
	hasher.Write([]byte(password))

	hp, err := bcrypt.GenerateFromPassword(hasher.Sum(nil), 10)

	if err != nil {
		return "", err
	}

	block, _ := aes.NewCipher([]byte(pepper))
	hpd := padData(hp)

	ct := make([]byte, aes.BlockSize+len(hpd))
	iv := ct[:aes.BlockSize]

	_, err = io.ReadFull(crand.Reader, iv)

	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ct[aes.BlockSize:], hpd)

	return removeBase64Padding(base64.URLEncoding.EncodeToString(ct)), nil
}

// Check compares password and encrypted hash
func Check(password, pepper, hash string) bool {
	if password == "" || hash == "" || !isValidPepper(pepper) {
		return false
	}

	block, _ := aes.NewCipher([]byte(pepper))
	hpd, err := base64.URLEncoding.DecodeString(addBase64Padding(hash))

	if err != nil {
		return false
	}

	hdpl := len(hpd)

	if hdpl < aes.BlockSize || (hdpl%aes.BlockSize) != 0 {
		return false
	}

	iv := hpd[:aes.BlockSize]
	hp := hpd[aes.BlockSize:]

	if len(hp) == 0 {
		return false
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(hp, hp)

	h, ok := unpadData(hp)

	if !ok {
		return false
	}

	hasher := sha512.New()
	hasher.Write([]byte(password))

	return bcrypt.CompareHashAndPassword(h, hasher.Sum(nil)) == nil
}

// GenPassword generates random password
func GenPassword(length, strength int) string {
	return getRandomPassword(length, between(strength, 0, 2))
}

// GetPasswordStrength returns password strength
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

// GenPasswordVariations generates password variants with possible
// typos fixes (case swap for all letters, first leter swap, password
// without last symbol)
func GenPasswordVariations(password string) []string {
	if len(password) < 6 {
		return nil
	}

	var result []string
	var variant string

	for i := 0; i < len(password); i++ {
		variant += flipCase(string(password[i]))
	}

	result = append(result, variant)
	result = append(result, flipCase(password[:1])+password[1:])
	result = append(result, password[:len(password)-1])

	return result
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

func padData(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padText...)
}

func unpadData(src []byte) ([]byte, bool) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, false
	}

	return src[:(length - unpadding)], true
}

func addBase64Padding(src string) string {
	m := len(src) % 4

	if m != 0 {
		src += strings.Repeat("=", 4-m)
	}

	return src
}

func removeBase64Padding(src string) string {
	return strings.TrimRight(src, "=")
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

func isValidPepper(pepper string) bool {
	switch len(pepper) {
	case 16, 24, 32:
		return true
	}

	return false
}

func flipCase(s string) string {
	sc := strings.ToLower(s)

	if s != sc {
		return sc
	}

	return strings.ToUpper(s)
}
