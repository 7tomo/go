package main

import (
	"errors"
	"fmt"
)

func main() {
	sum(2, 10)
	sum(8, 8)
	sum(10, 10)
	sum(11, 13)
	sum(5, 7)
}

func sum(a int, b int) (int, error) {

	var c = a + b
	if c != int(c) {
		return 0, errors.New("результат не целочисленный")
	}
	fmt.Println("Сумма чисел: ", c)
	return c, nil
}
