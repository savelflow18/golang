package main

import "fmt"

// Подсчет суммы двух чисел
// Дан массив чисел и целевое число. Проверьте, есть ли в массиве два числа, сумма которых равна целевому.
func main() {
	a := []int{1, 12, 3}
	n := 15
	m := make(map[int]int)
	for k, v := range a {
		m[k] = v
	}
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i]+m[j] == n {
				fmt.Println("true")
				break
			}
		}
	}
}
