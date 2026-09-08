package main

import "fmt"

func main() {
	// Это представляет текущий час в 24-часовом формате (0-23)
	timeOfDay := 14 // Это 2:00 PM
	
	// TODO: Завершите оператор switch без выражения
	// чтобы вывести соответствующее приветствие в зависимости от времени суток
	switch {
	// Добавьте свои cases здесь
	case timeOfDay < 11: 
		fmt.Println("Good morning!")
	case timeOfDay < 17: 
		fmt.Println("Good afternoon!")
	case timeOfDay < 21: 
		fmt.Println("Good evening!")
	default:
		fmt.Println("Hello!")
	}
}
