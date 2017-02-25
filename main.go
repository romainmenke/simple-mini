package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

FILE_ITERATOR:
	for _, f := range files {
		if !isFile(sourceDir + f.Name()) {
			fmt.Println(f.Name())
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

		content := readFile(sourceDir + f.Name())
		writeFile(content, f.Name(), extension, outDir)
	}
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

	err := ioutil.WriteFile(out+strings.TrimSuffix(fileName, extension)+"min."+extension, []byte(s), 0644)
	if err != nil {
		panic(err)
	}

}

type minifyWriter struct {
	io.Writer
}

func (m *minifyWriter) Write(p []byte) (n int, err error) {

	l := len(p)
	min := minify(p)

	l2, err := m.Writer.Write(min)
	if err != nil {
		return 0, err
	}

	_ = l2
	return l, err

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
