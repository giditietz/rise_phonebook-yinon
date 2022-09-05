package controller

import (
	"phonebook/server/entities"
	serverutils "phonebook/server/server-utils"
)

const (
	validationErrorString = "invalid Value"
)

type ValidationError struct{}

func (validError *ValidationError) Error() string {
	return validationErrorString
}

func ValidateContact(contact *entities.ContactRequestBody, isSaveOperation bool) error {
	if isSaveOperation {
		if serverutils.IsStringEmpty(contact.FirstName) {
			return &ValidationError{}
		}
		if serverutils.IsStringEmpty(contact.LastName) {
			return &ValidationError{}
		}
	}
	for _, address := range contact.AddressReq {
		if err := ValidateAddress(address); err != nil {
			return &ValidationError{}
		}
	}
	for _, phone := range contact.PhoneReq {
		if err := ValidatePhone(phone); err != nil {
			return &ValidationError{}
		}
	}
	return nil
}

func ValidateAddress(address entities.AddressRequestBody) error {
	if serverutils.IsStringEmpty(address.City) {
		return &ValidationError{}
	}
	return nil
}

func ValidatePhone(phone entities.PhoneRequestBody) error {
	if serverutils.IsStringEmpty(phone.PhoneNumber) {
		return &ValidationError{}
	}
	if !serverutils.IsValidPhoneNumber(phone.PhoneNumber) {
		return &ValidationError{}
	}
	return nil
}
