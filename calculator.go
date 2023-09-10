package main

import "fmt"

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
}
