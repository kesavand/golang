/*
sample file to test the logger
*/
package main

import (
	"github.com/kesavand/golang/pkg/log"
	"fmt"
)


func main() {
	fmt.Println("Ths exmple is used to verify the logging library")
	logger,err := log.NewLogger("zap");
	fmt.Println("NewLogger",err)

	if err != nil {
		fmt.Println("Unable to create logger",err)
		return
	}

	logger.Debug("The log librry is success")

}
