package main

import "fmt"

// Поиск дубликатов в слайсе
// Дан слайс nums := []int{1, 2, 3, 2, 4, 5, 4}. Проверьте, есть ли в нём дубликаты, используя мапу.
func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 4}
	m := make(map[int]bool)
	n := true
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			n = false
			m[v] = true
		} else {
			m[v] = false
		}
	}
	if n == true {
		fmt.Println("нету дубликатов")
	} else {
		fmt.Println("есть дубликаты")
	}
	fmt.Println(m)
}
