package main

import (
	"fmt"
	"github.com/usergo/fibonacci/fib"
)


func main() {
	n := 10
	fibo := fib.FibCache(n)
	fibor := fib.FibRec(n)
	fmt.Println(fibo)
	fmt.Println(fibor)
}
