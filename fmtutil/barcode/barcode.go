// Package barcode provides methods to generate colored representation of unique data
package barcode

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"hash/crc32"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// dots is used for dots generation
var dots = []string{
	"⠮", "⠉", "⠒", "⠤", "⠝", "⠳", "⠪", "⠭", "⠕", "⠔", "⠮", "⠌", "⠡", "⠢", "⠜", "⠵",
}

// lines is used for lines generation
var lines = []string{
	"╬", "╪", "╩", "╧", "╦", "╤", "╣", "╡", "╠", "╞", "╝", "╛", "╚", "╘", "╕", "╔",
}

// boxes is used for boxes generation
var boxes = []string{
	"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁",
}

// colors contains ASCII color codes
var colors = []int{
	1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 13, 14, 2, 12, 5, 14,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Dots returns a 16-symbol ANSI-colored barcode string composed of braille dot
// characters. The glyph and foreground color of each symbol are derived from the
// CRC32-IEEE checksum of data, so any two distinct byte slices will, with high
// probability, produce visually different barcodes.
//
// If data is empty, Dots returns a uniform 16-character red error indicator
// ("⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒") to signal that no meaningful input was provided.
func Dots(data []byte) string {
	if len(data) == 0 {
		return "\033[38;5;1m⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒\033[0m"
	}

	return generate(data, dots)
}

// Lines returns a 16-symbol ANSI-colored barcode string composed of box-drawing
// line characters. The glyph and foreground color of each symbol are derived
// from the CRC32-IEEE checksum of data, providing a consistent visual
// fingerprint for any unique byte sequence.
//
// If data is empty, Lines returns a uniform 16-character red error indicator
// ("╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬") to signal that no meaningful input was provided.
func Lines(data []byte) string {
	if len(data) == 0 {
		return "\033[38;5;1m╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬\033[0m"
	}

	return generate(data, lines)
}

// Boxes returns a 16-symbol ANSI-colored barcode string composed of block
// element characters. The fill level and foreground color of each symbol are
// derived from the CRC32-IEEE checksum of data, giving a compact visual hash
// that is easy to compare at a glance in a terminal.
//
// If data is empty, Boxes returns a uniform 16-character red error indicator
// ("████████████████") to signal that no meaningful input was provided.
func Boxes(data []byte) string {
	if len(data) == 0 {
		return "\033[38;5;1m████████████████\033[0m"
	}

	return generate(data, boxes)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// generate computes the CRC32-IEEE checksum of data, splits it into four bytes
// (most-significant first), and concatenates the colored symbol segment for
// each byte into a single 16-symbol ANSI string
func generate(data []byte, alpha []string) string {
	var result strings.Builder

	c := crc32.ChecksumIEEE(data)

	c1, c2, c3, c4 := (c>>24)&0xFF, (c>>16)&0xFF, (c>>8)&0xFF, c&0xFF

	result.WriteString(barcode(c1, alpha))
	result.WriteString(barcode(c2, alpha))
	result.WriteString(barcode(c3, alpha))
	result.WriteString(barcode(c4, alpha))

	return result.String()
}

// barcode maps nibble-range indexes extracted from byte c to ANSI-colored
// terminal symbols using the colors and alpha lookup tables
func barcode(c uint32, alpha []string) string {
	var result strings.Builder

	lastColor := -1

	for _, i := range []uint32{(c >> 6) & 0x3, (c >> 4) & 0x3, (c >> 2) & 0x3, c & 0x3} {
		color := colors[i]

		if color != lastColor {
			if lastColor != -1 {
				result.WriteString("\033[0m")
			}

			fmt.Fprintf(&result, "\033[38;5;%dm", color)

			lastColor = color
		}

		result.WriteString(alpha[i])
	}

	if lastColor != -1 {
		result.WriteString("\033[0m")
	}

	return result.String()
}
