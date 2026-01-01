package main

import "fmt"

func main() {
sum(2,10)
}

func sum(a int, b int){
	var c = a + b
	fmt.Println("Сумма чисел: ", c)
}