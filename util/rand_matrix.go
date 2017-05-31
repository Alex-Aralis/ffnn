package util

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/matrix/mat64"
)

var source *rand.Rand

// NewRandMatrix builds a randomely initallized matrix of size m by n.
func NewRandMatrix(m, n int) *mat64.Dense {
	if source == nil {
		rg := rand.NewSource(time.Now().UnixNano())

		source = rand.New(rg)
	}

	backing := make([]float64, m*n)

	for i := range backing {
		backing[i] = source.Float64()
	}

	return mat64.NewDense(m, n, backing)
}
