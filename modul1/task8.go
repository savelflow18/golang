package main

import "fmt"

// Реализуйте программу, которая вычисляет сумму всех элементов в массиве.
func main() {
	a := [5]int{5, 6, 3, 8, 2}
	b := a[0]
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	fmt.Println(b)
}
