// Package passwd contains methods for working with passwords
package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
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

type Strength uint8

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	STRENGTH_WEAK   Strength = 0 // Only lowercase English alphabet characters
	STRENGTH_MEDIUM Strength = 1 // Lowercase and uppercase English alphabet characters, digits
	STRENGTH_STRONG Strength = 2 // Lowercase and uppercase English alphabet characters, digits, special symbols
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SYMBOLS_WEAK   = "abcdefghijklmnopqrstuvwxyz"
	_SYMBOLS_MEDIUM = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_SYMBOLS_STRONG = "!\";%:?*()_+=-~/\\<>,.[]{}"
)

var (
	ErrEmptyPassword = errors.New("Password can't be empty")
	ErrEmptyPepper   = errors.New("Pepper can't be empty")
	ErrInvalidPepper = errors.New("Pepper has invalid size")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Encrypt creates hash and encrypts it with salt and pepper
// Deprecated: Use Hash method instead
func Encrypt(password, pepper string) (string, error) {
	return Hash(password, pepper)
}

// Hash creates hash and encrypts it with salt and pepper
func Hash(password, pepper string) (string, error) {
	hash, err := HashBytes([]byte(password), []byte(pepper))

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// HashBytes creates hash and encrypts it with salt and pepper
func HashBytes(password, pepper []byte) ([]byte, error) {
	switch {
	case len(password) == 0:
		return nil, ErrEmptyPassword
	case len(pepper) == 0:
		return nil, ErrEmptyPepper
	}

	if !isValidPepper(pepper) {
		return nil, ErrInvalidPepper
	}

	hasher := sha512.New()
	hasher.Write(password)

	hp, err := bcrypt.GenerateFromPassword(hasher.Sum(nil), 10)

	if err != nil {
		return nil, err
	}

	block, _ := aes.NewCipher(pepper)
	hpd := padData(hp)

	ct := make([]byte, aes.BlockSize+len(hpd))
	iv := ct[:aes.BlockSize]

	_, err = io.ReadFull(crand.Reader, iv)

	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ct[aes.BlockSize:], hpd)

	buf := make([]byte, base64.URLEncoding.EncodedLen(len(ct)))
	base64.URLEncoding.Encode(buf, ct)

	return removeBase64Padding(buf), nil
}

// Check compares password with encrypted hash
func Check(password, pepper, hash string) bool {
	return CheckBytes([]byte(password), []byte(pepper), []byte(hash))
}

// CheckBytes compares password with encrypted hash
func CheckBytes(password, pepper, hash []byte) bool {
	if len(password) == 0 || len(hash) == 0 || !isValidPepper(pepper) {
		return false
	}

	hs := addBase64Padding(hash)
	hpd := make([]byte, base64.URLEncoding.DecodedLen(len(hs)))
	block, _ := aes.NewCipher(pepper)
	n, err := base64.URLEncoding.Decode(hpd, hs)

	if err != nil {
		return false
	}

	hpd = hpd[:n]
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
	hasher.Write(password)

	return bcrypt.CompareHashAndPassword(h, hasher.Sum(nil)) == nil
}

// GenPassword generates random password
func GenPassword(length int, strength Strength) string {
	return string(GenPasswordBytes(length, strength))
}

// GenPassword generates random password
func GenPasswordBytes(length int, strength Strength) []byte {
	return getRandomPasswordBytes(length, getStrength(strength))
}

// GetPasswordStrength returns password strength
func GetPasswordStrength(password string) Strength {
	return GetPasswordBytesStrength([]byte(password))
}

// GetPasswordBytesStrength returns password strength
func GetPasswordBytesStrength(password []byte) Strength {
	if len(password) == 0 {
		return STRENGTH_WEAK
	}

	var conditions int

	if bytes.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
		bytes.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		conditions++
	}

	if bytes.ContainsAny(password, "1234567890") {
		conditions++
	}

	if bytes.ContainsAny(password, _SYMBOLS_STRONG) {
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
	}

	return STRENGTH_WEAK
}

// GenPasswordVariations generates password variants with possible
// typos fixes (case swap for all letters, first leter swap, password
// without last symbol)
func GenPasswordVariations(password string) []string {
	if len(password) < 6 {
		return nil
	}

	var result []string

	passwordBytes := []byte(password)

	result = append(result, string(genVariantFlipAll(passwordBytes)))
	result = append(result, string(genVariantFlipFirst(passwordBytes)))
	result = append(result, string(genVariantTrimLast(passwordBytes)))

	return result
}

// GenPasswordBytesVariations generates password variants with possible
// typos fixes (case swap for all letters, first leter swap, password
// without last symbol)
func GenPasswordBytesVariations(password []byte) [][]byte {
	if len(password) < 6 {
		return nil
	}

	var result [][]byte

	result = append(result, genVariantFlipAll(password))
	result = append(result, genVariantFlipFirst(password))
	result = append(result, genVariantTrimLast(password))

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getStrength(s Strength) Strength {
	if s > STRENGTH_STRONG {
		return STRENGTH_STRONG
	}

	return s
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

func addBase64Padding(src []byte) []byte {
	m := len(src) % 4

	if m == 0 {
		return src
	}

	buf := make([]byte, len(src)+(4-len(src)%4))
	copy(buf, src)

	for i := len(src); i < len(buf); i++ {
		buf[i] = '='
	}

	return buf
}

func removeBase64Padding(src []byte) []byte {
	return bytes.TrimRight(src, "=")
}

func getRandomPasswordBytes(length int, strength Strength) []byte {
	if length == 0 {
		return nil
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

	ls := len(symbols)
	buf := make([]byte, length)

	for {
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < length; i++ {
			buf[i] = symbols[rand.Intn(ls)]
		}

		if GetPasswordBytesStrength(buf) == strength {
			return buf
		}
	}
}

func isValidPepper(pepper []byte) bool {
	switch len(pepper) {
	case 16, 24, 32:
		return true
	}

	return false
}

func genVariantFlipAll(password []byte) []byte {
	result := make([]byte, len(password))

	for i := 0; i < len(password); i++ {
		result[i] = flipCase(password[i])
	}

	return result
}

func genVariantFlipFirst(password []byte) []byte {
	result := make([]byte, len(password))

	copy(result, password)

	result[0] = flipCase(password[0])

	return result
}

func genVariantTrimLast(password []byte) []byte {
	return append(password[:0:0], password[:len(password)-1]...)
}

func flipCase(b byte) byte {
	s := string(b)
	sc := strings.ToLower(s)

	if s != sc {
		return byte(sc[0])
	}

	return byte(strings.ToUpper(s)[0])
}
