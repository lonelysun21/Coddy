package main

import "fmt"

func main() {
	// Эта переменная представляет день недели
	day := "Saturday"
	
	// TODO: Завершите оператор switch, чтобы проверить, является ли день будним или выходным
	switch day {
	// Добавьте варианты для выходных дней (суббота и воскресенье) и выведите "It's the weekend!"
	case "Saturday":
		fmt.Println("It's the weekend!")
    case "Sunday":
		fmt.Println("It's the weekend!")
	// Добавьте вариант по умолчанию для будних дней и выведите "It's a weekday."
	default:
		fmt.Println("It's a weekday.")
	}
}
