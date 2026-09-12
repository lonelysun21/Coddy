package main

import "fmt"

// getGreeting возвращает приветствие в зависимости от часа дня
func getGreeting(hour int) string {
    // TODO: Вернуть "Good morning", если час меньше 12
    // TODO: Вернуть "Good afternoon", если час от 12 до 17 (включительно)
    // TODO: Вернуть "Good evening" для всех остальных часов
    
    // Ваш код здесь
    if hour < 12 {
        return "Good morning"
    } else if hour > 12 && hour <= 17 {
        return "Good afternoon"
    } else {
        return "Good evening"
    }
}

func main() {
    // Протестируйте функцию с разным временем
    morningHour := 8
    afternoonHour := 15
    eveningHour := 20
    
    fmt.Println("At", morningHour, "hours:", getGreeting(morningHour))
    fmt.Println("At", afternoonHour, "hours:", getGreeting(afternoonHour))
    fmt.Println("At", eveningHour, "hours:", getGreeting(eveningHour))
}
