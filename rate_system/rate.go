package main

import ("fmt")

func main() {
	rate(94)
	rate(88)
	rate(30)
	rate(70)
	rate(40)
	rate(-42)

}

func rate(rate int) {
	if rate >= 90 && rate <= 100 {
		fmt.Println("A")
	} else if rate >= 80 && rate <= 89 {
		fmt.Println("B")
	} else if rate >= 70 && rate <= 79 {
		fmt.Println("C")
	} else if rate >= 60 && rate <= 69 {
		fmt.Println("D")
	} else if rate < 60 && rate >= 0 {
		fmt.Println("F")
	} else {
		fmt.Errorf("Недопустимое значение" )
	}
}
