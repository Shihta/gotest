//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Hello, World!");
//}

package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func myfunc() int {
	fmt.Println("this is myfunc")
	return 10
}

func main() {
	fmt.Println("Hello, World111!")
	fmt.Println("asdfijij")
	a := myfunc()
	fmt.Println(string(a) + " asdasda")
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}
