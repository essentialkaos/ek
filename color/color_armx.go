// +build arm, arm64

// Package color provides methods for working with colors
package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"

	"pkg.re/essentialkaos/ek.v11/mathutil"
)

// RGB2Hex convert RGB color to Hex
func RGB2Hex(r, g, b int64) int64 {
	return r<<16 | g<<8 | b
}

// Hex2RGB convert Hex color to RGB
func Hex2RGB(h int64) (int64, int64, int64) {
	if IsRGBA(h) {
		return 0xFF, 0xFF, 0xFF
	}

	return h >> 16 & 0xFF, h >> 8 & 0xFF, h & 0xFF
}

// RGBA2Hex convert RGBA color to Hex
func RGBA2Hex(r, g, b, a int64) int64 {
	return r<<24 | g<<16 | b<<8 | a
}

// Hex2RGBA convert Hex color to RGBA
func Hex2RGBA(h int64) (int64, int64, int64, int64) {
	if h >= 0xFFFFFF {
		return h >> 24 & 0xFF, h >> 16 & 0xFF, h >> 8 & 0xFF, h & 0xFF
	}

	return h >> 16 & 0xFF, h >> 8 & 0xFF, h & 0xFF, 0
}

// RGB2HSB convert RGB color to HSB (HSV)
func RGB2HSB(r, g, b int64) (int64, int64, int64) {
	if r+g+b == 0 {
		return 0, 0, 0
	}

	max := mathutil.Max64(mathutil.Max64(r, g), b)
	min := mathutil.Min64(mathutil.Min64(r, g), b)

	var (
		h     int64
		s, bb float64
	)

	switch max {
	case min:
		h = 0
	case r:
		h = (60*(g-b)/(max-min) + 360) % 360
	case g:
		h = (60*(b-r)/(max-min) + 120)
	case b:
		h = (60*(r-g)/(max-min) + 240)
	}

	bb = math.Ceil((float64(max) / 255.0) * 100.0)

	if max != 0 {
		fmax, fmin := float64(max), float64(min)
		s = math.Ceil(((fmax - fmin) / fmax) * 100.0)
	}

	return h, int64(s), int64(bb)
}

// HSB2RGB convert HSB (HSV) color to RGB
func HSB2RGB(h, s, b int64) (int64, int64, int64) {
	var r, g, bb float64

	if h+s+b == 0 {
		return 0, 0, 0
	}

	ts := float64(s) / 100.0
	tb := float64(b) / 100.0

	f := float64(h)/60.0 - math.Floor(float64(h)/60.0)
	p := (tb * (1 - ts))
	q := (tb * (1 - f*ts))
	t := (tb * (1 - (1-f)*ts))

	switch (h / 60) % 6 {
	case 0:
		r, g, bb = tb, t, p
	case 1:
		r, g, bb = q, tb, p
	case 2:
		r, g, bb = p, tb, t
	case 3:
		r, g, bb = p, q, tb
	case 4:
		r, g, bb = t, p, tb
	case 5:
		r, g, bb = tb, p, q
	}

	return int64(mathutil.Round(r*0xFF, 0)),
		int64(mathutil.Round(g*0xFF, 0)),
		int64(mathutil.Round(bb*0xFF, 0))
}

// IsRGBA if Hex coded color has alpha channel info
func IsRGBA(h int64) bool {
	return h > 0xFFFFFF
}

// RGB2Term convert rgb color to terminal color code
// https://misc.flogisoft.com/bash/tip_colors_and_formatting#colors1
func RGB2Term(r, g, b int64) int64 {
	// grayscale
	if r == g && g == b {
		if r == 175 {
			return 145
		}

		return (r / 10) + 232
	}

	return 36*(r/51) + 6*(g/51) + (b / 51) + 16
}
