// version 2018-10-07

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// inDir is dir to start. must be full path
const inDir = "/Users/xah/xx_manual/"

var dirsToSkip = []string{".git"}

type frPair struct {
	fs string // find string
	rs string // replace string
}

var frPairs = []frPair{

	frPair{
		fs: `<code class="[a-z]">tttttt`,
		rs: `ppppp`,
	},
}

// ext is file extension, with the dot.
// only these are searched
const fnameRegex = `\.html$`

const writeToFile = false

// doBackup. when writeToFile is also false, no backup is made
const doBackup = true
const backupSuffix = "~~"

// scriptPath returns the current running script path
// version 2018-10-07
func scriptPath() string {
	name, errPath := os.Executable()
	if errPath != nil {
		panic(errPath)
	}
	return name
}

// stringMatchAny return true if x equals any of y
func stringMatchAny(x string, y []string) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

func doFile(path string) error {
	content, er := ioutil.ReadFile(path)
	if er != nil {
		panic(er)
	}

	// var contentOriginal = content
	var changed = false
	for _, pair := range frPairs {
		var re = regexp.MustCompile(regexp.QuoteMeta(pair.fs))
		var matched = re.Match(content)
		if matched {
			content = re.ReplaceAllLiteral(content, []byte(pair.rs))
			changed = true
		}
	}

	if changed {
		fmt.Printf("changed: 〘%v〙\n", path)
		if writeToFile {
			if doBackup {
				err := os.Rename(path, path+backupSuffix)
				if err != nil {
					panic(err)
				}
			}
			err2 := ioutil.WriteFile(path, content, 0644)
			if err2 != nil {
				panic("write file problem")
			}
		}
	}
	return nil
}

func main() {
	// need to print date, find string, rep string, and root dir, extension

	fmt.Println("-*- coding: utf-8; mode: xah-find-output -*-")
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("Script: %v\n", filepath.Base(scriptPath()))
	fmt.Printf("in dir: %v\n", inDir)
	fmt.Printf("file regex filter: %v\n", fnameRegex)
	fmt.Printf("Find replace pairs: 「%#v」\n", frPairs)
	fmt.Println()

	var pWalker = func(pathX string, infoX os.FileInfo, errX error) error {
		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return errX
		}
		if infoX.IsDir() {
			if stringMatchAny(filepath.Base(pathX), dirsToSkip) {
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
