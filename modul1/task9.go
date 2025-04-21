package main

//Напишите функцию, которая генерирует первые N чисел Фибоначчи.
import (
	"fmt"
)

func fibonacciIndex(A int) {
	f0, f1 := 1, 1
	for i := 0; i < A; i++ {
		f0, f1 = f1, f0+f1
		fmt.Println(f1)
	}

}
func main() {
	var A int
	fmt.Scan(&A)
	fibonacciIndex(A)
}
