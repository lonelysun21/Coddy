package main

import "fmt"

func main() {
	// Перебрать числа от 1 до 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("Checking number: %d\n", i)
		
		// TODO: Проверить, делится ли текущее число на 3
		// Если да, вывести "Found it: [number]!" и использовать break для выхода из цикла
		if i%3==0 {
			fmt.Printf("Found it: %d!\n", i)
			break
		}
	}
	
	fmt.Println("Search complete")
}
