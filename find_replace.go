package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	// inDir is dir to start. must be full path
	inDir        = "/Users/xah/xx_manual/"
	fnameRegex   = `\.html$`
	writeToFile  = false
	doBackup     = true
	backupSuffix = "~~"
)

var dirsToSkip = []string{".git"}

var frPairs = []frPair{

	frPair{
		fs: `emacs`,
		rs: `hhhhhhh`,
	},
}

type frPair struct {
	fs string // find string
	rs string // replace string
}

// ------------------------------------------------------------

// pass return false if x equals any of y
func pass(x string, y []string) bool {
	for _, v := range y {
		if x == v {
			return false
		}
	}
	return true
}

func doFile(path string) error {
	contentBytes, er := ioutil.ReadFile(path)
	if er != nil {
		panic(er)
	}
	var content = string(contentBytes)
	var changed = false
	for _, pair := range frPairs {
		var found = strings.Index(content, pair.fs)
		if found != -1 {
			content = strings.Replace(content, pair.fs, pair.rs, -1)
			changed = true
		}
	}
	if changed {
		fmt.Printf("changed: %v\n", path)
		if writeToFile {
			if doBackup {
				err := os.Rename(path, path+backupSuffix)
				if err != nil {
					panic(err)
				}
			}
			err2 := ioutil.WriteFile(path, []byte(content), 0644)
			if err2 != nil {
				panic("write file problem")
			}
		}
	}
	return nil
}

func main() {
	scriptName, errPath := os.Executable()
	if errPath != nil {
		panic(errPath)
	}

	fmt.Printf("%v\n", time.Now())
	fmt.Printf("Script: %v\n", filepath.Base(scriptName))
	fmt.Printf("In dir: %v\n", inDir)
	fmt.Printf("File regex filter: %v\n", fnameRegex)
	fmt.Printf("Write to file: %v\n", writeToFile)
	fmt.Printf("Do backup: %v\n", doBackup)
	fmt.Printf("Find replace pairs: 「%#v」\n", frPairs)
	fmt.Println()

	var pWalker = func(pathX string, infoX os.FileInfo, errX error) error {
		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return errX
		}
		if infoX.IsDir() {
			if !pass(filepath.Base(pathX), dirsToSkip) {
				return filepath.SkipDir
			}
		} else {
			var x, err = regexp.MatchString(fnameRegex, filepath.Base(pathX))
			if err != nil {
				panic("stupid MatchString error 59767")
			}
			if x {
				doFile(pathX)
			}
		}
		return nil
	}

	err := filepath.Walk(inDir, pWalker)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", inDir, err)
	}

	fmt.Printf("%v\n", "Done.")
}
