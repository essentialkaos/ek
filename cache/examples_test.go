package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNew() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")

	fmt.Println(store.Get("test"))
	// Output: ABCD
}

func ExampleStore_Set() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")

	fmt.Println(store.Get("test"))
	// Output: ABCD
}

func ExampleStore_Has() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")

	fmt.Println(store.Has("test"))
	// Output: true
}

func ExampleStore_Get() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")

	fmt.Println(store.Get("test"))
	// Output: ABCD
}

func ExampleStore_GetWithExpiration() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")

	item, exp := store.GetWithExpiration("test")

	fmt.Println(item, exp.String())
}

func ExampleStore_Delete() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")
	store.Delete("test")

	fmt.Println(store.Get("test"))
	// Output: <nil>
}

func ExampleStore_Flush() {
	store := New(time.Second, time.Minute)

	store.Set("test", "ABCD")
	store.Flush()

	fmt.Println(store.Get("test"))
	// Output: <nil>
}
