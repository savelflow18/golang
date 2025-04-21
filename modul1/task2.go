package main

import "fmt"

func main() {
	var a int
	fmt.Println("first numb:")
	fmt.Scan(&a)
	var b int
	fmt.Println("second numb:")
	fmt.Scan(&b)
	fmt.Println(a + b)
}
