package service

import (
	"phonebook/server/entities"
	serverutils "phonebook/server/server-utils"
	"phonebook/setup"
	"strconv"
)

type PhoneService interface {
	FindContactPhones(contactId int) ([]entities.PhoneQuery, error)
	SaveBulk(contactID int, phones []entities.PhoneRequestBody) error
	Save(contactID int, phone *entities.PhoneRequestBody) error
	Edit(phone *entities.PhoneRequestBody) error
}

type phoneService struct {
}

func NewPhoneService() *phoneService {
	return &phoneService{}
}

func (phoneService *phoneService) FindContactPhones(contactId int) ([]entities.PhoneQuery, error) {
	db := setup.GetDBConn()

	getContactPhoneQuery, _ := serverutils.GetQuery(sqlQueryGetContactPhones)
	var phones []entities.PhoneQuery

	phonesRows, err := db.Query(getContactPhoneQuery, contactId)
	if err != nil {
		return nil, err
	}

	defer phonesRows.Close()

	for phonesRows.Next() {
		var phone entities.PhoneQuery
		if err := phonesRows.Scan(&phone.PhoneID, &phone.ContactID, &phone.Description,
			&phone.PhoneNumber); err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}

	return phones, nil
}

func (phoneService *phoneService) SaveBulk(contactID int, phones []entities.PhoneRequestBody) error {
	db := setup.GetDBConn()
	insertPhoneQuery, _ := serverutils.GetQuery(sqlQueryInsertPhone)
	for _, phone := range phones {
		_, err := db.Exec(insertPhoneQuery, contactID, phone.Description, phone.PhoneNumber)
		if err != nil {
			return err
		}
	}

	return nil
}

func (phoneService *phoneService) Save(contactID int, phone *entities.PhoneRequestBody) error {
	db := setup.GetDBConn()
	insertPhoneQuery, _ := serverutils.GetQuery(sqlQueryInsertPhone)
	_, err := db.Exec(insertPhoneQuery, contactID, phone.Description, phone.PhoneNumber)

	return err
}

func (phoneService *phoneService) Edit(phone *entities.PhoneRequestBody) error {
	db := setup.GetDBConn()
	editPhoneQuery := preparePhoneUpdateQuery(phone)
	editPhoneQuery += getWhereCond(phoneIdFieldInDB, strconv.FormatInt(int64(phone.PhoneID), 10))
	_, err := db.Exec(editPhoneQuery)

	return err
}

func preparePhoneUpdateQuery(phone *entities.PhoneRequestBody) string {
	ret, _ := serverutils.GetQuery(sqlQueryEditPhone)
	var isSeparatorNeeded bool
	if phone.Description != "" {
		ret += serverutils.AddValuesToQuery(descriptionFieldInDB, phone.Description)
		isSeparatorNeeded = true
	}
	if phone.PhoneNumber != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
		}
		ret += serverutils.AddValuesToQuery(phoneNumberFieldInDB, phone.PhoneNumber)
		isSeparatorNeeded = true
	}
	return ret
}
