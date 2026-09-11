package main

import (
    "fmt"
)

// TODO: Завершите функцию greet, которая принимает параметр name
// и возвращает строку приветствия в формате: "Hello, " + name + "!"
// Пример: если name равно "Gopher", верните точный текст из инструкций
func greet(name string) string {
    return "Hello, " + name + "!"
}

func main() {
    // Протестируйте функцию с предопределённым именем
    name := "Gopher"
    message := greet(name)
    fmt.Println(message)
    
    // Ожидаемый вывод: Hello, Gopher!
}
