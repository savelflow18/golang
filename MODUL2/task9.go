package main

import "fmt"

// Вставка элемента по индексу
// Напишите функцию, которая вставляет элемент в срез по указанному индексу с сохранением исходного capacity.
// Пример:
// Вход: slice = []int{10, 20, 30}, index = 1, value = 99
// Выход: [10 99 20 30]
func main() {
	var a int
	fmt.Println("введите индекс:")
	fmt.Scan(&a)
	var b int
	fmt.Println("введите значение нового элемента:")
	fmt.Scan(&b)
	//c := make([]int, 4, 100)
	c := []int{12, 5, 7, 4}
	if a > len(c) {
		a = len(c)
	}
	fmt.Println("начальный слайс", c)
	d := make([]int, len(c)+1)
	copy(d, c[:a])
	d[a] = b
	copy(d[a+1:], c[a:])
	fmt.Println(d)
}
