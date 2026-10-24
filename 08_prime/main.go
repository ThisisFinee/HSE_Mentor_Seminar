// Задание 8. Проверка простого числа функцией isPrime.
package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for d := 3; d <= n/d; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int

	fmt.Print("Введите число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: ожидалось целое число")
		return
	}

	if isPrime(n) {
		fmt.Printf("%d — простое число\n", n)
	} else {
		fmt.Printf("%d — не является простым числом\n", n)
	}
}
