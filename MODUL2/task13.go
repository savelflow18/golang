package main

//Проверка на анаграммы
//Напишите функцию, которая проверяет, являются ли две строки анаграммами (содержат одинаковые символы в разном порядке), используя мапу для подсчета символов.
//Пример:
//Вход: "listen", "silent"
//Выход: true

import (
	"fmt"
)

func main() {
	str1 := "listen"
	str2 := "silent"
	a := true
	b := make(map[rune]int)
	for _, v := range str1 {
		b[v]++
	}
	for _, v := range str2 {
		b[v]--
		if b[v] < 0 {
			a = false
		}
	}
	fmt.Println(a) // Вывод: true
}
