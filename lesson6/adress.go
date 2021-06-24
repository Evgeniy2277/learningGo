package main

import "fmt"

func main() {

	n := 1
	mapa := map[string]int{
		"one": 1,
		"two": 2,
	}
	slice := []int{1, 2}
	an := &n
	amapa := &mapa
	amapaOne := mapa["one"]
	amapaTwo := mapa["two"]
	aslice := &slice
	aslice0 := &slice[0]
	aslice1 := &slice[1]
	fmt.Println(n)
	fmt.Println(an)
	fmt.Println(*an)
	fmt.Println(mapa)
	fmt.Println(amapa)
	fmt.Println(amapaOne)
	fmt.Println(amapaTwo)
	fmt.Println(slice)
	fmt.Println(aslice)
	fmt.Println(*aslice)
	fmt.Println(aslice0)
	fmt.Println(*aslice0)
	fmt.Println(aslice1)
	fmt.Println(*aslice1)
}
