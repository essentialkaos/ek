package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"time"

	"github.com/essentialkaos/ek/v12/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type githubRelease struct {
	Tag       string    `json:"tag_name"`
	Published time.Time `json:"published_at"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var githubAPI = "https://api.github.com"

// ////////////////////////////////////////////////////////////////////////////////// //

// GitHubChecker checks new releases on GitHub
func GitHubChecker(app, version, data string) (string, time.Time, bool) {
	if version == "" || data == "" || !isUpdateCheckRequired() {
		return "", time.Time{}, false
	}

	release := getLatestGitHubRelease(app, version, data)

	if release == nil {
		return "", time.Time{}, false
	}

	return strings.TrimLeft(release.Tag, "v"), release.Published, true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLatestGitHubRelease fetches the latest release from GitHub
func getLatestGitHubRelease(app, version, repository string) *githubRelease {
	engine := req.Engine{}

	engine.SetDialTimeout(3)
	engine.SetRequestTimeout(3)
	engine.SetUserAgent(app, version, "GoEK.v12")

	response, err := engine.Get(req.Request{
		URL:         githubAPI + "/repos/" + repository + "/releases/latest",
		Headers:     req.Headers{"X-GitHub-Api-Version": "2022-11-28"},
		AutoDiscard: true,
	})

	if err != nil || response.StatusCode != 200 {
		return nil
	}

	if response.Header.Get("X-RateLimit-Remaining") == "0" {
		return nil
	}

	release := &githubRelease{}
	err = response.JSON(release)

	if err != nil {
		return nil
	}

	return release
}
