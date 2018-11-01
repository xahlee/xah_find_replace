// 2018-10-07
// given a dir, find any occurence of
// <code class="python">
// </code>
// that contains
// & < >

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

// inDir is dir to start. must be full path
var inDir = "/Users/xah/web/xahlee_info/python/"

var dirsToSkip = []string{".git"}

// stringMatchAny return true if x equals any of y
func stringMatchAny(x string, y []string) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// scriptPath returns the current running script path
// version 2018-10-07
func scriptPath() string {
	name, errPath := os.Executable()
	if errPath != nil {
		panic(errPath)
	}
	return name
}

func doFile(path string) error {
	// open file, get content
	// search content

	// <code class="python">
	// </code>
	// that contains
	// & < >

	var textBytes, er = ioutil.ReadFile(path)
	if er != nil {
		panic(er)
	}

	var re = regexp.MustCompile(`<code class="[-_A-Za-z0-9]+">((?s).+?)</code>`)

	var reresult = re.FindAllSubmatch(textBytes, -1)

	for i, v := range reresult {
		var meat = string(v[1])
		meat = strings.Replace(meat, "&amp;", ``, -1)
		meat = strings.Replace(meat, "&lt;", ``, -1)
		meat = strings.Replace(meat, "&gt;", ``, -1)

		var isFound, err = regexp.MatchString(`[<>&]`, meat)
		if err != nil {
			panic(err)
		}

		if isFound {
			fmt.Printf("=================================\n")
			fmt.Printf("〘%v〙\n\n", path)

			fmt.Printf("%v,〖%v〗\n\n", i, string(meat))

		}

	}

	return nil
}

func main() {

	fmt.Println("-*- coding: utf-8; mode: xah-find-output -*-")
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("Script: %v\n", filepath.Base(scriptPath()))
	fmt.Printf("in dir: %v\n", inDir)
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
			var x, err = regexp.MatchString(`\.html$`, filepath.Base(pathX))
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
