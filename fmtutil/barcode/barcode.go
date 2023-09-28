// Package barcode provides methods to generate colored representation of unique data
package barcode

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"hash/crc32"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// dots is used for dots generation
var dots = []string{
	"⠮", "⠉", "⠒", "⠤", "⠝", "⠳", "⠪", "⠭",
	"⠕", "⠔", "⠮", "⠌", "⠡", "⠢", "⠜", "⠵",
}

// lines is used for lines generation
var lines = []string{
	"╬", "╪", "╩", "╧", "╦", "╤", "╣", "╡",
	"╠", "╞", "╝", "╛", "╚", "╘", "╕", "╔",
}

// boxes is used for boxes generation
var boxes = []string{
	"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█",
	"█", "▇", "▆", "▅", "▄", "▃", "▂", "▁",
}

// colors contains ASCII color codes
var colors = []int{
	1, 2, 3, 4, 5, 6, 9, 10,
	11, 12, 13, 14, 2, 12, 5, 14,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Dots generates dots
func Dots(data []byte) string {
	if len(data) == 0 {
		return fmt.Sprintf("\033[38;5;1m%s\033[0m", "⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒")
	}

	return generate(data, dots)
}

// Lines generates lines
func Lines(data []byte) string {
	if len(data) == 0 {
		return fmt.Sprintf("\033[38;5;1m%s\033[0m", "╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬")
	}

	return generate(data, lines)
}

// Boxes generates boxes
func Boxes(data []byte) string {
	if len(data) == 0 {
		return fmt.Sprintf("\033[38;5;1m%s\033[0m", "████████████████")
	}

	return generate(data, boxes)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// generate generates barcode for given data
func generate(data []byte, alpha []string) string {
	c := crc32.ChecksumIEEE(data)
	c1, c2, c3, c4 := (c>>24)&0xFF, (c>>16)&0xFF, (c>>8)&0xFF, c&0xFF
	return barcode(c1, alpha) + barcode(c2, alpha) + barcode(c3, alpha) + barcode(c4, alpha)
}

// barcode generates barcode part
func barcode(c uint32, alpha []string) string {
	var result string

	for _, i := range []uint32{c >> 4 & 0xF, c >> 2 & 0xF, c >> 1 & 0xF, c & 0xF} {
		result += fmt.Sprintf("\033[38;5;%dm%s\033[0m", colors[i], alpha[i])
	}

	return result
}
