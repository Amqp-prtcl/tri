package main

const (
	brandIndex = 1
	libIndex   = 4
	modIndex   = 5
)

var (
	InputFilename  = `/Users/temp/Desktop/tri/db sort (CSV Backup)/csv/dialyse explo.csv`
	OutputFilename = `test.csv`

	file *File
)

func main() {
	//Test()
	//return
	ClearScreenHard()
	file = LoadCSV(InputFilename)
	//MatchAndAskCol(file, modIndex)
	//ExtractAndSaveLibMod(file)
	//ExtractSimAndSaveLibMod(file)
	ExtractSimAndAskMod(file)
	//file.Save(outputFilename)
}
func ExtractAndSaveLibMod(file *File) {
	var dictMod = ExtractMod(file)
	//var dictLib = ExtractMod(file, 4)
	//SaveJSON("libelle.json", SortMap(dictLib))
	SaveJSON("modele.json", SortMap(dictMod))
}

func ExtractMod(f *File) map[string]*struct {
	brand string
	occ   int
} {
	var ret = map[string]*struct {
		brand string
		occ   int
	}{}

	var modKey = f.ColToKey(modIndex)
	var brandKey = f.ColToKey(brandIndex)
	file.Foreach(func(i int, row Row) {
		v, ok := ret[row[modKey]]
		if ok && v.brand == row[brandKey] {
			//
			v.occ++
			return
		}
		v = &struct {
			brand string
			occ   int
		}{
			brand: row[brandKey],
			occ:   1,
		}
		ret[row[modKey]] = v
	})
	return ret
}

type Sim struct {
	Name       string
	Brand      string
	Occurences int
	Similar    map[string]*struct {
		brand string
		occ   int
	}

	matched bool
}

func ExtractSimAndSaveMod(file *File) {
	var dictMod = ExtractMod(file)
	var mods = SortMap(dictMod)
	var m = SimMod(file, mods)
	SaveJSON("test.json", m)
}

func ExtractSimAndAskMod(file *File) {
	var dictMod = ExtractMod(file)
	var mods = SortMap(dictMod)
	var m = SimMod(file, mods)
	HandleSims(file, m) // perfomrs the asking

	file.Save(OutputFilename)
}

func MatchAndAskCol(file *File, col int) {
	m := GetAllMatchIndexCol(file, col, Matches[0])
	HandleMatches(file, m)
}

func SimMod(file *File, mods MapSort) map[string]*Sim {
	var m = map[string]*Sim{}

	for i := len(mods) - 1; i >= 0; i-- {
		m[mods[i].Name] = &Sim{
			Name:       mods[i].Name,
			Brand:      mods[i].Brand,
			Occurences: mods[i].Occurences,
			Similar: map[string]*struct {
				brand string
				occ   int
			}{},
			matched: false,
		}
	}
	for i := len(mods) - 1; i >= 0; i-- {
		var modi = mods[i]
		if m[modi.Name].matched {
			//delete(m, sDitcMod[i].Name)
			continue
		}
		for i2 := len(mods) - 1; i2 >= 0; i2-- {
			var modi2 = mods[i2]
			if i == i2 {
				continue
			}
			if AreSim(modi.Name, modi2.Name) {
				sim, ok := m[modi.Name].Similar[modi2.Name]
				if !ok {
					sim = &struct {
						brand string
						occ   int
					}{}
					m[modi.Name].Similar[modi2.Name] = sim
				}
				sim.occ = modi2.Occurences
				sim.brand = modi2.Brand
				if modi.Occurences > modi2.Occurences {
					m[modi2.Name].matched = true
				}
			}
		}
	}

	return m
}
