package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите название товара")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	keyboard:= "Клавиатура JZ9"
	headphones:= "Наушники N45"
	phone:= "Смартфон S10"

	keyboardLow := strings.ToLower(keyboard)
	headphonesLow := strings.ToLower(headphones)
	phoneLow := strings.ToLower(phone)

	input := scanner.Text()
	inputLow := strings.ToLower(input)

	if strings.Contains(keyboardLow, inputLow) {
		fmt.Println("Клавиатура JZ9: 19200")
	} else if strings.Contains(headphonesLow, inputLow) {
		fmt.Println("Наушники N45: 9600")
	} else if strings.Contains(phoneLow, inputLow) {
		fmt.Println("Смартфон S10: 55000")

	} else {
		fmt.Printf("Товар \"%s\" не найден.", input)
	}

}
