package util

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	fmt.Println(IsEmpty(""))
	fmt.Println(IsEmpty(" "))
	fmt.Println(IsEmpty("\n"))
	fmt.Println(IsEmpty("\t"))

	fmt.Println(IsBlank(""))
	fmt.Println(IsBlank(" "))
	fmt.Println(IsBlank("\n"))
	fmt.Println(IsBlank("\t"))
}
