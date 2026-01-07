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
	"strings"

	"github.com/essentialkaos/ek/v13/uuid"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// encoder is base64 encoder
var encoder = base64.StdEncoding.WithPadding(base64.NoPadding)

// ////////////////////////////////////////////////////////////////////////////////// //

// Encode creates base64-encoded prefixed UUID
func Encode(prefix string, uuid uuid.UUID) string {
	if prefix == "" || uuid.IsZero() {
		return ""
	}

	return prefix + "." + encoder.EncodeToString(uuid)
}

// Decode decodes base64-encoded prefixed UUID and returns prefix and UUID
func Decode(prefixedUUID string) (string, uuid.UUID, error) {
	prefix, data, ok := strings.Cut(prefixedUUID, ".")

	switch {
	case !ok:
		return "", nil, errors.New("Prefixed UUID has no prefix")
	case prefix == "":
		return "", nil, errors.New("Prefixed UUID has empty prefix")
	case data == "":
		return "", nil, errors.New("Prefixed UUID has no UUID data")
	}

	u, err := encoder.DecodeString(data)

	if err != nil {
		return "", nil, errors.New("Can't decode UUID data: " + err.Error())
	}

	return prefix, uuid.UUID(u), nil
}
