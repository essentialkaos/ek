package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type gitlabRelease struct {
	Tag      string    `json:"tag_name"`
	Released time.Time `json:"released_at"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var gitlabAPI = "https://gitlab.com/api/v4"

// ////////////////////////////////////////////////////////////////////////////////// //

// GitLabChecker checks new releases on GitLab
func GitLabChecker(app, version, data string) (string, time.Time, bool) {
	if version == "" || data == "" || !isUpdateCheckRequired() {
		return "", time.Time{}, false
	}

	release := getLatestGitLabRelease(app, version, data)

	if release == nil {
		return "", time.Time{}, false
	}

	return strings.TrimLeft(release.Tag, "v"), release.Released, true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLatestGitLabRelease fetches the latest release from GitLab
func getLatestGitLabRelease(app, version, repository string) *gitlabRelease {
	engine := req.Engine{}

	engine.SetDialTimeout(3)
	engine.SetRequestTimeout(3)

	if strings.Contains(repository, "/") {
		repository = strings.ReplaceAll(repository, "/", "%2F")
	}

	resp, err := engine.Get(req.Request{
		URL:         gitlabAPI + "/projects/" + repository + "/releases/permalink/latest",
		AutoDiscard: true,
	})

	if err != nil || resp.StatusCode != 200 {
		return nil
	}

	if resp.Header.Get("RateLimit-Remaining") == "0" {
		return nil
	}

	release := &gitlabRelease{}
	err = resp.JSON(release)

	if err != nil {
		return nil
	}

	return release
}
