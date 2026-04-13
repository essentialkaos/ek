//go:build linux

// Package sensors provide methods for collecting sensors information
package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v14/fsutil"
	"github.com/essentialkaos/ek/v14/sortutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MAX_TEMP_SENSORS is the upper bound for hwmon temperature sensor indices.
// The Linux hwmon ABI does not define a hard limit; 128 is a practical ceiling.
const MAX_TEMP_SENSORS = 128

// ////////////////////////////////////////////////////////////////////////////////// //

// hwmonDir is a path to hwmon directory
var hwmonDir = "/sys/class/hwmon"

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect reads sensor data from all hwmon devices and returns the collected results.
// See https://www.kernel.org/doc/Documentation/hwmon/sysfs-interface for the sysfs
// layout.
func Collect() ([]*Device, error) {
	err := fsutil.ValidatePerms("DR", hwmonDir)

	if err != nil {
		return nil, fmt.Errorf("can't read sensors information: %w", err)
	}

	var result []*Device

	deviceDir := fsutil.List(hwmonDir, false)

	sortutil.StringsNatural(deviceDir)

	for _, deviceDir := range deviceDir {
		sensorsDir := path.Join(hwmonDir, deviceDir)

		if !hasTempSensorsData(sensorsDir) {
			continue
		}

		device, err := collectDeviceInfo(sensorsDir)

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

// Temperature returns the minimum, maximum, and average current temperature across
// all temperature sensors of the device
func (d *Device) Temperature() (float64, float64, float64) {
	if len(d.TempSensors) == 0 {
		return 0.0, 0.0, 0.0
	}

	min, max, tot := math.MaxFloat64, -math.MaxFloat64, 0.0

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

// String returns a human-readable representation of the sensor readings
func (s TempSensor) String() string {
	return fmt.Sprintf(
		"{Name:%s Cur:+%g°C Min:+%g°C Max:+%g°C Crit:+%g°C}",
		s.Name, s.Cur, s.Min, s.Max, s.Crit,
	)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// hasTempSensorsData checks if directory contains temperature sensors data
func hasTempSensorsData(dir string) bool {
	return fsutil.IsExist(path.Join(dir, "temp1_input"))
}

// collectDeviceInfo collects information about device sensors
func collectDeviceInfo(dir string) (*Device, error) {
	var err error

	device := &Device{}

	device.Name, err = readSensorLabel(path.Join(dir, "name"))

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

// collectTempSensorsInfo collects information about temperature sensors
func collectTempSensorsInfo(dir string) ([]TempSensor, error) {
	var result []TempSensor

	for i := 1; i < MAX_TEMP_SENSORS; i++ {
		filePrefix := path.Join(dir, "temp"+strconv.Itoa(i))

		if !fsutil.IsExist(filePrefix + "_input") {
			continue
		}

		sensor, err := readTempSensors(filePrefix)

		if err != nil {
			return nil, err
		}

		result = append(result, sensor)
	}

	return result, nil
}

// readTempSensors reads temperature sensors data from files
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

// readSensorLabel reads sensor label from file
func readSensorLabel(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", fmt.Errorf("can't read data from %s: %w", file, err)
	}

	return strings.Trim(string(data), "\r\n"), nil
}

// readTempSensorValue reads temperature sensor value from file
func readTempSensorValue(file string) (float64, error) {
	if !fsutil.IsExist(file) {
		return 0.0, nil
	}

	data, err := os.ReadFile(file)

	if err != nil {
		return 0.0, fmt.Errorf("can't read sensor data from %s: %w", file, err)
	}

	value, err := strconv.ParseFloat(strings.Trim(string(data), "\r\n"), 64)

	if err != nil {
		return 0.0, fmt.Errorf("can't parse sensor data from %s: %w", file, err)
	}

	return value / 1000.0, nil
}
