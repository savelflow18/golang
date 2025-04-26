package main

import "fmt"

// Напишите функцию, которая создает срез целых чисел от 1 до n,
// затем добавляет к нему числа от n+1 до 2n с помощью append.
// Выведите исходный и измененный срез.
// Пример:
// Вход: n = 3
// Выход:
// Исходный: [1 2 3]
// После append: [1 2 3 4 5 6]
func main() {
	var n int
	fmt.Println("write your number")
	fmt.Scan(&n)
	a := []int{}
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	fmt.Println("исходный слайс", a)
	for i := n; i <= 2*n; i++ {
		a = append(a, i)
	}
	fmt.Println("новый слайс", a)
}
