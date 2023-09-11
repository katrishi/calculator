package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, operator, b, c string

	fmt.Println("введите значение в формате: 1 + 5")
	fmt.Scanln(&a, &operator, &b, &c)

	if c != "" {
		panic("не удовлетворяет заданию - только два операнда и один оператор (+, -, /, *)")
	}

	if a == "" || operator == "" || b == "" {
		panic("не является математической операцией")
	}

	_a, okA := strconv.Atoi(a)
	_b, okB := strconv.Atoi(b)

	if okA == nil && okB == nil {
		if (_a > 0 && _a <= 10) && (_b > 0 && _b <= 10) {
			res := strconv.Itoa(operate(_a, _b, operator))
			fmt.Println(res)
		} else {
			fmt.Println("только целые положительные числа от 1 до 10 включительно")
		}
	}
}

func operate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	default:
		panic("только +, -, /, *")
	}
}
