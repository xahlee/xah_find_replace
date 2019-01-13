// find replace string pairs in a dir
// no regex
// version 2019-01-13

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
	inDir        = "/Users/xah/web/ergoemacs_org/ttt"
	fnameRegex   = `\.html`
	writeToFile  = true
	doBackup     = true
	backupSuffix = "~~"
)

var dirsToSkip = []string{".git"}

// fileList if not empty, only these are processed. Each element is a full path
var fileList = []string{

"/Users/xah/web/ergoemacs_org/emacs/elisp_count-region.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_run_current_file.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_delete-current-file.html",
"/Users/xah/web/ergoemacs_org/emacs/emacs_copy_file_path.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_convert_line_ending.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_make-backup.html",
"/Users/xah/web/ergoemacs_org/emacs/emacs_copy_rectangle_text_to_clipboard.html",
"/Users/xah/web/ergoemacs_org/emacs/emacs_dired_convert_images.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_dired_rename_space_to_underscore.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_python_2to3.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_move_code_to_files.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_update_atom.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_update_pagetag.html",
"/Users/xah/web/ergoemacs_org/emacs/elisp_copy-paste_register_1.html",

}

// frPairs is a slice of frPair struct.
var frPairs = []frPair{

	frPair{
		fs: `<div class="buyxahemacs97449">`,
		rs: `<div class="topic_xl">
<h4>Elisp Commands Do thing-at-point</h4>
<ol>
<li><a href="elisp_count-region.html">Count Words</a></li>
<li><a href="elisp_run_current_file.html">Run Current File</a></li>
<li><a href="elisp_delete-current-file.html">Delete Current File</a></li>
<li><a href="emacs_copy_file_path.html">Copy File Path</a></li>
<li><a href="elisp_convert_line_ending.html">Convert Line Ending</a></li>
<li><a href="elisp_make-backup.html">Make Backup</a></li>
<li><a href="emacs_copy_rectangle_text_to_clipboard.html">Copy Column of Text to Clipboard</a></li>
<li><a href="emacs_dired_convert_images.html">Convert Image File</a></li>
<li><a href="elisp_dired_rename_space_to_underscore.html">Dired Rename Space to Hyphen</a></li>
<li><a href="elisp_python_2to3.html">Python 2to3 Wrapper</a></li>
<li><a href="elisp_move_code_to_files.html">Move Code to Files</a></li>
<li><a href="elisp_update_atom.html">Update Web Feed</a></li>
<li><a href="elisp_update_pagetag.html">Updating Page Tags</a></li>
<li><a href="elisp_copy-paste_register_1.html">Single Key Copy/Paste Register</a></li>
</ol>
</div>

<div class="buyxahemacs97449">`,
	},

	frPair{
		fs: `</div> <!-- end main -->`,
		rs: `</div> <!-- end main -->

<div class="panel_stick_7hpgn5">
<ul>
<li><a href="../comp/comp_lang_tutorials_index.html">Lang tutorials</a></li>
<li><a href="../linux/linux_index.html">Practical Linux</a></li>
<li><a href="../linux/computer_networking_index.html">TCP/IP</a></li>
<li><a href="../comp/lisp_index.html">On Lisp</a></li>
<li><a href="../comp/syntax_soup_index.html">Syntax Soup</a></li>
<li><a href="../comp/semantic_noodle_index.html">Semantic Noodle</a></li>
<li><a href="../comp/comp_lang_doc_problems.html">doc by dummies</a></li>
<li><a href="../comp/python_problems.html">Why Python Suck</a></li>
<li><a href="../comp/comp_jargons_index.html">Jargons</a></li>
<li><a href="../js/web_design_index.html">Web Design</a></li>
<li><a href="../js/web_html_validation_index.html">HTML History</a></li>
<li><a href="../w/spam_index.html">Web Spam Scam</a></li>
<li><a href="../UnixResource_dir/writ/anti_hacker_2k_index.html">Anti Hacker y2k</a></li>
<li><a href="../comp/technological_musing.html">Tech, Past n Future</a></li>
<li><a href="../kbd/keyboarding.html">Keyboard Guide</a></li>
</ul>
</div>`,
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
