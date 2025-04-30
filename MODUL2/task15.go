package main

import "fmt"

// Преобразование мапы в срез ключей/значений
// Напишите функцию, которая принимает map[int]string и возвращает два среза: один с ключами, другой со значениями.
// Пример:
// Вход: map[int]string{1: "a", 2: "b"}
// Выход: []int{1, 2}, []string{"a", "b"}
func main() {
	a := map[int]string{1: "a", 2: "b", 3: "c"}
	b := []int{}
	c := []string{}
	fmt.Println("начальная мапа ", a)
	for k, v := range a {
		b = append(b, k)
		c = append(c, v)
	}
	fmt.Println(b)
	fmt.Println(c)
}
