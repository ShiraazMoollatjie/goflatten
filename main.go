package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var rootDir = flag.String("rootDir", ".", "The root directory to start from.")
var skipHidden = flag.Bool("skipHidden", true, "True if hidden files and directories should be skipped.")
var skipDupes = flag.Bool("skipDupes", true, "True if duplicate files should be skipped.")
var dryRun = flag.Bool("dryRun", true, "A dryrun prints out all the file moves instead of actually moving them.")

var dupes = make(map[string]bool)

func flattenDirectories(path string, info os.FileInfo, _ error) error {
	if info.IsDir() {
		n := info.Name()
		if len(n) > 1 && strings.HasPrefix(n, ".") {
			return filepath.SkipDir
		}
		return nil
	}
	np := fmt.Sprintf("%s%c%s", *rootDir, os.PathSeparator, info.Name())

	if maybeSkipDupes(np) {
		return nil
	}

	return moveFile(path, np, *dryRun)
}

// maybeSkipDupes returns true if the newPath should be skipped.
func maybeSkipDupes(newPath string) bool {
	_, ok := dupes[newPath]
	if ok {
		log.Printf("WARNING: Duplicate file %s found.\n", newPath)
		if *skipDupes {
			return true
		}
	}

	dupes[newPath] = true
	return false
}

func moveFile(old, new string, dry bool) error {
	if dry {
		log.Printf("MOVING: %v to %v\n", old, new)
		return nil
	}

	return os.Rename(old, new)
}

func main() {
	flag.Parse()
	err := filepath.Walk(*rootDir, flattenDirectories)
	if err != nil {
		panic(err)
	}
}
