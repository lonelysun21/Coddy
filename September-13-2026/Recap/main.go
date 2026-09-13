package main

import "fmt"

// calculateCalories определяет, сколько калорий сжигается во время активности
// Параметры:
// - activity: тип упражнения ("running", "swimming", "cycling" или любая другая активность)
// - duration: как долго выполнялась активность (в минутах)
// - intensity: множитель, представляющий интенсивность тренировки (1.0 — обычная)
// Возвращает:
// - общее количество сожженных калорий
func calculateCalories(activity string, duration int, intensity float64) (result float64) {
    switch activity {
        case "running":
            result = 10 * intensity * float64(duration)
        case "swimming":
            result = 8 * intensity * float64(duration)
        case "cycling":
            result = 7 * intensity * float64(duration)
        default:
            result = 5 * intensity * float64(duration)
    }
    return 
}

func main() {
    // Протестируйте функцию с различными активностями
    fmt.Println("Running for 30 minutes at intensity 1.2:", calculateCalories("running", 30, 1.2))
    fmt.Println("Swimming for 45 minutes at intensity 1.0:", calculateCalories("swimming", 45, 1.0))
    fmt.Println("Cycling for 60 minutes at intensity 1.5:", calculateCalories("cycling", 60, 1.5))
    fmt.Println("Yoga for 60 minutes at intensity 0.8:", calculateCalories("yoga", 60, 0.8))
}
