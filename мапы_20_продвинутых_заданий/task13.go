package main

import "fmt"

// Проверить, является ли массив подмножеством другого
// Даны два массива. Проверьте, все ли элементы первого массива присутствуют во втором.
func main() {
	m1 := []int{1, 2, 3, 4, 5}
	m2 := []int{1, 2, 3, 4, 6, 7, 5}
	m := make(map[int]int)
	b := false
	for _, v := range m1 {
		m[v]++
	}
	for _, v := range m2 {
		m[v]--

	}
	for _, v := range m {
		if v == 1 {
			b = false
			break
		} else if v != 1 {
			b = true
		}
	}
	if b == true {
		fmt.Println("все элементы есть")
	} else if b == false {
		fmt.Println("не все элементы")
	}
}
