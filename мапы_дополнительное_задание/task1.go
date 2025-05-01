package main

import "fmt"

// Создание и вывод мапы
// Создайте мапу map[string]int с ключами "apple", "banana", "cherry" и значениями 5, 10, 15. Выведите её.
func main() {
	m := map[string]int{}
	m["apple"] = 5
	m["banana"] = 10
	m["cherry"] = 15
	fmt.Println("мапа:")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
