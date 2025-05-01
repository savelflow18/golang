package main

import "fmt"

// Подсчёт частоты символов в строке
// Дана строка "hello". Создайте мапу, где ключи — символы, а значения — сколько раз они встречаются.
func main() {
	s := "hello"
	fmt.Println("начальная строка", s)
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	fmt.Println(m)
}
