package service

import (
	"phonebook/server/entities"
	serverutils "phonebook/server/server-utils"
	"phonebook/server/setup"
	"strconv"
)

type AddressService interface {
	FindContactAddresses(contactId int) ([]entities.AddressResponseBody, error)
	SaveBulk(contactID int, addresses []entities.AddressRequestBody) error
	Save(contactID int, address *entities.AddressRequestBody) error
	Edit(address *entities.AddressRequestBody) error
}

type addressService struct {
}

func NewAddressService() *addressService {
	return &addressService{}
}

func (addressService *addressService) FindContactAddresses(contactId int) ([]entities.AddressQuery, error) {
	db := setup.GetDBConn()

	getContactAddressQuery, err := serverutils.GetQuery(sqlQueryGetContactAddress)
	if err != nil {
		return nil, err
	}
	var addresses []entities.AddressQuery

	addressRows, err := db.Query(getContactAddressQuery, contactId)
	if err != nil {
		return nil, err
	}

	defer addressRows.Close()

	for addressRows.Next() {
		var address entities.AddressQuery
		if err := addressRows.Scan(&address.AddressID, &address.ContactID, &address.Description,
			&address.City, &address.Street, &address.HomeNumber, &address.Apartment); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (addressService *addressService) SaveBulk(contactID int, addresses []entities.AddressRequestBody) error {
	db := setup.GetDBConn()

	insertAddressQuery, err := serverutils.GetQuery(sqlQueryInsertAddress)
	if err != nil {
		return err
	}

	for _, address := range addresses {
		_, err := db.Exec(insertAddressQuery, contactID, address.Description, address.City,
			address.Street, address.HomeNumber, address.Apartment)
		if err != nil {
			return err
		}
	}
	return nil
}

func (addressService *addressService) Save(contactID int, address *entities.AddressRequestBody) error {
	db := setup.GetDBConn()
	insertAddressQuery, err := serverutils.GetQuery(sqlQueryInsertAddress)
	if err != nil {
		return err
	}

	_, err = db.Exec(insertAddressQuery, contactID, address.Description,
		address.City, address.Street,
		address.HomeNumber, address.Apartment)

	return err
}

func (addressService *addressService) Edit(address *entities.AddressRequestBody) error {
	db := setup.GetDBConn()

	editAddressQuery, err := prepareAddressUpdateQuery(address)
	if err != nil {
		return err
	}
	whereQuery, err := getWhereCond(addressIdFieldInDB, strconv.FormatInt(int64(address.AddressID), 10))
	if err != nil {
		return err
	}
	editAddressQuery += whereQuery
	_, err = db.Exec(editAddressQuery)

	return err
}

func prepareAddressUpdateQuery(address *entities.AddressRequestBody) (string, error) {
	ret, err := serverutils.GetQuery(sqlQueryEditAddress)
	if err != nil {
		return "", err
	}

	var isSeparatorNeeded bool
	if address.Description != "" {
		ret += serverutils.AddValuesToQuery(descriptionFieldInDB, address.Description)
		isSeparatorNeeded = true
	}
	if address.City != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
			isSeparatorNeeded = true
		}
		ret += serverutils.AddValuesToQuery(cityFieldInDB, address.City)
	}
	if address.Street != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
			isSeparatorNeeded = true
		}
		ret += serverutils.AddValuesToQuery(streetFieldInDB, address.Street)
	}
	if address.HomeNumber != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
			isSeparatorNeeded = true
		}
		ret += serverutils.AddValuesToQuery(homeNumberFieldInDB, address.HomeNumber)
	}
	if address.Apartment != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
			isSeparatorNeeded = true
		}
		ret += serverutils.AddValuesToQuery(apartmentFieldInDB, address.Apartment)
	}

	return ret, nil
}
