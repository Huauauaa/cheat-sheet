package golang

import (
	"fmt"
	"testing"
)

func TestGolang(t *testing.T) {
	fmt.Println(1)
	fmt.Println(2)
	var a = 1
	var b float64 = 2
	// fmt.Println(a + b)
	fmt.Println(a + 2)

	fmt.Printf("%v %T", 2, 2)

	if float64(a) > b {
		fmt.Println("a > b")
	}
}
