package main

import "fmt"

// Изменение in-place
// Напишите функцию, которая принимает срез целых чисел и умножает каждый его элемент на 2 без возврата нового среза (изменя исходный).
// Пример:
// Вход: []int{1, 2, 3}
// Выход: [2 4 6]
func main() {
	a := []int{2, 3, 4}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * 2
	}
	fmt.Println(a)
}
