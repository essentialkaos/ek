package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// BounceIn accelerating from zero velocity
func BounceIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c - BounceOut(d-t, 0.0, c, d) + b
}

// BounceOut Decelerating to zero velocity
func BounceOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d

	switch {
	case t < 1/2.75:
		return c*(7.5625*t*t) + b
	case t < 2/2.75:
		t -= 1.5 / 2.75
		return c*(7.5625*t*t+0.75) + b
	case t < 2.5/2.75:
		t -= 2.25 / 2.75
		return c*(7.5625*t*t+0.9375) + b
	}

	t -= 2.625 / 2.75
	return c*(7.5625*t*t+0.984375) + b
}

// BounceInOut acceleration until halfway, then deceleration
func BounceInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	if t < d/2 {
		return BounceIn(t*2, 0.0, c, d)*0.5 + b
	}

	return BounceOut(t*2-d, 0.0, c, d)*0.5 + c*0.5 + b
}
