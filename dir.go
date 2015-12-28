// Package dir contains a function to return the most recently modified file
// found at a Glob pattern and a supporting sort utility.
package dir

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// LastModifiedFile returns a string path to the most recently modified file
// at a glob pattern.
func LastModifiedFile(pattern string) string {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	var files []file
	for _, f := range filenames {
		file := file{path: f, mtime: getModTime(f)}
		files = append(files, file)
	}
	sort.Sort(byTime(files))

	return files[len(files)-1].path
}

func getModTime(path string) time.Time {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.ModTime()
}

type file struct {
	path  string
	mtime time.Time
}

// type byTime and its functions Len, Less and Swap are supporting
// implementations to allow for sorting a slice of files by their
// modification time.
// https://gobyexample.com/sorting-by-functions
type byTime []file

func (bt byTime) Len() int {
	return len(bt)
}

func (bt byTime) Less(i, j int) bool {
	return bt[i].mtime.Before(bt[j].mtime)
}

func (bt byTime) Swap(i, j int) {
	bt[i], bt[j] = bt[j], bt[i]
}
