package memory

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v13/cache"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNew() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Get("test"))
	// Output: ABCD
}

func ExampleCache_Set() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")
	c.Set("test", "ABCD", 15*cache.MINUTE)

	fmt.Println(c.Get("test"))
	// Output: ABCD
}

func ExampleCache_Has() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Has("test"))
	// Output: true
}

func ExampleCache_Get() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Get("test"))
	// Output: ABCD
}

func ExampleCache_Size() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Size())
	// Output: 1
}

func ExampleCache_Expired() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Expired())
	// Output: 0
}

func ExampleCache_GetWithExpiration() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")

	item, exp := c.GetWithExpiration("test")

	fmt.Println(item, exp.String())
}

func ExampleCache_Delete() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")
	c.Delete("test")

	fmt.Println(c.Get("test"))
	// Output: <nil>
}

func ExampleCache_Flush() {
	c, _ := New(Config{
		DefaultExpiration: cache.SECOND,
		CleanupInterval:   cache.MINUTE,
	})

	c.Set("test", "ABCD")
	c.Flush()

	fmt.Println(c.Get("test"))
	// Output: <nil>
}
