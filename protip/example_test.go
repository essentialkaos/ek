package protip

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleAdd() {
	// Add simple tip
	Add(&Tip{
		Title:   "Tip #1",
		Message: `Tip message`,
	})

	// Add tip with custom weight (≠ 0.5)
	Add(&Tip{
		Title:   "Tip #2",
		Message: `Tip message`,
		Weight:  0.1,
	})

	// Add tip with custom weight (≠ 0.5) and color
	Add(&Tip{
		Title:    "Tip #3",
		Message:  `Tip message`,
		Weight:   0.8,
		ColorTag: "{b}",
	})
}

func ExampleShow() {
	// Add simple tip
	Add(&Tip{
		Title:   "Tip #1",
		Message: `Tip message`,
	})

	// Increase default probability to 50%
	Probability = 0.5

	// Set default color to green
	ColorTag = "{g}"

	// Try to show tip
	Show(false)

	// Force show
	Show(true)
}
