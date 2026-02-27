package main

import "fmt"

func main() {

	var check bool

	check = ((age >= 18 && role == "user") || ((role == "admin" || role == "moderator") && age < 18) || (role == "admin" || role == "moderator") || age >= 18) && !(role == "officer") && (!(status == "inactive") || !(role == "user"))

	fmt.Println(check)
}
