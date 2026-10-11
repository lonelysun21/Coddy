package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // Чтение входных данных
    var dimensions string
    var targetStr string
    fmt.Scanln(&dimensions)
    fmt.Scanln(&targetStr)
    
    // Разбор размеров
    dimParts := strings.Split(dimensions, ",")
    rows, _ := strconv.Atoi(dimParts[0])
    cols, _ := strconv.Atoi(dimParts[1])
    
    // Разбор целевого числа
    target, _ := strconv.Atoi(targetStr)
    
    // Предопределенная сетка
    grid := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    found := false
    // TODO: Напишите свой код ниже
    // Используйте вложенные циклы с именованным break для поиска цели
    outer:
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == target {
                fmt.Printf("Found %d at position (%d, %d)\n", target, i, j)
                found = true
                break outer
            }
        }
    }
    if !found {
        fmt.Printf("Target %d not found\n", target)
    }
}
