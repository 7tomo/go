package main

import (
	"fmt"
	"math"
)

var weight int
var height int

var imt float64

func main() {
	fmt.Print("Введите ваш вес (кг): ")
	fmt.Scanln(&weight)
	fmt.Print("Введите ваш рост: (см): ")
	fmt.Scanln(&height)

	var height_conv float64
	height_conv = float64(height) / 100

	imt = float64(weight) / (math.Pow(height_conv, 2))
	fmt.Println(imt)
	fmt.Printf("Ваш ИМТ: %.2f\n", imt)

	if imt < 18.5 {
		fmt.Println("Категория: Недостаточный вес")
	} else if imt >= 30 {
		fmt.Println("Категория: Ожирение")

	} else if 18.5 <= imt && imt < 25 {
		fmt.Println("Категория: Нормальный вес")
	} else {
		fmt.Println("Категория: Избыточный вес")
	}

}
