package main

import (
	"fmt"

	"github.com/alex.aralis/ffnn/util"
)

func main() {
	fmt.Printf("hello world.")

	a := util.NewRandMatrix(4, 2)

	fmt.Print(a)
}
