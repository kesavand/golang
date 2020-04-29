package main

import (
	"fmt"
	"sync"
	"time"
)

type mydata struct {
	lock sync.Mutex
	val  int
}

var x mydata

func main() {
	var wg sync.WaitGroup
		wg.Add(1)
	go func() {
		x.lock.Lock()
		x.val++
		x.lock.Unlock()
		wg.Done()

	}()
		wg.Wait()
	x.lock.Lock()
	fmt.Println("The value of x is ", x.val)
	x.lock.Unlock()
}
