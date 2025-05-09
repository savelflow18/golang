package main

import "fmt"

// Подсчет частоты слов
// Дана строка с словами. Посчитайте, сколько раз каждое слово встречается в строке.
func main() {
	s := []string{"me", "yse", "me", "world"}
	m := make(map[string]int)
	for _, v := range s {
		m[v]++
	}
	fmt.Println(m)
}
