package main

// Goroutines: Lightweight threads (non-blocking)

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing Task:",id)
}

func main(){
	for i:=0;i<=10;i++{
		go task(i)
		// inline anonymous fx
		go func (i int)  {
			fmt.Println("Internal GoRoutine: ",i)
		}(i)
	}

	time.Sleep(time.Second*2) // pasuing it for 2 seconds (simulation)
}