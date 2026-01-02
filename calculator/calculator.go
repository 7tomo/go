package main

import "fmt"

func main(){
	var choose int
	var a int
	var b int
	
	fmt.Println("Выберите операцию: 1 - Сложение, 2 - Вычитание, 3 - Деление, 4 - Умножение")
	fmt.Scanln(&choose)
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	
	switch choose{
		case 1:
			sum(a,b)
		
		case 2:
			substract(a,b)
		
		case 3:
			divide(a,b)
		
		case 4:
			multiply(a,b)
		
	}
}

func sum(a int, b int){
	var c = a + b;
	fmt.Println("Сумма чисел: ", c)
}

func substract(a int, b int){
	var c = a - b;
	fmt.Println("Вычитание:", c)
}

func divide(a int, b int){
	var c = a / b;
	fmt.Println("Деление:", c)
}

func multiply(a int, b int){
	var c = a * b;
	fmt.Println("Умножение:", c)
}

