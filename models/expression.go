package models

type Expression struct {
	Numbers    []int
	Operations []string
}

func (e Expression) EmptyExpression() bool {
	return len(e.Numbers) == 0 && len(e.Operations) == 0
}

func (e Expression) NumbersDetectedButNotOperations() bool {
	return len(e.Numbers) != 0 && len(e.Operations) == 0
}
