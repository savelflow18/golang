package main

import "fmt"

// Подсчет гласных в строке
// Дана строка. Посчитайте количество каждой гласной буквы (a, e, i, o, u).
func main() {
	s := "jdhbgehuigbiowruuwivnejn"
	s2 := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			s2 += string(s[i])
		}
	}
	m := make(map[string]int)
	for _, v := range s2 {
		m[string(v)]++
	}
	for k, v := range m {
		fmt.Println("буква:", k, "сколько раз:", v)
	}
}
