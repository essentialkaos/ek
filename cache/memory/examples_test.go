package memory

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
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Get("test"))
	// Output: ABCD
}

func ExampleCache_Set() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")
	cache.Set("test", "ABCD", 15*time.Minute)

	fmt.Println(cache.Get("test"))
	// Output: ABCD
}

func ExampleCache_Has() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Has("test"))
	// Output: true
}

func ExampleCache_Get() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Get("test"))
	// Output: ABCD
}

func ExampleCache_Size() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Size())
	// Output: 1
}

func ExampleCache_Expired() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	fmt.Println(cache.Expired())
	// Output: 0
}

func ExampleCache_GetWithExpiration() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")

	item, exp := cache.GetWithExpiration("test")

	fmt.Println(item, exp.String())
}

func ExampleCache_Delete() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")
	cache.Delete("test")

	fmt.Println(cache.Get("test"))
	// Output: <nil>
}

func ExampleCache_Flush() {
	cache, _ := New(Config{
		DefaultExpiration: time.Second,
		CleanupInterval:   time.Minute,
	})

	cache.Set("test", "ABCD")
	cache.Flush()

	fmt.Println(cache.Get("test"))
	// Output: <nil>
}
