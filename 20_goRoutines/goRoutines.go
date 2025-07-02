package main

// Goroutines: Lightweight threads (non-blocking)

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing Task:",id)
}

func main(){
	var wg sync.WaitGroup
	for i:=0;i<=10;i++{
		wg.Add(1)
		go task(i,&wg)
		//inline anonymous fx
		// go func (i int)  {
		// 	fmt.Println("Internal GoRoutine: ",i)
		// }(i)
	}

	wg.Wait()

	//time.Sleep(time.Second*2) // pausing it for 2 seconds (simulation)
}

//! NOTE:
// A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls [WaitGroup.Add] to set the number of goroutines to wait for. Then each of the goroutines runs and calls [WaitGroup.Done] when finished. At the same time, [WaitGroup.Wait] can be used to block until all goroutines have finished.

// A WaitGroup must not be copied after first use.

// In the terminology of the Go memory model, a call to [WaitGroup.Done] “synchronizes before” the return of any Wait call that it unblocks.

// func (wg *sync.WaitGroup) Add(delta int)
// func (wg *sync.WaitGroup) Done()
// func (wg *sync.WaitGroup) Wait()


//Output:
// Doing Task: 1
// Doing Task: 10
// Doing Task: 6
// Doing Task: 7
// Doing Task: 8
// Doing Task: 9
// Doing Task: 3
// Doing Task: 5
// Doing Task: 2
// Doing Task: 0
// Doing Task: 4