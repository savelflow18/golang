package main

import "fmt"

// Подсчёт длины мапы
// Дана мапа scores := map[string]int{"math": 90, "physics": 85, "chemistry": 78}. Выведите количество элементов в ней.
func main() {
	scores := map[string]int{"math": 90, "physics": 85, "chemistry": 78}
	n := 0 //счетчик элементов
	for _, _ = range scores {
		n++
	}
	fmt.Println("количество элементов", n)
}
