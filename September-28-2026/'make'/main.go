package main

import "fmt"

func main() {
	// TODO: Создайте слайс строк длиной 3 и ёмкостью 5 с помощью make
	// var names = ...
	var names = make([]string, 3, 5)
	// Присвойте значения слайсу
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"
	
	// Выведите слайс
	fmt.Println("Names:", names)
	
	// Выведите длину и ёмкость слайса
	fmt.Printf("Length: %d, Capacity: %d\n", len(names), cap(names))
}
