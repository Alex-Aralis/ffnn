package util

import (
	"math/rand"

	"gonum.org/v1/gonum/matrix/mat64"
)

// NewRandMatrix builds a randomely initallized matrix of size m by n.
func NewRandMatrix(m, n int) mat64.Matrix {
	backing := make([]float64, m*n)

	for i := range backing {
		backing[i] = rand.Float64()
	}

	return mat64.NewDense(m, n, backing)
}
