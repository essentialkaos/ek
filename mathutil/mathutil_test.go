package mathutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type MathUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&MathUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *MathUtilSuite) TestBetween(c *C) {
	c.Assert(Between(5, 0, 100), Equals, 5)
	c.Assert(Between(10, 0, 5), Equals, 5)
	c.Assert(Between(-5, -10, 0), Equals, -5)
	c.Assert(Between(5, 10, 10), Equals, 10)

	c.Assert(Between8(5, 0, 100), Equals, int8(5))
	c.Assert(Between8(10, 0, 5), Equals, int8(5))
	c.Assert(Between8(-5, -10, 0), Equals, int8(-5))
	c.Assert(Between8(5, 10, 10), Equals, int8(10))

	c.Assert(Between16(5, 0, 100), Equals, int16(5))
	c.Assert(Between16(10, 0, 5), Equals, int16(5))
	c.Assert(Between16(-5, -10, 0), Equals, int16(-5))
	c.Assert(Between16(5, 10, 10), Equals, int16(10))

	c.Assert(Between32(5, 0, 100), Equals, int32(5))
	c.Assert(Between32(10, 0, 5), Equals, int32(5))
	c.Assert(Between32(-5, -10, 0), Equals, int32(-5))
	c.Assert(Between32(5, 10, 10), Equals, int32(10))

	c.Assert(Between64(5, 0, 100), Equals, int64(5))
	c.Assert(Between64(10, 0, 5), Equals, int64(5))
	c.Assert(Between64(-5, -10, 0), Equals, int64(-5))
	c.Assert(Between64(5, 10, 10), Equals, int64(10))

	c.Assert(BetweenU(5, 0, 100), Equals, uint(5))
	c.Assert(BetweenU(10, 0, 5), Equals, uint(5))
	c.Assert(BetweenU(5, 10, 10), Equals, uint(10))

	c.Assert(BetweenU8(5, 0, 100), Equals, uint8(5))
	c.Assert(BetweenU8(10, 0, 5), Equals, uint8(5))
	c.Assert(BetweenU8(5, 10, 10), Equals, uint8(10))

	c.Assert(BetweenU16(5, 0, 100), Equals, uint16(5))
	c.Assert(BetweenU16(10, 0, 5), Equals, uint16(5))
	c.Assert(BetweenU16(5, 10, 10), Equals, uint16(10))

	c.Assert(BetweenU32(5, 0, 100), Equals, uint32(5))
	c.Assert(BetweenU32(10, 0, 5), Equals, uint32(5))
	c.Assert(BetweenU32(5, 10, 10), Equals, uint32(10))

	c.Assert(BetweenU64(5, 0, 100), Equals, uint64(5))
	c.Assert(BetweenU64(10, 0, 5), Equals, uint64(5))
	c.Assert(BetweenU64(5, 10, 10), Equals, uint64(10))

	c.Assert(BetweenF(0.5, 0.1, 0.7), Equals, 0.5)
	c.Assert(BetweenF(0.01, 0.1, 0.7), Equals, 0.1)
	c.Assert(BetweenF(5.0, 0.1, 0.7), Equals, 0.7)

	c.Assert(BetweenF32(0.5, 0.1, 0.7), Equals, float32(0.5))
	c.Assert(BetweenF32(0.01, 0.1, 0.7), Equals, float32(0.1))
	c.Assert(BetweenF32(5.0, 0.1, 0.7), Equals, float32(0.7))

	c.Assert(BetweenF64(0.5, 0.1, 0.7), Equals, float64(0.5))
	c.Assert(BetweenF64(0.01, 0.1, 0.7), Equals, float64(0.1))
	c.Assert(BetweenF64(5.0, 0.1, 0.7), Equals, float64(0.7))
}

func (s *MathUtilSuite) TestMinMax(c *C) {
	c.Assert(Min(1, 10), Equals, 1)
	c.Assert(Min(-10, 10), Equals, -10)
	c.Assert(Min(10, -10), Equals, -10)
	c.Assert(Max(1, 10), Equals, 10)
	c.Assert(Max(-10, 10), Equals, 10)
	c.Assert(Max(10, -10), Equals, 10)

	c.Assert(Min8(1, 10), Equals, int8(1))
	c.Assert(Min8(-10, 10), Equals, int8(-10))
	c.Assert(Min8(10, -10), Equals, int8(-10))
	c.Assert(Max8(1, 10), Equals, int8(10))
	c.Assert(Max8(-10, 10), Equals, int8(10))
	c.Assert(Max8(10, -10), Equals, int8(10))

	c.Assert(Min16(1, 10), Equals, int16(1))
	c.Assert(Min16(-10, 10), Equals, int16(-10))
	c.Assert(Min16(10, -10), Equals, int16(-10))
	c.Assert(Max16(1, 10), Equals, int16(10))
	c.Assert(Max16(-10, 10), Equals, int16(10))
	c.Assert(Max16(10, -10), Equals, int16(10))

	c.Assert(Min32(1, 10), Equals, int32(1))
	c.Assert(Min32(-10, 10), Equals, int32(-10))
	c.Assert(Min32(10, -10), Equals, int32(-10))
	c.Assert(Max32(1, 10), Equals, int32(10))
	c.Assert(Max32(-10, 10), Equals, int32(10))
	c.Assert(Max32(10, -10), Equals, int32(10))

	c.Assert(Min64(1, 10), Equals, int64(1))
	c.Assert(Min64(-10, 10), Equals, int64(-10))
	c.Assert(Min64(10, -10), Equals, int64(-10))
	c.Assert(Max64(1, 10), Equals, int64(10))
	c.Assert(Max64(-10, 10), Equals, int64(10))
	c.Assert(Max64(10, -10), Equals, int64(10))

	c.Assert(MinU(0, 10), Equals, uint(0))
	c.Assert(MinU(10, 0), Equals, uint(0))
	c.Assert(MaxU(0, 10), Equals, uint(10))
	c.Assert(MaxU(10, 0), Equals, uint(10))

	c.Assert(MinU8(0, 10), Equals, uint8(0))
	c.Assert(MinU8(10, 0), Equals, uint8(0))
	c.Assert(MaxU8(0, 10), Equals, uint8(10))
	c.Assert(MaxU8(10, 0), Equals, uint8(10))

	c.Assert(MinU16(0, 10), Equals, uint16(0))
	c.Assert(MinU16(10, 0), Equals, uint16(0))
	c.Assert(MaxU16(0, 10), Equals, uint16(10))
	c.Assert(MaxU16(10, 0), Equals, uint16(10))

	c.Assert(MinU32(0, 10), Equals, uint32(0))
	c.Assert(MinU32(10, 0), Equals, uint32(0))
	c.Assert(MaxU32(0, 10), Equals, uint32(10))
	c.Assert(MaxU32(10, 0), Equals, uint32(10))

	c.Assert(MinU64(0, 10), Equals, uint64(0))
	c.Assert(MinU64(10, 0), Equals, uint64(0))
	c.Assert(MaxU64(0, 10), Equals, uint64(10))
	c.Assert(MaxU64(10, 0), Equals, uint64(10))
}

func (s *MathUtilSuite) TestAbs(c *C) {
	c.Assert(Abs(-10), Equals, 10)
	c.Assert(Abs(10), Equals, 10)

	c.Assert(Abs8(-10), Equals, int8(10))
	c.Assert(Abs8(10), Equals, int8(10))

	c.Assert(Abs16(-10), Equals, int16(10))
	c.Assert(Abs16(10), Equals, int16(10))

	c.Assert(Abs32(-10), Equals, int32(10))
	c.Assert(Abs32(10), Equals, int32(10))

	c.Assert(Abs64(-10), Equals, int64(10))
	c.Assert(Abs64(10), Equals, int64(10))

	c.Assert(AbsF(-10), Equals, 10.0)
	c.Assert(AbsF(10), Equals, 10.0)

	c.Assert(AbsF32(-10), Equals, float32(10.0))
	c.Assert(AbsF32(10), Equals, float32(10.0))

	c.Assert(AbsF64(-10), Equals, float64(10.0))
	c.Assert(AbsF64(10), Equals, float64(10.0))
}

func (s *MathUtilSuite) TestRound(c *C) {
	c.Assert(Round(5.49, 0), Equals, 5.0)
	c.Assert(Round(5.50, 0), Equals, 6.0)
	c.Assert(Round(5.51, 0), Equals, 6.0)
}
