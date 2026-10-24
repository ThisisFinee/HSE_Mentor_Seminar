// Задание 3. Сумма и среднее арифметическое чисел от 1 до N.
package main

import (
	"fmt"
)

func sumAndAverage(n int) (int, float64, error) {
	if n < 1 {
		return 0, 0, fmt.Errorf("N должно быть не меньше 1")
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum, float64(sum) / float64(n), nil
}

func main() {
	var n int

	fmt.Print("Введите число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: ожидалось целое число")
		return
	}
	sum, average, err := sumAndAverage(n)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", average)
}
