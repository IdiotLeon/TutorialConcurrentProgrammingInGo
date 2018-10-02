package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
 *	However, the data is mostly ordered,
 *	hinting the process is no long running in parallel,
 *	as much as it might first appear.
 *	This is caused by the locking and unlocking.
 *	As a matter of fact,
 *	this application actually performs more slowly than a single threaded model
 *	due to all of the synchronization.
 */
func main() {
	mutex := new(sync.Mutex)
	runtime.GOMAXPROCS(4)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// To obtain the lock on the mutex
			mutex.Lock()
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				// To unlock the mutext allowing the next goroutine to be processed				mutex.Unlock()
				mutex.Unlock()
			}()
		}
	}
}
