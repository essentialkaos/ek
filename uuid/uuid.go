// Package uuid contains methods for generating version 4 and 5 UUID's
package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Predefined namespace UUID's
var (
	// NsDNS is the predefined UUID namespace for fully qualified domain names
	NsDNS = UUID{107, 167, 184, 16, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}

	// NsURL is the predefined UUID namespace for URLs
	NsURL = UUID{107, 167, 184, 17, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}

	// NsOID is the predefined UUID namespace for ISO OIDs
	NsOID = UUID{107, 167, 184, 18, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}

	// NsX500 is the predefined UUID namespace for X.500 distinguished names
	NsX500 = UUID{107, 167, 184, 20, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
)

// ////////////////////////////////////////////////////////////////////////////////// //

// UUID is a 16-byte RFC 4122 universally unique identifier
type UUID [16]byte

// ////////////////////////////////////////////////////////////////////////////////// //

// randGenerator is function to generate random data
var randGenerator = rand.Read

// ////////////////////////////////////////////////////////////////////////////////// //

// UUID4 returns a randomly generated UUID (version 4)
func UUID4() UUID {
	var uuid UUID

	_, err := randGenerator(uuid[:])

	if err != nil {
		return UUID{}
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return UUID(uuid)
}

// UUID5 returns a UUID (version 5) derived from the SHA-1 hash of ns and name
func UUID5(ns UUID, name string) UUID {
	var uuid UUID

	hash := sha1.New()
	hash.Write(ns[:])
	hash.Write([]byte(name))

	copy(uuid[:], hash.Sum(nil))

	uuid[6] = (uuid[6] & 0x0F) | 0x50
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return UUID(uuid)
}

// UUID7 returns a UUID (version 7) with a millisecond-precision Unix timestamp in
// the high bits
func UUID7() UUID {
	var uuid UUID

	_, err := randGenerator(uuid[:])

	if err != nil {
		return UUID{}
	}

	ts := uint64(time.Now().UnixNano() / 1_000_000)

	uuid[0] = byte((ts >> 40) & 0xFF)
	uuid[1] = byte((ts >> 32) & 0xFF)
	uuid[2] = byte((ts >> 24) & 0xFF)
	uuid[3] = byte((ts >> 16) & 0xFF)
	uuid[4] = byte((ts >> 8) & 0xFF)
	uuid[5] = byte(ts & 0xFF)

	uuid[6] = (uuid[6] & 0x0F) | 0x70
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return uuid
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsZero reports whether the UUID is the all-zero value
func (u UUID) IsZero() bool {
	return u == UUID{}
}

// String returns the standard hyphenated lowercase hex representation of the UUID
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
