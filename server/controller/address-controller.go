package controller

import (
	"phonebook/server/entities"
	"phonebook/server/service"
)

type AddressController interface {
	GetContactAddresses(contactId int) ([]entities.AddressResponseBody, error)
	SaveBulk(contactID int, addresses []entities.AddressRequestBody) error
	Save(contactID int, address *entities.AddressRequestBody) error
	Edit(addresses *entities.AddressRequestBody) error
}

type addressController struct {
	service service.AddressService
}

func NewAddressController(service service.AddressService) *addressController {
	return &addressController{
		service: service,
	}
}

func (controller *addressController) GetContactAddresses(contactId int) ([]entities.AddressResponseBody, error) {
	addresses, err := controller.service.FindContactAddresses(contactId)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (controller *addressController) SaveBulk(contactID int, addresses []entities.AddressRequestBody) error {
	return controller.service.SaveBulk(contactID, addresses)
}

func (controller *addressController) Save(contactID int, address *entities.AddressRequestBody) error {
	return controller.service.Save(contactID, address)
}

func (controller *addressController) Edit(address *entities.AddressRequestBody) error {
	return controller.service.Edit(address)
}
