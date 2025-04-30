package main

import (
	"fmt"
)

// Создайте map, в которой необходимо хранить информацию о выданных читателю печатных изданиях: книгах и периодических изданиях.
// Тип ключей отображения является строкой, тип значений — отображением с ключами типа строка и значениями с типом слайс-строк.
// Добавьте несколько произвольных значений, моделирующих наличие изданий на руках у читателей.
// Выведите на экран количество читателей с изданиями на руках.
// Выведите на экран общее количество изданий на руках у каждого читателя.
// Для перебора мап используйте приведённый в разделе цикл по диапазону for … range {}
func add(a map[string]map[string][]string, name string, book string, pereodisd []string) {
	a[name] = make(map[string][]string)
	a[name][book] = pereodisd
}
func main() {
	a := make(map[string]map[string][]string)
	//n := 0
	c := 0
	//s := ""
	//i := ""
	add(a, "saveli", "Властелин колец", []string{"1_tom", "2_tom"})
	add(a, "ivan", "Параты карибского моря", []string{"1_tom", "2_tom", "3_tom"})
	/*for _, v := range a {
		for _, V := range v{
			for _, V2 := range V {
				if V2 != "" {
					n += 1
				}
			}
		}
	}
	for _, v := range a {
		for _, V := range v {
			for _, V2 := range V {
				if V2 != "" {
					n += 1
				}
			}
		}
	}*/
	for _, v := range a {
		for _, V := range v {
			if V != nil {
				c += 1
			}
		}
	}
	//for k, _ := range a {
	//	s = k
	//}
	fmt.Println("количество людей с периодическими изданиями", c)
	//fmt.Println(n, "периодических издания у ", s)
	fmt.Println(a)

}
