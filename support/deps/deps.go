// Package pkgs provides methods for collecting information about used dependencies
package deps

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"runtime/debug"
	"strings"

	"github.com/essentialkaos/ek/v13/support"

	"github.com/essentialkaos/depsy"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Extract extracts dependencies info from gomod data
func Extract(gomod []byte, withIndirect ...bool) []support.Dep {
	if len(withIndirect) > 0 && withIndirect[0] {
		return filterDeps(depsy.Extract(gomod, true))
	}

	return filterDeps(depsy.Extract(gomod, false))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// filterDeps filters dependencies from gomod using information from bundled build info
func filterDeps(deps depsy.Dependencies) []support.Dep {
	var result []support.Dep

	buildInfo, _ := debug.ReadBuildInfo()

	for _, dep := range deps {
		depInfo := support.Dep{
			Version: dep.Version,
			Path:    dep.PrettyPath(),
			Extra:   dep.Extra,
		}

		if buildInfo != nil {
			hasDep, version := hasBuiltDep(dep, buildInfo)

			if !hasDep {
				continue
			}

			if version != "" && strings.Contains(version, "(") {
				depInfo.Version = strings.Trim(version, "()")
			}
		}

		result = append(result, depInfo)
	}

	return result
}

// hasBuiltDep checks if given dependency is present in build info
func hasBuiltDep(dep depsy.Dependency, buildInfo *debug.BuildInfo) (bool, string) {
	for _, bDep := range buildInfo.Deps {
		if bDep.Path == dep.Path {
			if bDep.Replace != nil {
				return true, bDep.Replace.Version
			}

			return true, ""
		}
	}

	return false, ""
}
