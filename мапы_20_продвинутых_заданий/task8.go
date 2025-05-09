package main

import "fmt"

// Найти самый частый элемент
// Дан массив чисел. Найдите число, которое встречается чаще всего.
func main() {
	a := []int{1, 3, 9, 7, 5, 3, 2, 1, 6, 5, 7, 4, 5, 5, 5, 5, 5}
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	max := m[1]
	fmt.Println(m)
	for k, _ := range m {
		if max < m[k] {
			max = m[k]
			fmt.Println("число", k)
		}
	}
	fmt.Println("сколько раз", max)
}
