package util

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/matrix/mat64"
)

func TestEvalNN(t *testing.T) {
	layers := []*mat64.Dense{
		makeOnesMatrix(3, 4),
		makeOnesMatrix(3, 3),
		makeOnesMatrix(2, 3),
	}

	x := mat64.NewDense(4, 1, []float64{2, 1, 2, 3})

	y := EvalNN(x, layers)

	fmt.Println(y)
}

func makeOnesMatrix(r, c int) *mat64.Dense {
	m := mat64.NewDense(r, c, nil)

	m.Apply(
		func(i, j int, val float64) float64 {
			return 1
		},
		m,
	)

	return m
}
