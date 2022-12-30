package day07

import (
	"fmt"
)

type File struct {
	Name string
	size int
}

func NewFile(name string, size int) File {
	var file File
	file.Name = name
	file.size = size

	return file
}

func (f File) Size() int {
	return f.size
}

func (f File) Print() {
	fmt.Printf("%*s %s %d\n", indent*2, "-", f.Name, f.Size())
}
