package model

import (
	"github.com/alex.aralis/ffnn/util"
	"gonum.org/v1/gonum/matrix/mat64"
)

// FFNN is the interface all forward feed
// neural nets need to implement
type FFNN interface {
	Layers() []*mat64.Dense
	Activation(float64) float64
	DiffAct(float64) float64
	LearningRate() float64
}

// Eval the FFNN with the value x
func Eval(nn FFNN, x *mat64.Dense) *mat64.Dense {
	layers := nn.Layers()

	for _, m := range layers {
		var y mat64.Dense

		y.Mul(m, x)
		x = util.LogisticActivation(&y)
	}

	return x
}

// 
func BackPropogate(nn FFNN, y, t) {

}
