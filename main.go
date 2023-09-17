package main

import (
	"fmt"
	"strconv"
)

var result int
var n, romannumber string

func conteinelements(arabic [10]string, roman [10]string, num string) bool {

	for _, elem := range arabic {
		if elem == num {
			return true
		}
	}
	for _, elem := range roman {
		if elem == num {
			return true
		}
	}
	return false
}

func romantoarabic(arabic [10]string, roman [10]string, rom string) {

	for i := 0; i < len(roman); i++ {
		if rom == roman[i] {
			n = arabic[i]
		}
	}
}

func arabictoroman(number int) {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, conversion := range conversions {
		for number >= conversion.value {
			romannumber += conversion.digit
			number -= conversion.value
		}
	}
	return
}

func calculate(num1, num2, op string) {
	var a, b int
	a, _ = strconv.Atoi(num1)
	b, _ = strconv.Atoi(num2)

	if op == "+" {
		result = a + b
	}
	if op == "-" {
		result = a - b
	}
	if op == "*" {
		result = a * b
	}
	if op == "/" {
		result = a / b
	}
	return
}

func main() {
	var p int
	var num1, num2, op, ex string
	arabic := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	fmt.Print("Введите 2 числа и оператор через пробелы например(5 + 5) или (V + III). Числа должны быть в диапазоне от 1 до 10 включительно\n")
	fmt.Scanln(&num1, &op, &num2, &ex)
	if ex != "" {
		fmt.Print("Операция должна состоять из 2х операндов")
		return
	}
	if op == "-" || op == "+" || op == "*" || op == "/" {
		p = 1
	}
	if p != 1 {
		fmt.Print("Ошибка: Строка не является математической операцией или нарушен формат ввода")
		return
	}
	if !conteinelements(arabic, roman, num1) || !conteinelements(arabic, roman, num2) {
		fmt.Print("Операнд должен быть арабскими или римскими числами в диапазоне от 1 до 10 ")
	}
	for i := 0; i < len(arabic); i++ {
		for j := 0; j < len(arabic); j++ {
			if num1 == arabic[i] && num2 == arabic[j] {
				calculate(num1, num2, op)
				if p == 1 {
					fmt.Print(result)
				}
			}
			if num1 == roman[i] && num2 == roman[j] {
				romantoarabic(arabic, roman, num1)
				rom1 := n
				romantoarabic(arabic, roman, num2)
				rom2 := n
				calculate(rom1, rom2, op)
				if result <= 0 {
					fmt.Print("Ошибка: В римской системе счисления нет отрицательных чисел и нуля.")
					break
				}
				arabictoroman(result)
				fmt.Print(romannumber)
			}
			if (num1 == roman[i] && num2 == arabic[j]) || (num1 == arabic[i] && num2 == roman[j]) {
				fmt.Println("Ошибка: разные типы счисления")
			}
		}
	}
}
