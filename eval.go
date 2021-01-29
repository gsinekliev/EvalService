package main

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
	ErrorOther                   Status = "General Error"
)

func computeExpression(expression string) (float64, Status) {
	// What is 5?
	// What is 5 multiplied by 10?
	// What is 5 plus 6?
	// What is 6 minus 2?
	// What is 7 divided by 2?
	var template = regexp.MustCompile(`What is ([a-z A-Z0-9]+)\?$`)
	var numberExpression = template.FindStringSubmatch(expression)[1]

	numberExpression = strings.ReplaceAll(numberExpression, "multiplied by", "*")
	numberExpression = strings.ReplaceAll(numberExpression, "divided by", "/")
	numberExpression = strings.ReplaceAll(numberExpression, "minus", "-")
	numberExpression = strings.ReplaceAll(numberExpression, "plus", "+")

	var parts = strings.Split(numberExpression, " ")

	var result float64 = 0.0
	var operation = ""
	for _, part := range parts {
		if number, err := strconv.ParseFloat(part, 64); err == nil {
			if operation == "" {
				result = number
			} else {
				switch operation {
				case "+":
					result += number
				case "-":
					result -= number
				case "*":
					result *= number
				case "/":
					result /= number
				}
				operation = ""
			}
		} else {
			// part is operation
			operation = part
		}
	}

	return result, NoError
}

func validateExpression(expression string) Status {
	// What is <number>(<operator> <number>)* ?
	// <operator> = plus|minus|multiplied by|divided by
	// <number> = any integer
	var template = regexp.MustCompile(`What is ([a-z A-Z0-9]+)\?$`)
	var subTemplate = regexp.MustCompile(`\d+( (multiplied by|minus|plus|divided by) \d+)*`)
	var matches = template.FindStringSubmatch(expression)
	if len(matches) != 2 {
		return ErrorNoQuestion
	}

	subExpression := matches[1]
	var isMatch = subTemplate.MatchString(subExpression)
	if !isMatch {
		return ErrorInvalidNumberExpression
	}

	if strings.Contains(subExpression, "divided by 0") {
		return ErrorDivisionByZero
	}

	return NoError
}
