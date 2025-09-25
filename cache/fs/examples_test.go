package fs

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
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
		CleanupInterval:   cache.MINUTE,
		ValidationRegexp:  `^[a-fA-F0-9]{8}$`,
	})

	c.Set("045d01c1", "Test data")

	fmt.Println(c.Get("test"))
}

func ExampleCache_Set() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")
	c.Set("test", "ABCD", 15*cache.MINUTE)

	fmt.Println(c.Get("test"))
}

func ExampleCache_Has() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Has("test"))
}

func ExampleCache_Get() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Get("test"))
}

func ExampleCache_Size() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Size())
}

func ExampleCache_Expired() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")

	fmt.Println(c.Expired())
}

func ExampleCache_GetWithExpiration() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")

	item, exp := c.GetWithExpiration("test")

	fmt.Println(item, exp.String())
}

func ExampleCache_Invalidate() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")

	// ... some time after

	c.Invalidate()
}

func ExampleCache_Delete() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")
	c.Delete("test")

	fmt.Println(c.Get("test"))
}

func ExampleCache_Flush() {
	c, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: cache.DAY,
	})

	c.Set("test", "ABCD")
	c.Flush()

	fmt.Println(c.Get("test"))
}
