// Package sensors provide methods for collecting sensors information
package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

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
	return nil, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Temperature returns min, max and average temperature
func (d *Device) Temperature() (float64, float64, float64) {
	return 0.0, 0.0, 0.0
}

// String formats sensor data as a string
func (s *TempSensor) String() string {
	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //
