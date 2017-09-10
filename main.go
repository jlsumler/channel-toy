package main

import (
	"fmt"
	"sync"
)

var requests chan int64

const (
	// maxRequests is the max size of our thread pool, e.g. max number of worker
	// threads running at any one time
	maxRequests = 4
	// reqLimit is the size of the requests buffered channel
	reqLimit = 10
	// bigLimit is the number of requests we will create in the main function
	bigLimit = 10000
)

func main() {
	requests = make(chan int64, reqLimit)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go serve(requests, &wg)
	for i := 0; i < bigLimit; i++ {
		//fmt.Println("Requests: ", len(requests))
		requests <- int64(i)
	}
	close(requests)
	//fmt.Println("Waiting")
	wg.Wait()
	//fmt.Println("Done")
}

func serve(r chan int64, w *sync.WaitGroup) {
	gate := make(chan int, maxRequests)
	kids := sync.WaitGroup{}
	for req := range r {
		r := req
		//fmt.Println("Gate: ", len(gate))
		gate <- 1
		kids.Add(1)
		go process(r, gate, &kids)
	}
	// let any remaining process goroutines drain
	kids.Wait()
	// decrement the waitgroup so the main thread can exit
	w.Done()
}

func process(j int64, g chan int, k *sync.WaitGroup) {
	defer func() {
		k.Done()
		<-g
	}()

	fmt.Println(j)

}
