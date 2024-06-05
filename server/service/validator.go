package service

import (
	"Calculator/models"
	"Calculator/store"
	"errors"
	"strings"
)

type ValidateResultDTO struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

func Validate(expressionString string, store *store.Store, url string) ValidateResultDTO {
	_, err := Evaluate(expressionString, store, url)

	if err != nil {
		return ValidateResultDTO{Valid: false, Reason: err.Error()}
	} else {
		return ValidateResultDTO{Valid: true}
	}

}

func validateSupportedOperations(expression models.Expression) error {
	if expression.NumbersDetectedButNotOperations() {
		return errors.New("unsupported operation")
	}

	return nil
}

func validateCorrectPattern(expression string) error {
	if !beginningSentencePattern.MatchString(expression) {
		return errors.New("non-math question")
	}

	return nil
}

func validateMathQuestion(expression models.Expression) error {

	if expression.EmptyExpression() {
		return errors.New("non-math question")
	}

	return nil
}

func validateExpressionSyntax(expression models.Expression) error {
	if len(expression.Operations)+1 != len(expression.Numbers) {
		return errors.New("invalid syntax")
	}

	return nil
}

func validateStringSyntax(expression string) error {
	if !strings.HasSuffix(expression, "?") {
		return errors.New("invalid syntax")
	}

	return nil
}

func performPreExpressionValidations(expression string) error {
	// check question begins with "What is..."
	err := validateCorrectPattern(expression)

	if err == nil {
		err = validateStringSyntax(expression)
	}

	return err
}

func performPostExpressionValidations(expression models.Expression) error {
	err := validateExpressionSyntax(expression)

	return err
}
