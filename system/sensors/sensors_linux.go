// Package sensors provide methods for collecting sensors information
package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v11/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SYS_DIR = "/sys/class/hwmon"

// ////////////////////////////////////////////////////////////////////////////////// //

// Device contains info from different device sensors
type Device struct {
	Name         string
	TempSensors  []*TempSensor
	PowerSensors []*PowerSensor
}

// TempSensor contains info from temperature sensor
type TempSensor struct {
	Name string
	Cur  float64
	Min  float64
	Mid  float64
	Max  float64
	Crit float64
}

// PowerSensor contains info from power sensor
type PowerSensor struct {
	Name string
	Cur  float64
	Min  float64
	Max  float64
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects sensors information
func Collect() ([]*Device, error) {
	if !fsutil.CheckPerms("DR", _SYS_DIR) {
		return nil, fmt.Errorf("Can't read sensors information")
	}

	var result []*Device
	var coreTempCount int

	for _, deviceDir := range fsutil.List(_SYS_DIR, false) {
		device, err := collectDeviceInfo(_SYS_DIR + "/" + deviceDir)

		if err != nil {
			return nil, err
		}

		if device != nil {
			if device.Name == "coretemp" {
				device.Name += strconv.Itoa(coreTempCount)
				coreTempCount++
			}

			result = append(result, device)
		}
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Temperature returns min, max and average temperature
func (d *Device) Temperature() (float64, float64, float64) {
	if len(d.TempSensors) == 0 {
		return 0.0, 0.0, 0.0
	}

	min := 100.0
	max := 0.0
	tot := 0.0

	for _, v := range d.TempSensors {
		if v.Cur < min {
			min = v.Cur
		}

		if v.Cur > max {
			max = v.Cur
		}

		tot += v.Cur
	}

	return min, max, tot / float64(len(d.TempSensors))
}

// String formats sensor data as a string
func (s *TempSensor) String() string {
	return fmt.Sprintf(
		"[Name:%s Cur:%g Min:%g Mid:%g Max:%g Crit:%g]",
		s.Name, s.Cur, s.Min, s.Mid, s.Max, s.Crit,
	)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func hasSensorsData(dir string) bool {
	switch {
	case fsutil.IsExist(dir + "/temp1_input"),
		fsutil.IsExist(dir + "/temp2_input"):
		return true
	}

	return false
}

func collectDeviceInfo(dir string) (*Device, error) {
	var err error
	var dataDir string

	switch {
	case hasSensorsData(dir):
		dataDir = dir
	case hasSensorsData(dir + "/device"):
		dataDir = dir + "/device"
	default:
		return nil, nil
	}

	device := &Device{}

	device.Name, err = readSensorLabel(dataDir + "/name")

	if err != nil {
		return nil, err
	}

	device.TempSensors, err = collectTempInfo(dataDir)

	if err != nil {
		return nil, err
	}

	return device, nil
}

func collectTempInfo(dir string) ([]*TempSensor, error) {
	var sensors []*TempSensor

	for i := 1; i < 128; i++ {
		filePrefix := dir + "/temp" + strconv.Itoa(i)

		if !fsutil.IsExist(filePrefix + "_input") {
			break
		}

		info, err := getSensorInfo(filePrefix)

		if err != nil {
			return nil, err
		}

		sensors = append(sensors, info)
	}

	return sensors, nil
}

func getSensorInfo(filePrefix string) (*TempSensor, error) {
	var err error

	sensor := &TempSensor{}
	sensor.Name, err = readSensorLabel(filePrefix + "_label")

	sensor.Cur, err = readTempSensorValue(filePrefix + "_input")

	if err != nil {
		return nil, err
	}

	sensor.Min, err = readTempSensorValue(filePrefix + "_min")

	if err != nil {
		return nil, err
	}

	sensor.Mid, err = readTempSensorValue(filePrefix + "_mid")

	if err != nil {
		return nil, err
	}

	sensor.Max, err = readTempSensorValue(filePrefix + "_max")

	if err != nil {
		return nil, err
	}

	sensor.Crit, err = readTempSensorValue(filePrefix + "_crit")

	if err != nil {
		return nil, err
	}

	return sensor, nil
}

func readSensorLabel(file string) (string, error) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return "", fmt.Errorf("Can't read data from %s: %v", file, err)
	}

	return strings.Trim(string(data), "\r\n"), nil
}

func readTempSensorValue(file string) (float64, error) {
	if !fsutil.IsExist(file) {
		return 0.0, nil
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return 0.0, fmt.Errorf("Can't read sensor data from %s: %v", file, err)
	}

	value, err := strconv.ParseFloat(strings.Trim(string(data), "\r\n"), 64)

	if err != nil {
		return 0.0, fmt.Errorf("Can't parse sensor data from %s: %v", file, err)
	}

	return value / 1000.0, nil
}
