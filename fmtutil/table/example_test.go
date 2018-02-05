package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewTable() {
	t := NewTable()

	t.SetHeaders("id", "user", "balance")
	t.SetSizes(4, 12)
	t.SetAlignments(ALIGN_RIGHT, ALIGN_RIGHT, ALIGN_LEFT)

	t.Add(1, "{g}Bob{!}", 1.42)
	t.Add(2, "John", 73.1)
	t.Add(3, "Mary", 2.29)

	t.Render()
}
