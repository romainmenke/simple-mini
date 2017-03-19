package main

import "bytes"

// "\n"
var lineBreak = []byte{10}

// " "
var space = []byte{32}

// "  "
var doubleSpace = []byte{32, 32}

// "	"
var tab = []byte{9}

// ""
var empty = []byte{}

func minify(source []byte) []byte {
	// replace tabs with empty strings
	source = bytes.Replace(source, tab, empty, -1)

	// replace linebreak with a space
	source = bytes.Replace(source, lineBreak, space, -1)

	// replace all doubles spaces with single spaces
	for bytes.Contains(source, doubleSpace) {
		source = bytes.Replace(source, doubleSpace, space, -1)
	}

	// add a linebreak
	source = append(source, 10)

	return source
}
