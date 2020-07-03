package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const tagSymbol string = "+"
const indexFileName string = "_filetags.md"

var whiteSpaceDelim = [...]string{" ", "_"}

type tagFiles struct {
	tag  string
	file []string
}

type tags struct {
	tf []tagFiles
}

func (tt *tags) append(tag, fileName string) {
	tag = strings.ToLower(tag)
	//record found tag and filename
	tagFound := false
	for i := range tt.tf {
		if tt.tf[i].tag == tag {
			tagFound = true
			fileFound := false
			for _, cf := range tt.tf[i].file {
				if cf == fileName {
					fileFound = true
					break
				}
			}
			if !fileFound {
				tt.tf[i].file = append(tt.tf[i].file, fileName)
			}
			break
		}
	}
	if !tagFound {
		tt.tf = append(tt.tf, tagFiles{tag, []string{fileName}})
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func formatRes(tfs tags) string {
	var str strings.Builder
	for _, ctf := range tfs.tf {
		str.WriteString("<details markdown='1'><summary markdown='1'>" + ctf.tag + "</summary>\n")
		sort.Slice(ctf.file, func(i, j int) bool {
			return ctf.file[i] < ctf.file[j]
		})
		for _, cf := range ctf.file {
			str.WriteString("<li><a href=\"" + strings.Replace(cf, " ", "%20", -1) + "\">" + strings.TrimRight(strings.Replace(cf, "_", " ", -1), ".md") + "</a>\n")
		}
		str.WriteString("</details>\n\n")

	}
	return str.String()
}

func getTags(fn string) []string {
	res := make([]string, 0)
	fn = strings.ToLower(fn)
	arr := strings.Split(fn, tagSymbol)
	for i, elem := range arr {
		if elem == "" {
			continue
		}
		if i == 0 && fn[:1] != tagSymbol {
			continue
		}
		// take the first value before any delimiter
		elem += " "
		res = append(res, elem[:strings.IndexAny(elem, " _.[]")])
	}
	return res
}

func processDir(dir string) {
	var tfs tags

	lst, err := ioutil.ReadDir(dir)
	check(err)

	for _, file := range lst {
		if file.IsDir() {
			defer processDir(filepath.Join(dir, file.Name()))
			continue
		} else if filepath.Ext(strings.TrimSpace(file.Name())) != ".md" {
			continue
		} else if !strings.Contains(file.Name(), tagSymbol) {
			continue
		}

		for _, curTag := range getTags(file.Name()) {
			tfs.append(curTag, file.Name())
		}
	}

	// sort tags
	sort.Slice(tfs.tf, func(i, j int) bool {
		return tfs.tf[i].tag < tfs.tf[j].tag
	})

	// write result

	if len(tfs.tf) > 0 {
		f, err := os.Create(filepath.Join(dir, indexFileName))
		check(err)
		defer f.Close()
		_, err = f.WriteString(formatRes(tfs))
		check(err)
	}
}

func main() {
	targetDir := "."
	args := os.Args
	if len(args) > 1 {
		targetDir = args[1]
	}
	targetDir, err := filepath.Abs(targetDir)
	check(err)

	processDir(targetDir)
}
