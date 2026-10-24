// Задание 9. Мини-банкомат: просмотр, пополнение и снятие.
package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	showBalance = 1
	deposit     = 2
	withdraw    = 3
)

func isValidAmount(amount float64) bool {
	return amount > 0 && !math.IsNaN(amount) && !math.IsInf(amount, 0)
}

func processTransaction(balance float64, action int, amount float64) (float64, error) {
	if balance < 0 || math.IsNaN(balance) || math.IsInf(balance, 0) {
		return 0, errors.New("баланс должен быть неотрицательным конечным числом")
	}

	switch action {
	case showBalance:
		return balance, nil
	case deposit:
		if !isValidAmount(amount) {
			return 0, errors.New("сумма пополнения должна быть положительной")
		}
		return roundMoney(balance + amount), nil
	case withdraw:
		if !isValidAmount(amount) {
			return 0, errors.New("сумма снятия должна быть положительной")
		}
		if amount > balance {
			return 0, errors.New("недостаточно средств на балансе")
		}
		return roundMoney(balance - amount), nil
	default:
		return 0, errors.New("неизвестное действие")
	}
}

func roundMoney(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func main() {
	var balance float64

	fmt.Print("Введите баланс: ")
	if _, err := fmt.Scan(&balance); err != nil {
		fmt.Println("Ошибка: ожидалось число")
		return
	}
	if balance < 0 || math.IsNaN(balance) || math.IsInf(balance, 0) {
		fmt.Println("Ошибка: баланс не может быть отрицательным")
		return
	}

	fmt.Println("Выберите действие:")
	fmt.Println("1 — посмотреть баланс")
	fmt.Println("2 — пополнить баланс")
	fmt.Println("3 — снять средства")
	fmt.Print("Ваш выбор: ")

	var action int
	if _, err := fmt.Scan(&action); err != nil {
		fmt.Println("Ошибка: ожидался номер действия")
		return
	}

	if action == showBalance {
		newBalance, err := processTransaction(balance, action, 0)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			return
		}
		fmt.Printf("Текущий баланс: %.2f\n", newBalance)
		return
	}

	var amount float64
	switch action {
	case deposit:
		fmt.Print("Введите сумму пополнения: ")
		if _, err := fmt.Scan(&amount); err != nil {
			fmt.Println("Ошибка: ожидалось число")
			return
		}
	case withdraw:
		fmt.Print("Введите сумму снятия: ")
		if _, err := fmt.Scan(&amount); err != nil {
			fmt.Println("Ошибка: ожидалось число")
			return
		}
	default:
		fmt.Println("Ошибка: неизвестное действие")
		return
	}

	newBalance, err := processTransaction(balance, action, amount)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if action == deposit {
		fmt.Printf("Баланс после пополнения: %.2f\n", newBalance)
	} else {
		fmt.Printf("Баланс после снятия: %.2f\n", newBalance)
	}
}
