package main

import "fmt"

// Поиск уникальных элементов в слайсе
// Напишите функцию FindUnique, которая принимает слайс целых чисел и возвращает новый слайс, содержащий только уникальные элементы (встречающиеся 1 раз)
func FindUnique(s *[]int) []int {
	s1 := []int{}
	m := make(map[int]int)
	for _, v := range *s {
		m[v]++
	}
	for k, v := range m {
		v--
		if m[v] == 0 {
			s1 = append(s1, k)
		}
	}
	return s1
}
func main() {
	s := []int{1, 4, 6, 7, 4, 2, 5, 7, 3, 1, 4, 9}
	fmt.Println("начальный слайс", s)
	a := FindUnique(&s)
	fmt.Println("слайс после обработки", a)
}
