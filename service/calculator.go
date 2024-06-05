package service

import (
	"Calculator/models"
	"fmt"
)

func Calculate(expression models.Expression) int {
	var result int

	if len(expression.Operations) == 0 {
		return expression.Numbers[0]
	}

	for _, operation := range expression.Operations {
		numbers := expression.Numbers

		x := numbers[0]
		y := numbers[1]

		switch operation {
		case "+":
			result = add(x, y)
			numbers[1] = result
		case "-":
			result = subtract(x, y)
			numbers[1] = result
		case "*":
			result = multiply(x, y)
			numbers[1] = result
		case "/":
			result = divide(x, y)
			numbers[1] = result
		default:
			fmt.Println("unsupported operation")
		}

		expression.Numbers = expression.Numbers[1:]
	}

	return result
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
