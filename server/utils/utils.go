package utils

import (
	"Calculator/models"
	"strconv"
	"strings"
)

func ConvertOperations(expression string) string {

	for i := range allowedOperations() {
		operation := allowedOperations()[i]

		if strings.Contains(expression, operation.StringRepresentation) {
			expression = strings.Replace(expression,
				operation.StringRepresentation, operation.MathematicalRepresentation,
				-1)
		}
	}

	return expression
}

func ConvertToIntegers(numbersInString []string) []int {
	numbers := make([]int, len(numbersInString))

	for i := range numbersInString {
		numbers[i], _ = strconv.Atoi(numbersInString[i])
	}

	return numbers
}

func allowedOperations() []models.Operation {
	return []models.Operation{
		{"plus", "+"},
		{"minus", "-"},
		{"multiplied by", "*"},
		{"divided by", "/"}}
}
