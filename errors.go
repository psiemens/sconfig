package sconfig

import (
	"errors"
	"fmt"
)

var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

type ErrInvalidField struct {
	Field string
	Err   error
}

func (e *ErrInvalidField) Error() string {
	return fmt.Sprintf("invalid field %s: %s", e.Field, e.Err.Error())
}

type ErrUnsupportedFieldType struct {
	Type string
}

func (e *ErrUnsupportedFieldType) Error() string {
	return fmt.Sprintf("%s is an unsupported type", e.Type)
}

type ErrInvalidFlagFormat struct {
	Format string
}

func (e *ErrInvalidFlagFormat) Error() string {
	return fmt.Sprintf("invalid flag format \"%s\"", e.Format)
}
