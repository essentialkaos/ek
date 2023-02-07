//go:build linux
// +build linux

// Package sensors provide methods for collecting sensors information
package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/sortutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Device contains info from different device sensors
type Device struct {
	Name        string
	TempSensors []TempSensor
}

// TempSensor contains info from temperature sensor
type TempSensor struct {
	Name string
	Cur  float64
	Min  float64
	Max  float64
	Crit float64
}

// ////////////////////////////////////////////////////////////////////////////////// //

var hwmonDir = "/sys/class/hwmon"

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects sensors information
// https://www.kernel.org/doc/Documentation/hwmon/sysfs-interface
func Collect() ([]*Device, error) {
	if !fsutil.CheckPerms("DR", hwmonDir) {
		return nil, fmt.Errorf("Can't read sensors information")
	}

	var result []*Device

	deviceDir := fsutil.List(hwmonDir, false)

	sortutil.StringsNatural(deviceDir)

	for _, deviceDir := range deviceDir {
		if !hasSensorsData(hwmonDir + "/" + deviceDir) {
			continue
		}

		device, err := collectDeviceInfo(hwmonDir + "/" + deviceDir)

		if err != nil {
			return nil, err
		}

		if device != nil {
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

	min, max, tot := 100.0, 0.0, 0.0

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
func (s TempSensor) String() string {
	return fmt.Sprintf(
		"{Name:%s Cur:+%g°C Min:+%g°C Max:+%g°C Crit:+%g°C}",
		s.Name, s.Cur, s.Min, s.Max, s.Crit,
	)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func hasSensorsData(dir string) bool {
	switch {
	case hasTempSensorsData(dir):
		return true
	}

	return false
}

func hasTempSensorsData(dir string) bool {
	return fsutil.IsExist(dir + "/temp1_input")
}

func collectDeviceInfo(dir string) (*Device, error) {
	var err error

	device := &Device{}

	device.Name, err = readSensorLabel(dir + "/name")

	if err != nil {
		return nil, err
	}

	if hasTempSensorsData(dir) {
		device.TempSensors, err = collectTempSensorsInfo(dir)

		if err != nil {
			return nil, err
		}
	}

	return device, nil
}

func collectTempSensorsInfo(dir string) ([]TempSensor, error) {
	var result []TempSensor

	for i := 1; i < 128; i++ {
		filePrefix := dir + "/temp" + strconv.Itoa(i)

		if !fsutil.IsExist(filePrefix + "_input") {
			break
		}

		sensor, err := readTempSensors(filePrefix)

		if err != nil {
			return nil, err
		}

		result = append(result, sensor)
	}

	return result, nil
}

func readTempSensors(filePrefix string) (TempSensor, error) {
	var err error

	sensor := TempSensor{}

	if fsutil.IsExist(filePrefix + "_label") {
		sensor.Name, err = readSensorLabel(filePrefix + "_label")

		if err != nil {
			return TempSensor{}, err
		}
	}

	sensor.Cur, err = readTempSensorValue(filePrefix + "_input")

	if err != nil {
		return TempSensor{}, err
	}

	sensor.Min, err = readTempSensorValue(filePrefix + "_min")

	if err != nil {
		return TempSensor{}, err
	}

	sensor.Max, err = readTempSensorValue(filePrefix + "_max")

	if err != nil {
		return TempSensor{}, err
	}

	sensor.Crit, err = readTempSensorValue(filePrefix + "_crit")

	if err != nil {
		return TempSensor{}, err
	}

	return sensor, nil
}

func readSensorLabel(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", fmt.Errorf("Can't read data from %s: %w", file, err)
	}

	return strings.Trim(string(data), "\r\n"), nil
}

func readTempSensorValue(file string) (float64, error) {
	if !fsutil.IsExist(file) {
		return 0.0, nil
	}

	data, err := os.ReadFile(file)

	if err != nil {
		return 0.0, fmt.Errorf("Can't read sensor data from %s: %w", file, err)
	}

	value, err := strconv.ParseFloat(strings.Trim(string(data), "\r\n"), 64)

	if err != nil {
		return 0.0, fmt.Errorf("Can't parse sensor data from %s: %w", file, err)
	}

	return value / 1000.0, nil
}
