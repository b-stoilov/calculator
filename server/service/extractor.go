package service

import (
	"Calculator/models"
	"Calculator/utils"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var beginningSentencePattern = regexp.MustCompile(`What is`)
var unsupportedOperations = []string{"squared", "cubed", "square root"}

func ExtractExpression(expressionString string) (models.Expression, error) {
	expressionString = utils.ConvertOperations(expressionString)

	err := performValidations(expressionString)
	if err != nil {
		return models.Expression{}, err
	}

	expressionString = strings.TrimSuffix(expressionString, "?")
	trimmedExpressionArray := strings.Fields(expressionString)[2:]

	// check after what is only ints and operations
	numbers, operations, err := extractNumbersAndOperations(trimmedExpressionArray)

	if err != nil {
		return models.Expression{}, err
	}

	expression := models.Expression{Numbers: numbers, Operations: operations}

	return expression, nil
}

func extractNumbersAndOperations(wordsArray []string) ([]int, []string, error) {
	var numbers []int
	var operations []string

	for i, word := range wordsArray {
		if i%2 == 0 {
			num, err := strconv.Atoi(word)
			if err != nil {
				return nil, nil, errors.New("invalid syntax")
			}

			numbers = append(numbers, num)
		} else {
			err := validateSupportedOperation(word)
			if err != nil {
				return nil, nil, err
			}

			operations = append(operations, word)
		}
	}

	return numbers, operations, nil
}

func validateSupportedOperation(operation string) error {
	for _, unsupportedOperation := range unsupportedOperations {
		if operation == unsupportedOperation {
			return errors.New("unsupported operation")
		}
	}

	return nil
}

func splitExpressionToWords(expression string) []string {
	var words []string

	for _, word := range strings.Fields(expression) {
		words = append(words, word)
	}

	return words
}
