package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Чтение ввода
    var rowsInput string
    var skipConditionInput string
    fmt.Scanln(&rowsInput)
    fmt.Scanln(&skipConditionInput)
    
    // Разбор входных данных
    numRows, _ := strconv.Atoi(rowsInput)
    skipCondition, _ := strconv.Atoi(skipConditionInput)
    
    // Предоставленный двумерный массив данных
    data := [][]int{
        {1, 2, 3, 4, 5},
        {6, 7, 8, 9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    }
    
    // TODO: Напишите ваш код ниже
    // Используйте вложенные циклы с помеченным continue для обработки данных
    // Проверьте каждую строку на число условия пропуска
    // Выведите соответствующие сообщения в зависимости от того, найдено ли условие
    outer:
    for i := 0; i < numRows; i++ {
        for j := 0; j < 5; j++ {
            if data[i][j] == skipCondition {
                fmt.Printf("Skipping row %d due to condition\n", i)
                continue outer
            }
        }
        fmt.Printf("Processing row %d: %d %d %d %d %d\n", i, data[i][0], data[i][1], data[i][2], data[i][3], data[i][4])
    }





}
