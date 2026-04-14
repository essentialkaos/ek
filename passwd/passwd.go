// Package passwd contains methods for working with passwords
package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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

	"golang.org/x/crypto/bcrypt"

	"github.com/essentialkaos/ek/v14/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Strength represents password strength
type Strength uint8

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	// STRENGTH_WEAK indicates a password composed only of lowercase ASCII letters
	STRENGTH_WEAK Strength = 0

	// STRENGTH_MEDIUM indicates a password with mixed-case letters and digits
	STRENGTH_MEDIUM Strength = 1

	// STRENGTH_STRONG indicates a password with mixed-case letters, digits, and special symbols
	STRENGTH_STRONG Strength = 2
)

// MIN_PASSWORD_LEN is the minimal password length
const MIN_PASSWORD_LEN = 6

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SYMBOLS_WEAK   = "abcdefghijklmnopqrstuvwxyz"
	_SYMBOLS_MEDIUM = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_SYMBOLS_STRONG = "!\";%:?*()_+=-~/\\<>,.[]{}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyPassword is returned by [Hash] and [HashBytes] when the password
	// argument is empty
	ErrEmptyPassword = errors.New("password can't be empty")

	// ErrEmptyPepper is returned by [Hash] and [HashBytes] when the pepper argument
	// is empty
	ErrEmptyPepper = errors.New("pepper can't be empty")

	// ErrInvalidPepper is returned when the pepper length is not 16, 24, or 32 bytes
	ErrInvalidPepper = errors.New("pepper has invalid size")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Hash bcrypt-hashes the password, encrypts the result with the pepper using AES-CFB,
// and returns a URL-safe base64-encoded string
func Hash(password, pepper string) (string, error) {
	hash, err := HashBytes([]byte(password), []byte(pepper))

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// HashBytes is the byte-slice variant of [Hash]
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

// Check reports whether password matches the hash produced by [Hash] using
// the given pepper
func Check(password, pepper, hash string) bool {
	return CheckBytes([]byte(password), []byte(pepper), []byte(hash))
}

// CheckBytes is the byte-slice variant of [Check]
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
	hdpLen := len(hpd)

	if hdpLen < aes.BlockSize || (hdpLen%aes.BlockSize) != 0 {
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

// GenPassword returns a random password of the given length and strength level
func GenPassword(length int, strength Strength) string {
	return string(GenPasswordBytes(length, strength))
}

// GenPasswordBytes returns a random password as a byte slice of the given length
// and strength level
func GenPasswordBytes(length int, strength Strength) []byte {
	return getRandomPasswordBytes(
		length, mathutil.Between(strength, STRENGTH_WEAK, STRENGTH_STRONG),
	)
}

// GetPasswordStrength evaluates and returns the strength of the given password
// string
func GetPasswordStrength(password string) Strength {
	return GetPasswordBytesStrength([]byte(password))
}

// GetPasswordBytesStrength evaluates and returns the strength of the given password
// byte slice
func GetPasswordBytesStrength(password []byte) Strength {
	if len(password) == 0 {
		return STRENGTH_WEAK
	}

	hasLower := bytes.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	hasUpper := bytes.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasDigit := bytes.ContainsAny(password, "1234567890")
	hasSymbol := bytes.ContainsAny(password, _SYMBOLS_STRONG)
	longEnough := len(password) >= MIN_PASSWORD_LEN

	switch {
	case longEnough && hasLower && hasUpper && hasDigit && hasSymbol:
		return STRENGTH_STRONG
	case longEnough && hasDigit && (hasLower || hasUpper):
		return STRENGTH_MEDIUM
	}

	return STRENGTH_WEAK
}

// GenPasswordVariations returns up to three variants of the password with common
// typo corrections applied: full case-flip, first-character case-flip, and last
// character removed
func GenPasswordVariations(password string) []string {
	if len(password) < MIN_PASSWORD_LEN {
		return nil
	}

	passwordBytes := []byte(password)

	return []string{
		string(genVariantFlipAll(passwordBytes)),
		string(genVariantFlipFirst(passwordBytes)),
		string(genVariantTrimLast(passwordBytes)),
	}
}

// GenPasswordBytesVariations is the byte-slice variant of [GenPasswordVariations]
func GenPasswordBytesVariations(password []byte) [][]byte {
	if len(password) < MIN_PASSWORD_LEN {
		return nil
	}

	return [][]byte{
		genVariantFlipAll(password),
		genVariantFlipFirst(password),
		genVariantTrimLast(password),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// padData adds padding to data to make its length a multiple of the AES block size
func padData(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padText...)
}

// unpadData removes padding from data
func unpadData(src []byte) ([]byte, bool) {
	length := len(src)

	if length == 0 {
		return nil, false
	}

	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, false
	}

	return src[:(length - unpadding)], true
}

// addBase64Padding adds padding to base64 encoded data
func addBase64Padding(src []byte) []byte {
	m := len(src) % 4

	if m == 0 {
		return src
	}

	buf := make([]byte, len(src)+(4-m))

	copy(buf, src)

	for i := len(src); i < len(buf); i++ {
		buf[i] = '='
	}

	return buf
}

// removeBase64Padding removes padding from base64 encoded data
func removeBase64Padding(src []byte) []byte {
	return bytes.TrimRight(src, "=")
}

// getRandomPasswordBytes generates a random password of the specified length and strength
func getRandomPasswordBytes(length int, strength Strength) []byte {
	if length == 0 {
		return nil
	}

	if strength == STRENGTH_STRONG && length < MIN_PASSWORD_LEN {
		length = MIN_PASSWORD_LEN
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
		for i := range length {
			buf[i] = symbols[rand.Intn(ls)]
		}

		if GetPasswordBytesStrength(buf) == strength {
			return buf
		}
	}
}

// isValidPepper checks if the pepper has a valid size for AES encryption
func isValidPepper(pepper []byte) bool {
	switch len(pepper) {
	case 16, 24, 32:
		return true
	}

	return false
}

// genVariantFlipAll generates a password variant with all letters
// flipped (case swap)
func genVariantFlipAll(password []byte) []byte {
	result := make([]byte, len(password))

	for i := range len(password) {
		result[i] = flipCase(password[i])
	}

	return result
}

// genVariantFlipFirst generates a password variant with the first letter
// flipped (case swap)
func genVariantFlipFirst(password []byte) []byte {
	result := make([]byte, len(password))

	copy(result, password)

	result[0] = flipCase(password[0])

	return result
}

// genVariantTrimLast generates a password variant with the last symbol trimmed
func genVariantTrimLast(password []byte) []byte {
	return bytes.Clone(password[:len(password)-1])
}

// flipCase flips the case of a single byte
func flipCase(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}

	if b >= 'A' && b <= 'Z' {
		return b + 32
	}

	return b
}
