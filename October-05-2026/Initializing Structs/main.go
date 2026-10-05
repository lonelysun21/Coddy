package main

import "fmt"

// Определение структуры Person
type Person struct {
	Name       string
	Age        int
	IsEmployed bool
}

func main() {
	// TODO: Initialize a Person struct with name "Alice", age 28, and isEmployed true
	// Используйте либо имена полей, либо синтаксис литерала структуры
	var alice Person
	alice = Person{Name: "Alice", Age: 28, IsEmployed: true}
	// Выведите информацию о человеке
	fmt.Printf("Name: %s\n", alice.Name)
	fmt.Printf("Age: %d\n", alice.Age)
	fmt.Printf("Employed: %t\n", alice.IsEmployed)
}
