/*
This package is used to illustrate the channel in golang.
This file creates two channels of type int, creates one produer and one conusmer
the producer sends values 1 to 10 in the chan and receviver prints the values received in
the channel.
*/

package channel

import (
	"fmt"
)

//Producer Produces values from 1 to 10
func Producer(x chan<- int) {

	go func() {

		for i := 0; i < 10; i++ {
			x <- i
		}

	}()

}

//Consumer consumes the values received in the channel x
func Consumer(x <-chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Value received from the channel", x)
		}
	}()
}
