package util

import "gonum.org/v1/gonum/matrix/mat64"

// Shape is an array of matricies used to
// specify the weights of an ffnn.
type Shape []*mat64.Dense

// NewShape Makes a Shape that has appropriate
// layer outputs based on the input int array.
// The layers[0] is the input size, and layers[len(layers)]
// is the output size.
func NewShape(layers []int) Shape {
	shape := Shape{}
	prevSize := layers[0]
	layers = layers[1:]

	for _, size := range layers {
		shape = append(shape, NewRandMatrix(size, prevSize))
	}

	return shape
}
