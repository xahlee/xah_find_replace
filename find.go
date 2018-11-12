// version 2018-10-07

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
	"unicode/utf8"
)

// inDir is dir to start. must be full path
var inDir = "/Users/xah/web/xahlee_info/comp/"

var dirsToSkip = []string{
	".git"}

// const findStr = `</mark>
// → integers.`

const findStr = `U+2124: DOUBLE-STRUCK CAPITAL Z">`

// fnameRegex. only these are searched
const fnameRegex = `\.html$`

// number of chars (actually bytes) to show before the found string
const before = 1
const after = 100

const fileSep = "ff━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"

const occurSep = "oo────────────────────────────────────────────────────────────\n"

const occurBracketL = '〖'
const occurBracketR = '〗'

const posBracketL = '❪'
const posBracketR = '❫'

// scriptPath returns the current running script path
// version 2018-10-07
func scriptPath() string {
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
	var textBytes, er = ioutil.ReadFile(path)
	if er != nil {
		panic(er)
	}
	var re = regexp.MustCompile(regexp.QuoteMeta(findStr))
	var indexes = re.FindAllIndex(textBytes, -1)

	var bytesLength = len(textBytes)

	if len(indexes) != 0 {
		for _, k := range indexes {
			var foundStart = k[0]
			var foundEnd = k[1]
			var showStart = max(foundStart-before, 0)

			for !utf8.RuneStart(textBytes[showStart]) {
				showStart = max(showStart-1, 0)
			}

			var showEnd = min(foundEnd+after, bytesLength)
			for !utf8.RuneStart(textBytes[showEnd]) {
				showEnd = min(showEnd+1, bytesLength)
			}

			// 			fmt.Printf("%s〖%s〗%s\n", textBytes[showStart:foundStart],
			fmt.Printf("%c%d%c %s%c%s%c%s\n",
				posBracketL,
				utf8.RuneCount(textBytes[0:foundStart+1]),
				posBracketR,
				textBytes[showStart:foundStart],
				occurBracketL,
				textBytes[foundStart:foundEnd],
				occurBracketR,
				textBytes[foundEnd:showEnd])
			fmt.Println(occurSep)
		}
		fmt.Printf("%v 〘%v〙\n", len(indexes), path)
		fmt.Println(fileSep)
	}
	return nil
}

func main() {

	fmt.Println("-*- coding: utf-8; mode: xah-find-output -*-")
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("Script: %v\n", filepath.Base(scriptPath()))
	fmt.Printf("in dir: %v\n", inDir)
	fmt.Printf("file regex filter: %v\n", fnameRegex)
	fmt.Printf("Find string: 「%v」\n", findStr)
	fmt.Println()
	fmt.Println(fileSep)

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
	fmt.Println("\nDone.")
}
