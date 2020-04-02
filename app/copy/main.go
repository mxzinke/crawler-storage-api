package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var filePath string
var copyPath string

func init() {
	if len(os.Args) < 3 {
		log.Fatalln("Please specify a path to copy and a result path")
	}

	realFile := os.Args[1]
	copyFile := os.Args[2]

	absDir, err := filepath.Abs(os.Args[0])
	check(err)

	_Dir := path.Dir(absDir)

	if filepath.IsAbs(realFile) {
		filePath = realFile
	} else {
		filePath = filepath.Join(_Dir, realFile)
	}

	if filepath.IsAbs(copyFile) {
		copyPath = copyFile
	} else {
		copyPath = filepath.Join(_Dir, copyFile)
	}
}

func main() {
	f, err := ioutil.ReadFile(filePath)
	check(err)

	fmt.Printf("bytes copying: %d\n", len(f))

	data := []byte(f)
	err = ioutil.WriteFile(copyPath, data, 0644)
	check(err)
}
