package main

import "fmt"

// Передача среза по указателю
// Напишите функцию, которая принимает указатель на срез строк и добавляет строку "modified" в конец.
// Пример:
// Вход: &[]string{"a", "b"}
// Выход: ["a" "b" "modified"]
func main() {
	a := []string{"a", "b", "c"}
	b := &a
	d := "d"
	*b = append(*b, d)
	fmt.Println(*b)
}
