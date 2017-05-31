package util

import (
	"math"

	"gonum.org/v1/gonum/matrix/mat64"
)

// LogisticActivation applies a logistic activation function to every entry
// memeber of the povided vector.
func LogisticActivation(net *mat64.Dense) *mat64.Dense {
	var o mat64.Dense

	o.Apply(
		func(i, j int, val float64) float64 {
			return Logistic(val)
		},
		net,
	)

	return &o
}

// Logistic returns the standard logistic function evaluated at x
func Logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
