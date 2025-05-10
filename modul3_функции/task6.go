package main

import "fmt"

// 6. Разделение слайса на пакеты (чанки)
// Напишите функцию ChunkSlice, которая принимает слайс целых чисел и размер чанка, а возвращает слайс слайсов, разбитый на части указанного размера.
func ChunkSlice(s []int, n *int) [][]int {
	s1 := [][]int{}
	for i := 0; i < len(s); i += *n {
		end := i + *n
		s1 = append(s1, s[i:end])
	}
	return s1

}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var n int
	fmt.Println("введите размер чанка")
	fmt.Scan(&n)
	a := ChunkSlice(s, &n)
	fmt.Println(a)
}
