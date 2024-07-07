package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"time"

	"github.com/essentialkaos/ek.v13/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ReleaseInfo contains info about the latest version of application
type ReleaseInfo struct {
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// UpdateChecker checks new releases on custom storage
func UpdateChecker(app, version, data string) (string, time.Time, bool) {
	if version == "" || data == "" || !isUpdateCheckRequired() {
		return "", time.Time{}, false
	}

	release := getLastReleaseInfo(app, version, data)

	if release == nil {
		return "", time.Time{}, false
	}

	return release.Version, release.Date, release.Version != version
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLastReleaseInfo fetches info about the latest version
func getLastReleaseInfo(app, version, storage string) *ReleaseInfo {
	engine := req.Engine{}

	engine.SetDialTimeout(3)
	engine.SetRequestTimeout(3)
	engine.SetUserAgent(app, version, "GoEK.v13")

	response, err := engine.Get(req.Request{
		URL:         storage + "/latest.json",
		Accept:      req.CONTENT_TYPE_JSON,
		AutoDiscard: true,
	})

	if err != nil || response.StatusCode != 200 {
		return nil
	}

	release := &ReleaseInfo{}
	err = response.JSON(release)

	if err != nil {
		return nil
	}

	return release
}

// isUpdateCheckRequired checks if update check is required
func isUpdateCheckRequired() bool {
	return !(os.Getenv("CI") == "true" && os.Getenv("EK_TEST_PORT") == "")
}
