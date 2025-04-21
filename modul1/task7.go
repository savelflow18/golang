package main

//Напишите функцию, которая проверяет, является ли строка палиндромом.
import (
	"fmt"
)

func polydrom(a string) int {
	len := 0
	b := 0
	for range a {
		len++
	}
	for i := 0; i < len/2; i++ {
		if a[i] == a[len-i+1] {
			b = 1
		}
	}
	return b
}
func main() {
	var input string
	fmt.Println("write a text")
	fmt.Scanln(&input)
	a := polydrom(input)
	if a != 0 {
		fmt.Println("палидром")
	} else if a == 0 {
		fmt.Println("обычное слово")
	}
}
