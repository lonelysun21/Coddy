package main

import "fmt"

// updateQuantity обновляет количество товара и возвращает новое количество
// и закончился ли товар (quantity = 0)
func updateQuantity(quantity *int, change int) (int, bool) {
    *quantity = *quantity + change
    if *quantity <= 0 {
        *quantity = 0
    } else {
        return *quantity, false
    }
    // TODO: Обновить указатель quantity и вернуть новое значение и статус отсутствия на складе
    return 0, true
}

// calculateValue вычисляет общую стоимость товара на основе quantity и price
// It also updates the mostValuableItem if this item is more valuable
func calculateValue(itemName string, quantity int, price float64, mostValuableItem *string, highestValue *float64) float64 {
    // TODO: Calculate the total value, update mostValuableItem if needed, and return the total value
    totalValue := float64(quantity) * price
    if totalValue > *highestValue {
        *highestValue = totalValue
        *mostValuableItem = itemName
    }
    return totalValue
}

// displayInventory выводит сведения об инвентаре
// Она принимает указатели, чтобы функция могла показывать данные в реальном времени
func displayInventory(apples, oranges, bananas *int, applePrice, orangePrice, bananaPrice float64) {
    fmt.Printf("Apples: %d (Value: $%.2f)\n", *apples, float64(*apples) * applePrice)
    fmt.Printf("Oranges: %d (Value: $%.2f)\n", *oranges, float64(*oranges) * orangePrice)
    fmt.Printf("Bananas: %d (Value: $%.2f)\n", *bananas, float64(*bananas) * bananaPrice)
}

func main() {
    // Инициализация инвентаря
    apples := 10
    oranges := 15
    bananas := 8
    
    // Цены
    applePrice := 0.5  // $0.50 каждый
    orangePrice := 0.7 // $0.70 каждый
    bananaPrice := 0.3 // $0.30 каждый
    
    // Track the most valuable item
    var mostValuableItem string
    var highestValue float64
    
    // Display initial inventory
    fmt.Println("Initial Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Calculate initial values and find most valuable item
    appleValue := calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue := calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue := calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)
    
    // Симуляция некоторых продаж
    fmt.Println("Processing sales...")
    _, applesOutOfStock := updateQuantity(&apples, -4) // Продать 4 яблока
    _, orangesOutOfStock := updateQuantity(&oranges, -8) // Продать 8 апельсинов
    _, bananasOutOfStock := updateQuantity(&bananas, -10) // Попытаться продать 10 бананов (больше, чем у нас есть)
    
    // Check if any items are out of stock
    if applesOutOfStock {
        fmt.Println("Apples are out of stock!")
    }
    if orangesOutOfStock {
        fmt.Println("Oranges are out of stock!")
    }
    if bananasOutOfStock {
        fmt.Println("Bananas are out of stock!")
    }
    
    // Display updated inventory
    fmt.Println("\nUpdated Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Reset most valuable tracking for recalculation
    mostValuableItem = ""
    highestValue = 0
    
    // Пересчитать значения
    appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)
    
    // Пополнить запасы
    fmt.Println("Restocking...")
    updateQuantity(&apples, 5)  // Добавить 5 яблок
    updateQuantity(&oranges, 10) // Добавить 10 апельсинов
    updateQuantity(&bananas, 12) // Добавить 12 бананов
    
    // Display final inventory
    fmt.Println("\nFinal Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Reset most valuable tracking for final calculation
    mostValuableItem = ""
    highestValue = 0
    
    // Финальный расчёт стоимости
    appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n", mostValuableItem)
}
