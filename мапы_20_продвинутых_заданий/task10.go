package main

import "fmt"

// Проверить, являются ли строки изоморфными
// Две строки называются изоморфными, если можно заменить символы одной строки на символы другой (с сохранением порядка). Например, "egg" и "add" изоморфны.
func main() {
	s1 := "egf"
	s2 := "dd"
	b := false
	if len(s1) != len(s2) {
		b = false
	}
	m2 := make(map[string]int)
	m := make(map[string]int)
	for _, v := range s1 {
		m[string(v)]++
	}
	for _, v := range s2 {
		m2[string(v)]++
	}
	fmt.Println(m2)
	fmt.Println(m)
	for _, v := range m {
		if m[string(v)] == m2[string(v)] {
			b = true
		} else if m[string(v)] != m2[string(v)] {
			b = false
			break
		}
	}
	if b == false {
		fmt.Println(s1, "и", s2, "-", "не изоморфны")
	} else if b == true {
		fmt.Println(s1, "и", s2, "-", "изоморфны")
	}
}
