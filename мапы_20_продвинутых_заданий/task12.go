package main

import "fmt"

// Найти пропущенное число
// Дан массив чисел от 1 до n с одним пропущенным числом. Найдите его, используя хешмапу.
func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	m := make(map[int]int)
	n := 0
	for k, v := range mas {
		m[k] = v
		if v-k > 1 {
			n = v - 1
			break
		}
	}
	fmt.Println("пропавшее число", n)
}
