package main

import "fmt"

//Добавление элемента в мапу
//Добавьте новый ключ "orange" со значением 20 в мапу из предыдущей задачи и выведите обновлённую мапу.

func main() {
	m := map[string]int{}
	m["apple"] = 5
	m["banana"] = 10
	m["cherry"] = 15
	m["orange"] = 20
	fmt.Println("мапа:")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
