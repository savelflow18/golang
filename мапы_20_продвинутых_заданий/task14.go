package main

import (
	"fmt"
)

// Найти индекс первого уникального символа
// Дана строка. Верните индекс первого символа, который встречается только один раз.
func main() {
	s := "34568314567392"
	var n int
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	for k, v := range m {
		if v == 1 {
			fmt.Println("число", m[k])
			n = m[k]
			break
		}
	}
	//for i := 0; i < len(s); i++ {
	//	if int(s[i]) == (n) {
	//		fmt.Println("индекс числа", i)
	//		break
	//	}
	//}
	for k, v := range s {
		if int(v)-48 == n {
			fmt.Println("индекс в строке", k)
		}
	}
}
