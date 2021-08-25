package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleLinear() {
	fmt.Println(Linear(5, 0, 10, 10))
	// Output: 5
}

func ExampleQuadIn() {
	fmt.Println(QuadIn(5, 0, 10, 10))
	// Output: 2.5
}

func ExampleQuadOut() {
	fmt.Println(QuadOut(5, 0, 10, 10))
	// Output: 7.5
}

func ExampleQuadInOut() {
	fmt.Println(QuadInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleCubicIn() {
	fmt.Println(CubicIn(5, 0, 10, 10))
	// Output: 1.25
}

func ExampleCubicOut() {
	fmt.Println(CubicOut(5, 0, 10, 10))
	// Output: 8.75
}

func ExampleCubicInOut() {
	fmt.Println(CubicInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleQuintIn() {
	fmt.Println(QuintIn(5, 0, 10, 10))
	// Output: 0.3125
}

func ExampleQuintOut() {
	fmt.Println(QuintOut(5, 0, 10, 10))
	// Output: 9.6875
}

func ExampleQuintInOut() {
	fmt.Println(QuintInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleSineIn() {
	fmt.Println(SineIn(5, 0, 10, 10))
	// Output: 3.0978992192751917
}

func ExampleSineOut() {
	fmt.Println(SineOut(5, 0, 10, 10))
	// Output: 7.236090437019012
}

func ExampleSineInOut() {
	fmt.Println(SineInOut(5, 0, 10, 10))
	// Output: 4.999999999999999
}

func ExampleExpoIn() {
	fmt.Println(ExpoIn(5, 0, 10, 10))
	// Output: 0.3125
}

func ExampleExpoOut() {
	fmt.Println(ExpoOut(5, 0, 10, 10))
	// Output: 9.6875
}

func ExampleExpoInOut() {
	fmt.Println(ExpoInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleCircIn() {
	fmt.Println(CircIn(5, 0, 10, 10))
	// Output: 1.339745962155614
}

func ExampleCircOut() {
	fmt.Println(CircOut(5, 0, 10, 10))
	// Output: 8.660254037844386
}

func ExampleCircInOut() {
	fmt.Println(CircInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleElasticIn() {
	fmt.Println(ElasticIn(5, 0, 10, 10))
	// Output: -0.15625000000000044
}

func ExampleElasticOut() {
	fmt.Println(ElasticOut(5, 0, 10, 10))
	// Output: 10.15625
}

func ExampleElasticInOut() {
	fmt.Println(ElasticInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleBackIn() {
	fmt.Println(BackIn(5, 0, 10, 10))
	// Output: -0.9655673136318949
}

func ExampleBackOut() {
	fmt.Println(BackOut(5, 0, 10, 10))
	// Output: 10.965567313631894
}

func ExampleBackInOut() {
	fmt.Println(BackInOut(5, 0, 10, 10))
	// Output: 5
}

func ExampleBounceIn() {
	fmt.Println(BounceIn(5, 0, 10, 10))
	// Output: 2.34375
}

func ExampleBounceOut() {
	fmt.Println(BounceOut(5, 0, 10, 10))
	// Output: 7.65625
}

func ExampleBounceInOut() {
	fmt.Println(BounceInOut(5, 0, 10, 10))
	// Output: 5
}
