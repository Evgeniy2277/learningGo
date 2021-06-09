package main

import (
	"fmt"
	"math"
)

func main() {
// вычисление площади прямоугольника
	var width int64
	var height int64
	fmt.Print("Введите ширину прямоугольника ")
	fmt.Scanln(&width)
	fmt.Print("Введите высоту прямоугольника ")
	fmt.Scanln(&height)
	fmt.Println("Площадь прямоугольника с шириной", width, "и высотой", height, "равна ", width * height)
// вычисление диаметра и длинны окружности по площади круга
	var area float64
	fmt.Print("введите площадь круга ")
	fmt.Scanln(&area)
	var d float64 = math.Sqrt(area/math.Pi) * 2
	var c float64 = d * math.Pi
	fmt.Println("диаметр круга с площадью ", area, "равен ", d, "длинна его окружности равна ", c)
// вывод сотен десятков и едениц введенного числа
	var number int64
	var hundreds int64
	var tens int64
	var unit int64

	fmt.Print("Введите трехзначное целое число ")
	fmt.Scanln(&number)
	hundreds = number / 100
	tens = (number - (hundreds * 100))  / 10
	unit = number - (hundreds * 100) - (tens * 10)
	fmt.Println("Вы ввели ", hundreds, "сотен", tens, "десятков и ", unit, "едениц")

}
