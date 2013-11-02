package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var file = flag.String("f", "", "Bibtex File. Just cat the file for now.")

var out = flag.String("o", "list.csv", "File with names path and MD5")

var ls = flag.Bool("ls", false, "List Files")

var ver = flag.Bool("ver", false, "Show version info.")
var dbg = flag.Bool("d", false, "Print debug information.")
var dbgfile = flag.String("df", "", "File to print debug information.")

var dbgoutput *os.File

func main() {
	flag.Parse()

	if *dbgfile != "" {
		var err error
		dbgoutput, err = os.Create(*dbgfile)
		if err != nil {
			panic(err)
		}
		defer dbgoutput.Close()
	}

	debug("Running Debug Mode...")

	if *ver {
		fmt.Println(`obo, Version Tag 0.1
Copyright (C) Valentin Mayer-Eichberger
License GPLv2+: GNU GPL version 2 or later <http://gnu.org/licenses/gpl.html>
There is NO WARRANTY, to the extent permitted by law.`)
		return
	}

	if *file != "" {
		parse(*file)
	}

	if *ls {
		showFiles()
	}

}

func showFiles() {

	var f filepath.WalkFunc

	f = func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if !strings.HasPrefix(info.Name(),".") {

            debug("name",info.Name(),"path:", path)
        } 
		return nil
	}

	filepath.Walk(".", f)
}

func debug(arg ...interface{}) {
	if *dbg {
		if *dbgfile == "" {
			fmt.Print("dbg: ")
			for _, s := range arg {
				fmt.Print(s, " ")
			}
			fmt.Println()
		} else {
			ss := "dbg: "
			for _, s := range arg {
				ss += fmt.Sprintf("%v", s) + " "
			}
			ss += "\n"

			if _, err := dbgoutput.Write([]byte(ss)); err != nil {
				panic(err)
			}
		}
	}
}

type Entry struct {
	path  string
	md5   int
	texId string
}

func parse(filename string) {

	input, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Please specifiy correct path to instance. File does not exist: ", filename)
		panic(err)
	}

	output, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	lines := strings.Split(string(input), "\n")

	for _, l := range lines {
		fmt.Println(l)
	}

	return
}
