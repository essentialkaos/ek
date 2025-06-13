package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type githubRelease struct {
	Tag       string    `json:"tag_name"`
	Published time.Time `json:"published_at"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var githubAPI = "https://api.github.com"

// ////////////////////////////////////////////////////////////////////////////////// //

// GitHubChecker checks for updates on GitHub
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

// getLatestGitHubRelease fetches the latest release from a GitHub repository
func getLatestGitHubRelease(app, version, repository string) *githubRelease {
	var auth req.Auth

	engine := req.Engine{}

	engine.SetDialTimeout(3)
	engine.SetRequestTimeout(3)

	if os.Getenv("GH_TOKEN") != "" {
		auth = req.AuthBearer{os.Getenv("GH_TOKEN")}
	}

	resp, err := engine.Get(req.Request{
		URL:         githubAPI + "/repos/" + repository + "/releases/latest",
		Headers:     req.Headers{"X-GitHub-Api-Version": "2022-11-28"},
		Auth:        auth,
		AutoDiscard: true,
	})

	if err != nil || resp.StatusCode != 200 {
		return nil
	}

	if resp.Header.Get("X-RateLimit-Remaining") == "0" {
		return nil
	}

	release := &githubRelease{}
	err = resp.JSON(release)

	if err != nil {
		return nil
	}

	return release
}
