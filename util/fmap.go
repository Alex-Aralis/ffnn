package util

import (
	"github.com/bradfitz/iter"
	"gonum.org/v1/gonum/matrix/mat64"
)

// FMap takes a matrix and a function and performes the function elementwise
// to produce a new matrix
func FMap(m mat64.Matrix) *mat64.Dense {
	r, c := m.Dims()

	backing := make([]float64, r*c)

	for i := range iter.N(r) {
		for j := range iter.N(c) {
			backing[i*c+j] = 2 * m.At(i, j)
		}
	}

	return mat64.NewDense(r, c, backing)
}
