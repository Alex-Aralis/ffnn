package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/matrix/mat64"

	"github.com/alex.aralis/ffnn/util"
)

func main() {
	fmt.Printf("hello world.")

	a := util.NewRandMatrix(4, 2)

	var b mat64.Dense

	b.Apply(
		rand.Float64,
		a,
	)

	fmt.Println(a)
	fmt.Println(b)
}
