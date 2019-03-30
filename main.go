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

type tf struct {
	Tag  string
	File []string
}

func check(err error) {
	if err != nil {
		panic(err)
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

func processDir(dir string) {
	lst, err := ioutil.ReadDir(dir)
	check(err)

	tfs := make([]tf, 0)

	for _, file := range lst {

		if file.IsDir() {
			defer processDir(filepath.Join(dir, file.Name()))
			continue
		} else if filepath.Ext(strings.TrimSpace(file.Name())) != ".md" {
			continue
		} else if !strings.Contains(file.Name(), tagSymbol) {
			continue
		}
		fileName := strings.ToLower(file.Name())
		arr := strings.Split(fileName, tagSymbol)
		for i, elem := range arr {
			if elem == "" {
				continue
			}
			if i == 0 && fileName[:1] != tagSymbol {
				continue
			}
			// take the first value before any delimiter
			elem += " "
			curTag := elem[:strings.IndexAny(elem, " _.")]

			//record found tag and filename
			found := false
			for i := range tfs {
				if tfs[i].Tag == curTag {
					tfs[i].File = append(tfs[i].File, file.Name())
					found = true
					break
				}
			}
			if !found {
				tfs = append(tfs, tf{curTag, []string{fileName}})
			}
		}
	}

	// sort tags
	sort.Slice(tfs, func(i, j int) bool {
		return tfs[i].Tag < tfs[j].Tag
	})

	// result
	if len(tfs) > 0 {
		f, err := os.Create(filepath.Join(dir, indexFileName))
		check(err)
		defer f.Close()
		_, err = f.WriteString("---\nheading:accordion\n...\n\n")
		check(err)

		for _, ctf := range tfs {
			_, err := f.WriteString("### " + ctf.Tag + "\n|   |\n|---|\n")
			check(err)
			sort.Slice(ctf.File, func(i, j int) bool {
				return ctf.File[i] < ctf.File[j]
			})
			for _, cf := range ctf.File {
				_, err := f.WriteString("| [" + strings.Replace(cf, "_", " ", -1) + "](" + strings.Replace(cf, " ", "%20", -1) + ") |\n")
				check(err)
			}
		}
	}
}
