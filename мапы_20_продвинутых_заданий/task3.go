package main

import "fmt"

// Поиск дубликатов в массиве
// Дан массив чисел. Верните true, если в массиве есть дубликаты, и false в противном случае.
func main() {
	a := []int{1, 2, 3, 4, 55, 7, 9, 0}
	m := make(map[int]int)
	b := false
	for _, v := range a {
		if _, ok := m[v]; !ok {
			b = false
		} else {
			b = true
		}
	}
	fmt.Println(b)
}
