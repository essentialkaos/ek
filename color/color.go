// Package color provides methods for working with colors
package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"strconv"

	"github.com/essentialkaos/ek/v13/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Hex struct {
	v uint32 // Hex color 0x00000000 - 0xFFFFFFFF
	a bool   // alpha flag
}

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
	A float64 // Alpha
}

type HSL struct {
	H float64 // Hue
	S float64 // Saturation
	L float64 // Value
	A float64 // Alpha
}

// ////////////////////////////////////////////////////////////////////////////////// //

// colors is colors keywords
var colors = map[string]uint32{
	"aliceblue":            0xf0f8ff,
	"antiquewhite":         0xfaebd7,
	"aqua":                 0x00ffff,
	"aquamarine":           0x7fffd4,
	"azure":                0xf0ffff,
	"beige":                0xf5f5dc,
	"bisque":               0xffe4c4,
	"black":                0x000000,
	"blanchedalmond":       0xffebcd,
	"blue":                 0x0000ff,
	"blueviolet":           0x8a2be2,
	"brown":                0xa52a2a,
	"burlywood":            0xdeb887,
	"cadetblue":            0x5f9ea0,
	"chartreuse":           0x7fff00,
	"chocolate":            0xd2691e,
	"coral":                0xff7f50,
	"cornflowerblue":       0x6495ed,
	"cornsilk":             0xfff8dc,
	"crimson":              0xdc143c,
	"cyan":                 0x00ffff,
	"darkblue":             0x00008b,
	"darkcyan":             0x008b8b,
	"darkgoldenrod":        0xb8860b,
	"darkgray":             0xa9a9a9,
	"darkgreen":            0x006400,
	"darkgrey":             0xa9a9a9,
	"darkkhaki":            0xbdb76b,
	"darkmagenta":          0x8b008b,
	"darkolivegreen":       0x556b2f,
	"darkorange":           0xff8c00,
	"darkorchid":           0x9932cc,
	"darkred":              0x8b0000,
	"darksalmon":           0xe9967a,
	"darkseagreen":         0x8fbc8f,
	"darkslateblue":        0x483d8b,
	"darkslategray":        0x2f4f4f,
	"darkslategrey":        0x2f4f4f,
	"darkturquoise":        0x00ced1,
	"darkviolet":           0x9400d3,
	"deeppink":             0xff1493,
	"deepskyblue":          0x00bfff,
	"dimgray":              0x696969,
	"dimgrey":              0x696969,
	"dodgerblue":           0x1e90ff,
	"firebrick":            0xb22222,
	"floralwhite":          0xfffaf0,
	"forestgreen":          0x228b22,
	"fuchsia":              0xff00ff,
	"gainsboro":            0xdcdcdc,
	"ghostwhite":           0xf8f8ff,
	"gold":                 0xffd700,
	"goldenrod":            0xdaa520,
	"gray":                 0x808080,
	"green":                0x008000,
	"greenyellow":          0xadff2f,
	"grey":                 0x808080,
	"honeydew":             0xf0fff0,
	"hotpink":              0xff69b4,
	"indianred":            0xcd5c5c,
	"indigo":               0x4b0082,
	"ivory":                0xfffff0,
	"khaki":                0xf0e68c,
	"lavender":             0xe6e6fa,
	"lavenderblush":        0xfff0f5,
	"lawngreen":            0x7cfc00,
	"lemonchiffon":         0xfffacd,
	"lightblue":            0xadd8e6,
	"lightcoral":           0xf08080,
	"lightcyan":            0xe0ffff,
	"lightgoldenrodyellow": 0xfafad2,
	"lightgray":            0xd3d3d3,
	"lightgreen":           0x90ee90,
	"lightgrey":            0xd3d3d3,
	"lightpink":            0xffb6c1,
	"lightsalmon":          0xffa07a,
	"lightseagreen":        0x20b2aa,
	"lightskyblue":         0x87cefa,
	"lightslategray":       0x778899,
	"lightslategrey":       0x778899,
	"lightsteelblue":       0xb0c4de,
	"lightyellow":          0xffffe0,
	"lime":                 0x00ff00,
	"limegreen":            0x32cd32,
	"linen":                0xfaf0e6,
	"magenta":              0xff00ff,
	"maroon":               0x800000,
	"mediumaquamarine":     0x66cdaa,
	"mediumblue":           0x0000cd,
	"mediumorchid":         0xba55d3,
	"mediumpurple":         0x9370db,
	"mediumseagreen":       0x3cb371,
	"mediumslateblue":      0x7b68ee,
	"mediumspringgreen":    0x00fa9a,
	"mediumturquoise":      0x48d1cc,
	"mediumvioletred":      0xc71585,
	"midnightblue":         0x191970,
	"mintcream":            0xf5fffa,
	"mistyrose":            0xffe4e1,
	"moccasin":             0xffe4b5,
	"navajowhite":          0xffdead,
	"navy":                 0x000080,
	"oldlace":              0xfdf5e6,
	"olive":                0x808000,
	"olivedrab":            0x6b8e23,
	"orange":               0xffa500,
	"orangered":            0xff4500,
	"orchid":               0xda70d6,
	"palegoldenrod":        0xeee8aa,
	"palegreen":            0x98fb98,
	"paleturquoise":        0xafeeee,
	"palevioletred":        0xdb7093,
	"papayawhip":           0xffefd5,
	"peachpuff":            0xffdab9,
	"peru":                 0xcd853f,
	"pink":                 0xffc0cb,
	"plum":                 0xdda0dd,
	"powderblue":           0xb0e0e6,
	"purple":               0x800080,
	"rebeccapurple":        0x663399,
	"red":                  0xff0000,
	"rosybrown":            0xbc8f8f,
	"royalblue":            0x4169e1,
	"saddlebrown":          0x8b4513,
	"salmon":               0xfa8072,
	"sandybrown":           0xf4a460,
	"seagreen":             0x2e8b57,
	"seashell":             0xfff5ee,
	"sienna":               0xa0522d,
	"silver":               0xc0c0c0,
	"skyblue":              0x87ceeb,
	"slateblue":            0x6a5acd,
	"slategray":            0x708090,
	"slategrey":            0x708090,
	"snow":                 0xfffafa,
	"springgreen":          0x00ff7f,
	"steelblue":            0x4682b4,
	"tan":                  0xd2b48c,
	"teal":                 0x008080,
	"thistle":              0xd8bfd8,
	"tomato":               0xff6347,
	"turquoise":            0x40e0d0,
	"violet":               0xee82ee,
	"wheat":                0xf5deb3,
	"white":                0xffffff,
	"whitesmoke":           0xf5f5f5,
	"yellow":               0xffff00,
	"yellowgreen":          0x9acd32,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewHex creates new hex color from uint32 number
func NewHex(c uint32) Hex {
	return Hex{v: c}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsRGBA returns true if color contains info about alpha channel
func (c Hex) IsRGBA() bool {
	return c.v > 0xFFFFFF || c.a
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

// ToHSL converts RGBA color to HSL
func (c RGBA) ToHSL() HSL {
	return RGBA2HSL(c)
}

// ToHSV converts RGBA color to HSV
func (c RGBA) ToHSV() HSV {
	return RGBA2HSV(c)
}

// ToRGB converts CMYK color to RGB
func (c CMYK) ToRGB() RGB {
	return CMYK2RGB(c)
}

// ToRGB converts HSV color to RGB
func (c HSV) ToRGB() RGB {
	return HSV2RGB(c)
}

// ToRGBA converts HSV color to RGBA
func (c HSV) ToRGBA() RGBA {
	return HSV2RGBA(c)
}

// ToRGB converts HSL color to RGB
func (c HSL) ToRGB() RGB {
	return HSL2RGB(c)
}

// ToRGBA converts HSL color to RGBA
func (c HSL) ToRGBA() RGBA {
	return HSL2RGBA(c)
}

// ToRGB converts hex color to RGB
func (c Hex) ToRGB() RGB {
	return Hex2RGB(c)
}

// ToRGB converts hex color to RGBA
func (c Hex) ToRGBA() RGBA {
	return Hex2RGBA(c)
}

// WithAlpha returns color with given alpha value
func (c RGBA) WithAlpha(alpha float64) RGBA {
	c.A = uint8(255.0 * mathutil.Between(alpha, 0, 1))
	return c
}

// ToWeb converts hex color to notation used in web (#RGB/#RGBA/#RRGGBB/#RRGGBBAA)
func (c Hex) ToWeb(useCaps, allowShorthand bool) string {
	var k string

	switch {
	case c.IsRGBA() && useCaps:
		k = fmt.Sprintf("%08X", uint32(c.v))
	case c.IsRGBA() && !useCaps:
		k = fmt.Sprintf("%08x", uint32(c.v))
	case useCaps:
		k = fmt.Sprintf("%06X", uint32(c.v))
	default:
		k = fmt.Sprintf("%06x", uint32(c.v))
	}

	if allowShorthand {
		if !c.IsRGBA() {
			if k[0] == k[1] && k[2] == k[3] && k[4] == k[5] {
				k = k[0:1] + k[2:3] + k[4:5]
			}
		} else {
			if k[0] == k[1] && k[2] == k[3] && k[4] == k[5] && k[6] == k[7] {
				k = k[0:1] + k[2:3] + k[4:5] + k[6:7]
			}
		}
	}

	return "#" + k
}

// String returns string representation of hex color
func (c Hex) String() string {
	return c.ToWeb(true, false)
}

// String returns string representation of RGB color
func (c RGB) String() string {
	return fmt.Sprintf("%d,%d,%d", c.R, c.G, c.B)
}

// GoString returns Go representation of RGB color
func (c RGB) GoString() string {
	return fmt.Sprintf("RGB{R:%d, G:%d, B:%d}", c.R, c.G, c.B)
}

// String returns string representation of RGBA color
func (c RGBA) String() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", c.R, c.G, c.B, float64(c.A)/255.0)
}

// GoString returns Go representation of RGBA color
func (c RGBA) GoString() string {
	if c.A == 0 {
		return fmt.Sprintf("RGBA{R:%d, G:%d, B:%d}", c.R, c.G, c.B)
	}

	return fmt.Sprintf("RGBA{R:%d, G:%d, B:%d, A:%d}", c.R, c.G, c.B, c.A)
}

// String returns string representation of CMYK color
func (c CMYK) String() string {
	return fmt.Sprintf(
		"%.0f%%,%.0f%%,%.0f%%,%.0f%%",
		c.C*100.0, c.M*100.0, c.Y*100.0, c.K*100.0,
	)
}

// GoString returns Go representation of CMYK color
func (c CMYK) GoString() string {
	return fmt.Sprintf("CMYK{C:%g, M:%g, Y:%g, K:%g}", c.C, c.M, c.Y, c.K)
}

// String returns string representation of HSV color
func (c HSV) String() string {
	return fmt.Sprintf(
		"%.0f°,%.0f%%,%.0f%%,%.0f%%",
		c.H*360, c.S*100, c.V*100, c.A*100,
	)
}

// GoString returns Go representation of HSV color
func (c HSV) GoString() string {
	if c.A == 0 {
		return fmt.Sprintf("HSV{H:%g, S:%g, V:%g}", c.H, c.S, c.V)
	}

	return fmt.Sprintf("HSV{H:%g, S:%g, V:%g, A:%g}", c.H, c.S, c.V, c.A)
}

// String returns string representation of HSL color
func (c HSL) String() string {
	return fmt.Sprintf(
		"%.0f°,%.0f%%,%.0f%%,%.0f%%",
		c.H*360, c.S*100, c.L*100, c.A*100,
	)
}

// GoString returns Go representation of HSL color
func (c HSL) GoString() string {
	if c.A == 0 {
		return fmt.Sprintf("HSL{H:%g, S:%g, L:%g}", c.H, c.S, c.L)
	}

	return fmt.Sprintf("HSL{H:%g, S:%g, L:%g, A:%g}", c.H, c.S, c.L, c.A)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Parse parses color
func Parse(c string) (Hex, error) {
	if colors[c] != 0 {
		return Hex{v: colors[c]}, nil
	}

	if c != "" && c[0] == '#' {
		c = c[1:]
	}

	var hasAlpha bool

	switch len(c) {
	case 0:
		return Hex{}, fmt.Errorf("Color is empty")

	// Shorthand #RGB
	case 3:
		c = c[0:1] + c[0:1] + c[1:2] + c[1:2] + c[2:3] + c[2:3]

	// Shorthand #RGBA
	case 4:
		hasAlpha = true
		c = c[0:1] + c[0:1] + c[1:2] + c[1:2] + c[2:3] + c[2:3] + c[3:4] + c[3:4]
	}

	k, err := strconv.ParseUint(c, 16, 32)

	return Hex{v: uint32(k), a: k > 0xFFFFFF || hasAlpha}, err
}

// RGB2Hex converts RGB color to Hex
func RGB2Hex(c RGB) Hex {
	return Hex{v: (uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B))}
}

// Hex2RGB converts Hex color to RGB
func Hex2RGB(h Hex) RGB {
	return RGB{uint8(h.v >> 16 & 0xFF), uint8(h.v >> 8 & 0xFF), uint8(h.v & 0xFF)}
}

// RGBA2Hex converts RGBA color to Hex
func RGBA2Hex(c RGBA) Hex {
	return Hex{v: (uint32(c.R)<<24 | uint32(c.G)<<16 | uint32(c.B)<<8 | uint32(c.A)), a: true}
}

// Hex2RGBA converts Hex color to RGBA
func Hex2RGBA(h Hex) RGBA {
	if h.IsRGBA() {
		return RGBA{uint8(h.v>>24) & 0xFF, uint8(h.v>>16) & 0xFF, uint8(h.v>>8) & 0xFF, uint8(h.v) & 0xFF}
	}

	return RGBA{uint8(h.v>>16) & 0xFF, uint8(h.v>>8) & 0xFF, uint8(h.v) & 0xFF, 0}
}

// RGB2Term converts RGB color to terminal color code
// https://github.com/essentialkaos/ek/tree/master/fmtc#88256-colors
func RGB2Term(c RGB) int {
	R, G, B := int(c.R), int(c.G), int(c.B)

	// Grayscale
	if R == G && G == B {
		if R == 175 {
			return 145
		}

		return (R / 10) + 232
	}

	return 36*(R/51) + 6*(G/51) + (B / 51) + 16
}

// Term2RGB converts terminal color code (0-255) to RGB color
// https://github.com/essentialkaos/ek/tree/master/fmtc#88256-colors
func Term2RGB(c uint8) RGB {
	// Grayscale
	if c > 231 {
		c = ((c - 232) * 10) + 8
		return RGB{c, c, c}
	}

	// Default colors (not standardized, average values)
	switch c {
	case 0:
		return RGB{}
	case 1:
		return RGB{255, 0, 0}
	case 2:
		return RGB{0, 255, 0}
	case 3:
		return RGB{255, 255, 0}
	case 4:
		return RGB{0, 0, 255}
	case 5:
		return RGB{255, 0, 255}
	case 6:
		return RGB{0, 255, 255}
	case 7:
		return RGB{191, 191, 191}
	case 8:
		return RGB{64, 64, 64}
	case 9:
		return RGB{255, 127, 127}
	case 10:
		return RGB{127, 255, 127}
	case 11:
		return RGB{255, 255, 127}
	case 12:
		return RGB{127, 127, 255}
	case 13:
		return RGB{255, 127, 255}
	case 14:
		return RGB{127, 255, 255}
	case 15:
		return RGB{127, 127, 127}
	}

	var r, g, b, ir, ig, ib uint8

	ir = (c - 16) / 36
	ig = ((c - 16) % 36) / 6
	ib = (c - 16) % 6

	if ir > 0 {
		r = (ir * 40) + 55
	}

	if ig > 0 {
		g = (ig * 40) + 55
	}

	if ib > 0 {
		b = (ib * 40) + 55
	}

	return RGB{r, g, b}
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
	C := mathutil.Between(c.C, 0.0, 1.0)
	M := mathutil.Between(c.M, 0.0, 1.0)
	Y := mathutil.Between(c.Y, 0.0, 1.0)
	K := mathutil.Between(c.K, 0.0, 1.0)

	return RGB{
		uint8(255 * (1 - C) * (1 - K)),
		uint8(255 * (1 - M) * (1 - K)),
		uint8(255 * (1 - Y) * (1 - K)),
	}
}

// RGB2HSV converts RGB color to HSV (HSB)
func RGB2HSV(c RGB) HSV {
	h, s, v := rgbToHsv(c.R, c.G, c.B)
	return HSV{h, s, v, 0.0}
}

// RGBA2HSV converts RGBA color to HSV (HSB)
func RGBA2HSV(c RGBA) HSV {
	h, s, v := rgbToHsv(c.R, c.G, c.B)
	return HSV{h, s, v, float64(c.A) / 255.0}
}

// HSV2RGB converts HSV (HSB) color to RGB
func HSV2RGB(c HSV) RGB {
	r, g, b := hsvToRgb(c.H, c.S, c.V)
	return RGB{r, g, b}
}

// HSV2RGBA converts HSV (HSB) color to RGBA
func HSV2RGBA(c HSV) RGBA {
	r, g, b := hsvToRgb(c.H, c.S, c.V)
	return RGBA{r, g, b, uint8(255.0 * mathutil.Between(c.A, 0, 1))}
}

// RGB2HSL converts RGB color to HSL
func RGB2HSL(c RGB) HSL {
	h, s, l := rgbToHsl(c.R, c.G, c.B)
	return HSL{h, s, l, 0.0}
}

// RGBA2HSL converts RGBA color to HSL
func RGBA2HSL(c RGBA) HSL {
	h, s, l := rgbToHsl(c.R, c.G, c.B)
	return HSL{h, s, l, float64(c.A) / 255.0}
}

// HSL2RGB converts HSL color to RGB
func HSL2RGB(c HSL) RGB {
	r, g, b := hsl2Rgb(c.H, c.S, c.L)
	return RGB{r, g, b}
}

// HSL2RGBA converts HSL color to RGBA
func HSL2RGBA(c HSL) RGBA {
	r, g, b := hsl2Rgb(c.H, c.S, c.L)
	return RGBA{r, g, b, uint8(255.0 * mathutil.Between(c.A, 0, 1))}
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

func rgbToHsv(r, g, b uint8) (float64, float64, float64) {
	R, G, B := float64(r)/255.0, float64(g)/255.0, float64(b)/255.0

	max := math.Max(math.Max(R, G), B)
	min := math.Min(math.Min(R, G), B)

	h, s, v := 0.0, 0.0, max

	if max != min {
		d := max - min
		s = d / max
		h = calcHUE(max, R, G, B, d)
	}

	return h, s, v
}

func hsvToRgb(h, s, v float64) (uint8, uint8, uint8) {
	i := (h * 360.0) / 60.0
	f := i - math.Floor(i)

	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	var r, g, b float64

	switch int(h*6) % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return uint8(r * 0xFF), uint8(g * 0xFF), uint8(b * 0xFF)
}

func rgbToHsl(r, g, b uint8) (float64, float64, float64) {
	R, G, B := float64(r)/255.0, float64(g)/255.0, float64(b)/255.0

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

	return h, s, l
}

func hsl2Rgb(h, s, l float64) (uint8, uint8, uint8) {
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

	return uint8(R * 255), uint8(G * 255), uint8(B * 255)
}

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
