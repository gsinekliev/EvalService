package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenValidExpressionValidateReturnsTrue(t *testing.T) {
	status := validateExpression("What is 5 multiplied by 7?")
	assert.Equal(t, status, NoError)
}

func TestDivisionByZero(t *testing.T) {
	status := validateExpression("What is 5 multiplied by 7 divided by 0?")
	assert.Equal(t, status, ErrorDivisionByZero)
}

func TestComputeExpression(t *testing.T) {
	result, status := computeExpression("What is 5 multiplied by 7?")
	assert.Equal(t, status, NoError)
	assert.Equal(t, result, 35.0)
}

func TestComputeExpressionFloat(t *testing.T) {
	result, status := computeExpression("What is 5 multiplied by 7 divided by 2?")
	assert.Equal(t, status, NoError)
	assert.Equal(t, result, 17.5)
}
