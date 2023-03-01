package math

import (
	"fmt"
	"testing"
)

func Testadd(T *testing.T) {
	result := add(1, 2)
	if result != 3 {
		fmt.Println("hello world")
	} else {
		fmt.Println("hi")
	}
}

func Testsub(T *testing.T) {
	result := sub(2, 1)
	if result != 1 {
		fmt.Println("hello world")
	} else {
		fmt.Println("hi")
	}
}
