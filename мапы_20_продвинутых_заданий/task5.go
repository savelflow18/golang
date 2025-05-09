package main

import (
	"fmt"
	"os"
)

// Проверка анаграмм
// Даны две строки. Проверьте, являются ли они анаграммами (состоят из одних и тех же символов в разном порядке).
func main() {
	s1 := "make"
	s2 := "ekam"
	if len(s1) != len(s2) {
		fmt.Println("не анограмма")
		os.Exit(0)
	}
	m1 := make(map[string]int)
	for _, v := range s1 {
		m1[string(v)]++
	}
	for _, v := range s2 {
		m1[string(v)]--
		if m1[string(v)] < 0 {
			fmt.Println("не анограмма")
			os.Exit(0)
		}
	}
	fmt.Println("анограмма")
}
