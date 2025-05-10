package main

import "fmt"

// Объединение двух мап
// Напишите функцию MergeMaps, которая принимает две мапы (map[string]int) и возвращает новую мапу, объединяющую их. Если ключ присутствует в обеих мапах, выбирается значение из второй.
func MergeMaps(m1 map[string]int, m2 map[string]int) map[string]int {
	m := make(map[string]int)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			m[k] = v
		} else if _, ok := m1[k]; ok {
			m[k] = v
		}
	}
	return m
}
func main() {
	m1 := map[string]int{"a": 1, "b": 2, "d": 3}
	m2 := map[string]int{"a": 3, "b": 7, "c": 3}
	m := MergeMaps(m1, m2)
	fmt.Println(m)
}
