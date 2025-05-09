package main

import "fmt"

// Найти пересечение двух массивов
// Даны два массива чисел. Верните массив их пересечения (общие элементы).
func main() {
	s1 := []int{3, 5, 6, 2}
	s2 := []int{5, 2, 7, 9}
	s0 := []int{}
	m1 := make(map[int]int)
	for _, v := range s1 {
		m1[v]++
	}
	for _, v := range s2 {
		m1[v]--
	}
	for k, v := range m1 {
		if v == 0 {
			s0 = append(s0, k)
		}
	}
	fmt.Println("объединение двух массивов", s0)
}
