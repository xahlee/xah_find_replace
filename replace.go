// find replace string pairs in a dir
// no regex
// version 2018-10-15

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
	inDir        = "/Users/xah/web/xahlee_info/powershell"
	fnameRegex   = `\.html$`
	writeToFile  = false
	doBackup     = true
	backupSuffix = "~~"
)

var dirsToSkip = []string{".git"}

// fileList if not empty, only these are processed. Each element is a full path
var fileList = []string{

"/Users/xah/web/xahlee_info/js/svg_basic_examples.html",
"/Users/xah/web/xahlee_info/js/svg_path_spec.html",
"/Users/xah/web/xahlee_info/js/svg_path_ellipse_arc.html",
"/Users/xah/web/xahlee_info/js/svg_specify_style.html",
"/Users/xah/web/xahlee_info/js/svg_shape_styles.html",
"/Users/xah/web/xahlee_info/js/svg_viewport.html",
"/Users/xah/web/xahlee_info/js/svg_viewBox.html",
"/Users/xah/web/xahlee_info/js/svg_transformation.html",
"/Users/xah/web/xahlee_info/js/svg_text_element.html",
"/Users/xah/web/xahlee_info/js/svg_font_size.html",
"/Users/xah/web/xahlee_info/js/svg_structure_elements.html",
"/Users/xah/web/xahlee_info/js/js_scritping_svg_basics.html",
"/Users/xah/web/xahlee_info/js/svg_clock.html",
"/Users/xah/web/xahlee_info/js/svg_clock.js",
"/Users/xah/web/xahlee_info/js/svg_animation.html",

}

var frPairs = []frPair{

	frPair{
		fs: `<nav class="nav-back-85230"><a href="../js/index.html">Web Dev Tutorials</a></nav>`,
		rs: `<nav class="nav-back-85230"><a href="svg.html">SVG Tutorial</a></nav>`,
	},
}

// ------------------------------------------------------------

type frPair struct {
	fs string // find string
	rs string // replace string
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
	fmt.Printf("fileList: %#v\n", fileList)
	fmt.Printf("Find replace pairs: 「%#v」\n", frPairs)
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
