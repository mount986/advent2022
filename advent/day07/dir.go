package day07

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Dir struct {
	Name   string
	Parent *Dir
	Subs   map[string]Dir
	Files  map[string]File
	size   int
}

var indent = 0

var ROOT = rootDir()

func rootDir() Dir {
	var dir Dir
	dir.Name = "/"
	dir.Subs = make(map[string]Dir, 0)
	dir.Files = make(map[string]File, 0)
	dir.size = 0

	return dir
}

func NewDir(name string, parent Dir) Dir {
	var dir Dir
	dir.Name = name
	dir.Parent = &parent
	dir.Subs = make(map[string]Dir, 0)
	dir.Files = make(map[string]File, 0)
	dir.size = 0

	return dir
}

func (d Dir) NewSub(name string) Dir {
	d.Subs[name] = NewDir(name, d)

	return d.Subs[name]
}

func (d Dir) NewFile(name string, size int) File {
	d.Files[name] = NewFile(name, size)

	return d.Files[name]
}

func (d Dir) ContainsSub(name string) bool {
	return slices.Contains(maps.Keys(d.Subs), name)
}

func (d Dir) ContainsFile(name string) bool {
	return slices.Contains(maps.Keys(d.Files), name)
}

func (d Dir) Size() int {
	if d.size != 0 {
		return d.size
	}

	for _, s := range d.Subs {
		d.size += s.Size()
	}

	for _, f := range d.Files {
		d.size += f.Size()
	}

	return d.size
}

func (d Dir) Print() {
	fmt.Printf("%*s %s %d\n", indent*2, "-", d.Name, d.Size())
	indent++

	for _, s := range d.Subs {
		s.Print()
	}

	for _, f := range d.Files {
		f.Print()
	}

	indent--
}
