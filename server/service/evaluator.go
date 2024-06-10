package service

import (
	"server/models"
	"server/store"
)

type EvaluateResultDTO struct {
	Result int `json:"result,omitempty"`
}

func Evaluate(expressionString string, store *store.Store, url string) (EvaluateResultDTO, error) {
	expression, err := ExtractExpression(expressionString)

	if err != nil {
		store.Set(
			expressionString,
			models.Error{Expression: expressionString, URL: url, Frequency: 1, Type: err.Error()})

		return EvaluateResultDTO{}, err
	}

	result, err := Calculate(expression)

	if err != nil {
		return EvaluateResultDTO{}, err
	}

	return EvaluateResultDTO{result}, nil
}
