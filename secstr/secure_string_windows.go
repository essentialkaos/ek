// Package secstr provides methods and structs for working with protected (secure) strings
package secstr

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ String contains protected data
type String struct {
	Data []byte
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ NewSecureString creates new secure string
func NewSecureString(data interface{}) (*String, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Destroy destroys data
func (s *String) Destroy() error {
	panic("UNSUPPORTED")
	return nil
}
