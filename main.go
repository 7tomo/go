package main

import (
	"fmt"
)

func main() {
	isPalindrome([1, 2, 3, 4, 5, 4, 3, 2, 1, 1])
}

func isPalindrome(nums [10]int) {
	for i := 0; i < 5; i++ {
		if nums[i] != nums[len(nums)-i] {
			fmt.Println("Не палиндром!")
		}
	}

}
