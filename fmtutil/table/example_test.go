package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewTable() {
	t := NewTable().SetHeaders("id", "user", "balance").SetSizes(4)

	t.HeaderCapitalize = true

	t.Add(1, "{g}Bob{!}", 1.42)
	t.Add(2, "John", 73.1)
	t.Add(3, "Mary", 2.29)

	t.Render()
}
