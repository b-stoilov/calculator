package service

import (
	"Calculator/models"
	"Calculator/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var re = regexp.MustCompile("[0-9]+")

func ExtractExpression(expressionString string) models.Expression {
	numbers := extractNumbers(expressionString)
	operations := extractOperations(expressionString)

	expression := models.Expression{Numbers: numbers, Operations: operations}

	return expression
}

func extractNumbers(expressionString string) []int {
	expressionString = utils.ConvertOperations(expressionString)
	fmt.Println(expressionString)

	numbersInString := re.FindAllString(expressionString, -1)
	numbers := utils.ConvertToIntegers(numbersInString)

	return numbers
}

func extractOperations(expression string) []string {
	var operations []string

	allowedOperations := []string{"+", "-", "/", "*"}

	for _, word := range splitExpressionToWords(expression) {
		if slices.Contains(allowedOperations, word) {
			operations = append(operations, word)
		}
	}

	return operations
}

func splitExpressionToWords(expression string) []string {
	var words []string

	for _, word := range strings.Fields(expression) {
		words = append(words, word)
	}

	return words
}
