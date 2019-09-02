package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
