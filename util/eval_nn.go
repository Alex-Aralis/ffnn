package util

import "gonum.org/v1/gonum/matrix/mat64"

// EvalNN takes a vector y of input values and
// runs it through the neral net represented by
// layers.
func EvalNN(y *mat64.Dense, layers []*mat64.Dense) *mat64.Dense {
	for _, m := range layers {
		y.Mul(m.T(), y)
	}

	return y
}
