package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "strings"

// ////////////////////////////////////////////////////////////////////////////////// //

// Headers is a map[string]string used for headers
type Headers map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// Set sets header value
func (h Headers) Set(name, value string) bool {
	if h == nil || name == "" || value == "" {
		return false
	}

	h[strings.ToLower(name)] = value

	return true
}

// SetIf sets header value if given condition is true
func (h Headers) SetIf(cond bool, name, value string) bool {
	if h == nil || name == "" || !cond {
		return false
	}

	h[strings.ToLower(name)] = value

	return true
}

// Get returns header with given name
func (h Headers) Get(name string) string {
	if h == nil || name == "" {
		return ""
	}

	name = strings.ToLower(name)

	for hn, hv := range h {
		if strings.ToLower(hn) == name {
			return hv
		}
	}

	return ""
}

// Has returns true if header is set
func (h Headers) Has(name string) bool {
	if h == nil || name == "" {
		return false
	}

	name = strings.ToLower(name)

	for hn := range h {
		if strings.ToLower(hn) == name {
			return true
		}
	}

	return false
}

// Delete deletes parameter with given name
func (h Headers) Delete(name string) bool {
	if h == nil || name == "" {
		return false
	}

	name = strings.ToLower(name)

	for hn := range h {
		if strings.ToLower(hn) == name {
			delete(h, strings.ToLower(hn))
			return true
		}
	}

	return false
}

// DeleteIf deletes parameter with given name if condition is true
func (h Headers) DeleteIf(cond bool, name string) bool {
	if h == nil || name == "" || !cond {
		return false
	}

	return h.Delete(name)
}

// ////////////////////////////////////////////////////////////////////////////////// //
