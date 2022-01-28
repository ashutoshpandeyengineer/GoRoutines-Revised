package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)      //wg.Add(2)----> denotes that 2 goroutines are  called at the same time ....
		go sayhello()  // directly calling them laeds the goroutines tom race against each other..
		go increment() // thats why providing irregular output
	}
	wg.Wait()
}

func sayhello() {
	fmt.Printf("Hello world %v\n", counter)
	wg.Done()
}
func increment() {
	counter++
	wg.Done()
}
