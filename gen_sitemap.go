// version 2018-09-02

// given a dir, generate a sitemap.xml file for all its html files
// prints to stdout

// sitemapp file looks like this

// <?xml version="1.0" encoding="UTF-8"?>
// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
// <url><loc>http://xaharts.org/xamsi_calku/venus_comb/venus_comb.html</loc></url>
// <url><loc>http://xaharts.org/xamsi_calku/tusk/tusk.html</loc></url>
// </urlset>

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// inDir is dir to start. must be full path
var inDir = "/Users/xah/web/xaharts_org/arts/"

// var inDir = "/Users/xah/web/ergoemacs_org/"

var dirsToSkip = []string{".git"}

// fnameRegex. only these are searched
const fnameRegex = `\.html$`

var fnameRegexToSkip = []string{
	`^xx`,
	`^403error.html`,
	`^404error.html`,
}

var dirRegexToSkip = []string{
	`^\.git$`,
	`^xx`,
}

var dirPathToUrl = map[string]string{
	"/Users/xah/web/ergoemacs_org":    "http://ergoemacs.org",
	"/Users/xah/web/wordyenglish_com": "http://wordyenglish.com",
	"/Users/xah/web/xaharts_org":      "http://xaharts.org",
	"/Users/xah/web/xahlee_info":      "http://xahlee.info",
	"/Users/xah/web/xahlee_org":       "http://xahlee.org",
	"/Users/xah/web/xahmusic_org":     "http://xahmusic.org",
	"/Users/xah/web/xahporn_org":      "http://xahporn.org",
	"/Users/xah/web/xahsl_org":        "http://xahsl.org",
}

var pathToUrlReplacePair = getMatched(inDir, dirPathToUrl)

// getMatched return the pair from mm, whose key is a prefix in ss. If none, panic.
// version 2018-09-02
var getMatched = func(ss string, mm map[string]string) []string {
	var bb = []string{``, ``}
	for k, v := range mm {
		if strings.HasPrefix(ss, k) {
			bb[0] = k
			bb[1] = v
			return bb
		}
	}
	panic("logic error. 83580")
	return nil
}

// equalAny return true if x equals any of y
// version 2018-09-02
func equalAny(x string, y []string) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// matchAny return true if ss is matched by any of regex regexes.
// version 2018-09-01
func matchAny(ss string, regexes []string) bool {
	for _, re := range regexes {
		result, err := regexp.MatchString(re, ss)
		if err != nil {
			panic(err)
		}
		if result {
			return true
		}
	}
	return false
}

// getHeadBytes return the first n bytes in file at path
// version 2018-09-02
func getHeadBytes(path string, n int) []byte {
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	headBytes := make([]byte, n)
	m, err := file.Read(headBytes)
	if err != nil {
		log.Fatal(err)
	}
	return headBytes[:m]
}

func doFile(path string) error {
	var firstLinish = getHeadBytes(path, 200)

	var pmoved, err = regexp.Match("page_moved_64598", firstLinish)
	if err != nil {
		panic(err)
	}

	if !pmoved {

		fmt.Printf("<url><loc>%v</loc></url>\n", strings.Replace(path, pathToUrlReplacePair[0], pathToUrlReplacePair[1], 1))
	}

	// /Users/xah/web/ergoemacs_org/emacs/elisp_menu_for_major_mode.html
	// fmt.Printf("%q\n", x)

	return nil
}

func genSiteMap(dirX string) {

	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Println(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	var pWalker = func(pathX string, infoX os.FileInfo, errX error) error {
		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return errX
		}
		if infoX.IsDir() {
			if equalAny(filepath.Base(pathX), dirsToSkip) || matchAny(filepath.Base(pathX), dirRegexToSkip) {
				return filepath.SkipDir
			}

		} else {
			var fname = filepath.Base(pathX)
			var goodExtension, err = regexp.MatchString(fnameRegex, fname)
			if err != nil {
				panic("stupid MatchString error 59767")
			}
			if goodExtension && !matchAny(fname, fnameRegexToSkip) {
				doFile(pathX)
			}
		}
		return nil
	}
	err := filepath.Walk(inDir, pWalker)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", inDir, err)
	}

	fmt.Println(`</urlset>`)

}

func main() {
	genSiteMap(inDir)
}
