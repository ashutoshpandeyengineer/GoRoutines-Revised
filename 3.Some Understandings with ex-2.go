package main

import (
	"fmt"
	"time"
)

// 1st main function m pehle  toh main function ka goroutine  chlegaa ...
//jab main function ka gouroutine  anonymous function k goroutine p pahunchega tabb wo encounter nahi krega.
//  kyuu ki abhi uss wlle go routine p  time nhi hai .
// Jab time p pahunch jaega firr anonymous goroutine ko execute karne ki kosis  karega.
//uss time tak value n ki value Good bye  hogayi rahegii fir goodbye print kar dega ..

//func main() {
//	var n = "Hello"
//	go func() {
//		fmt.Println(n)
//	}()
//	n = "goodbye"
//	time.Sleep(10 * time.Millisecond)
//}

func main() {
	var n = "Hello"
	go func(n string) {
		fmt.Println(n)
	}(n)
	n = "Goodbye"
	time.Sleep(10 * time.Millisecond)
}

// 2nd function  m pehle mainthread execute hooggaa fir raste m goroutine dekkeag are usme argument dekh k argument value le legga.
//ab jab time k pass pahunchegaa toh go routine execute karegaa .
//lekin argument to pehle hi le chukka h isiliye argument ki value " Hello" hi execute karega.
