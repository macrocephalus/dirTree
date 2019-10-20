package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err := tree(out, path, "", printFiles)
	return err
}

func tree(out io.Writer, root string, indent string, printFiles bool) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	if !fi.IsDir() {
		if printFiles {
			if fi.Size() > 0 {
				fmt.Fprintf(out, "%s (%v%s)\n", fi.Name(), fi.Size(), "b")
			} else {
				fmt.Fprintf(out, "%s %s\n", fi.Name(), "(empty)")
			}

		}
		return nil
	} else {
		if fi.Name()[0] != '.' && indent != "" {
			fmt.Fprintln(out, fi.Name())
		}
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}
	//fmt.Println(fis)

	var names []string
	for _, fi := range fis {
		if fi.IsDir() {
			names = append(names, fi.Name())
		} else {
			if printFiles {

				names = append(names, fi.Name())
			}
		}
		//names = append(names, fi.Name())

	}
	sort.Strings(names)

	for i, name := range names {
		add := "│\t" //нужно розабраться!!!
		if i == len(names)-1 {

			fmt.Fprintf(out, indent+"└───")
			add = "\t"
		} else {
			fmt.Fprintf(out, indent+"├───")
		}

		if err := tree(out, filepath.Join(root, name), indent+add, printFiles); err != nil {
			return err
		}
	}

	return nil
}
