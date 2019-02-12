package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRun() {
	err := Run("/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func ExampleSudo() {
	err := Sudo("/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func ExampleRunAsUser() {
	// run echo as user some user and redirect output to /var/log/output.log
	err := RunAsUser("someuser", "/var/log/output.log", "/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
