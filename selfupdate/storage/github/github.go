package github

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/req"
	"github.com/essentialkaos/ek/v13/selfupdate"
	"github.com/essentialkaos/ek/v13/selfupdate/storage"
	"github.com/essentialkaos/ek/v13/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Storage represents a GitHub-based storage using organization and repository releases
type Storage struct {
	Org        string // Org is the GitHub organization or user name
	Repository string // Repository is the GitHub repository name
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validate Storage interface
var _ storage.Storage = (*Storage)(nil)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyOrg indicates that the GitHub organization name is empty
	ErrEmptyOrg = fmt.Errorf("organization name is empty")

	// ErrEmptyRepo indicates that the GitHub repository name is empty
	ErrEmptyRepo = fmt.Errorf("repository name is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewStorage creates a new GitHub Storage using the given organization and repository
func NewStorage(org, repo string) *Storage {
	return &Storage{Org: org, Repository: repo}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Check queries GitHub releases for a newer version of the given application
func (s *Storage) Check(appName, appVersion string) (selfupdate.Update, bool, error) {
	switch {
	case s == nil:
		return selfupdate.Update{}, false, storage.ErrNilStorage
	case s.Org == "":
		return selfupdate.Update{}, false, ErrEmptyOrg
	case s.Repository == "":
		return selfupdate.Update{}, false, ErrEmptyRepo
	case appName == "":
		return selfupdate.Update{}, false, storage.ErrEmptyName
	case appVersion == "":
		return selfupdate.Update{}, false, storage.ErrEmptyVersion
	}

	curVersion, err := version.Parse(appVersion)

	if err != nil {
		return selfupdate.Update{}, false, fmt.Errorf("can't parse current version: %w", err)
	}

	latestVersionStr, err := s.getLatestVersion()

	if err != nil {
		return selfupdate.Update{}, false, fmt.Errorf("can't get the latest version info: %w", err)
	}

	latestVersion, err := version.Parse(latestVersionStr)

	if err != nil {
		return selfupdate.Update{}, false, fmt.Errorf("can't parse latest version: %w", err)
	}

	binaryURL := fmt.Sprintf(
		"https://github.com/%s/%s/releases/download/v%s/%s-%s-%s",
		s.Org, s.Repository, latestVersionStr,
		appName, storage.OSName(), storage.ArchName(),
	)

	return selfupdate.Update{
		BinaryURL:    binaryURL,
		SignatureURL: binaryURL + ".sig",
		Version:      latestVersionStr,
	}, latestVersion.Greater(curVersion), nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLatestVersion fetches the latest version from GitHub releases
func (s *Storage) getLatestVersion() (string, error) {
	var auth req.Auth

	ghToken := os.Getenv("GH_TOKEN")

	if ghToken != "" {
		auth = req.AuthBearer{ghToken}
	}

	resp, err := req.Request{
		URL:         fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", s.Org, s.Repository),
		Headers:     req.Headers{"X-GitHub-Api-Version": "2022-11-28"},
		Timeout:     5 * time.Second,
		Auth:        auth,
		AutoDiscard: true,
	}.Get()

	if err != nil {
		return "", fmt.Errorf("can't send request to check the latest version: %w", err)
	}

	if resp.Header.Get("X-RateLimit-Remaining") == "0" {
		return "", fmt.Errorf("rate limit reached")
	}

	if resp.StatusCode != req.STATUS_OK {
		return "", fmt.Errorf("storage returned non-ok status code (%d)", resp.StatusCode)
	}

	release := &struct {
		Tag string `json:"tag_name"`
	}{}

	err = resp.JSON(release)

	if err != nil {
		return "", fmt.Errorf("can't parse the latest version info: %w", err)
	}

	return strings.TrimLeft(release.Tag, "v"), nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
