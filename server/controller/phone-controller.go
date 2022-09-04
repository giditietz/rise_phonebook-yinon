package controller

import (
	"phonebook/server/entities"
	"phonebook/server/service"
)

type PhoneController interface {
	GetContactPhones(contactId int) ([]entities.PhoneQuery, error)
	SaveBulk(contactID int, phones []entities.PhoneRequestBody) error
	Save(contactID int, phone *entities.PhoneRequestBody) error
	Edit(phone *entities.PhoneRequestBody) error
}

type phoneController struct {
	service service.PhoneService
}

func NewPhoneController(service service.PhoneService) *phoneController {
	return &phoneController{
		service: service,
	}
}

func (controller *phoneController) GetContactPhones(contactId int) ([]entities.PhoneQuery, error) {
	return controller.service.FindContactPhones(contactId)
}

func (controller *phoneController) SaveBulk(contactID int, phones []entities.PhoneRequestBody) error {
	return controller.service.SaveBulk(contactID, phones)
}

func (controller *phoneController) Save(contactID int, phone *entities.PhoneRequestBody) error {
	return controller.service.Save(contactID, phone)
}

func (controller *phoneController) Edit(phone *entities.PhoneRequestBody) error {
	return controller.service.Edit(phone)
}
