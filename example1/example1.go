package main

import "fmt"

func main(){
	var name string
	var surname string
	var age int

	fmt.Print("Введите ваше имя и фамилию: ")
	fmt.Scanln(&name, &surname)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)

	agediff := age - 5

	fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %d. Как молоды мы были!", name, surname, agediff)
	

}
