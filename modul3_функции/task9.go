package main

import "fmt"

// Группировка строк по длине
// Напишите функцию GroupByLength, которая принимает слайс строк и возвращает мапу, где ключ — длина строки, а значение — слайс строк этой длины.
func GroupByLength(s *[]string) map[int][]string {
	m := make(map[int][]string)
	for k, v := range *s {
		k = len(v)
		m[(k)] = append(m[(k)], v)
	}
	return m
}
func main() {
	s := []string{"lol", "recs", "yes", "save", "house"}
	a := GroupByLength(&s)
	fmt.Println(a)
}
