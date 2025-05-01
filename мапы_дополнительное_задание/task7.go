package main

import "fmt"

// Объединение двух мап
// Даны две мапы:
//
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Создайте новую мапу, объединив их (при конфликте берётся значение из m2).
func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m := make(map[string]int)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		if _, ok := m[k]; ok { //проверочка на наличие значений по ключу
			m[k] += v
		} else {
			m[k] = v
		}
	}
	fmt.Println(m)

}
