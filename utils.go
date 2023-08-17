package main

import (
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
	"golang.org/x/exp/constraints"
)

func mapToArray(keys []string, m map[string]string) []string {
	var res = []string{}
	for _, v := range keys {
		res = append(res, m[v])
	}
	return res
}

func Ask() bool {
	var str string
	fmt.Scan(&str)
	return str == "y" || str == "yes"
}

func AreSim(a string, b string) bool {
	l := levenshtein.ComputeDistance(a, b)
	m := len(a)
	if len(b) > m {
		m = len(b)
	}

	return (1.0-float64(l)/float64(m))*100.0 > 60
}

func lerp(a float64, b float64, t float64) float64 {
	return (1-t)*(a) + t*b
}

func lerpInt(a int, b int, t float64) int {
	return int(lerp(float64(a), float64(b), t))
}
func Max[T constraints.Ordered](a T, b T) T {
	if b > a {
		return b
	}
	return a
}

func rend(str string) string {
	return str
}

func trimLineEnd(str string) string {
	return strings.TrimSuffix(strings.TrimSuffix(str, "\n"), "\r")
}
