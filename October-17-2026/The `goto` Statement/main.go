package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Чтение входных данных
    var maxRetriesStr string
    var successAttemptStr string
    fmt.Scanln(&maxRetriesStr)
    fmt.Scanln(&successAttemptStr)
    
    // Преобразование входных данных в целые числа
    maxRetries, _ := strconv.Atoi(maxRetriesStr)
    successAttempt, _ := strconv.Atoi(successAttemptStr)
    
    // Инициализация счетчика попыток
    attempt := 1
    
    // TODO: Напишите свой код ниже
    // Используйте оператор goto с меткой для реализации логики повторных попыток
    tryFind:
        if attempt != successAttempt && attempt <= maxRetries{
            fmt.Printf("Attempt %d failed\n", attempt)
            attempt++
            goto tryFind
        } else if attempt == successAttempt && attempt <= maxRetries{
            fmt.Printf("Attempt %d succeeded\n", attempt)
        } else if attempt > maxRetries{
            fmt.Printf("All attempts failed\n")
        }




}
