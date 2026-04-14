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

	"github.com/essentialkaos/ek/v14/fmtc"
	"github.com/essentialkaos/ek/v14/lscolors"
	"github.com/essentialkaos/ek/v14/sortutil"
	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Tree represents a file tree built from a flat list of paths
type Tree struct {
	Path     string
	MaxDepth int
	Root     *Dir
}

// Dir represents a single directory node containing files and subdirectories
type Dir struct {
	Files []string
	Dirs  DirIndex
	level int
}

// DirIndex maps directory names to their corresponding [Dir] nodes
type DirIndex map[string]*Dir

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// TreeColorTag is the fmtc color tag applied to tree branch and connector lines
	TreeColorTag = "{s-}"

	// DirColorTag is the fmtc color tag applied to directory name labels
	DirColorTag = "{*}"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// pathSeparator is path separator
var pathSeparator = string(filepath.Separator)

// ////////////////////////////////////////////////////////////////////////////////// //

// FromPaths builds a Tree from the given slice of file paths, deduplicating
// nested entries
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

		if prevPath == p {
			continue
		}

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

// IsEmpty returns true if the tree is nil or contains no files or directories
func (t *Tree) IsEmpty() bool {
	return t == nil || t.Root.IsEmpty()
}

// Render prints the file tree to stdout using fmtc color formatting
func (t *Tree) Render() {
	if t.IsEmpty() {
		return
	}

	t.Root.render(strutil.Q(t.Path, pathSeparator), true, true)
}

// IsEmpty returns true if the directory is nil or contains no files or subdirectories
func (d *Dir) IsEmpty() bool {
	return d == nil || (len(d.Files) == 0 && len(d.Dirs) == 0)
}

// Names returns a naturally sorted slice of directory names in the index
func (d DirIndex) Names() []string {
	var result []string

	for n := range d {
		result = append(result, n)
	}

	sortutil.StringsNatural(result)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// addFile inserts a file path into the tree, creating intermediate directories
// as needed
func (t *Tree) addFile(filePath string) {
	filePath = strings.TrimLeft(filePath, pathSeparator)

	if strings.TrimSpace(filePath) == "" {
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

// createDir traverses or creates the directory chain described by path and returns
// the leaf [Dir]
func (d *Dir) createDir(path []string) *Dir {
	cwd := d

	for _, p := range path {
		dd, ok := cwd.Dirs[p]

		if !ok {
			cwd.Dirs[p] = &Dir{Dirs: DirIndex{}, level: cwd.level + 1}
			dd = cwd.Dirs[p]
		}

		cwd = dd
	}

	return cwd
}

// render recursively prints this directory and its contents with tree-drawing
// connectors
func (d *Dir) render(name string, isFirst, isLast bool) {
	if d.IsEmpty() {
		return
	}

	if isFirst {
		fmtc.Printfn(TreeColorTag+"┌{!} "+DirColorTag+"%s{!}", name)
	} else {
		fmtc.Printfn(
			TreeColorTag+strings.Repeat("│  ", d.level-1)+"%s{!} "+DirColorTag+"%s{!}",
			strutil.B(isLast && len(d.Files) == 0, "└─", "├─"), name,
		)
	}

	for i, n := range d.Dirs.Names() {
		d.Dirs[n].render(n, false, i+1 == len(d.Dirs) && len(d.Files) == 0)
	}

	for i, f := range d.Files {
		fmtc.Printfn(
			TreeColorTag+strings.Repeat("│  ", d.level)+"%s{!} %s",
			strutil.B(i+1 == len(d.Files), "└─", "├─"), lscolors.Colorize(f),
		)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getRootDir returns the longest common directory prefix shared by all paths
// in the slice
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

	idx := strings.LastIndex(rootDir, pathSeparator)

	if idx >= 0 {
		rootDir = rootDir[:idx]
	}

	return rootDir
}

// getLongestPath returns the path with the greatest number of separators from
// the given slice
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
