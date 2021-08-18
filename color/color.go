// Package color provides methods for working with colors
package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"

	"pkg.re/essentialkaos/ek.v12/mathutil"
)

// RGB2Hex converts RGB color to Hex
func RGB2Hex(r, g, b int) int {
	r = mathutil.Between(r, 0, 255)
	g = mathutil.Between(g, 0, 255)
	b = mathutil.Between(b, 0, 255)

	return r<<16 | g<<8 | b
}

// Hex2RGB converts Hex color to RGB
func Hex2RGB(h int) (int, int, int) {
	h = mathutil.Between(h, 0x000000, 0xFFFFFF)
	return h >> 16 & 0xFF, h >> 8 & 0xFF, h & 0xFF
}

// RGBA2Hex converts RGBA color to Hex
func RGBA2Hex(r, g, b, a int) int64 {
	r = mathutil.Between(r, 0, 255)
	g = mathutil.Between(g, 0, 255)
	b = mathutil.Between(b, 0, 255)
	a = mathutil.Between(a, 0, 255)

	return int64(r)<<24 | int64(g)<<16 | int64(b)<<8 | int64(a)
}

// Hex2RGBA converts Hex color to RGBA
func Hex2RGBA(h int64) (int, int, int, int) {
	h = mathutil.Between64(h, 0x000000, 0xFFFFFFFF)

	if h >= 0xFFFFFF {
		return int(h>>24) & 0xFF, int(h>>16) & 0xFF, int(h>>8) & 0xFF, int(h) & 0xFF
	}

	return int(h) >> 16 & 0xFF, int(h>>8) & 0xFF, int(h) & 0xFF, 0
}

// RGB2HSB converts RGB color to HSB (HSV)
func RGB2HSB(r, g, b int) (int, int, int) {
	if r+g+b == 0 {
		return 0, 0, 0
	}

	r = mathutil.Between(r, 0, 255)
	g = mathutil.Between(g, 0, 255)
	b = mathutil.Between(b, 0, 255)

	max := mathutil.Max(mathutil.Max(r, g), b)
	min := mathutil.Min(mathutil.Min(r, g), b)

	var h int
	var s, bb float64

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

	return h, int(s), int(bb)
}

// HSB2RGB converts HSB (HSV) color to RGB
func HSB2RGB(h, s, b int) (int, int, int) {
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

	return int(mathutil.Round(r*0xFF, 0)),
		int(mathutil.Round(g*0xFF, 0)),
		int(mathutil.Round(bb*0xFF, 0))
}

// RGB2CMYK converts RGB color to CMYK
func RGB2CMYK(r, g, b int) (float64, float64, float64, float64) {
	r = mathutil.Between(r, 0, 255)
	g = mathutil.Between(g, 0, 255)
	b = mathutil.Between(b, 0, 255)

	R, G, B := float64(r)/255, float64(g)/255, float64(b)/255
	K := 1.0 - math.Max(math.Max(R, G), B)

	return calcCMYKColor(R, K),
		calcCMYKColor(G, K),
		calcCMYKColor(B, K),
		K
}

// CMYK2RGB converts CMYK color to RGB
func CMYK2RGB(c, m, y, k float64) (int, int, int) {
	c = mathutil.BetweenF(c, 0.0, 1.0)
	m = mathutil.BetweenF(m, 0.0, 1.0)
	y = mathutil.BetweenF(y, 0.0, 1.0)
	k = mathutil.BetweenF(k, 0.0, 1.0)

	return int(255 * (1 - c) * (1 - k)),
		int(255 * (1 - m) * (1 - k)),
		int(255 * (1 - y) * (1 - k))
}

// RGB2HSL converts RGB color to HSL
func RGB2HSL(r, g, b int) (float64, float64, float64) {
	r = mathutil.Between(r, 0, 255)
	g = mathutil.Between(g, 0, 255)
	b = mathutil.Between(b, 0, 255)

	R, G, B := float64(r)/255, float64(g)/255, float64(b)/255
	max := math.Max(math.Max(R, G), B)
	min := math.Min(math.Min(R, G), B)

	h, s, l := 0.0, 0.0, (min+max)/2.0

	if max != min {
		diff := max - min

		if l > 0.5 {
			s = diff / (2.0 - max - min)
		} else {
			s = diff / (max + min)
		}

		switch max {
		case R:
			if G < B {
				h = (G-B)/diff + 6.0
			} else {
				h = (G - B) / diff
			}
		case G:
			h = (B-R)/diff + 2.0
		case B:
			h = (R-G)/diff + 4.0
		}

		h /= 6
	}

	return h, s, l
}

// HSL2RGB converts HSL color to RGB
func HSL2RGB(h, s, l float64) (int, int, int) {
	h = mathutil.BetweenF(h, 0.0, 1.0)
	s = mathutil.BetweenF(s, 0.0, 1.0)
	l = mathutil.BetweenF(l, 0.0, 1.0)

	R, G, B := l, l, l

	if s != 0 {
		var q float64

		if l > 0.5 {
			q = l + s - (l * s)
		} else {
			q = l * (1.0 + s)
		}

		p := (2.0 * l) - q

		R = HUE2RGB(p, q, h+1.0/3.0)
		G = HUE2RGB(p, q, h)
		B = HUE2RGB(p, q, h-1.0/3.0)
	}

	return int(R * 255), int(G * 255), int(B * 255)
}

// HUE2RGB calculates HUE value for given RGB color
func HUE2RGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}

	if t > 1 {
		t -= 1
	}

	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6.0*t
	case t < 1.0/2.0:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*(2.0/3.0-t)*6.0
	}

	return p
}

// IsRGBA checks if Hex coded color has alpha channel info
func IsRGBA(h int64) bool {
	return h > 0xFFFFFF
}

// RGBLuminance returns relative luminance for RGB color
func RGBLuminance(r, g, b int) float64 {
	r = mathutil.Between(r, 0, 255)
	g = mathutil.Between(g, 0, 255)
	b = mathutil.Between(b, 0, 255)

	R := calcLumColor(float64(r) / 255)
	G := calcLumColor(float64(g) / 255)
	B := calcLumColor(float64(b) / 255)

	return 0.2126*R + 0.7152*G + 0.0722*B
}

// HEXLuminance returns relative luminance for HEX color
func HEXLuminance(h int) float64 {
	r, g, b := Hex2RGB(h)
	return RGBLuminance(r, g, b)
}

// Contrast calculates contrast ratio of foreground and background colors
func Contrast(fg, bg int) float64 {
	fg = mathutil.Between(fg, 0x000000, 0xFFFFFF)
	bg = mathutil.Between(bg, 0x000000, 0xFFFFFF)

	L1 := HEXLuminance(fg) + 0.05
	L2 := HEXLuminance(bg) + 0.05

	if L1 > L2 {
		return L1 / L2
	}

	return L2 / L1
}

// RGB2Term convert rgb color to terminal color code
// https://misc.flogisoft.com/bash/tip_colors_and_formatting#colors1
func RGB2Term(r, g, b int) int {
	// grayscale
	if r == g && g == b {
		if r == 175 {
			return 145
		}

		return (r / 10) + 232
	}

	return 36*(r/51) + 6*(g/51) + (b / 51) + 16
}

// ////////////////////////////////////////////////////////////////////////////////// //

func calcCMYKColor(c, k float64) float64 {
	if c == 0 && k == 1 {
		return 0
	}

	return (1 - c - k) / (1 - k)
}

func calcLumColor(c float64) float64 {
	if c <= 0.03928 {
		return c / 12.92
	}

	return math.Pow(((c + 0.055) / 1.055), 2.4)
}
