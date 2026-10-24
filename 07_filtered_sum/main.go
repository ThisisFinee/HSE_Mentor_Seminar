// Задание 7. Сумма чисел от 1 до N, которые чётные и делятся на 3.
package main

import "fmt"

func sumEvenDivisibleByThree(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("N должно быть не меньше 1")
	}

	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	return sum, nil
}

func main() {
	var n int

	fmt.Print("Введите число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: ожидалось целое число")
		return
	}
	sum, err := sumEvenDivisibleByThree(n)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Сумма подходящих чисел от 1 до %d: %d\n", n, sum)
}
