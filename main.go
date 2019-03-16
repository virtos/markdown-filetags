package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const tagSymbol string = "-"

var whiteSpaceDelim = [...]string{" ", "_"}

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

	targetDir, err := filepath.Abs(filepath.Dir(targetDir))
	check(err)

	processDir(targetDir)
}

func processDir(dir string) {
	fmt.Println("Processing directory ", dir, "..")
	lst, err := ioutil.ReadDir(dir)
	check(err)
	tags := make(map[string]string)

	for _, file := range lst {

		if file.IsDir() {
			defer processDir(filepath.Join(dir, file.Name()))
			// } else if filepath.Ext(strings.TrimSpace(file.Name())) != ".md" {
			// 	continue
		} else if !strings.Contains(file.Name(), tagSymbol) {
			continue
		}

		arr := strings.Split(file.Name(), tagSymbol)
		for _, elem := range arr {
			if string(elem[1]) == file.Name() {
				break
			}
			// take the first value before any delimiter
			elem += " "
			curTag := elem[:strings.IndexAny(elem, " _.")]
			//tags[file.Name()] = curTag
			tags[curTag] = file.Name()
		}

		fmt.Println("File :", file.Name())
		fmt.Println(tags)
	}
}
