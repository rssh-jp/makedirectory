package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	defer func(t time.Time) {
		log.Println(time.Now().Sub(t))
	}(time.Now())

	var outdir string
	var depth, count int
	var isFileCreate bool
	flag.StringVar(&outdir, "o", "", "target directory")
	flag.IntVar(&depth, "d", 1, "directory depth")
	flag.IntVar(&count, "c", 10, "directory count")
	flag.BoolVar(&isFileCreate, "f", false, "create file flag of deep depth")

	flag.Parse()

	if outdir == "" {
		log.Fatal("Not specified target directory")
	}

	err := os.MkdirAll(outdir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(filepath.Abs(outdir))

	err = recursiveMakeDirectory(depth, outdir, count, isFileCreate)
	if err != nil {
		log.Fatal(err)
	}
}

func makeFile(dir string) error {
	fd, err := os.Create(dir)
	if err != nil && !os.IsExist(err) {
		return err
	}

	defer fd.Close()

	_, err = fd.WriteString(dir)
	if err != nil {
		return err
	}

	return nil
}

func recursiveMakeDirectory(depth int, dir string, count int, isFileCreate bool) error {
	if depth == 0 {
		if isFileCreate {
			return makeFile(filepath.Join(dir, "file"))
		} else {
			return nil
		}
	}

	for i := 0; i < count; i++ {
		p := filepath.Join(dir, strconv.Itoa(i))
		err := os.Mkdir(p, 0777)
		if err != nil && !os.IsExist(err) {
			return err
		}

		err = recursiveMakeDirectory(depth-1, p, count, isFileCreate)
		if err != nil {
			return err
		}
	}

	return nil

}
