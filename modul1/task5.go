package main

import "fmt"

func main() {
	var a string
	fmt.Println("введите строку")
	fmt.Scanln(&a)
	var b string
	for i := len(a) - 1; i >= 0; i-- {
		b += string(uint64(a[i]))
	}
	fmt.Println(b)
}
