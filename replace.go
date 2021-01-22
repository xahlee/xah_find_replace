// find replace string pairs in a dir
// no regex
// version 2019-01-13
// website: http://xahlee.info/golang/goland_find_replace.html

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
	inDir        = "/Users/xah/web/"
	fnameRegex   = `\.html$`
	writeToFile  = true
	doBackup     = true
	backupSuffix = "~~"
)

var dirsToSkip = []string{".git"}

// fileList if not empty, only these are processed. Each element is a full path
var fileList = []string{}

// frPairs is a slice of frPair struct.
var frPairs = []frPair{

	frPair{
		findStr:    `<div class="lpanel_h7h547">
<h4>Emacs</h4>
<ul>
<li><a href="emacs/emacs_modernization.html">Emacs Modernization</a></li>
<li><a href="emacs_fun_index.html">Emacs Fun</a></li>
<li><a href="emacs/blog.html">Blog</a></li>
</ul>
</div>`,
		replaceStr: ``,
	},
}

// ------------------------------------------------------------

type frPair struct {
	findStr    string // find string
	replaceStr string // replace string
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

func printSliceStr(sliceX []string) error {
	for k, v := range sliceX {
		fmt.Printf("%v %v\n", k, v)
	}
	return nil
}

func doFile(path string) error {
	contentBytes, er := ioutil.ReadFile(path)
	if er != nil {
		fmt.Printf("processing %v\n", path)
		panic(er)
	}
	var content = string(contentBytes)
	var changed = false
	for _, pair := range frPairs {
		var found = strings.Index(content, pair.findStr)
		if found != -1 {
			content = strings.Replace(content, pair.findStr, pair.replaceStr, -1)
			changed = true
		}
	}
	if changed {
		fmt.Printf("〘%v〙\n", path)

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

func main() {
	scriptPath, errPath := os.Executable()
	if errPath != nil {
		panic(errPath)
	}

	fmt.Println("-*- coding: utf-8; mode: xah-find-output -*-")
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("Script: %v\n", filepath.Base(scriptPath))
	fmt.Printf("In dir: %v\n", inDir)
	fmt.Printf("File regex filter: %v\n", fnameRegex)
	fmt.Printf("Write to file: %v\n", writeToFile)
	fmt.Printf("Do backup: %v\n", doBackup)
	fmt.Printf("fileList:\n")
	printSliceStr(fileList)
	fmt.Printf("Find replace pairs: 「%v」\n", frPairs)
	fmt.Println()

	if len(fileList) >= 1 {
		for _, v := range fileList {
			doFile(v)
		}
	} else {
		err := filepath.Walk(inDir, pWalker)
		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", inDir, err)
		}
	}

	fmt.Println()

	if !writeToFile {
		fmt.Printf("Note: writeToFile is %v\n", writeToFile)
	}

	fmt.Printf("%v\n", "Done.")
}
