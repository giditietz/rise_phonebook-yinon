package serverutils

import "regexp"

func IsStringEmpty(s string) bool {
	return s == ""
}

func IsValidPhoneNumber(phoneNumber string) bool {
	validPhone := regexp.MustCompile(`^\+?[0-9]{3}-?[0-9]{6,12}$`)
	return validPhone.MatchString(phoneNumber)
}

func IsOnlyNumber(s string) bool {
	validNumber := regexp.MustCompile(`^[0-9]*$`)
	return validNumber.MatchString(s)
}
