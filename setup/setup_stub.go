//go:build !linux
// +build !linux

// Package setup provides methods to install/unistall application as a service on the
// system
package setup

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "os"

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ App contains basic application configuration
type App struct {
	Name       string   // Application name
	Options    []string // List of options
	DocsURL    string   // Documentation URL
	User       string   // Service user
	Identifier string   // Syslog identifier
	WorkingDir string   // Working dir

	StopSignal   string // Stop signal
	ReloadSignal string // Reload signal

	WithLog            bool // Create directory for logs
	WithoutPrivateTemp bool // Disable private temp

	Configs []Config // Configuration files
}

// ❗ Config contains configuration file data
//
// Note that all configurations are stored in /etc
type Config struct {
	Name string      // File name
	Data []byte      // Data
	Mode os.FileMode // File mode
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Install installs or reinstalls application on the system
func (app App) Install() error {
	panic("UNSUPPORTED")
}

// ❗ Uninstall uninstall unistalls application from the system
func (app App) Uninstall(full bool) error {
	panic("UNSUPPORTED")
}
