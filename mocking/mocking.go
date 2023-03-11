package mocking

import (
	"fmt"
)

type ISomeService interface {
	SomeServiceMethod(value int) bool
}

type SomeService struct{}

type MyService struct {
	someService ISomeService
}

func (service SomeService) SomeServiceMethod(value int) bool {
	fmt.Printf("SomeServiceMethod was called: value is %d\n", value)
	return true
}

func (service MyService) MySomeServiceMethod(value int) error {
	fmt.Printf("MySomeServiceMethod was called: value is %d\n", value)
	service.someService.SomeServiceMethod(value)
	return nil
}
