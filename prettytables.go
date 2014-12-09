package prettytables

import (
	"fmt"
	"strings"
)

type Tableable interface {
	Table() [][]string
}

type Table [][]string

func (t Table) Table() [][]string {
	return t
}

// PrintTable takes a slice of slices and "Pretty Prints" it in table form.
// The length of the first slice is assumed to be the length of the rest and
// serves as the header
func PrintTable(obj Tableable) {
	tab := obj.Table()
	if tab == nil || len(tab) < 1 {
		return
	}
	numFields := len(tab[0])

	// max length for each field
	fieldLengths := make([]int, numFields)
	for _, v := range tab {
		for j, w := range v {
			split := strings.Split(w, "\n")
			groupMax := 0
			for _, y := range split {
				if len(y) > groupMax {
					groupMax = len(y)
				}
			}
			if groupMax > fieldLengths[j] {
				fieldLengths[j] = groupMax
			}
		}
	}

	PrintSep(fieldLengths)
	PrintFields(fieldLengths, tab[0])
	PrintSep(fieldLengths)
	for _, v := range tab[1:] {
		PrintFields(fieldLengths, v)
		//PrintSep(fieldLengths)
	}
	PrintSep(fieldLengths)
}

func PrintSep(fieldLengths []int) {
	for _, v := range fieldLengths {
		fmt.Print("+")
		for i := 0; i < v+2; i++ {
			fmt.Print("-")
		}
	}
	fmt.Print("+\n")
}

func PrintFields(fieldLengths []int, fields []string) {
	maxLines := 0
	for _, v := range fields {
		lines := len(strings.Split(v, "\n"))
		if lines > maxLines {
			maxLines = lines
		}
	}

	multiLine := make([][]string, maxLines)
	for i, _ := range multiLine {
		multiLine[i] = make([]string, len(fields))
	}

	for i, v := range fields {
		lines := strings.Split(v, "\n")
		for j, w := range lines {
			multiLine[j][i] = w
		}
	}

	for _, line := range multiLine {
		for i, v := range fieldLengths {
			fmt.Print("| ")
			padding := v - len(line[i])
			fmt.Print(line[i])
			for i := 0; i < padding+1; i++ {
				fmt.Print(" ")
			}
		}
		fmt.Print("|\n")
	}
}

func FromMaps(fields []string, tabMap []map[string]string) [][]string {
	tab := make([][]string, len(tabMap)+1)
	tab[0] = make([]string, len(fields))
	for i, v := range fields {
		tab[0][i] = strings.Replace(strings.Title(v), "_", " ", -1)
		for j, w := range tabMap {
			tab[j][i] = w[v]
		}

	}

	return tab
}
