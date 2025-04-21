package main

import "fmt"

func main() {
	var a int
	fmt.Println("write a numb")
	fmt.Scan(&a)
	c := 1
	for i := 1; i <= a; i++ {
		c *= i
	}
	fmt.Println("факториал числа", a, "=", c)
}
