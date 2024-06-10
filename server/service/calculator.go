package service

import (
	"errors"
	"fmt"
	"server/models"
)

func Calculate(expression models.Expression) (int, error) {
	var result int

	if len(expression.Operations) == 0 {
		return expression.Numbers[0], nil
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
			var err error
			result, err = divide(x, y)

			if err != nil {
				return 0, err
			}

			numbers[1] = result
		default:
			fmt.Println("unsupported operation")
		}

		expression.Numbers = expression.Numbers[1:]
	}

	return result, nil
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

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("can't divide by 0")
	}

	return x / y, nil
}
