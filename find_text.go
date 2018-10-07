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
const inDir = "/Users/xah/web/xahlee_info/python/"

var dirsToSkip = []string{".git"}

const findStr = `<code class="python">`

// fnameRegex. only these are searched
const fnameRegex = `\.html$`

// number of chars (actually bytes) to show before the found string
const before = 100
const after = 100

// scriptName returns the current running script path
func scriptName() string {
	name, errPath := os.Executable()
	if errPath != nil {
		panic(errPath)
	}
	return name
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
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
	var re = regexp.MustCompile(regexp.QuoteMeta(findStr))

	var textBytes, er = ioutil.ReadFile(path)
	if er != nil {
		panic(er)
	}

	var indexes = re.FindAllIndex(textBytes, -1)

	if len(indexes) != 0 {
		fmt.Println("\n==================================================")
		fmt.Printf("%v ⦉%v⦊ \n", len(indexes), path)

		for _, k := range indexes {
			fmt.Println("-------------------------")
			var foundStart = k[0]
			var foundEnd = k[1]
			var showStart = max(foundStart-before, 0)
			var showEnd = min(foundEnd+after, len(textBytes))

			fmt.Printf("%s「%s」%s\n", textBytes[showStart:foundStart],
				textBytes[foundStart:foundEnd],
				textBytes[foundEnd:showEnd])
		}
	}
	return nil
}

func main() {

	fmt.Println("-*- coding: utf-8; mode: xah-find-output -*-")
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("Script: %v\n", filepath.Base(scriptName()))
	fmt.Printf("in dir: %v\n", inDir)
	fmt.Printf("file regex filter: %v\n", fnameRegex)
	fmt.Printf("Find string: 「%v」\n", findStr)
	fmt.Println()

	var pWalker = func(pathX string, infoX os.FileInfo, errX error) error {
		// first thing to do, check error. and decide what to do about it
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
	fmt.Println("Done.")
}
