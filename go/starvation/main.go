package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for {
			fmt.Println("Running parallel")
		}
	}()

	fmt.Println("Outside")
	for {
	}
}
