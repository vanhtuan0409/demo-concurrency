package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("Hello from the 1st loop")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		fmt.Println("Hello from the 2nd loop")
		time.Sleep(100 * time.Millisecond)
	}
}
