package ek

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/crypto/bcrypt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// worthless is used as dependency fix
func worthless() {
	bcrypt.Cost(nil)
}
