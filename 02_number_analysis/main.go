// Задание 2. Анализ числа: знак, чётность, делимость на 3 и на 5.
package main

import "fmt"

type NumberAnalysis struct {
	Sign         string
	Even         bool
	DivisibleBy3 bool
	DivisibleBy5 bool
}

func analyzeNumber(n int) NumberAnalysis {
	analysis := NumberAnalysis{
		Even:         n%2 == 0,
		DivisibleBy3: n%3 == 0,
		DivisibleBy5: n%5 == 0,
	}
	switch {
	case n > 0:
		analysis.Sign = "положительное"
	case n < 0:
		analysis.Sign = "отрицательное"
	default:
		analysis.Sign = "равно нулю"
	}
	return analysis
}

func main() {
	var n int

	fmt.Print("Введите число N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: ожидалось целое число")
		return
	}

	analysis := analyzeNumber(n)
	fmt.Printf("Число %s\n", analysis.Sign)

	if analysis.Even {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	if analysis.DivisibleBy3 && analysis.DivisibleBy5 {
		fmt.Println("Число делится и на 3, и на 5")
	} else if analysis.DivisibleBy3 {
		fmt.Println("Число делится на 3, но не делится на 5")
	} else if analysis.DivisibleBy5 {
		fmt.Println("Число делится на 5, но не делится на 3")
	} else {
		fmt.Println("Число не делится ни на 3, ни на 5")
	}
}
