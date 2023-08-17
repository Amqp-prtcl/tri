package main

type File struct {
	Headers           []string
	Rows              []Row
	longestHeaderSize int
}

type Row map[string]string

func (f *File) Save(filename string) {
	SaveCSV(f, filename)
}

func (f *File) ColToKey(col int) string {
	return f.Headers[col]
}

func (f *File) Foreach(fu func(int, Row)) {
	for i := 0; i < len(f.Rows); i++ {
		fu(i, f.Rows[i])
	}
}

func (f *File) ReplaceAllInCol(col int, old string, new string) int {
	var key = file.Headers[col]
	var ret int
	for _, line := range f.Rows {
		if line[key] == old {
			line[key] = new
			ret++
		}
	}
	return ret
}
