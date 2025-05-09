package main

import "fmt"

// Сгруппировать слова по длине
// Дан список слов. Сгруппируйте их в хешмапу, где ключ — длина слова, а значение — список слов такой длины.
func main() {
	s := []string{"aboba", "lol", "me", "apply", "house", "yes", "no", "and"}
	m := make(map[int][]string)
	for k, v := range s {
		k = len(v)
		m[k] = append(m[k], v)
	}
	fmt.Println(m)

}
