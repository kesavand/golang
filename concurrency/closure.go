/*
This example shows the anonymous function and the closure

1. Anonymous functions are the functions that doesn't have a name
2. cloures are functions that closes/retains  the conext of the parent function
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("anonymous function",anonymous())
	fmt.Println("anonymous function",anonymous())
	fmt.Println("anonymous function",anonymous())
	x := closure()
	fmt.Println("Closure Function ",x())
	fmt.Println("Closure Function ",x())
	fmt.Println("Closure Function ",x())

}
func anonymous() (int){
	var x int
	//Below is the anonymous function
	func() {
		x++
	}()
	return x
}

func closure() (func()(int)) {
	var x int
	increment := func() (int){
		x++
		return x
	}
	return increment
}
