package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InsertSort(A []int64) []int64 {
	//копируем полученный функцией слайс
	slice := make([]int64, len(A))
	copy(slice, A)
	// сортитуем простой вставкой
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		lst := i
		for j := i - 1; j > -1; j-- {
			if slice[j] < key {break}
			slice[j+1] = slice[j]
			lst = j
		}
		slice[lst] = key
	}
	return slice
}

func main() {
	//читаем ввод в inputNums
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}
	//получаем отсортированный ввод вызовом функции
	sortInputNums := InsertSort(inputNums)

	for i := 0; i < len(sortInputNums); i++ {
		fmt.Println(sortInputNums[i])
	}
}

