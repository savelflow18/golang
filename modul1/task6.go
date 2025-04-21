package main

//Создайте программу, которая находит максимальное число в массиве целых чисел.

import "fmt"

func main() {
	a := [5]int{3, 5, 2, 7, 3}
	b := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > b {
			b = a[i]
		}
	}
	fmt.Println(b)
}
