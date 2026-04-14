// Package prefixed contains methods for encoding/decoding prefixed UUID
package prefixed

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/essentialkaos/ek/v14/uuid"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNoPrefix    = errors.New("prefixed UUID has no prefix")
	ErrEmptyPrefix = errors.New("prefixed UUID has empty prefix")
	ErrEmptyUUID   = errors.New("prefixed UUID has no UUID data")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// encoder is base64 encoder
var encoder = base64.StdEncoding.WithPadding(base64.NoPadding)

// ////////////////////////////////////////////////////////////////////////////////// //

// Encode returns a base64-encoded UUID string prefixed with prefix and a dot
// separator. An empty string is returned if prefix is empty or UUID is the zero value.
func Encode(prefix string, uuid uuid.UUID) string {
	if prefix == "" || uuid.IsZero() {
		return ""
	}

	return prefix + "." + encoder.EncodeToString(uuid[:])
}

// Decode parses a prefixed UUID string and returns the prefix, UUID, and parsing
// error
func Decode(prefixedUUID string) (string, uuid.UUID, error) {
	prefix, data, ok := strings.Cut(prefixedUUID, ".")

	switch {
	case !ok:
		return "", uuid.UUID{}, ErrNoPrefix
	case prefix == "":
		return "", uuid.UUID{}, ErrEmptyPrefix
	case data == "":
		return "", uuid.UUID{}, ErrEmptyUUID
	}

	u, err := encoder.DecodeString(data)

	if err != nil {
		return "", uuid.UUID{}, fmt.Errorf("can't decode UUID data: %w", err)
	}

	return prefix, uuid.UUID(u), nil
}
