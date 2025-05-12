package apps

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "fmt"

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleExtractVersion() {
	// Get Ruby version from 'ruby --version' command (line: 0, field: 1)
	rubyVersion := ExtractVersion([]string{"ruby", "--version"}, 0, 1)
	fmt.Println(rubyVersion)

	// Get Java version from 'java -version' command (line: 1, field: 3)
	javaVersion := ExtractVersion([]string{"java", "-version"}, 1, 3)
	fmt.Println(javaVersion)
}
