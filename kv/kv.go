// Package kv provides simple key-value structs
package kv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type kvSlice []KV

// KV is key-value struct
type KV struct {
	Key   string
	Value interface{}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s kvSlice) Len() int           { return len(s) }
func (s kvSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s kvSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }

// ////////////////////////////////////////////////////////////////////////////////// //

// Sort sorts slice with kv structs by key
func Sort(slice []KV) {
	sort.Sort(kvSlice(slice))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String return value as string
func (k KV) String() string {
	return k.Value.(string)
}

// Int return value as int
func (k KV) Int() int {
	return k.Value.(int)
}

// Float return value as float64
func (k KV) Float() float64 {
	return k.Value.(float64)
}
