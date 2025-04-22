package main

import (
	"fmt"
	"math"
)

// Создайте новую директорию и новый файл main.go. Напишите код, в котором:
// объявите два новых типа AmericanVelocity и EuropeanVelocity;
// выполните преобразование скорости 120.4 м/сек в км/ч и присвойте результат переменной с типом EuropeanVelocity;
// выполните преобразование скорости 130 м/с в миль/ч и присвойте результат переменной с типом AmericanVelocity;
// примечание: 1 миля = 1.609 км. Если потребуется, округлите значение до 2 знаков после запятой, для округления обратитесь к пакету math.
func main() {
	Europeanvelocity := 120.4 * 1000 / 3600
	fmt.Println(math.Round(Europeanvelocity))
	Americanvelocity := 130 * 1000 / 3600 / 1.609
	fmt.Println(math.Round(Americanvelocity))
}
