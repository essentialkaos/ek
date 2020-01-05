// Package update contains update checkers for different services
package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v11/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type githubRelease struct {
	Tag       string    `json:"tag_name"`
	Published time.Time `json:"published_at"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GitHubChecker checks new releases on GitHub
func GitHubChecker(app, version, data string) (string, time.Time, bool) {
	if version == "" || data == "" {
		return "", time.Time{}, false
	}

	ghRelease := getLatestGitHubRelease(app, version, data)

	if ghRelease == nil {
		return "", time.Time{}, false
	}

	return strings.TrimLeft(ghRelease.Tag, "v"), ghRelease.Published, true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLatestRelease fetches the latest release from GitHub
func getLatestGitHubRelease(app, version, repository string) *githubRelease {
	engine := req.Engine{}

	engine.SetDialTimeout(1)
	engine.SetRequestTimeout(1)
	engine.SetUserAgent(app, version, "go.ek")

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
