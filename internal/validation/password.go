package validation

import (
	"errors"
	"unicode"
)

func ValidateUserPassword(password string) error {
	if len(password) < 10 {
		return errors.New("password must be at least 10 characters")
	}

	var hasLetter, hasDigit, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsLetter(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasLetter {
		return errors.New("password must contain at least one letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
