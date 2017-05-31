package util

import (
	"testing"

	"gonum.org/v1/gonum/matrix/mat64"
)

func TestLogisticActivation(t *testing.T) {
	backing := []float64{0.4, 4, 2}

	net := mat64.NewDense(1, len(backing), backing)

	o := LogisticActivation(net)

	for i, val := range backing {
		if Logistic(val) != o.At(0, i) {
			t.Error("Logistic funciton not being correctly calculated.")
		}
	}
}
