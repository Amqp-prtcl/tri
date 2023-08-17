package main

import (
	"sort"
)

type tSortable interface {
	~[]any

	Less(i, j int) bool
}

type tSort[T tSortable] struct {
	sort T
}

func (a tSort[T]) Len() int           { return len(a.sort) }
func (a tSort[T]) Swap(i, j int)      { a.sort[i], a.sort[j] = a.sort[j], a.sort[i] }
func (a tSort[T]) Less(i, j int) bool { return a.sort.Less(i, j) }

func TypSort[T tSortable](s T) {
	var t = tSort[T]{s}
	sort.Sort(t)
}

type fSort[T any] struct {
	sort []T
	less func(int, int) bool
}

func (a fSort[T]) Len() int           { return len(a.sort) }
func (a fSort[T]) Swap(i, j int)      { a.sort[i], a.sort[j] = a.sort[j], a.sort[i] }
func (a fSort[T]) Less(i, j int) bool { return a.less(i, j) }

func funSort[T any](s []T, less func(int, int) bool) {
	var t = fSort[T]{s, less}
	sort.Sort(t)
}

type SortEntry struct {
	Name       string
	Brand      string
	Occurences int
}

type MapSort []SortEntry

func (a MapSort) Len() int           { return len(a) }
func (a MapSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MapSort) Less(i, j int) bool { return a[i].Occurences < a[j].Occurences }

func SortMap(m map[string]*struct {
	brand string
	occ   int
}) MapSort {
	var ar = MapSort{}
	for k, v := range m {
		ar = append(ar, SortEntry{k, v.brand, v.occ})
	}
	sort.Sort(ar)
	return ar
}
