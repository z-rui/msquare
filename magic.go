package main

// https://en.wikipedia.org/wiki/Siamese_method
func siamese(sq Square, offset int) {
	dim := sq.Dim()
	if dim%2 == 0 {
		panic("square is even")
	}
	sq.Clear()
	i, j := 0, dim/2
	for n := 1; n < dim*dim; n++ {
		sq[i][j] = n + offset
		i1, j1 := i, j

		if i1 == 0 {
			i1 = dim
		}
		i1--
		j1++
		if j1 == dim {
			j1 = 0
		}
		if sq[i1][j1] != 0 {
			i++
		} else {
			i, j = i1, j1
		}
		if sq[i][j] != 0 {
			panic("don't know what to do")
		}
	}
	sq[i][j] = dim*dim + offset
}

func FillOddSquare(sq Square) {
	siamese(sq, 0)
}

// https://en.wikipedia.org/wiki/Magic_square#A_method_of_constructing_a_magic_square_of_doubly_even_order
func FillDoublyEvenSquare(sq Square) {
	dim := sq.Dim()
	if dim%4 != 0 {
		panic("square is not doubly even")
	}
	sq.Clear()

	// first pass -- truth table
	subDim := dim / 4
	for i := 0; i < dim; i += subDim {
		for j := 0; j < dim; j += subDim {
			sub := sq.SubSquare(i, j, subDim)
			if i == j || i+j+subDim == dim {
				// on the diagonals
				sub.Fill(1)
			}
		}
	}

	// second pass -- fill the numbers
	inc, dec := 1, dim*dim
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if sq[i][j] == 1 {
				sq[i][j] = inc
			} else {
				sq[i][j] = dec
			}
			inc++
			dec--
		}
	}
}

// https://en.wikipedia.org/wiki/Strachey_method_for_magic_squares
func FillSinglyEvenSquare(sq Square) {
	dim := sq.Dim()
	if dim%4 != 2 {
		panic("square is not singly even")
	}

	subDim := dim / 2

	A := sq.SubSquare(0, 0, subDim)
	B := sq.SubSquare(subDim, subDim, subDim)
	C := sq.SubSquare(0, subDim, subDim)
	D := sq.SubSquare(subDim, 0, subDim)

	for i, sub := range [...]Square{A, B, C, D} {
		siamese(sub, i*subDim*subDim)
	}

	n := dim / 4

	// exchange the leftmost n columns in A and D
	for j := 0; j < n; j++ {
		for i := 0; i < subDim; i++ {
			A[i][j], D[i][j] = D[i][j], A[i][j]
		}
	}

	// exchange the rightmost n-1 columns in C and B
	for j := subDim - (n - 1); j < subDim; j++ {
		for i := 0; i < subDim; i++ {
			C[i][j], B[i][j] = B[i][j], C[i][j]
		}
	}

	// exchange the middle cell of the leftmost column of A and D
	// exchange the central cell in A and D
	mid := subDim / 2
	A[mid][0], D[mid][0] = D[mid][0], A[mid][0]
	A[mid][mid], D[mid][mid] = D[mid][mid], A[mid][mid]
}
