package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".js") && !strings.HasSuffix(f.Name(), "min.js") {
			minifyFileWithExtension(f.Name(), ".js")
		}

		if strings.HasSuffix(f.Name(), ".css") && !strings.HasSuffix(f.Name(), "min.css") {
			minifyFileWithExtension(f.Name(), ".css")
		}
	}
}

func minifyFileWithExtension(fileName string, extension string) {

	buf := bytes.NewBuffer(nil)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	file.Close()
	s := string(buf.Bytes())

	s = minify(s)

	err = ioutil.WriteFile(strings.TrimSuffix(fileName, extension)+".min"+extension, []byte(s), 0644)
	if err != nil {
		panic(err)
	}

}

func minify(source string) string {

	mini := strings.Replace(source, `
`, " ", -1)
	mini = strings.Replace(mini, "  ", " ", -1)
	mini = strings.Replace(mini, `	`, "", -1)
	return mini

}
