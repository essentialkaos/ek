package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// StringToInterface convert slice with strings to slice with interface{}
func StringToInterface(data []string) []interface{} {
	result := []interface{}{}

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// IntToInterface convert slice with ints to slice with interface{}
func IntToInterface(data []int) []interface{} {
	result := []interface{}{}

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// Contains check if string slice contains some value
func Contains(slice []string, value string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
