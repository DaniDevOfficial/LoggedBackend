package validation

import (
	"fmt"
	"strconv"
)

const (
	minLength = 8
)

func IsValidPassword(password string) (bool, error) {
	if len(password) < minLength {
		return false, fmt.Errorf("The Password needs to be at least " + strconv.Itoa(minLength) + " letters long")
	}

	if !checkForCharacters(password) {
		return false, fmt.Errorf("password needs Upper, lower and special characters and at least one number")
	}

	return true, nil
}

func checkForCharacters(password string) bool {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		ascii := int(char)

		if !hasUpper && ascii >= 65 && ascii <= 90 {
			hasUpper = true
		} else if !hasLower && ascii >= 97 && ascii <= 122 {
			hasLower = true
		} else if !hasDigit && ascii >= 48 && ascii <= 57 {
			hasDigit = true
		} else if !hasSpecial && ((ascii >= 33 && ascii <= 47) || (ascii >= 58 && ascii <= 64) || (ascii >= 91 && ascii <= 96) || (ascii >= 123 && ascii <= 126)) {
			hasSpecial = true
		}
		if hasUpper && hasLower && hasDigit && hasSpecial {
			break
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}
