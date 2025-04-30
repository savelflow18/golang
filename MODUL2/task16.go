package main

import "fmt"

// Удаление дубликатов из среза с помощью мапы
// Напишите функцию, которая принимает срез строк и возвращает новый срез без дубликатов, используя мапу для проверки уникальности.
// Пример:
// Вход: []string{"a", "b", "a", "c"}
// Выход: []string{"a", "b", "c"}
func main() {
	a := []string{"a", "b", "c", "a", "d", "b", "f"}
	fmt.Println("начальный срез", a)
	c := []string{}
	b := make(map[string]bool)
	for _, v := range a {
		if _, ok := b[v]; !ok {
			b[v] = true
			c = append(c, v)
		}
	}
	fmt.Println("срез с уникальными элементами", c)
}
