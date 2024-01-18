package phonenumberutil

import (
	"fmt"
	"regexp"
)

const (
	PHONE_MIN_LENGTH       = 10
	PHONE_MAX_LENGTH       = 13
	PHONE_NON_NUMBER_REGEX = `[^0-9]`
)

func ParsePhoneNumber(number string) (string, error) {
	number = stripPhoneNumber(number)
	numLength := len(number)
	if !(numLength >= PHONE_MIN_LENGTH && numLength < PHONE_MAX_LENGTH) {
		return "", fmt.Errorf("invalid phone number")
	}
	return number, nil
}

func stripPhoneNumber(number string) string {
	phoneSymbolRegex := regexp.MustCompile(PHONE_NON_NUMBER_REGEX)
	number = phoneSymbolRegex.ReplaceAllString(number, "")
	return number
}
