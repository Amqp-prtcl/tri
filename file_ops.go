package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const sep = ";"

func LoadCSV(filename string) *File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	str, err := r.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	if str == "" {
		panic("cannot read header")
	}

	var file = &File{
		Headers:           strings.Split(trimLineEnd(str), sep),
		Rows:              []Row{},
		longestHeaderSize: 0,
	}

	var l = len(file.Headers[0])
	for _, head := range file.Headers[1:] {
		l = Max(l, len(head))
	}
	file.longestHeaderSize = l

	var run = true
	for run {
		str, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			run = false
		}
		if str == "" {
			continue
		}
		var row = Row{}
		var cells = strings.Split(trimLineEnd(str), sep)
		for i := range cells {
			row[file.Headers[i]] = cells[i]
		}
		file.Rows = append(file.Rows, row)
	}
	return file
}

func SaveCSV(file *File, filename string) {
	if filename == "" {
		panic(fmt.Errorf("empty output filename"))
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(strings.Join(file.Headers, sep))

	for _, row := range file.Rows {
		f.WriteString(strings.Join(mapToArray(file.Headers, row), sep))
	}
}

func SaveJSON(filename string, v interface{}) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}
