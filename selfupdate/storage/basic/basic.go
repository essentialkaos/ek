package basic

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"

	"github.com/essentialkaos/ek/v13/req"
	"github.com/essentialkaos/ek/v13/selfupdate"
	"github.com/essentialkaos/ek/v13/selfupdate/storage"
	"github.com/essentialkaos/ek/v13/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Storage is basic storage client
type Storage struct {
	URL string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validate Storage interface
var _ storage.Storage = (*Storage)(nil)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyURL is returned if storage URL is empty
	ErrEmptyURL = fmt.Errorf("Storage URL is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewStorage creates a new Storage instance
func NewStorage(url string) *Storage {
	return &Storage{URL: url}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Check checks for updates in the storage
func (s *Storage) Check(appName, appVersion string) (selfupdate.Update, bool, error) {
	switch {
	case s == nil:
		return selfupdate.Update{}, false, storage.ErrNilStorage
	case s.URL == "":
		return selfupdate.Update{}, false, ErrEmptyURL
	case appName == "":
		return selfupdate.Update{}, false, storage.ErrEmptyName
	case appVersion == "":
		return selfupdate.Update{}, false, storage.ErrEmptyVersion
	}

	curVersion, err := version.Parse(appVersion)

	if err != nil {
		return selfupdate.Update{}, false, fmt.Errorf("Can't parse current version: %w", err)
	}

	latestVersionStr, err := s.getLatestVersion(appName, appVersion)

	if err != nil {
		return selfupdate.Update{}, false, fmt.Errorf("Can't get the latest version info: %w", err)
	}

	latestVersion, err := version.Parse(latestVersionStr)

	if err != nil {
		return selfupdate.Update{}, false, fmt.Errorf("Can't parse latest version: %w", err)
	}

	binaryURL := fmt.Sprintf(
		"%s/%s/%s/%s/%s/%s",
		s.URL, appName, latestVersionStr,
		storage.OSName(), storage.ArchName(), appName,
	)

	return selfupdate.Update{
		BinaryURL:    binaryURL,
		SignatureURL: binaryURL + ".sig",
		Version:      latestVersionStr,
	}, latestVersion.Greater(curVersion), nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLatestVersion fetches the latest version of the application from the storage
func (s *Storage) getLatestVersion(appName, appVersion string) (string, error) {
	resp, err := req.Request{
		URL:         s.URL + "/" + appName + "/latest.json",
		Timeout:     5 * time.Second,
		Query:       req.Query{"current": appVersion},
		AutoDiscard: true,
	}.Get()

	if err != nil {
		return "", fmt.Errorf("Can't send request to check the latest version: %w", err)
	}

	if resp.StatusCode != req.STATUS_OK {
		return "", fmt.Errorf("Storage returned non-ok status code (%d)", resp.StatusCode)
	}

	versionInfo := &struct {
		Version string `json:"version"`
	}{}

	err = resp.JSON(versionInfo)

	if err != nil {
		return "", fmt.Errorf("Can't parse the latest version info: %w", err)
	}

	return versionInfo.Version, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
