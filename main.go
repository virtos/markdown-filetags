package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const tagSymbol string = "+"
const indexFileName string = "_filetags.md"

var whiteSpaceDelim = [...]string{" ", "_", ".", "[", "]"}

type tagFiles struct {
	tag  string
	file []string
}

type tags struct {
	tf []tagFiles
}

func initialize() (dir string, short bool) {
	flag.BoolVar(&short, "s", true, "Include files from subdirectories")
	flag.StringVar(&dir, "t", ".", "Target directory path")
	flag.Parse()
	dir, err := filepath.Abs(dir)
	check(err)
	return
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

func getTags(source string) []string {
	res := make([]string, 0)
	source = " " + strings.ToLower(source)

	flds := strings.FieldsFunc(source, func(r rune) bool {
		if contains(whiteSpaceDelim[:], string(r)) {
			return true
		}
		return false
	})
	for _, fld := range flds {
		fld = strings.Trim(fld, " ")
		if fld == "" || fld == tagSymbol {
			continue
		}
		if string([]rune(fld)[0:1]) == tagSymbol {
			tag := string([]rune(fld)[1:])
			if !contains(res, tag) {
				res = append(res, tag)
			}
		}
	}
	return res
}

func appendSubdirTags(source, newTags tags, subdir string) tags {
	for _, curTag := range newTags.tf {
		for _, curFile := range curTag.file {
			source.append(curTag.tag, subdir+"/"+curFile)
			source.append(filepath.Base(subdir), subdir+"/"+curFile)
		}
	}
	return source
}

func writeIndexFile(tfs tags, dir string) {
	dstFName := filepath.Join(dir, indexFileName)
	if len(tfs.tf) == 0 {
		os.Remove(dstFName)
		return
	}

	newInfo := formatRes(tfs)

	dstClean := false
	s, err := os.Stat(dstFName)
	if err == nil && s.Size() == int64(len(newInfo)) {
		if err == nil {
			rawExs, _ := ioutil.ReadFile(dstFName)
			if newInfo == string(rawExs) {
				dstClean = true
			}
		}
	}

	if !dstClean {
		os.Remove(dstFName)
		f, err := os.Create(dstFName)
		check(err)
		defer f.Close()
		_, err = f.WriteString(formatRes(tfs))
		check(err)
	}
}

func processDir(dir string, includeSubdirTags bool) tags {
	var tfs tags

	lst, err := ioutil.ReadDir(dir)
	check(err)

	for _, file := range lst {
		if file.IsDir() {
			subdirTags := processDir(filepath.Join(dir, file.Name()), includeSubdirTags)
			if includeSubdirTags {
				tfs = appendSubdirTags(tfs, subdirTags, file.Name())
			}
			continue
		} else if filepath.Ext(strings.TrimSpace(file.Name())) != ".md" || file.Name() == indexFileName {
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

	writeIndexFile(tfs, dir)

	return tfs
}

func main() {
	processDir(initialize())
}
