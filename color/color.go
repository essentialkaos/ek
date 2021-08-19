// Package color provides methods for working with colors
package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"

	"pkg.re/essentialkaos/ek.v12/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Hex uint32 // Hex color 0x00000000 - 0xFFFFFFFF

type RGB struct {
	R uint8 // Red
	G uint8 // Green
	B uint8 // Blue
}

type RGBA struct {
	R uint8 // Red
	G uint8 // Green
	B uint8 // Blue
	A uint8 // Alpha
}

type CMYK struct {
	C float64 // Cyan
	M float64 // Magenta
	Y float64 // Yellow
	K float64 // Key (black)
}

type HSV struct {
	H float64 // Hue
	S float64 // Saturation
	V float64 // Lightness
}

type HSL struct {
	H float64 // Hue
	S float64 // Saturation
	L float64 // Value
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsRGBA returns true if color contains info about alpha channel
func (c Hex) IsRGBA() bool {
	return c > 0xFFFFFF
}

// ToHex converts RGB color to hex
func (c RGB) ToHex() Hex {
	return RGB2Hex(c)
}

// ToCMYK converts RGB color to CMYK
func (c RGB) ToCMYK() CMYK {
	return RGB2CMYK(c)
}

// ToHSV converts RGB color to HSV
func (c RGB) ToHSV() HSV {
	return RGB2HSV(c)
}

// ToHSL converts RGB color to HSL
func (c RGB) ToHSL() HSL {
	return RGB2HSL(c)
}

// ToTerm converts RGB color to terminal color code
func (c RGB) ToTerm() int {
	return RGB2Term(c)
}

// ToHex converts RGBA color to hex
func (c RGBA) ToHex() Hex {
	return RGBA2Hex(c)
}

// ToRGB converts CMYK color to RGB
func (c CMYK) ToRGB() RGB {
	return CMYK2RGB(c)
}

// ToRGB converts HSV color to RGB
func (c HSV) ToRGB() RGB {
	return HSV2RGB(c)
}

// ToRGB converts HSL color to RGB
func (c HSL) ToRGB() RGB {
	return HSL2RGB(c)
}

// ToRGB converts hex color to RGB
func (c Hex) ToRGB() RGB {
	return Hex2RGB(c)
}

// ToRGB converts hex color to RGBA
func (c Hex) ToRGBA() RGBA {
	return Hex2RGBA(c)
}

// ToWeb converts hex color notation used in web (#RRGGBB/#RRGGBBAA)
func (c Hex) ToWeb(caps bool) string {
	if caps {
		return fmt.Sprintf("#%X", uint32(c))
	}

	return fmt.Sprintf("#%x", uint32(c))
}

// String returns string representation of RGB color
func (c RGB) String() string {
	return fmt.Sprintf(
		"RGB{R:%d G:%d B:%d}",
		c.R, c.G, c.B,
	)
}

// String returns string representation of RGBA color
func (c RGBA) String() string {
	return fmt.Sprintf(
		"RGBA{R:%d G:%d B:%d A:%.2f}",
		c.R, c.G, c.B, float64(c.A)/255.0,
	)
}

// String returns string representation of hex color
func (c Hex) String() string {
	return fmt.Sprintf("Hex{#%X}", uint32(c))
}

// String returns string representation of CMYK color
func (c CMYK) String() string {
	return fmt.Sprintf(
		"CMYK{C:%.0f%% M:%.0f%% Y:%.0f%% K:%.0f%%}",
		c.C*100.0, c.M*100.0, c.Y*100.0, c.K*100.0,
	)
}

// String returns string representation of HSV color
func (c HSV) String() string {
	return fmt.Sprintf(
		"HSV{H:%.0f° S:%.0f%% V:%.0f%%}",
		c.H*360, c.S*100, c.V*100,
	)
}

// String returns string representation of HSL color
func (c HSL) String() string {
	return fmt.Sprintf(
		"HSL{H:%.0f° S:%.0f%% L:%.0f%%}",
		c.H*360, c.S*100, c.L*100,
	)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// RGB2Hex converts RGB color to Hex
func RGB2Hex(c RGB) Hex {
	return Hex(int(c.R)<<16 | int(c.G)<<8 | int(c.B))
}

// Hex2RGB converts Hex color to RGB
func Hex2RGB(h Hex) RGB {
	return RGB{uint8(h >> 16 & 0xFF), uint8(h >> 8 & 0xFF), uint8(h & 0xFF)}
}

// RGBA2Hex converts RGBA color to Hex
func RGBA2Hex(c RGBA) Hex {
	return Hex(int64(c.R)<<24 | int64(c.G)<<16 | int64(c.B)<<8 | int64(c.A))
}

// Hex2RGBA converts Hex color to RGBA
func Hex2RGBA(h Hex) RGBA {
	if h >= 0xFFFFFF {
		return RGBA{uint8(h>>24) & 0xFF, uint8(h>>16) & 0xFF, uint8(h>>8) & 0xFF, uint8(h) & 0xFF}
	}

	return RGBA{uint8(h) >> 16 & 0xFF, uint8(h>>8) & 0xFF, uint8(h) & 0xFF, 0}
}

// RGB2Term convert rgb color to terminal color code
// https://misc.flogisoft.com/bash/tip_colors_and_formatting#colors1
func RGB2Term(c RGB) int {
	R, G, B := int(c.R), int(c.G), int(c.B)

	// grayscale
	if R == G && G == B {
		if R == 175 {
			return 145
		}

		return (R / 10) + 232
	}

	return 36*(R/51) + 6*(G/51) + (B / 51) + 16
}

// RGB2CMYK converts RGB color to CMYK
func RGB2CMYK(c RGB) CMYK {
	R, G, B := float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0
	K := 1.0 - math.Max(math.Max(R, G), B)

	return CMYK{
		calcCMYKColor(R, K),
		calcCMYKColor(G, K),
		calcCMYKColor(B, K),
		K,
	}
}

// CMYK2RGB converts CMYK color to RGB
func CMYK2RGB(c CMYK) RGB {
	C := mathutil.BetweenF(c.C, 0.0, 1.0)
	M := mathutil.BetweenF(c.M, 0.0, 1.0)
	Y := mathutil.BetweenF(c.Y, 0.0, 1.0)
	K := mathutil.BetweenF(c.K, 0.0, 1.0)

	return RGB{
		uint8(255 * (1 - C) * (1 - K)),
		uint8(255 * (1 - M) * (1 - K)),
		uint8(255 * (1 - Y) * (1 - K)),
	}
}

// RGB2HSV converts RGB color to HSV (HSB)
func RGB2HSV(c RGB) HSV {
	R, G, B := float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0

	max := math.Max(math.Max(R, G), B)
	min := math.Min(math.Min(R, G), B)

	h, s, v := 0.0, 0.0, max

	if max != min {
		d := max - min
		s = d / max
		h = calcHUE(max, R, G, B, d)
	}

	return HSV{h, s, v}
}

// HSV2RGB converts HSV (HSB) color to RGB
func HSV2RGB(c HSV) RGB {
	i := (c.H * 360.0) / 60.0
	f := i - math.Floor(i)

	p := c.V * (1 - c.S)
	q := c.V * (1 - f*c.S)
	t := c.V * (1 - (1-f)*c.S)

	var R, G, B float64

	switch int(c.H*6) % 6 {
	case 0:
		R, G, B = c.V, t, p
	case 1:
		R, G, B = q, c.V, p
	case 2:
		R, G, B = p, c.V, t
	case 3:
		R, G, B = p, q, c.V
	case 4:
		R, G, B = t, p, c.V
	case 5:
		R, G, B = c.V, p, q
	}

	return RGB{uint8(R * 0xFF), uint8(G * 0xFF), uint8(B * 0xFF)}
}

// RGB2HSL converts RGB color to HSL
func RGB2HSL(c RGB) HSL {
	R, G, B := float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0

	max := math.Max(math.Max(R, G), B)
	min := math.Min(math.Min(R, G), B)

	h, s, l := 0.0, 0.0, (min+max)/2.0

	if max != min {
		d := max - min

		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}

		h = calcHUE(max, R, G, B, d)
	}

	return HSL{h, s, l}
}

// HSL2RGB converts HSL color to RGB
func HSL2RGB(c HSL) RGB {
	R, G, B := c.L, c.L, c.L

	if c.S != 0 {
		var q float64

		if c.L > 0.5 {
			q = c.L + c.S - (c.L * c.S)
		} else {
			q = c.L * (1.0 + c.S)
		}

		p := (2.0 * c.L) - q

		R = HUE2RGB(p, q, c.H+1.0/3.0)
		G = HUE2RGB(p, q, c.H)
		B = HUE2RGB(p, q, c.H-1.0/3.0)
	}

	return RGB{uint8(R * 255), uint8(G * 255), uint8(B * 255)}
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

// Luminance returns relative luminance for RGB color
func Luminance(c RGB) float64 {
	R := calcLumColor(float64(c.R) / 255)
	G := calcLumColor(float64(c.G) / 255)
	B := calcLumColor(float64(c.B) / 255)

	return 0.2126*R + 0.7152*G + 0.0722*B
}

// Contrast calculates contrast ratio of foreground and background colors
func Contrast(fg, bg Hex) float64 {
	L1 := Luminance(fg.ToRGB()) + 0.05
	L2 := Luminance(bg.ToRGB()) + 0.05

	if L1 > L2 {
		return L1 / L2
	}

	return L2 / L1
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

func calcHUE(max, r, g, b, d float64) float64 {
	var h float64

	switch max {
	case r:
		if g < b {
			h = (g-b)/d + 6.0
		} else {
			h = (g - b) / d
		}
	case g:
		h = (b-r)/d + 2.0
	case b:
		h = (r-g)/d + 4.0
	}

	return h / 6
}
