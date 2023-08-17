package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

var (
	stringFormat = Bold() + "'%s'" + ClearFormat()

	askFormat  = "found potential matches for " + stringFormat + " of brand " + stringFormat + " (%d occu.):\n\n"
	propFormat = "%s%d. " + stringFormat + " of brand " + stringFormat + " (%d occu.)\n"
)

func Test() {
	//terminal.ToggleEcho()
	//terminal.ToggleCanonicalMode()
}

func Printbanner(progress int, total int) {
	const bannerSize = 90
	var percentage = float64(progress) / float64(total) * 100
	var l = int(math.Log10(float64(total)) + 1)

	var vals = fmt.Sprintf(" (%"+fmt.Sprint(l)+"d / %d) %5.2f%% ", progress, total, percentage)
	var pad = strings.Repeat("=", (bannerSize-len(vals))/2)

	fmt.Println(LerpColor(DarkRed, DarkGreen, percentage/100) + pad + vals + pad + ClearFormat())
}

func HandleSims(file *File, sims map[string]*Sim) {
	var replaced = map[string]bool{}
	var ar = []*Sim{}
	for _, v := range sims {
		if len(v.Similar) != 0 {
			ar = append(ar, v)
		}
	}

	for i, v := range ar {
		if replaced[v.Name] {
			continue
		}
		ClearScreen()
		Printbanner(i, len(ar))
		fmt.Println()
		printSimBody(file, v, replaced)
		i++
	}
	ClearScreenHard()
}

func printSimBody(file *File, sim *Sim, replaced map[string]bool) {
	const tab = "    "
	var i int = 0
	var resp int
	var m []string

dofunc:
	fmt.Printf(askFormat, sim.Name, sim.Brand, sim.Occurences)

	m = make([]string, len(sim.Similar))
	i = 0
	for k, v := range sim.Similar {
		m[i] = k
		fmt.Printf(propFormat, tab, i+1, k, v.brand, v.occ)
		i++
	}
	//fmt.Scanln(&d)
	fmt.Printf("\nUse 0 to ignore value.\nYour choice (0 to %d): ", len(sim.Similar))
	resp = -1
retry:
	fmt.Scan(&resp)
	if resp < 0 || resp > len(sim.Similar) {
		fmt.Printf("\nPlease enter a valid number in the range 0 to %d: ", len(sim.Similar))
		goto retry
	}
	if resp == 0 {
		return
	}

	if o := file.ReplaceAllInCol(modIndex, m[resp-1], sim.Name); o != sim.Similar[m[resp-1]].occ {
		//panic(fmt.Errorf("not found expected amount of occurences (expected: %d, got: %d)", o, sim.Similar[m[resp-1]].occ))
		fmt.Printf("\n"+Red.Format()+"not found expected amount of occurences (expected: %d, got: %d)\n"+ClearFormat(), o, sim.Similar[m[resp-1]].occ)
		time.Sleep(5 * time.Second)
	}
	replaced[m[resp-1]] = true
	delete(sim.Similar, m[resp-1])
	if len(sim.Similar) != 0 {
		Goto(3, 1)
		ClearScreenAfterCursor()
		goto dofunc
	}
}

func HandleMatches(file *File, m []MatchRes) {
	for i, v := range m {
		ClearScreen()
		Printbanner(i, len(m))
		fmt.Println()
		printMatchBody(file, v)
		i++
	}
	ClearScreenHard()
}

func printMatchBody(file *File, m MatchRes) {
	const tab = "    "
	fmt.Printf("line: %d\n", m.Index+2)
	for _, head := range file.Headers {
		fmt.Printf("%s%"+fmt.Sprintf("%d", file.longestHeaderSize)+"s: "+stringFormat+"\n", tab, rend(head), m.Row[head])
	}
	fmt.Scanln()
}

// 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

/*

================  ( 44 / 243) 00.18%  ================

found potential matches for 'bablabla', brand: (1 occu.)

It can be corrected to:
	1. (48 occu.) "blablabla"
	2. (2 occu.) "blabbla"
	3. (1 occu.) "blablbl"

use -1 to not change anything.
your choice (-1 to 3):
*/

/*

================  ( 44 / 243) 00.18%  ================

line 423 matched request:
	Numero de serie :
	Marque :
*/
