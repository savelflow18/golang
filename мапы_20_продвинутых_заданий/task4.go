package main

import "fmt"

// Найти первый неповторяющийся символ
// Дана строка. Найдите первый символ, который встречается только один раз.
func main() {
	s := "3114621871458174687"
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++

	}
	for k, v := range s {
		if m[string(v)] == 1 {
			fmt.Println("первый символ", string(s[k]))
			break
		}

	}
	fmt.Println(m)
}
