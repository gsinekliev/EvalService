package eval

import (
	"regexp"
	"strconv"
	"strings"
)

type Status string

const (
	NoError                      Status = "Ok"
	ErrorNoQuestion              Status = "Expression must have the template 'What is <>?'"
	ErrorInvalidNumberExpression Status = "Invalid number expression"
	ErrorDivisionByZero          Status = "Error: Division by zero"
)

type IEvaluator interface {
	ComputeExpression(string) (float64, Status)
	ValidateExpression(string) bool
}

type Evaluator struct{}

func (evaluator Evaluator) ComputeExpression(expression string) (float64, Status) {
	// <expression> = What is <number> (<operator> <number>)*?
	// <operator> = plus|minus|multiplied by|divided by
	// <number> = any integer

	if status := evaluator.ValidateExpression(expression); status != NoError {
		return 0.0, status
	}

	var template = regexp.MustCompile(`What is ([a-z A-Z0-9]+)\?$`)
	var numberExpression = template.FindStringSubmatch(expression)[1]

	numberExpression = strings.ReplaceAll(numberExpression, "multiplied by", "*")
	numberExpression = strings.ReplaceAll(numberExpression, "divided by", "/")
	numberExpression = strings.ReplaceAll(numberExpression, "minus", "-")
	numberExpression = strings.ReplaceAll(numberExpression, "plus", "+")

	var parts = strings.Split(numberExpression, " ")

	var result, _ = strconv.ParseFloat(parts[0], 64)

	for i := 1; i < len(parts); i += 2 {
		operation := parts[i]
		operand, _ := strconv.ParseFloat(parts[i+1], 64)
		switch operation {
		case "+":
			result += operand
		case "-":
			result -= operand
		case "*":
			result *= operand
		case "/":
			result /= operand
		}
	}

	return result, NoError
}

func (evaluator Evaluator) ValidateExpression(expression string) Status {
	// <expression> = <number> (<operator> <number>)*
	// <operator> = plus|minus|multiplied by|divided by
	// <number> = any integer
	var template = regexp.MustCompile(`What is ([a-z A-Z0-9]+)\?$`)
	var subTemplate = regexp.MustCompile(`\d+( (multiplied by|minus|plus|divided by) \d+)*`)
	var matches = template.FindStringSubmatch(expression)
	if len(matches) != 2 || len(matches[0]) != len(expression) {
		return ErrorNoQuestion
	}

	subExpression := matches[1]
	var subExpressionMatches = subTemplate.FindStringSubmatch(subExpression)
	if len(subExpressionMatches) < 1 || len(subExpressionMatches[0]) != len(subExpression) {
		return ErrorInvalidNumberExpression
	}

	if strings.Contains(subExpression, "divided by 0") {
		return ErrorDivisionByZero
	}

	return NoError
}
