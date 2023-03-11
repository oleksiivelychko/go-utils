package mocking

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type SomeServiceMock struct {
	mock.Mock
}

func (m *SomeServiceMock) SomeServiceMethod(value int) bool {
	fmt.Printf("Mocked SomeServiceMethod was called: value is %d\n", value)
	// Called tells the mock object that a method has been called, and gets an array
	// of arguments to return.  Panics if the call is unexpected (i.e. not preceded by
	// appropriate .On .Return() calls)
	args := m.Called(value)
	// Bool gets the argument at the specified index. Panics if there is no argument, or
	// if the argument is of the wrong type.
	return args.Bool(0)
}

func TestMockSomeServiceMethod(t *testing.T) {
	someServiceMock := new(SomeServiceMock)
	someServiceMock.On("SomeServiceMethod", 100).Return(true)

	myService := MyService{someServiceMock}
	_ = myService.MySomeServiceMethod(100)

	// verify that myService.MySomeServiceMethod called mocked SomeServiceMethod
	someServiceMock.AssertExpectations(t)
}
