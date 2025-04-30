package main

import "fmt"

// Объединение двух мап
// Напишите функцию, которая принимает две map[string]int и возвращает новую мапу, объединяя их. Если ключ присутствует в обеих мапах, сложите значения.
// Пример:
// Вход:
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Выход: map[a:1 b:5 c:4]
func main() {
	a := map[string]int{"a": 7, "b": 4, "c": 3}
	b := map[string]int{"a": 4, "c": 1}
	c := make(map[string]int)
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		if _, ok := c[k]; ok {
			c[k] = c[k] + v
		} else {
			c[k] = v
		}
	}
	fmt.Println(c)
}
