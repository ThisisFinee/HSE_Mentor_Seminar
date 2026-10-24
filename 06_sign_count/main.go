// Задание 6. Количество положительных чисел, отрицательных чисел и нулей.
package main

import "fmt"

type SignCounts struct {
	Positive int
	Negative int
	Zero     int
}

func countSigns(values []int) SignCounts {
	counts := SignCounts{}
	for _, value := range values {
		switch {
		case value > 0:
			counts.Positive++
		case value < 0:
			counts.Negative++
		default:
			counts.Zero++
		}
	}
	return counts
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
	counts := countSigns(values)

	fmt.Printf("Положительных: %d\n", counts.Positive)
	fmt.Printf("Отрицательных: %d\n", counts.Negative)
	fmt.Printf("Нулей: %d\n", counts.Zero)
}
