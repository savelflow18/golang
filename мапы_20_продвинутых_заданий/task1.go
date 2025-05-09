package main

import "fmt"

// Подсчет частоты символов
// Дана строка. Посчитайте, сколько раз каждый символ встречается в строке.
func main() {
	s := "1524578763453768"
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++

	}
	fmt.Println(m)
}
