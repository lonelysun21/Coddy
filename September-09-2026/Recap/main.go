package main

import "fmt"

func main() {
  // Заданные переменные
  score := 85
  grade := ""
  
  // TODO: Используйте оператор if, чтобы проверить, больше или равна ли score 90
  // Если истинно, установите grade в "A"
  if score >= 90 {
    grade = "A"
  } else if score >= 80 {
    grade = "B"
  } else if score >= 70 {
    grade = "C"
  } else {
    grade = "F"
  }
  // TODO: Выведите точный текст из инструкций, используя переменную grade
  fmt.Println("Grade:", grade)
  // TODO: Используйте оператор switch с переменной grade, чтобы вывести сообщение
  switch grade {
    case "A":
        fmt.Println("Excellent work!")
    case "B":
        fmt.Println("Good job!")
    case "C":
        fmt.Println("Satisfactory.")
    default:
        fmt.Println("You need to improve.")
  }
  // Для "A" выведите "Excellent work!"
  // Для "B" выведите точный текст из инструкций
  // Для "C" выведите "Satisfactory."
  // Для любой другой оценки выведите "You need to improve."

}
