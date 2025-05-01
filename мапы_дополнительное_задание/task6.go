package main

import "fmt"

// Перебор мапы
// Дана мапа colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}. Выведите все ключи и значения в формате "red: #FF0000".
func main() {
	colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	for k, v := range colors {
		fmt.Println(k, ":", v, ".")
	}
}
