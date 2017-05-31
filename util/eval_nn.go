package util

import "gonum.org/v1/gonum/matrix/mat64"

// EvalNN takes a vector x of input values and
// runs it through the neral net represented by
// layers.
func EvalNN(x *mat64.Dense, layers []*mat64.Dense) *mat64.Dense {
	for _, m := range layers {
		var y mat64.Dense

		y.Mul(m.T(), x)
		x = LogisticActivation(&y)
	}

	return x
}
