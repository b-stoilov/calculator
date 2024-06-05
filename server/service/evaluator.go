package service

import (
	"Calculator/models"
	"Calculator/store"
)

type EvaluateResultDTO struct {
	Result int `json:"result,omitempty"`
}

func Evaluate(expressionString string, store *store.Store, url string) (EvaluateResultDTO, error) {
	rawExpression := expressionString

	expression, err := ExtractExpression(expressionString)

	if err != nil {
		store.Set(rawExpression, models.Error{Expression: rawExpression, URL: url, Type: err.Error()})

		return EvaluateResultDTO{}, err
	}

	result, err := Calculate(expression)

	if err != nil {
		return EvaluateResultDTO{}, err
	}

	return EvaluateResultDTO{result}, nil
}
