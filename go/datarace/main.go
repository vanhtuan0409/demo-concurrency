package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	useMutexLock := false

	for i := 0; i < 20; i++ {
		dataRaceTest(useMutexLock)
	}
}

func dataRaceTest(useLock bool) {
	x := 0
	var wg sync.WaitGroup // Initialize WaitGroup
	var lock *sync.Mutex  // Initialize mutext lock
	if useLock {
		lock = &sync.Mutex{}
	}
	wg.Add(2)

	// Running 1st goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			increase(&x, lock)
		}
	}()

	// Running 2nd goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			increase(&x, lock)
		}
	}()

	// Wait for both goroutine to finish
	wg.Wait()
	fmt.Println(x)
}

func increase(target *int, lock *sync.Mutex) {
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	*target = *target + 1
}
