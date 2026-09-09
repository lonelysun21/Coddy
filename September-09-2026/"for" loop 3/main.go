package main

import "fmt"

func main() {
	// Эти переменные уже настроены для вас
	keepRunning := true
	count := 0
	
	// TODO: Завершите цикл for, который использует только условие (как цикл while)
	// Цикл должен продолжаться до тех пор, пока keepRunning имеет значение true
	for keepRunning{
		// Вывести текущее значение count
		fmt.Println(count)
		
		// Увеличить count
		count++
		
		// Проверить, равен ли count 5, и если да, установить keepRunning в false
		if count >= 5 {
			keepRunning = false
		}
	}
}
