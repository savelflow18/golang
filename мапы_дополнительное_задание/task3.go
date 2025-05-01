package main

import "fmt"

// Проверка наличия ключа
// Дана мапа ages := map[string]int{"Alice": 25, "Bob": 30}. Проверьте, есть ли в ней ключ "Bob", и выведите "Exists" или "Not exists".
func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30}
	a := 0
	for k, _ := range ages {
		if k == "Bob" {
			a += 1
		}
	}
	if a != 0 {
		fmt.Println("Exists")
	} else if a == 0 {
		fmt.Println("Not Exists")
	}
}
