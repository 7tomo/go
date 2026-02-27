package main

import "fmt"

func frame_counter(frames_count int) {
	for i := 1; i <= frames_count; i++ {
		fmt.Println(i, "-й фрейм")
	}
}

func main() {
	fmt.Println("Боулинг")
	var frames [10]int = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 10; i++ {
		fmt.Printf("Введите количество очков за %d-й фрейм: ", i+1)
		fmt.Scanln(&frames[i])

	}
	frame_counter(10)
}
