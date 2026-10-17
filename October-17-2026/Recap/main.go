package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Чтение входных данных
    var threatLevel string
    var maxZonesStr string
    fmt.Scanln(&threatLevel)
    fmt.Scanln(&maxZonesStr)
    
    // Преобразование максимального количества зон в целое число
    maxZones, _ := strconv.Atoi(maxZonesStr)
    
    // Данные зон безопасности
    zones := [][]string{
        {"low", "medium", "low"},
        {"medium", "high", "low"},
        {"critical", "medium", "high"},
        {"low", "critical", "medium"},
    }
    
    // TODO: Напишите свой код ниже
    // Используйте вложенные циклы с именованным break для поиска угрозы
    // Выведите местоположение угрозы, если она найдена, или "Threat not found in searched zones", если не найдена
    // Затем используйте switch с fallthrough для оповещений безопасности
    isFind := false
    outer:
    for i := 0; i < maxZones; i++ {
        for j := 0; j < 3; j++ {
            if zones[i][j] == threatLevel {
                fmt.Printf("Threat found at zone %d position %d\n", i, j)
                isFind = true
                break outer
            }
        }
    }
    
    if !isFind {
        fmt.Println("Threat not found in searched zones")
    }

    switch threatLevel {
        case "critical":
            fmt.Println("CRITICAL: Lockdown initiated!")
            fallthrough
        case "high":
            fmt.Println("HIGH: Security breach detected!")
            fallthrough
        case "medium":
            fmt.Println("MEDIUM: Increased monitoring active")
        case "low":
            fmt.Println("LOW: Standard security protocols")
    }






}
