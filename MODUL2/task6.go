package main

import "fmt"

// Объединение двух срезов
// Напишите функцию, которая принимает два среза и возвращает новый срез,
// содержащий все элементы обоих без использования append (на базе make и копирования).
// Пример:
// Вход: []int{1, 2}, []int{3, 4}
// Выход: [1 2 3 4]
func main() {
	a := []int{2, 3, 4}
	b := []int{1, 5, 6}
	a1 := make([]int, len(a))
	copy(a1, a)
	b1 := make([]int, len(b))
	copy(b1, b)
	c := make([]int, len(a)+len(b))
	for i := range a1 {
		c[i] = a1[i]
	}
	for i := range b {
		c[a1[i]+1] = b1[i]
	}
	fmt.Println(c)
}
