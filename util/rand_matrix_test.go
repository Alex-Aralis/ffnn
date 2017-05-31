package util

import (
	"testing"

	"gonum.org/v1/gonum/matrix/mat64"

	"github.com/bradfitz/iter"
)

func TestNewRandMatrix(t *testing.T) {
	tc, tr := 4, 3

	m := NewRandMatrix(tc, tr)

	c, r := m.Dims()

	if c != tc || r != tr {
		t.Errorf("Generated matrix of wrong dimensions, expected %v x %v and go %v x %v.", tc, tr, c, r)
		return
	}

	for i := range iter.N(c) {
		for j := range iter.N(r) {
			if m.At(i, j) < 0 || m.At(i, j) > 1 {
				t.Errorf("Entry (%v, %v) = %v not contained by [0, 1].", i, j, m.At(i, j))
			}
		}
	}

	if mat64.Equal(m, NewRandMatrix(tc, tr)) {
		t.Errorf("Second generated matrix equal to the first. both values are %v", m)
	}
}
