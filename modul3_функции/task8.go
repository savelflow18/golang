package main

import "fmt"

// Подсчёт среднего значения в слайсе
// Напишите функцию Average, которая принимает слайс чисел и возвращает их среднее значение. Если слайс пуст, возвращает 0.
func Average(s *[]int) (c int) {
	for i := 0; i < len(*s); i++ {
		c += (*s)[i]
	}
	return c
}
func main() {
	s := []int{1, 2, 3, 4, 5, 4, 2, 9, 6}
	a := Average(&s)
	fmt.Println("сумма элементов =", a)
}
