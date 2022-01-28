package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var m = "Good Morning"
	wg.Add(1)
	go func(m string) {
		fmt.Println(m)
		wg.Done()
	}(m)
	m = "Good Afternoon"
	wg.Wait()
}

//time.Sleep(10 * time.Millisecond)---->  ye wlla Function isliye remove karna pada kyuki ye function real time m time le rhaa..
// means agar hamara programme kam time le raha hai tohh baaki ka time toh waste hogaa naaa.
//Isilye sync.waitgroup{} ka use karenge ...isme sbse pehle main ek goroutine m chlegaa ....jab main at a time wg.wait() par pahunchega
//lookout kregaa toh wg.add(1) dikhaai degga fir uske function ko execute karegaa.....wg.done() pahunchne k badd complete ho jaegaa..
//phir wg.add(1-1)...1 decrement ho jaegaa..
