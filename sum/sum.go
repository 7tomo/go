package main

import "fmt"

func main() {
	sum(2, 10)
	sum(8, 8)
	sum(10, 10)
	sum(11, 13)
}

func sum(a int, b int) {
	var c = a + b
	fmt.Println("Сумма чисел: ", c)
}
