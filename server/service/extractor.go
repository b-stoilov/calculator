package service

import (
	"errors"
	"regexp"
	"server/models"
	"server/utils"
	"strconv"
	"strings"
)

var beginningSentencePattern = regexp.MustCompile(`What is`)
var unsupportedOperations = []string{"squared", "cubed", "square root"}

func ExtractExpression(expressionString string) (models.Expression, error) {
	expressionString = utils.ConvertSupportedOperations(expressionString)

	// validate if pattern of question is correct
	err := performPreExpressionValidations(expressionString)
	if err != nil {
		return models.Expression{}, err
	}

	expressionString = strings.TrimSuffix(expressionString, "?")
	wordsArray := strings.Fields(expressionString)[2:]

	numbers, operations, err := extractNumbersAndOperations(wordsArray)
	if err != nil {
		return models.Expression{}, err
	}

	expression := models.Expression{Numbers: numbers, Operations: operations}

	err = performPostExpressionValidations(expression)
	if err != nil {
		return models.Expression{}, err
	}

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
