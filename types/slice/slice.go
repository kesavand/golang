/*
This example
1. example1 declares a slice of integers and prints the elements in the slice
*/

package main

import (
	"fmt"
	"go.uber.org/zap"
	"errors"
)

var x int = 10

func main() {
	example1()
}

func example1() {
	var ia []int
	ia = make([]int, 5)

	for i, _ := range ia {
		ia[i] = i
	}
	for i, _ := range ia {
		fmt.Printf("The value of a[%d] is %d\n", i, ia[i])
	}
}
