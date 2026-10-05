package main

import "fmt"

// Структура Address содержит информацию о местоположении
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Структура Person должна включать встроенную структуру Address
type Person struct {
	Name string
	Age  int
	Address
	// TODO: Встройте здесь структуру Address (всего одна строка)
}

func main() {
	// Создайте новый экземпляр Person с информацией об адресе
	person := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Wonderland",
			ZipCode: "12345",
		},
	}

	// Выведите информацию о Person, включая адрес
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	
	// Выведите поля адреса непосредственно из экземпляра Person
	fmt.Println("Street:", person.Street)
	fmt.Println("City:", person.City)
	fmt.Println("ZipCode:", person.ZipCode)
}
