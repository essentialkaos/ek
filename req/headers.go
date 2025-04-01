package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Headers is a map[string]string used for headers
type Headers map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// Set sets header value
func (h Headers) Set(name, value string) bool {
	if h == nil || name == "" || value == "" {
		return false
	}

	h[name] = value

	return true
}

// SetIf sets header value if given condition is true
func (h Headers) SetIf(cond bool, name, value string) bool {
	if h == nil || name == "" || !cond {
		return false
	}

	h[name] = value

	return true
}

// Get returns header with given name
func (h Headers) Get(name string) string {
	if h == nil || name == "" {
		return ""
	}

	return h[name]
}

// Delete deletes parameter with given name
func (h Headers) Delete(name string) bool {
	if h == nil || name == "" {
		return false
	}

	delete(h, name)

	return true
}

// DeleteIf deletes parameter with given name if condition is true
func (h Headers) DeleteIf(cond bool, name string) bool {
	if h == nil || name == "" || !cond {
		return false
	}

	if h[name] != "" {
		delete(h, name)
	}

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //
