package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var rootDir = flag.String("rootDir", ".", "The root directory to start from")
var skipHidden = flag.Bool("skipHidden", true, "True if hidden files and directories should be skipped.")
var dryRun = flag.Bool("dryRun", true, "A dryrun prints out all the file moves instead of actually moving them.")

func flattenDirectories(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		n := info.Name()
		if len(n) > 1 && strings.HasPrefix(n, ".") {
			return filepath.SkipDir
		}
		return nil
	}
	np := fmt.Sprintf("%s%c%s", *rootDir, os.PathSeparator, info.Name())
	return moveFile(path, np, *dryRun)
}

func moveFile(old, new string, dry bool) error {
	if dry {
		_, err := fmt.Printf("Moving %v to %v\n", old, new)
		return err
	}

	return os.Rename(old, new)
}

func main() {
	flag.Parse()
	filepath.Walk(*rootDir, flattenDirectories)
}
