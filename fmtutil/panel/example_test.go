package panel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleErrorPanel() {
	ErrorPanel(
		"Can't send data to remote server.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
		WRAP, BOTTOM_LINE,
	)
}

func ExampleWarnPanel() {
	WarnPanel(
		"Can't find user bob on system.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
		WRAP,
	)
}

func ExampleInfoPanel() {
	InfoPanel(
		"Auto-saving is enabled - data will be saved after editing.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
		WRAP, INDENT_OUTER,
	)
}

func ExamplePanel() {
	Panel(
		"Yolo", "{m}",
		"Excepteur sint occaecat cupidatat non proident.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
		WRAP, BOTTOM_LINE,
	)
}
