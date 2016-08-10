package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func CheckMagicSquare(sq Square) bool {
	dim := sq.Dim()
	expectedSum := (dim*dim + 1) * dim / 2

	// rows
	for i := 0; i < dim; i++ {
		sum := 0
		for j := 0; j < dim; j++ {
			sum += sq[i][j]
		}
		if sum != expectedSum {
			return false
		}
	}

	// columns
	for j := 0; j < dim; j++ {
		sum := 0
		for i := 0; i < dim; i++ {
			sum += sq[i][j]
		}
		if sum != expectedSum {
			return false
		}
	}

	// diagonals
	sum1, sum2 := 0, 0
	for i := 0; i < dim; i++ {
		sum1 += sq[i][i]
		sum2 += sq[i][dim-i-1]
	}
	if sum1 != expectedSum || sum2 != expectedSum {
		return false
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s n\n\tgenerate an n × n magic square.\n", os.Args[0])
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	if n <= 0 {
		log.Fatal("n must be positive")
	}
	if n == 2 {
		log.Fatal("cannot generate a 2×2 magic square")
	}
	sq := MakeSquare(n)
	switch sq.Dim() % 4 {
	case 0:
		FillDoublyEvenSquare(sq)
	case 2:
		FillSinglyEvenSquare(sq)
	case 1, 3:
		FillOddSquare(sq)
	}
	fmt.Println(sq)
	if !CheckMagicSquare(sq) {
		panic("not a magic square")
	}
}
