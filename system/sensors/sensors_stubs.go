//go:build !linux

// Package sensors provide methods for collecting sensors information
package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Collect reads sensor data from all hwmon devices and returns the collected results.
// See https://www.kernel.org/doc/Documentation/hwmon/sysfs-interface for the sysfs
// layout.
func Collect() ([]*Device, error) {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Temperature returns the minimum, maximum, and average current temperature across
// all temperature sensors of the device
func (d *Device) Temperature() (float64, float64, float64) {
	panic("UNSUPPORTED")
}

// ❗ String returns a human-readable representation of the sensor readings
func (s TempSensor) String() string {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //
