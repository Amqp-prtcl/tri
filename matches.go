package main

import "regexp"

var (
	Matches = []Match{
		{regexp.MustCompile(`..H..`), "any"},
	}
)

type Match struct {
	regex *regexp.Regexp
	value string
}

func (m *Match) Match(input string) (string, bool) {
	return m.value, m.regex.MatchString(input)
}

type MatchRes struct {
	Regex *regexp.Regexp
	Index int
	Col   int

	Row Row
}

func GetAllMatchIndexCol(file *File, col int, m Match) []MatchRes {
	var a = []MatchRes{}

	for index, row := range file.Rows {
		if _, ok := m.Match(row[file.ColToKey(col)]); ok {
			a = append(a, MatchRes{
				Regex: m.regex,
				Index: index,
				Col:   col,
				Row:   row,
			})
		}
	}

	funSort(a, func(i1, i2 int) bool {
		return a[i1].Index < a[i2].Index || (a[i1].Index == a[i2].Index && a[i1].Col < a[i2].Col)
	})
	return a
}
