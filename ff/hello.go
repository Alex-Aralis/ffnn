package main

import (
	"fmt"

	"gonum.org/v1/gonum/matrix/mat64"

	"github.com/alex.aralis/ffnn/util"
)

func main() {
	fmt.Printf("hello world.")

	a := util.NewRandMatrix(4, 2)

	var b mat64.Dense

	b.Apply(
		func(i, j int, v float64) float64 {
			return float64(i+j) + v
		},
		a,
	)

	fmt.Println(a)
	fmt.Println(b)
}
