package main

import "fmt"

func main() {
	// Срез фруктов
	fruits := []string{"apple", "banana", "orange"}
	
	// TODO: Добавьте "grape" и "kiwi" к срезу fruits
	// Напишите свой код здесь
	fruits = append(fruits, "grape", "kiwi")
	
	// Выведите обновлённый срез
	fmt.Println("My fruit collection:", fruits)
}
