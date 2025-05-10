package main

import "fmt"

// Фильтрация слайса по условию
// Напишите функцию FilterEven, которая принимает слайс целых чисел и возвращает новый слайс, содержащий только чётные числа.
func FilterEven(s *[]int) (s1 []int) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i]%2 == 0 {
			s1 = append(s1, (*s)[i])
		}
	}
	return s1
}
func main() {
	s := []int{3, 5, 6, 4, 4, 9, 8, 5, 2, 6, 7, 10, 3}
	a := FilterEven(&s)
	fmt.Println(a)
}
