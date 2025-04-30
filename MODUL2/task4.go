package main

import "fmt"

// Сравнение длины и capacity
// Напишите программу, которая:
// Создает срез с помощью make([]int, 3, 5),
// Добавляет элементы, пока len не превысит cap,
// Выводит len и cap после каждого добавления.
// Пример вывода:
// Исходный: len=3, cap=5
// После добавления 4: len=4, cap=5
// После добавления 6: len=6, cap=10 (удвоение capacity)
func main() {
	b := 1
	a := make([]int, b, 7)
	for i := int(len(a)); i <= cap(a)-1; i++ {
		a = append(a, 1)
		fmt.Println("len =", len(a), "cap =", cap(a))

	}
	a = append(a, 1)
	fmt.Println("len =", len(a), "cap =", cap(a))
}
