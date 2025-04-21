package main

//Создайте программу, которая запрашивает у пользователя два числа и выводит их сумму.
import "fmt"

func main() {
	var a int
	fmt.Println("first numb:")
	fmt.Scan(&a)
	var b int
	fmt.Println("second numb:")
	fmt.Scan(&b)
	fmt.Println(a + b)
}
