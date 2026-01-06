// Package filetree provides methods for printing file tree
package filetree

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/lscolors"
	"github.com/essentialkaos/ek/v13/sortutil"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Tree is file tree
type Tree struct {
	Path     string
	MaxDepth int
	Root     *Dir
}

// Dir contains info about directory
type Dir struct {
	Files []string
	Dirs  DirIndex
	level int
}

// DirIndex is map "name" → dir
type DirIndex map[string]*Dir

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// TreeColorTag is color tag of tree lines
	TreeColorTag = "{s-}"

	// DirColorTag is color tag for directory names
	DirColorTag = "{*}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// pathSeparator is path separator
var pathSeparator = string(filepath.Separator)

// ////////////////////////////////////////////////////////////////////////////////// //

// FromPaths builds tree from paths slice
func FromPaths(paths []string) *Tree {
	if len(paths) == 0 {
		return &Tree{}
	}

	sortutil.StringsNatural(paths)

	tree := &Tree{
		Path: getRootDir(paths),
		Root: &Dir{Dirs: DirIndex{}, level: 0},
	}

	var prevPath string

	for _, p := range paths {
		p = strings.TrimPrefix(p, tree.Path)

		if prevPath == "" || strings.HasPrefix(p, prevPath) {
			prevPath = p
			continue
		}

		tree.addFile(prevPath)

		prevPath = p
	}

	tree.addFile(prevPath)

	return tree
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns true if tree is empty
func (t *Tree) IsEmpty() bool {
	return t == nil || t.Root.IsEmpty()
}

// Render renders tree data
func (t *Tree) Render() {
	if t.IsEmpty() {
		return
	}

	t.Root.render(strutil.Q(t.Path, pathSeparator), true, true)
}

// IsEmpty returns true if dir is empty
func (d *Dir) IsEmpty() bool {
	return d == nil || (len(d.Files) == 0 && len(d.Dirs) == 0)
}

// IsEmpty returns sorted slice of directory names
func (d DirIndex) Names() []string {
	var result []string

	for n := range d {
		result = append(result, n)
	}

	sortutil.StringsNatural(result)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// addFile adds file to the tree
func (t *Tree) addFile(filePath string) {
	filePath = strings.TrimLeft(filePath, pathSeparator)

	if strings.ReplaceAll(filePath, " ", "") == "" {
		return
	}

	fileName := path.Base(filePath)

	if !strings.Contains(filePath, pathSeparator) {
		t.Root.Files = append(t.Root.Files, fileName)
		return
	}

	dirPath := strings.Split(path.Dir(filePath), pathSeparator)
	dir := t.Root.createDir(dirPath)
	dir.Files = append(dir.Files, fileName)

	t.MaxDepth = max(t.MaxDepth, len(dirPath))
}

// createDir creates new directory structure or returns existing one
func (d *Dir) createDir(path []string) *Dir {
	cwd := d

	for _, p := range path {
		dd, ok := cwd.Dirs[p]

		if !ok {
			cwd.Dirs[p] = &Dir{Dirs: DirIndex{}, level: cwd.level + 1}
			cwd = cwd.Dirs[p]
		} else {
			cwd = dd
		}
	}

	return cwd
}

// Render renders dir data
func (d *Dir) render(name string, isFirst, isLast bool) {
	if d.IsEmpty() {
		return
	}

	if isFirst {
		fmtc.Printfn(TreeColorTag+"┌{!} "+DirColorTag+"%s{!}", name)
	} else {
		if isLast && len(d.Files) == 0 {
			fmtc.Printfn(
				TreeColorTag+strings.Repeat("│  ", d.level-1)+"└─{!} "+DirColorTag+"%s{!}",
				name,
			)
		} else {
			fmtc.Printfn(
				TreeColorTag+strings.Repeat("│  ", d.level-1)+"├─{!} "+DirColorTag+"%s{!}",
				name,
			)
		}
	}

	for i, n := range d.Dirs.Names() {
		if i+1 == len(d.Dirs) {
			d.Dirs[n].render(n, false, true)
		} else {
			d.Dirs[n].render(n, false, false)
		}
	}

	for i, f := range d.Files {
		if i+1 == len(d.Files) {
			fmtc.Printfn(
				TreeColorTag+strings.Repeat("│  ", d.level)+"└─{!} %s",
				lscolors.Colorize(f),
			)
		} else {
			fmtc.Printfn(
				TreeColorTag+strings.Repeat("│  ", d.level)+"├─{!} %s",
				lscolors.Colorize(f),
			)
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getRootDir returns root dir of all paths in given slice
func getRootDir(paths []string) string {
	rootDir := getLongestPath(paths)

	for _, p := range paths {
		for i := range len(rootDir) {
			if i == len(p) || rootDir[i] != p[i] {
				rootDir = rootDir[:i]
				break
			}
		}
	}

	return strings.TrimRight(rootDir, pathSeparator)
}

// getLongestPath returns the most longest path from given slice
func getLongestPath(paths []string) string {
	var maxDepth int
	var longestPath string

	for _, p := range paths {
		depth := strings.Count(p, pathSeparator)

		if maxDepth < depth {
			maxDepth = depth
			longestPath = path.Dir(p)
		}
	}

	return longestPath
}
