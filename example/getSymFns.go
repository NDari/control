package main

import (
	"control"

	"github.com/NDari/gocrunch/mat"
)

func main() {
	m := control.New("olcao.skl")
	n := m.GetSymFns()
	mat.ToCSV(n, "SymFns.csv")
}
