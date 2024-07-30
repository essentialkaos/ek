package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNew() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Get("test"))
}

func ExampleCache_Set() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")
	cache.Set("test", "ABCD", 15*time.Minute)

	fmt.Println(cache.Get("test"))
}

func ExampleCache_Has() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Has("test"))
}

func ExampleCache_Get() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Get("test"))
}

func ExampleCache_Size() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Size())
}

func ExampleCache_Expired() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Expired())
}

func ExampleCache_GetWithExpiration() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	item, exp := cache.GetWithExpiration("test")

	fmt.Println(item, exp.String())
}

func ExampleCache_Delete() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")
	cache.Delete("test")

	fmt.Println(cache.Get("test"))
}

func ExampleCache_Flush() {
	cache, _ := New(Config{
		Dir:               "/path/to/cache",
		DefaultExpiration: time.Minute,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")
	cache.Flush()

	fmt.Println(cache.Get("test"))
}
