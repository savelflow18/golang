package main

import (
	"fmt"
	"strconv"
)

// Создайте новую директорию с файлом main.go. Напишите код, в котором:
// объявите две переменные, первая — строка со значением 104, вторая — целое число со значением 35;
// приведите строку к целому числу, а целое число — к строке;
func main() {
	var a string
	a = "104"
	var b int
	b = 35
	B := strconv.Itoa(b)
	fmt.Println(B)
	A, _ := strconv.Atoi(a)
	fmt.Println(A)
}
