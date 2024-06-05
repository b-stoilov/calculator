package service

import (
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

func validateCorrectPattern(expression string) error {
	if !beginningSentencePattern.MatchString(expression) {
		return errors.New("non-math question")
	}

	return nil
}

func validateSyntax(expression string) error {
	if !strings.HasSuffix(expression, "?") {
		return errors.New("invalid syntax")
	}

	return nil
}

func performValidations(expression string) error {
	// check question begins with "What is..."
	err := validateCorrectPattern(expression)

	if err == nil {
		err = validateSyntax(expression)
	}

	return err
}
