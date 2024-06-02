package service

import (
	"Calculator/models"
	"errors"
)

type ValidateResultDTO struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

func Validate(expressionString string) ValidateResultDTO {
	_, err := Evaluate(expressionString)

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

func validateMathQuestion(expression models.Expression) error {

	if expression.EmptyExpression() {
		return errors.New("non-math question")
	}

	return nil
}

func validateCorrectSyntax(expression models.Expression) error {
	if len(expression.Operations)+1 != len(expression.Numbers) {
		return errors.New("invalid syntax")
	}

	return nil
}

func performValidations(expression models.Expression) error {
	err := validateSupportedOperations(expression)

	if err == nil {
		err = validateMathQuestion(expression)
	}

	if err == nil {
		err = validateCorrectSyntax(expression)
	}

	return err
}
