package main

import (
	"fmt"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return b
	}
}

func main() {

	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil  {
		panic(err)
	}
	f := fibonacci()
	mapFib := map[int]int{}
	for i := 1; i <= n; i++ {
		fib := f()
		if fib > 0 {
			mapFib[i] = fib
		} else {
			fmt.Println(n, "too big")
			break
		}

		}

	if n > 2 {fmt.Println( "fibonacci", n, " =", mapFib[n])
	} else if (n == 2 || n == 1)  {
		fmt.Println("fibonacci", n, " = 1")
	} else if n == 0 {
		fmt.Println("fibonacci ", n, " =", 0)
	} else {
		fmt.Println(n, " negative")
	}

}
