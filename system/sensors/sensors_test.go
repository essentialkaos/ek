//go:build linux
// +build linux

package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SensorsSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SensorsSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SensorsSuite) TestParsing(c *C) {
	tmpDir := c.MkDir()
	createFakeFS(tmpDir)

	hwmonDir = tmpDir

	devices, err := Collect()

	c.Assert(err, IsNil)
	c.Assert(devices, NotNil)
	c.Assert(devices, HasLen, 2)
	c.Assert(devices[0].Name, Equals, "acpitz")
	c.Assert(devices[0].TempSensors[0].Name, Equals, "")
	c.Assert(devices[0].TempSensors[0].Cur, Equals, 8.2)
	c.Assert(devices[0].TempSensors[0].Min, Equals, 0.0)
	c.Assert(devices[0].TempSensors[0].Max, Equals, 0.0)
	c.Assert(devices[0].TempSensors[0].Crit, Equals, 25.6)
	c.Assert(devices[1].Name, Equals, "coretemp")
	c.Assert(devices[1].TempSensors[0].Name, Equals, "Core 0")
	c.Assert(devices[1].TempSensors[0].Cur, Equals, 38.0)
	c.Assert(devices[1].TempSensors[0].Min, Equals, 12.0)
	c.Assert(devices[1].TempSensors[0].Max, Equals, 81.0)
	c.Assert(devices[1].TempSensors[0].Crit, Equals, 91.0)

	t1, t2, t3 := devices[1].Temperature()
	c.Assert(t1, Equals, 38.0)
	c.Assert(t2, Equals, 41.0)
	c.Assert(t3, Equals, 39.5)

	c.Assert(devices[1].TempSensors[0].String(), Equals, "{Name:Core 0 Cur:+38°C Min:+12°C Max:+81°C Crit:+91°C}")
}

func (s *SensorsSuite) TestParsingErrors(c *C) {
	hwmonDir = "/_unknown_"

	_, err := Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't read sensors information`)

	d := &Device{}
	t1, t2, t3 := d.Temperature()

	c.Assert(t1, Equals, 0.0)
	c.Assert(t2, Equals, 0.0)
	c.Assert(t3, Equals, 0.0)

	tmpDir := c.MkDir()
	createFakeFS(tmpDir)

	hwmonDir = tmpDir

	os.WriteFile(tmpDir+"/hwmon1/temp1_crit", []byte("AAA"), 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse sensor data from .*/hwmon1/temp1_crit: strconv.ParseFloat: parsing "AAA": invalid syntax`)

	os.WriteFile(tmpDir+"/hwmon1/temp1_max", []byte("AAA"), 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse sensor data from .*/hwmon1/temp1_max: strconv.ParseFloat: parsing "AAA": invalid syntax`)

	os.WriteFile(tmpDir+"/hwmon1/temp1_min", []byte("AAA"), 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse sensor data from .*/hwmon1/temp1_min: strconv.ParseFloat: parsing "AAA": invalid syntax`)

	os.WriteFile(tmpDir+"/hwmon1/temp1_input", []byte("AAA"), 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse sensor data from .*/hwmon1/temp1_input: strconv.ParseFloat: parsing "AAA": invalid syntax`)

	os.Remove(tmpDir + "/hwmon1/temp1_input")
	os.Mkdir(tmpDir+"/hwmon1/temp1_input", 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't read sensor data from .*/hwmon1/temp1_input: read .*/hwmon1/temp1_input: is a directory`)

	os.Remove(tmpDir + "/hwmon1/temp1_label")
	os.Mkdir(tmpDir+"/hwmon1/temp1_label", 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't read data from .*/hwmon1/temp1_label: read .*/hwmon1/temp1_label: is a directory`)

	os.Remove(tmpDir + "/hwmon1/name")
	os.Mkdir(tmpDir+"/hwmon1/name", 0644)
	_, err = Collect()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't read data from .*/hwmon1/name: read .*/hwmon1/name: is a directory`)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createFakeFS(tmpDir string) {
	os.Mkdir(tmpDir+"/hwmon0", 0755)
	os.Mkdir(tmpDir+"/hwmon1", 0755)
	os.Mkdir(tmpDir+"/hwmon2", 0755)

	os.WriteFile(tmpDir+"/hwmon0/name", []byte("acpitz"), 0644)
	os.WriteFile(tmpDir+"/hwmon0/temp1_crit", []byte("25600"), 0644)
	os.WriteFile(tmpDir+"/hwmon0/temp1_input", []byte("8200"), 0644)

	os.WriteFile(tmpDir+"/hwmon1/name", []byte("coretemp"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp1_label", []byte("Core 0"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp1_crit", []byte("91000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp1_input", []byte("38000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp1_max", []byte("81000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp1_min", []byte("12000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp2_label", []byte("Core 0"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp2_crit", []byte("91000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp2_input", []byte("41000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp2_max", []byte("81000"), 0644)
	os.WriteFile(tmpDir+"/hwmon1/temp2_min", []byte("12000"), 0644)
}
