// Package update contains update checkers
package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v8/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type githubRelease struct {
	Tag       string    `json:"tag_name"`
	Published time.Time `json:"published_at"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GitHubChecker check new releases on GitHub
func GitHubChecker(app, version, data string) (string, time.Time, bool) {
	ghRelease := getLatestGitHubRelease(app, version, data)

	if ghRelease == nil {
		return "", time.Time{}, false
	}

	return strings.TrimLeft(ghRelease.Tag, "v"), ghRelease.Published, true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLatestRelease fetch the latest release from GitHub
func getLatestGitHubRelease(app, version, repository string) *githubRelease {
	engine := req.Engine{}

	engine.SetDialTimeout(1)
	engine.SetRequestTimeout(1)
	engine.SetUserAgent(app, version, "go.ek/6")

	response, err := engine.Get(req.Request{
		URL:         "https://api.github.com/repos/" + repository + "/releases/latest",
		AutoDiscard: true,
	})

	if err != nil || response.StatusCode != 200 {
		return nil
	}

	if response.Header.Get("X-RateLimit-Remaining") == "0" {
		return nil
	}

	var ghRelease = &githubRelease{}

	err = response.JSON(ghRelease)

	if err != nil {
		return nil
	}

	return ghRelease
}
