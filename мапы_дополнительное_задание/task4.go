package main

import "fmt"

// Удаление элемента из мапы
// Удалите ключ "banana" из мапы fruits := map[string]int{"apple": 5, "banana": 10, "cherry": 15} и выведите результат.
func main() {
	fruits := map[string]int{"apple": 5, "banana": 10, "cherry": 15}
	fruits2 := map[string]int{}
	fmt.Println("начальная мапа", fruits)
	for k, v := range fruits {
		if k != "banana" {
			fruits2[k] = v
		}
	}
	fmt.Println("после обработки ", fruits2)
}
