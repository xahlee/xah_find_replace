package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// inDir is dir to start. must be full path
// var inDir = "/Users/xah/xx_manual/"
const inDir = "/Users/xah/web/xahlee_info/"

var dirsToSkip = []string{".git"}

type frPair struct {
	fs string // find string
	rs string // replace string
}

var frPairs = []frPair{

	frPair{
		fs: `haskell`,
		rs: `ppppp`,
	},
}

// ext is file extension, with the dot.
// only these are searched
const fnameRegex = `^blog.+\.html$`

const writeToFile = false

// doBackup. when writeToFile is also false, no backup is made
const doBackup = true
const backupSuffix = "~~"

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
		fmt.Printf("changed: %v\n", path)
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
