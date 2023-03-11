package validator_helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type TestStruct struct {
	Count int `json:"count" validate:"required,countValidator"`
}

func TestValidation(t *testing.T) {
	v := NewValidation()

	err := v.Validate(TestStruct{Count: 1})
	if err == nil {
		t.Error("TestStruct.Count validation failed")
	}

	err = v.Validate(TestStruct{Count: -1})
	if err != nil {
		fmt.Printf("TestStruct.Count validation failed: %s\n", err.Errors())
	} else {
		t.Error("TestStruct.Count validation did not invoke")
	}
}

func NewValidation() *Validation {
	newValidator := validator.New()
	_ = newValidator.RegisterValidation("countValidator", validateCount)
	return &Validation{newValidator}
}

func validateCount(field validator.FieldLevel) bool {
	return field.Field().Int() >= 0
}
