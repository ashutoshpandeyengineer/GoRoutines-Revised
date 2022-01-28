package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(5)
	fmt.Printf("Threads : %v\n", runtime.GOMAXPROCS(-1))    //used to change the threads in the OS.....by default 8 threads are their in my Pc.
}
