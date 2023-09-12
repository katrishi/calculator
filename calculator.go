package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, operator, b, c string

	fmt.Println("введите значение в формате: 1 + 5 или V * VII")
	fmt.Scanln(&a, &operator, &b, &c)

	if c != "" {
		panic("не удовлетворяет заданию - только два операнда и один оператор (+, -, /, *)")
	}

	if a == "" || operator == "" || b == "" {
		panic("не является математической операцией")
	}

	if strings.ContainsRune(a, '.') || strings.ContainsRune(b, '.') {
		panic("только целые числа")
	}

	_a, okA := strconv.Atoi(a)
	_b, okB := strconv.Atoi(b)

	if okA != nil && okB != nil {
		_a = romanToArabic(a)
		_b = romanToArabic(b)
		res := operate(_a, _b, operator)

		if res < 1 {
			panic(errors.New("в римской системе нет отрицательных чисел"))
		} else {
			fmt.Println(arabicToRome(res))
		}
	} else if okA == nil && okB == nil {
		if _a > 0 && _a <= 10 && _b > 0 && _b <= 10 {
			res := strconv.Itoa(operate(_a, _b, operator))
			fmt.Println(res)
		} else {
			fmt.Println("только целые положительные числа от 1 до 10 включительно")
		}
	} else {
		panic("только арабские или римские цифры")
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

func romanToArabic(str string) int {
	switch strings.ToUpper(str) {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		panic("неккоректный ввод - римские числа от I до X")
	}
}

func arabicToRome(num int) string {
	res := ""
	for num > 0 {
		if num == 100 {
			num -= 100
			res += "C"
			continue
		}
		if num >= 90 {
			num -= 90
			res += "XC"
			continue
		}
		if num >= 50 {
			num -= 50
			res += "l"
			continue
		}
		if num >= 40 {
			num -= 40
			res += "XL"
			continue
		}
		if num >= 10 {
			num -= 10
			res += "X"
			continue
		}
		if num == 9 {
			num -= 9
			res += "IX"
			continue
		}
		if num >= 5 {
			num -= 5
			res += "V"
			continue
		}
		if num == 4 {
			num -= 4
			res += "IV"
			continue
		}
		if num >= 1 {
			num -= 1
			res += "I"
			continue
		}
	}
	return res
}
