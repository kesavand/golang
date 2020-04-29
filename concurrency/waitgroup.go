package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1) // add to the waitgroup before invoking the go routine
	go func() {
		fmt.Println("Hello World 1")
		wg.Done() // siganl the completion of the go routine
	}()
	wg.Add(1)
	go  func(){
		fmt.Println("Hello World 2")
		wg.Done() // siganl the completion of the go routine
	}()
	wg.Wait()
	fmt.Println("Finished the main go routine")
}
