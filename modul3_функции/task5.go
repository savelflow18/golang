package main

import (
	"fmt"
	"os"
)

// Напишите функцию InvertMap, которая принимает мапу map[string]int и возвращает новую мапу map[int]string, где ключи и значения поменялись местами. Если в исходной мапе есть дублирующиеся значения, вернуть ошибку.
func InvertMap(m1 map[string]int) map[int]string {
	m2 := make(map[int]string)
	for k, v := range m1 {
		m2[v] = k
	}
	return m2
}
func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "f": 4}
	m0 := map[string]int{}
	for _, v := range m1 {
		m0[string(v)]++
	}
	for _, v := range m0 {
		v--
		if m0[string(v)] > 0 {
			fmt.Println("ошибка")
			os.Exit(0)
		}
	}
	m := InvertMap(m1)
	fmt.Println(m)
}
