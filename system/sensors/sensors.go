// Package sensors provides methods for collecting hardware sensor information
package sensors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Device contains sensor information collected from a single hwmon device
type Device struct {
	Name        string
	TempSensors []TempSensor
}

// TempSensor contains temperature readings from a single hwmon sensor input
type TempSensor struct {
	Name string
	Cur  float64
	Min  float64
	Max  float64
	Crit float64
}

// ////////////////////////////////////////////////////////////////////////////////// //
