package main

import "fmt"

/*
Удаление дубликатов из слайса
Напишите функцию RemoveDuplicates, которая принимает слайс строк и возвращает новый слайс без дубликатов.
*/
func RemoveDuplicates(s *[]string) (s1 []string) {
	m := make(map[string]int)
	for _, v := range *s {
		m[v]++
	}
	//for k, v := range m {
	//	if v == 1 {
	//		s1 = append(s1, k)
	//	}
	//}
	for k, _ := range m { //это если без дубликатов, а то что выше это если только уникальные
		s1 = append(s1, k)
	}
	return s1
}
func main() {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "a", "d", "g"}
	fmt.Println("начальный слайс", s)
	a := RemoveDuplicates(&s)
	fmt.Println("конечный слайс", a)
}
