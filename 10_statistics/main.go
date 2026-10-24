// Задание 10. Статистика последовательности из N чисел.
package main

import (
	"errors"
	"fmt"
)

type Statistics struct {
	Sum     int
	Average float64
	Min     int
	Max     int
	Even    int
	Odd     int
}

func calculateStatistics(values []int) (Statistics, error) {
	if len(values) == 0 {
		return Statistics{}, errors.New("последовательность не должна быть пустой")
	}

	stats := Statistics{Min: values[0], Max: values[0]}
	for _, value := range values {
		stats.Sum += value
		if value < stats.Min {
			stats.Min = value
		}
		if value > stats.Max {
			stats.Max = value
		}
		if value%2 == 0 {
			stats.Even++
		} else {
			stats.Odd++
		}
	}
	stats.Average = float64(stats.Sum) / float64(len(values))
	return stats, nil
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
	stats, _ := calculateStatistics(values)

	fmt.Printf("Сумма: %d\n", stats.Sum)
	fmt.Printf("Среднее значение: %.2f\n", stats.Average)
	fmt.Printf("Максимальное значение: %d\n", stats.Max)
	fmt.Printf("Минимальное значение: %d\n", stats.Min)
	fmt.Printf("Чётных чисел: %d\n", stats.Even)
	fmt.Printf("Нечётных чисел: %d\n", stats.Odd)
}
