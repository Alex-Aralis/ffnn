package main

import (
	"fmt"

	"github.com/alex.aralis/ffnn/util"
)

func main() {
	fmt.Printf("hello world.")

	a := util.NewRandMatrix(4, 2)

	b := util.FMap(a)

	fmt.Println(a)
	fmt.Println(b)
}
