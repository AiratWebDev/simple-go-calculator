package main

import (
	"fmt"
)

func main() {
	operator, digits := chatScanner()

	result := calculator(operator, digits)
	fmt.Println("Полученный ответ:" + fmt.Sprintf("%f", result))
}

func chatScanner() (string, [2]float64) {
	var firstOperand float64
	var secondOperand float64
	var symbol string

	fmt.Println("Укажите первое число")
	fmt.Scan(&firstOperand)

	fmt.Println("Укажите знак")
	fmt.Scan(&symbol)

	fmt.Println("Укажите второе число")
	fmt.Scan(&secondOperand)

	digits := [2]float64{firstOperand, secondOperand}

	return symbol, digits
}

func calculator(operator string, digits [2]float64) float64 {

	firstDigital := digits[0]
	secondDigital := digits[1]
	var result float64

	switch operator {
	case "+":
		result = firstDigital + secondDigital
	case "-":
		result = firstDigital - secondDigital
	case "*":
		result = firstDigital * secondDigital
	case "/":
		result = firstDigital / secondDigital
	default:
		fmt.Println("Введен неверный оператор")
	}

	return result
}
