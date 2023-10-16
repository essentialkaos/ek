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

// GenUUID generates v4 UUID (Universally Unique Identifier)
//
// Deprecated: Use method UUID instead
func GenUUID() string {
	return UUID4()
}

// UUID4 generates v4 UUID (Universally Unique Identifier)
func UUID() string {
	return UUID4()
}

// GenUUID4 generates random generated UUID
//
// Deprecated: Use method UUID4 instead
func GenUUID4() string {
	return UUID4()
}

// UUID4 generates random generated UUID
//
// Deprecated: Use method UUID4 instead
func UUID4() string {
	uuid := make([]byte, 16)

	rand.Read(uuid)

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return toString(uuid)
}

// GenUUID5 generates UUID based on SHA-1 hash of namespace UUID and name
//
// Deprecated: Use method UUID5 instead
func GenUUID5(ns []byte, name string) string {
	return UUID5(ns, name)
}

// UUID5 generates UUID based on SHA-1 hash of namespace UUID and name
func UUID5(ns []byte, name string) string {
	uuid := make([]byte, 16)

	hash := sha1.New()
	hash.Write(ns)
	hash.Write([]byte(name))

	copy(uuid, hash.Sum(nil))

	uuid[6] = (uuid[6] & 0x0F) | 0x50
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return toString(uuid)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func toString(uuid []byte) string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])

	return string(buf)
}
