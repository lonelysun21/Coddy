package main

import "fmt"

func main() {
	// Срез фруктов
	fruits := []string{"Apple", "Banana", "Cherry", "Dragon fruit", "Elderberry"}
	
	// TODO: Завершите цикл for, используя range для итерации по срезу fruits
	// Выведите каждый фрукт с его позицией, см. инструкции для точного формата
	for index, value := range fruits {// Добавьте ваш цикл range здесь 
		// Добавьте ваш код здесь, чтобы вывести каждый фрукт с его позицией
		fmt.Printf("%d. %s\n", index+1, value)
	}
}
