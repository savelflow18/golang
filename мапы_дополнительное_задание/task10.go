package main

import "fmt"

// Инвертирование мапы
// Дана мапа m := map[int]string{1: "one", 2: "two"}. Создайте новую мапу, где ключи и значения поменяются местами (map[string]int).
func main() {
	m := map[int]string{1: "one", 2: "two"}
	m1 := make(map[string]int)
	for k, v := range m {
		m1[v] = k
	}
	fmt.Println("начальная мапа", m)
	fmt.Println("конечная мапа", m1)
}
