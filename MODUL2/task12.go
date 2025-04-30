package main

import "fmt"

// Фильтрация мапы по значению
// Напишите функцию, которая принимает map[string]int и возвращает новую мапу, содержащую только элементы со значениями больше заданного числа n.
// Пример:
// Вход: map[string]int{"a": 10, "b": 5, "c": 20}, n = 15
// Выход: map[c:20]
func main() {
	var n int
	fmt.Println("введите ваше число")
	fmt.Scan(&n)
	a := map[string]int{"a": 10, "b": 7, "c": 6, "d": 4}
	fmt.Println("начальная мапа", a)
	b := make(map[string]int)
	for k, v := range a {
		if v > n {
			b[k] = v
		}
	}
	fmt.Println("конечная мапа", b)

}
