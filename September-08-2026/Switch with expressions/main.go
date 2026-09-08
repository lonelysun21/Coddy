package main

import "fmt"

func main() {
	// Название дня уже определено
	dayName := "Wednesday"
	
	// Переменная для хранения описания
	var description string
	
	// TODO: Напишите оператор switch, который устанавливает description на основе dayName
	// Для понедельника: "Start of the week"
	// Для среды: "Midweek"
	// Для пятницы: "Almost weekend"
	// Для субботы или воскресенья: "Weekend"
	// Для любого другого дня: "Regular day"
	switch dayName {
		case "Monday": {
			description = "Start of the week"
		}
		case "Wednesday": {
			description = "Midweek"
		}
		case "Friday": {
			description = "Almost weekend"
		}
		case "Saturday": {
			description = "Weekend"
		}
		case "Sunday": {
			description = "Weekend"
		}
		default: {
			description = "Regular day"
		}
	}
	
	// Выведите результат
	fmt.Printf("%s is %s\n", dayName, description)

}
