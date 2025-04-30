package main

import "fmt"

// Удаление элемента по индексу
// Напишите функцию, которая удаляет элемент из среза по заданному индексу in-place (без возврата нового среза).
// Пример:
// Вход: slice = []int{10, 20, 30, 40}, index = 1
// Выход: [10 30 40]
func main() {
	a := []int{2, 3, 4, 5}
	var index int
	fmt.Println("введите индекс для удаления")
	fmt.Scan(&index)
	a = append(a[:index], a[index+1:]...)
	fmt.Println(a)
}
