package main

import "fmt"

const (
	esc = "\x1b"
	csi = esc + "["
	m   = "m"
)

type Color struct {
	R int
	G int
	B int
}

func (c Color) Format() string {
	return Rgbc(c)
}

var (
	Red   = Color{255, 0, 0}
	Green = Color{0, 255, 0}

	DarkRed   = Color{204, 0, 0}
	DarkGreen = Color{78, 154, 6}
)

func ClearScreenAfterCursor() {
	fmt.Print(csi + "J")
}

func ClearScreen() {
	fmt.Print(csi + "H" + csi + "2J")
}

func ClearScreenHard() {
	fmt.Print(csi + "H" + csi + "3J")
}

func Goto(row int, col int) {
	fmt.Printf("%s%d;%dH", csi, row, col)
}

func ClearFormat() string {
	return csi + m
}

func Bold() string {
	return csi + "1" + m
}

func Rgb(r int, g int, b int) string {
	return csi + fmt.Sprintf("38;2;%d;%d;%d", r, g, b) + m
}

func Rgbc(c Color) string {
	return csi + fmt.Sprintf("38;2;%d;%d;%d", c.R, c.G, c.B) + m
}

func LerpColor(a Color, b Color, t float64) string {
	return Rgbc(Color{
		R: lerpInt(a.R, b.R, t),
		G: lerpInt(a.G, b.G, t),
		B: lerpInt(a.B, b.B, t),
	})
}
