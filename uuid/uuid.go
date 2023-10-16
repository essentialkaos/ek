// Package uuid contains methods for generating version 4 and 5 UUID's
package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Predefined namespace UUID's
var (
	NsDNS  = []byte{107, 167, 184, 16, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NsURL  = []byte{107, 167, 184, 17, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NsOID  = []byte{107, 167, 184, 18, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NsX500 = []byte{107, 167, 184, 20, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
)

// ////////////////////////////////////////////////////////////////////////////////// //

// UUID contains UUID data
type UUID []byte

// ////////////////////////////////////////////////////////////////////////////////// //

// UUID4 generates random generated UUID v4
func UUID4() UUID {
	uuid := make(UUID, 16)

	rand.Read(uuid)

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return UUID(uuid)
}

// UUID5 generates UUID v5 based on SHA-1 hash of namespace UUID and name
func UUID5(ns []byte, name string) UUID {
	uuid := make(UUID, 16)

	hash := sha1.New()
	hash.Write(ns)
	hash.Write([]byte(name))

	copy(uuid, hash.Sum(nil))

	uuid[6] = (uuid[6] & 0x0F) | 0x50
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return UUID(uuid)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string representation of UUID
func (u UUID) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GenUUID generates v4 UUID (Universally Unique Identifier)
//
// Deprecated: Use method UUID4.String() instead
func GenUUID() string {
	return UUID4().String()
}

// GenUUID4 generates random generated UUID
//
// Deprecated: Use method UUID4.String() instead
func GenUUID4() string {
	return UUID4().String()
}

// GenUUID5 generates UUID based on SHA-1 hash of namespace UUID and name
//
// Deprecated: Use method UUID5.String() instead
func GenUUID5(ns []byte, name string) string {
	return UUID5(ns, name).String()
}

// ////////////////////////////////////////////////////////////////////////////////// //
