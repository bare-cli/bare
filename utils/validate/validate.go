package validate

import "errors"

var (
	ErrUnexpectedArgs = errors.New("Unexpected arguments")
	ErrNotEnoughArgs  = errors.New("Not enough arguments")
)

func ValidateArgCount(expectedArgNo, argNo int) error {
	switch {
	case expectedArgNo < argNo:
		return ErrUnexpectedArgs
	case expectedArgNo > argNo:
		return ErrNotEnoughArgs
	}

	return nil
}
