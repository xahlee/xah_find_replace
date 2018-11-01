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
	writeToFile  = true
	doBackup     = true
	backupSuffix = "~~"
)

var dirsToSkip = []string{".git"}

// fileList if not empty, only these are processed. Each element is a full path
var fileList = []string{

"/Users/xah/web/xahlee_info/powershell/install_and_help.html",
"/Users/xah/web/xahlee_info/powershell/powershell_help.html",
"/Users/xah/web/xahlee_info/powershell/commands.html",
"/Users/xah/web/xahlee_info/powershell/aliases.html",
"/Users/xah/web/xahlee_info/powershell/piping_output_input.html",
"/Users/xah/web/xahlee_info/powershell/environment_variables.html",
"/Users/xah/web/xahlee_info/powershell/automatic_variables.html",
"/Users/xah/web/xahlee_info/powershell/scripts.html",
"/Users/xah/web/xahlee_info/powershell/PowerShell_for_unixer.html",

}

var frPairs = []frPair{

	frPair{
		fs: `</div> <!-- end main -->`,
		rs: `</div> <!-- end main -->

<div class="panel_stick_7hpgn5">
<h4>PowerShell</h4>
<ol>
<li><a href="install_and_help.html">Install</a></li>
<li><a href="powershell_help.html">Help Command</a></li>
<li><a href="commands.html">PowerShell as cmd.exe</a></li>
<li><a href="aliases.html">list Alias, find Alias</a></li>
<li><a href="piping_output_input.html">Piping Output and Input</a></li>
<li><a href="environment_variables.html">Environment Variables</a></li>
<li><a href="automatic_variables.html">Predefined Variables</a></li>
<li><a href="scripts.html">Creating PowerShell Scripts</a></li>
<li><a href="PowerShell_for_unixer.html">PowerShell vs Bash</a></li>
</ol>
</div>`,
	},
}

type frPair struct {
	fs string // find string
	rs string // replace string
}

// ------------------------------------------------------------

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
