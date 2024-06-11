package utils

import (
	"server/models"
	"strings"
)

func ConvertSupportedOperations(expression string) string {

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

func allowedOperations() []models.Operation {
	return []models.Operation{
		{"plus", "+"},
		{"minus", "-"},
		{"multiplied by", "*"},
		{"divided by", "/"}}
}
