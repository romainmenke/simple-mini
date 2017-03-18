package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {

	source := flag.String("source", "./", "source directory")
	out := flag.String("out", "./", "output directory")
	flag.Parse()

	sourceDir := strings.TrimSuffix(*source, "/") + "/"
	outDir := strings.TrimSuffix(*out, "/") + "/"
	createIfMissing(outDir)

	exclude := flag.Args()

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

FILE_ITERATOR:
	for _, f := range files {
		if !isFile(sourceDir + f.Name()) {
			continue FILE_ITERATOR
		}

		for _, exc := range exclude {
			if strings.Contains(f.Name(), exc) {
				continue FILE_ITERATOR
			}
		}

		if strings.Contains(f.Name(), "min") {
			continue FILE_ITERATOR
		}

		nameComponents := strings.Split(f.Name(), ".")
		extension := nameComponents[len(nameComponents)-1]

		j := job{
			sourceDir: sourceDir,
			outDir:    outDir,
			fileName:  f.Name(),
			extension: extension,
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			execute(j)
		}()
	}

	wg.Wait()
}

type job struct {
	sourceDir string
	outDir    string
	fileName  string
	extension string
}

func execute(j job) {
	content := readFile(j.sourceDir + j.fileName)
	writeFile(content, j.fileName, j.extension, j.outDir)
}

func readFile(name string) []byte {
	buf := bytes.NewBuffer(nil)
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	file.Close()
	return buf.Bytes()
}

func writeFile(content []byte, fileName string, extension string, out string) {

	s := minify(content)

	err := ioutil.WriteFile(out+strings.TrimSuffix(fileName, extension)+"min."+extension, s, 0644)
	if err != nil {
		panic(err)
	}

}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return !fileInfo.IsDir()
}

func createIfMissing(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
