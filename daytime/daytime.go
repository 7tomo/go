package main

import "fmt"

var time int

func main() {
	fmt.Println("Введите время:")
	fmt.Scanln(&time)

	switch time {
	case 6, 7, 8, 9, 10, 11:
		fmt.Printf("Сейчас %d ч. - утро", time)
	case 12, 13, 14, 15, 16, 17:
		fmt.Printf("Сейчас %d ч. - день", time)
	case 18, 19, 20, 21, 22:
		fmt.Printf("Сейчас %d ч. - вечер", time)
	case 23, 0, 1, 2, 3, 4, 5:
		fmt.Printf("Сейчас %d ч. - ночь", time)
	default:
		fmt.Println("Неверно задано время")
	}

}
	