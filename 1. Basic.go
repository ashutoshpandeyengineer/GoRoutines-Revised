package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(10 * time.Millisecond)

}
func hello() {
	fmt.Println("hello world")
}
