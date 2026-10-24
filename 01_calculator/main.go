// Задание 1. Калькулятор с проверкой деления на ноль.
package main

import (
	"errors"
	"fmt"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестная операция %q", op)
	}
}

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Ошибка: ожидалось число")
		return
	}

	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка: ожидалось число")
		return
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("Ошибка: ожидалась операция")
		return
	}

	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Printf("Результат: %g\n", result)
}
