// Задание 4. Таблица умножения через функцию printTable.
package main

import "fmt"

func printTable(n int) {
	for _, line := range multiplicationTable(n) {
		fmt.Println(line)
	}
}

func multiplicationTable(n int) []string {
	table := make([]string, 0, 10)
	for i := 1; i <= 10; i++ {
		table = append(table, fmt.Sprintf("%d x %d = %d", n, i, n*i))
	}
	return table
}

func main() {
	var n int

	fmt.Print("Введите число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: ожидалось целое число")
		return
	}

	printTable(n)
}
