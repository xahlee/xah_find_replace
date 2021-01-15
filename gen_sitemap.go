// given a dir, generate a sitemap.xml file for all its html files
// version 2018-11-04, 2021-01-14

// http://xahlee.info/golang/golang_gen_sitemap.html

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var xDomains = []string{
	"ergoemacs.org",
	"wordyenglish.com",
	"xaharts.org",
	"xahlee.info",
	"xahlee.org",
	"xahmusic.org",
	"xahsl.org",
}

var domainRootdirMap = map[string]string{
	"ergoemacs.org":    "/Users/xah/web/ergoemacs_org",
	"wordyenglish.com": "/Users/xah/web/wordyenglish_com",
	"xaharts.org":      "/Users/xah/web/xaharts_org",
	"xahlee.info":      "/Users/xah/web/xahlee_info",
	"xahlee.org":       "/Users/xah/web/xahlee_org",
	"xahmusic.org":     "/Users/xah/web/xahmusic_org",
	"xahsl.org":        "/Users/xah/web/xahsl_org",
}

var domainToSitemapHtmlFilenameMap = map[string]string{
	"ergoemacs.org":    "sitemap.html",
	"wordyenglish.com": "sitemap.html",
	"xaharts.org":      "sitemap.html",
	"xahlee.info":      "sitemap.html",
	"xahlee.org":       "sitemap.html",
	"xahmusic.org":     "index.html",
	"xahsl.org":        "sitemap.html",
}

var sitemapXmlFilename = "sitemap.xml"

var dirsToSkip = []string{
	".git",
	"emacs_manual",
	"REC-SVG11-20110816",
	"clojure-doc-1.8",
	"css_2.1_spec",
	"css_transitions",
	"js_es2011",
	"js_es2015",
	"js_es2015_orig",
	"js_es2016",
	"js_es2018",
	"node_api",
}

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

// getMatched return the pair from map1, whose key is a prefix of str1. If none, panic.
// version 2018-09-02
var getMatched = func(str1 string, map1 map[string]string) []string {
	var result = []string{``, ``}
	for k, v := range map1 {
		if strings.HasPrefix(str1, k) {
			result[0] = k
			result[1] = v
			return result
		}
	}
	panic("logic error. 83580")
	return nil
}

type FileInfo struct {
	path  string
	url   string
	title string
	moved bool
}

var fileList []FileInfo

// getFileInfo opens a file returns struct FileInfo
// if first line of file contains string page_moved_64598, then return nothing
func getFileInfo(path string) FileInfo {
	var header1 = getHeadBytes(path, 6000)
	var pmoved, err = regexp.Match("page_moved_64598", header1)
	if err != nil {
		panic(err)
	}
	var title = ""
	if !pmoved {
		var re = regexp.MustCompile(`<title>([^<]+)</title>`)
		var result = re.FindSubmatch(header1)
		if result == nil {
			fmt.Printf("%v\n", path)
			panic("no title found")
		} else {
			title = string(result[1])
		}
	}
	return FileInfo{
		path:  path,
		title: title,
		moved: pmoved,
	}
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

// matchAny return true if str1 is matched by any of regex regexes.
// version 2018-09-01
func matchAny(str1 string, regexes []string) bool {
	for _, re := range regexes {
		result, err := regexp.MatchString(re, str1)
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

// initFileList fills fileList.
func initFileList(dirX string) {
	fileList = make([]FileInfo, 0, 9000)
	var pWalker = func(pathX string, infoX os.FileInfo, errX error) error {
		if errX != nil {
			panic(fmt.Sprintf("error 「%v」 at a path 「%q」\n", errX, pathX))
		}
		if infoX.IsDir() {
			if equalAny(filepath.Base(pathX), dirsToSkip) || matchAny(filepath.Base(pathX), dirRegexToSkip) {
				return filepath.SkipDir
			}
		} else {
			var fname = filepath.Base(pathX)
			var goodExtension, err = regexp.MatchString(fnameRegex, fname)
			if err != nil {
				panic("stupid golang MatchString error")
			}
			if goodExtension {
				if !matchAny(fname, fnameRegexToSkip) {
					var fInfo = getFileInfo(pathX)
					if !fInfo.moved {
						fileList = append(fileList, fInfo)
					}
				}
			}
		}
		return nil
	}
	err := filepath.Walk(dirX, pWalker)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dirX, err)
	}
}

// writeToFile saves contentX into file at path pathX
func writeToFile(contentX []byte, pathX string) {
	var fileH, err = os.Create(pathX)
	if err != nil {
		panic(err)
	}
	defer fileH.Close()
	var _, errW = fileH.Write(contentX)
	if errW != nil {
		panic(errW)
	}
}

// take a file list and return a string of them as sitemap xml
func fileListToSitemapXml(fileList []FileInfo, rootDir string, domainUrl string) []byte {
	var output = make([]byte, 0, 10000)
	output = append(output, (`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
`)...)
	for _, fileinfo := range fileList {
		var fUrl = strings.Replace(filepath.ToSlash(fileinfo.path), rootDir, domainUrl, 1)
		output = append(output, fmt.Sprintf("<url><loc>%v</loc></url>\n", fUrl)...)
	}
	output = append(output, (`</urlset>` + "\n")...)
	return output
}

func updateSitemapHtmlFile(domain string, rootDir string) {
	var sitemapHtmlFilePath = rootDir + "/" + domainToSitemapHtmlFilenameMap[domain]
	var searchStrStart = `<ol id="sitemapList58821" style="margin-left:16px;">`
	var searchStrEnd = `</ol>`

	myTextB, myErr := ioutil.ReadFile(sitemapHtmlFilePath)
	if myErr != nil {
		panic(myErr)
	}

	var searchStrStartPos = bytes.Index(myTextB, []byte(searchStrStart))
	var searchStrEndPos = bytes.Index(myTextB, []byte(searchStrEnd))

	if searchStrStartPos == -1 {
		panic("searchStrStart not found")
	}
	if searchStrEndPos == -1 {
		panic("searchStrEnd not found")
	}

	var linesBytes = make([]byte, 0, 200000)

	for _, filex := range fileList {
		var path96864 = filex.path
		var title43608 = filex.title
		var relativePath, err = filepath.Rel(rootDir, path96864)
		if err != nil {
			panic(err)
		}
		linesBytes = append(linesBytes, fmt.Sprintf("<li><a href=\"%s\" target=\"_blank\">%s</a></li>\n", filepath.ToSlash(relativePath), title43608)...)

	}

	var beginChunck = myTextB[:searchStrStartPos+len(searchStrStart)]
	var endChunck = myTextB[searchStrEndPos:]

	var newTextB = append([]byte{}, beginChunck...)
	newTextB = append(newTextB, linesBytes...)
	newTextB = append(newTextB, endChunck...)

	err := ioutil.WriteFile(sitemapHtmlFilePath, newTextB, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Modified %v\n", sitemapHtmlFilePath)

}

func main() {
	for _, domain := range xDomains {
		var rootDir = domainRootdirMap[domain]
		initFileList(rootDir)
		var saveToPath = filepath.Join(rootDir, sitemapXmlFilename)
		var output = fileListToSitemapXml(fileList, rootDir, "http://"+domain)
		writeToFile(output, saveToPath)
		fmt.Printf("file saved to: %v\n", saveToPath)

		updateSitemapHtmlFile(domain, rootDir)

		fmt.Println()

	}
	fmt.Println("Done")
}
