package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(10 * time.Millisecond)

}

// Main function ka bhi ek apna thread h .sbse pehle wo main function wlla thread execute   hogaa.......wo encounter to krege lekin jbtak time allotment milega nahi tabtak wo kaam nhi kreagaa..
