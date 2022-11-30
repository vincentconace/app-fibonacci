package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciOk(t *testing.T) {
	init := 0
	limit := 20
	index := 6

	expectedResult := 8

	service := NewService()

	fNumber, err := service.CalculateFibonacci(init, limit, index)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, fNumber)
}

func TestFibonacciInitNotNegative(t *testing.T) {
	init := -1
	limit := 20
	index := 6

	expectedResult := 0

	service := NewService()

	fNumber, err := service.CalculateFibonacci(init, limit, index)

	assert.Error(t, err)
	assert.Equal(t, expectedResult, fNumber)
}

func TestFibonacciLimitNotNegative(t *testing.T) {
	init := 0
	limit := -1
	index := 6

	expectedResult := 0

	service := NewService()

	fNumber, err := service.CalculateFibonacci(init, limit, index)

	assert.Error(t, err)
	assert.Equal(t, expectedResult, fNumber)
}

func TestFibonacciLimitLessThanInit(t *testing.T) {
	init := 4
	limit := 2
	index := 6

	expectedResult := 0

	service := NewService()

	fNumber, err := service.CalculateFibonacci(init, limit, index)

	assert.Error(t, err)
	assert.Equal(t, expectedResult, fNumber)
}

func TestFibonacciIndexOutOfRange(t *testing.T) {
	init := 0
	limit := 20
	index := 21

	expectedResult := 0

	service := NewService()

	fNumber, err := service.CalculateFibonacci(init, limit, index)

	assert.Error(t, err)
	assert.Equal(t, expectedResult, fNumber)
}
