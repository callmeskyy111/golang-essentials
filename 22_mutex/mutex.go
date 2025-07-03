package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) incView(wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock() //before the modification
	p.views++
	
}

func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i:=0;i<100;i++{
		wg.Add(1)
		go myPost.incView(&wg)
	}

	// myPost.incView()
	// myPost.incView()

	wg.Wait()
	fmt.Println(myPost.views)

}