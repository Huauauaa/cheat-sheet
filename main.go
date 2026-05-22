package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	n := 10
	fmt.Printf("%d\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	a := 8.2
	b := 3.8
	fmt.Println(decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)), a-b)
}
