package main

import (
	"fmt"
)

// Вложенные мапы (словарь словарей)
// Создайте map[string]map[string]int, где внешний ключ — имя студента, а внутренняя мапа хранит предметы и оценки. Напишите функцию для добавления оценки студенту.
// Пример:
// Вход:
//
// db := map[string]map[string]int{}
// addGrade(db, "Alice", "Math", 90)
// Выход: map[Alice:map[Math:90]]
func add(a map[string]map[string]int, name string, pred string, grad int) {
	a[name] = make(map[string]int)
	a[name][pred] = grad
}
func main() {
	a := make(map[string]map[string]int)
	add(a, "saveli", "math", 90)
	fmt.Println(a)
}
