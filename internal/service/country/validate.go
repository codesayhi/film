package country

import (
	"errors"
	"strings"
)

var (
	ErrNameRequired    = errors.New("NAME_REQUIRED")
	ErrNameTooShort    = errors.New("NAME_TOO_SHORT")
	ErrCodeRequired    = errors.New("CODE_REQUIRED")
	ErrCodeTooLong     = errors.New("CODE_TOO_LONG")
	ErrPositionInvalid = errors.New("POSITION_INVALID")
)

func validateCreate(in CreateCountryInput) error {
	name := strings.TrimSpace(in.Name)
	code := strings.TrimSpace(in.Code)

	if name == "" {
		return ErrNameRequired
	}
	if len(name) < 2 {
		return ErrNameTooShort
	}
	if code == "" {
		return ErrCodeRequired
	}
	if len(code) > 20 {
		return ErrCodeTooLong
	}
	if in.Position < 0 {
		return ErrPositionInvalid
	}
	return nil
}
