package main

import "fmt"

// Подсчёт частоты элементов в слайсе
// Напишите функцию CountFrequency, которая принимает слайс строк и возвращает мапу, где ключи — элементы слайса, а значения — количество их вхождений.
func CountFrequency(s *[]string) map[string]int {
	m := make(map[string]int)
	for _, v := range *s {
		m[v]++
	}
	return m
}
func main() {
	s := []string{"lol", "say", "me", "why"}
	a := CountFrequency(&s)
	fmt.Println(a)
}
