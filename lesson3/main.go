package main

import "fmt"

func main() {
	var a, b, r float64
	var op string

	fmt.Print("Input a: ")
	if _, err := fmt.Scanln(&a); err != nil {
		return
	}

	fmt.Print("Input b: ")
	if _, err := fmt.Scanln(&b); err != nil {
		return
	}

	fmt.Print("Input op: ")
	if _, err := fmt.Scanln(&op); err != nil {
		return
	}

	switch op {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		r = a / b
	default:
		fmt.Println(op, "?")
		return
	}

	fmt.Println(a, op, b, "=", r)

}
