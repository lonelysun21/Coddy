package main

import "fmt"

func main() {
    // TODO: Initialize an array called 'favoriteNumbers' with these 5 integers: 7, 42, 8, 13, 99
    favoriteNumbers := [5]int{7, 42, 8, 13, 99}
    // Вы можете использовать любой синтаксис:
    // favoriteNumbers := [5]int{7, 42, 8, 13, 99}
    // ИЛИ
    // var favoriteNumbers [5]int = [5]int{7, 42, 8, 13, 99}
    
    // Это выведет ваш массив
    fmt.Printf("My favorite numbers are: %v\n", favoriteNumbers)
}
