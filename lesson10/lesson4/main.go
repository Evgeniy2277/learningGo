package main

import (
	"fmt"
	"github.com/usergo/insertsort/insertsort"
)

func main() {

	Sl := []int{51, 12, 3, 2, 1, 55, 44}
	fmt.Println(Sl)
	sortSl := insertsort.InsertionSort(Sl)
	fmt.Println(sortSl)
}
