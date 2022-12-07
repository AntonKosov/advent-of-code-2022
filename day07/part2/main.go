package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() [][]string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	res := make([][]string, len(lines))
	for i, line := range lines {
		res[i] = strings.Split(line, " ")
	}

	return res
}

func process(data [][]string) int {
	const totalSpace = 70_000_000
	const requiredSpace = 30_000_000

	root := buildTree(data)
	usedSpace := countSize(root)

	unusedSpace := totalSpace - usedSpace
	toDelete := requiredSpace - unusedSpace

	candidate := root.size

	findSmallestCandidate(root, toDelete, &candidate)

	return candidate
}

func findSmallestCandidate(f *folder, toDelete int, candidate *int) {
	if f.size < toDelete {
		return
	}

	if f.size-toDelete < *candidate-toDelete {
		*candidate = f.size
	}

	for _, sf := range f.folders {
		findSmallestCandidate(sf, toDelete, candidate)
	}
}

func countSize(f *folder) int {
	for _, sf := range f.folders {
		f.size += countSize(sf)
	}

	for _, fl := range f.files {
		f.size += fl.size
	}

	return f.size
}

func buildTree(data [][]string) *folder {
	root := newFolder("", nil)
	cf := root
	for _, line := range data {
		switch tp := line[0]; tp {
		case "$":
			switch command := line[1]; command {
			case "cd":
				switch dest := line[2]; dest {
				case "/":
					cf = cf.root()
				case "..":
					cf = cf.parent
				default:
					cf = cf.folders[dest]
				}
			case "ls":
				// ignore
			default:
				panic(fmt.Sprintf("unknown command: %v", command))
			}
		case "dir":
			folderName := line[1]
			cf.folders[folderName] = newFolder(folderName, cf)
		default:
			size, fileName := aoc.StrToInt(tp), line[1]
			cf.files = append(cf.files, file{name: fileName, size: size})
		}
	}

	return root
}

type file struct {
	name string
	size int
}

type folder struct {
	parent  *folder
	name    string
	folders map[string]*folder
	files   []file
	size    int
}

func newFolder(name string, parent *folder) *folder {
	return &folder{
		name:    name,
		parent:  parent,
		folders: map[string]*folder{},
	}
}

func (f *folder) root() *folder {
	for f.parent != nil {
		f = f.parent
	}

	return f
}
