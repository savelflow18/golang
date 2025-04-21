package main

//Создайте программу, которая принимает массив и возвращает новый массив, содержащий только уникальные элементы.
import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 2, 5, 6}
	var c []int
	for i := 0; i < len(a); i++ {
		b := 1
		for j := 0; j < len(c); j++ {
			if a[i] == c[j] {
				b = 0
			}
		}
		if b == 1 {
			c = append(c, a[i])
		}
	}

	fmt.Println(c)
}
