// Задание 5. Максимум и минимум среди N введённых чисел.
package main

import (
	"errors"
	"fmt"
)

func minMax(values []int) (int, int, error) {
	if len(values) == 0 {
		return 0, 0, errors.New("последовательность не должна быть пустой")
	}
	minValue, maxValue := values[0], values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}
	return minValue, maxValue, nil
}

func main() {
	var count int

	fmt.Print("Введите количество чисел N: ")
	if _, err := fmt.Scan(&count); err != nil {
		fmt.Println("Ошибка: ожидалось целое число")
		return
	}
	if count < 1 {
		fmt.Println("Ошибка: N должно быть не меньше 1")
		return
	}

	values := make([]int, count)
	for i := range values {
		fmt.Print("Введите число: ")
		if _, err := fmt.Scan(&values[i]); err != nil {
			fmt.Println("Ошибка: ожидалось целое число")
			return
		}
	}
	minValue, maxValue, _ := minMax(values)

	fmt.Printf("Максимальное значение: %d\n", maxValue)
	fmt.Printf("Минимальное значение: %d\n", minValue)
}
