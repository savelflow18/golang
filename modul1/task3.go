package main

import "fmt"

func main() {
	var a int
	fmt.Println("write a numb")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("четное")
	} else if a%2 == 1 {
		fmt.Println("нечетное")
	}
}
