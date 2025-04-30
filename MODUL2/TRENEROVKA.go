package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("начальная", m)
	a := make(map[int]string)
	for k, v := range m {
		a[v] = k
	}
	fmt.Println("КОНЕЧНАЯ", a)
}
