package main

import "fmt"

// Фильтрация in-place
// Напишите функцию, которая удаляет из среза все четные числа in-place (без выделения нового среза).
// Пример:
// Вход: []int{1, 2, 3, 4, 5}
// Выход: [1 3 5]
func main() {
	a := []int{2, 0, 1, 6, 7, 3, 6, 5, 1, 8}
	fmt.Println("начальный слайс", a)
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			index := i
			a = append(a[:index], a[index+1:]...)
		}
	}
	if a[0]%2 == 0 {
		a = append(a[:0], a[1:]...)
	}
	if a[len(a)-1]%2 == 0 {
		a = a[:len(a)-1]
	}
	fmt.Println("конечный массив", a)
}
