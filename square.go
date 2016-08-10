package main

import (
	"strconv"
	"strings"
)

type Square [][]int

func MakeSquare(dim int) (sq Square) {
	mem := make([]int, dim*dim)
	sq = make([][]int, dim)
	for i := range sq {
		sq[i] = mem[i*dim : (i+1)*dim]
	}
	return
}

func (sq Square) Clear() {
	for _, row := range sq {
		for i := range row {
			row[i] = 0
		}
	}
}

func (sq Square) Fill(value int) {
	for _, row := range sq {
		for i := range row {
			row[i] = value
		}
	}
}

func (sq Square) SubSquare(i, j, dim int) (sub Square) {
	sub = make([][]int, dim)
	for k := range sub {
		sub[k] = sq[i+k][j : j+dim]
	}
	return
}

func (sq Square) Dim() int {
	return len(sq)
}

func (sq Square) String() string {
	dim := sq.Dim()
	lines := make([][]string, dim)
	widths := make([]int, dim)

	// first pass -- convert all numbers to strings
	for i := range lines {
		line := make([]string, dim)
		for j, n := range sq[i] {
			s := strconv.Itoa(n)
			line[j] = s
			if widths[j] < len(s) {
				widths[j] = len(s)
			}
		}
		lines[i] = line
	}

	// second pass -- convert every line from []string to string
	strLines := make([]string, dim)
	for i, line := range lines {
		// pad if necessary
		for j, s := range line {
			pad := widths[j] - len(s)
			if pad > 0 {
				line[j] = strings.Repeat(" ", pad) + s
			}
		}
		strLines[i] = strings.Join(lines[i], " ")
	}

	return strings.Join(strLines, "\n")
}
