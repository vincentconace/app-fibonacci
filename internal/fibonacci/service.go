package fibonacci

import (
	"errors"
	"fmt"
)

type Service interface {
	CalculateFibonacci(init, limit, index int) (int, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

var (
	errorLimit             = errors.New("the limit can't be a number negative or equal to 0")
	errorInit              = errors.New("the init can't be a number negative")
	errorInitLessThanLimit = errors.New("the init number must be a number less than the limit")
	errorIndex             = errors.New("the index must be between the range of the init number and the limit")
	errorIndexNotFound     = errors.New("index couldn't be found")
)

func (s *service) CalculateFibonacci(init, limit, index int) (int, error) {
	if limit <= 0 {
		return 0, errorLimit
	}

	if init < 0 {
		return 0, errorInit
	}

	if init >= limit {
		return 0, errorInitLessThanLimit
	}

	if index < init || index > limit {
		return 0, errorIndex
	}

	var fibonacci = []int{0}
	var aux, fib int = 1, 0

	for i := init; i <= limit; i++ {
		aux += fib
		fib = aux - fib
		fibonacci = append(fibonacci, fib)
	}
	fmt.Println(fibonacci)
	for k, v := range fibonacci {
		if k == index {
			return v, nil
		}
	}

	return 0, errorIndexNotFound
}
