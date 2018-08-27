package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// inDir is dir to start. must be full path
// const inDir = "/Users/xah/xx_manual/"
const inDir = "/Users/xah/web/wordyenglish_com/"

const fnameRegex = `^blog.+\.html$`

const writeToFile = false

// doBackup. when writeToFile is also false, no backup is made
const doBackup = true
const backupSuffix = "~~"

var dirsToSkip = []string{".git"}

type frPair struct {
	fs string // find string
	rs string // replace string
}

var frPairs = []frPair{

	frPair{
		fs: `tttttttttttttthhhhhtht71317<aside id="aside-right-89129">`,
		rs: `<aside id="aside-right-89129">
<ul>
<li><a href="blog.html">Wordy Blog</a></li>
<li><a href="blog_past_2016-08.html">2016-08</a></li>
<li><a href="blog_past_2016-04.html">2016-04</a></li>
<li><a href="blog_past_2016-01.html">2016-01</a></li>
<li><a href="blog_past_2015-10.html">2015-10</a></li>
<li><a href="blog_past_2015-09.html">2015-09</a></li>
<li><a href="blog_past_2015-04.html">2015-04</a></li>
<li><a href="blog_past_2015-01.html">2015-01</a></li>
<li><a href="blog_past_2014-11.html">2014-11</a></li>
<li><a href="blog_past_2014-10.html">2014-10</a></li>
<li><a href="blog_past_2014-09.html">2014-09</a></li>
<li><a href="blog_past_2014-07.html">2014-07</a></li>
<li><a href="blog_past_2014-05.html">2014-05</a></li>
<li><a href="blog_past_2014-03.html">2014-03</a></li>
<li><a href="blog_past_2014-01.html">2014-01</a></li>
<li><a href="blog_past_2013-12.html">2013-12</a></li>
<li><a href="blog_past_2013-11.html">2013-11</a></li>
<li><a href="blog_past_2013-10.html">2013-10</a></li>
<li><a href="blog_past_2013-09.html">2013-09</a></li>
<li><a href="blog_past_2013-08.html">2013-08</a></li>
<li><a href="blog_past_2013-07.html">2013-07</a></li>
<li><a href="blog_past_2013-06.html">2013-06</a></li>
<li><a href="blog_past_2013-04.html">2013-04</a></li>
<li><a href="blog_past_2013-03.html">2013-03</a></li>
<li><a href="blog_past_2013-02.html">2013-02</a></li>
<li><a href="blog_past_2013-01.html">2013-01</a></li>
<li><a href="blog_past_2012-12.html">2012-12</a></li>
<li><a href="blog_past_2012-11.html">2012-11</a></li>
<li><a href="blog_past_2012-10.html">2012-10</a></li>
<li><a href="blog_past_2012-09.html">2012-09</a></li>
<li><a href="blog_past_2012-08.html">2012-08</a></li>
<li><a href="blog_past_2012-07.html">2012-07</a></li>
<li><a href="blog_past_2012-04.html">2012-04</a></li>
<li><a href="blog_past_2012-02.html">2012-02</a></li>
<li><a href="blog_past_2011-12.html">2011-12</a></li>
<li><a href="blog_past_2011-11.html">2011-11</a></li>
<li><a href="blog_past_2011-10.html">2011-10</a></li>
<li><a href="blog_past_2011-06.html">2011-06</a></li>
<li><a href="blog_past_2011-08.html">2011-08</a></li>
<li><a href="blog_past_2011-09.html">2011-09</a></li>
</ul>`,
	},
}

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
