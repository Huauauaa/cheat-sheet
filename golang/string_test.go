package golang

import (
	"fmt"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	s := "你好, World!"

	for _, v := range s {
		fmt.Printf("%v(%v)", v, string(v))
	}
	fmt.Println()
	b, _ := strconv.ParseBool("false")
	fmt.Printf("%T %v", b, b)
}
