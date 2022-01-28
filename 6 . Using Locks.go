package main

import (
	"fmt"
	"sync"
)

// to overcome the concept of racing we introduce the new concept called Locks..
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayhello()
		m.Lock()
		go increment()

	}
	wg.Wait()
}

func sayhello() {
	fmt.Printf("Hello world %v \n", counter)
	m.RUnlock()
	wg.Done()

}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
