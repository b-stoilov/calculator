package service

type EvaluateResultDTO struct {
	Result int `json:"result,omitempty"`
}

func Evaluate(expressionString string) (EvaluateResultDTO, error) {
	expression := ExtractExpression(expressionString)

	err := performValidations(expression)

	if err != nil {
		return EvaluateResultDTO{}, err
	}

	result := Calculate(expression)

	return EvaluateResultDTO{result}, nil
}
