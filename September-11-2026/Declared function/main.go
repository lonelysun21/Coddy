package main

import "fmt"

// TODO: Реализуйте функцию greet, которая принимает параметр name
// и возвращает строку приветствия
func greet(name string) string {
    return "Hello, " + name + "!"
}

func main() {
    // Эти тестовые случаи уже подготовлены для вас
    name1 := "Alice"
    name2 := "Bob"
    
    // Тестирование функции greet
    message1 := greet(name1)
    message2 := greet(name2)
    
    fmt.Println(message1)
    fmt.Println(message2)
}
