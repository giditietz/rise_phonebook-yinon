package controller

import (
	"fmt"
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
			fmt.Println("first val")
			return &ValidationError{}
		}
		if serverutils.IsStringEmpty(contact.LastName) {
			fmt.Println("sec val")
			return &ValidationError{}
		}
	}
	for _, address := range contact.AddressReq {
		if err := ValidateAddress(address); err != nil {
			fmt.Println("add val")
			return &ValidationError{}
		}
	}
	for _, phone := range contact.PhoneReq {
		if err := ValidatePhone(phone); err != nil {
			fmt.Println("phone val")
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
