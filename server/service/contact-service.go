package service

import (
	"phonebook/server/entities"
	serverutils "phonebook/server/server-utils"
	"strconv"

	"phonebook/setup"
)

type ContactService interface {
	FindAll(offset int, limit int) (map[int]entities.ContactResponseBody, error)
	Save(newContact *entities.ContactRequestBody) (int, error)
	Delete(contactID int) error
	Edit(updateContact *entities.ContactRequestBody, contactID int) error
	Search(firstName string, lastName string, pageNum int) ([]entities.ContactResponseBody, error)
}

type contactService struct {
}

func NewContactService() *contactService {
	return &contactService{}
}

func (contactService *contactService) FindAll(offset int, limit int) (map[int]entities.ContactResponseBody, error) {
	db := setup.GetDBConn()
	getAllQuery, _ := serverutils.GetQuery(sqlQueryGetAll)

	getAllQuery += serverutils.GetLimitQuery(offset, limit)
	contactRows, err := db.Query(getAllQuery)
	if err != nil {
		return nil, err
	}

	defer contactRows.Close()

	contacts := make(map[int]entities.ContactResponseBody)

	for contactRows.Next() {
		var contact entities.ContactResponseBody

		if err := contactRows.Scan(&contact.ContactID, &contact.FirstName, &contact.LastName); err != nil {
			return nil, err
		}

		addressService := NewAddressService()
		phoneService := NewPhoneService()

		addressQuery, err := addressService.FindContactAddresses(contact.ContactID)
		if err != nil {
			return nil, err
		}
		for _, address := range addressQuery {
			addressResponse := parseAddressQueryToResponse(&address)
			contact.AddressRes = append(contact.AddressRes, *addressResponse)
		}
		phoneQuery, err := phoneService.FindContactPhones(contact.ContactID)
		if err != nil {
			return nil, err
		}
		for _, phone := range phoneQuery {
			phoneResponse := parsePhoneQueryToResponse(&phone)
			contact.PhoneRes = append(contact.PhoneRes, *phoneResponse)
		}

		updateContact(contacts, &contact)
	}

	return contacts, nil
}

func parseAddressQueryToResponse(address *entities.AddressQuery) *entities.AddressResponseBody {
	var ret entities.AddressResponseBody

	ret.AddressID = address.AddressID
	ret.Description = address.Description
	ret.City = address.City
	ret.Street = address.Street
	ret.HomeNumber = address.HomeNumber
	ret.Apartment = address.Apartment

	return &ret
}

func parsePhoneQueryToResponse(phone *entities.PhoneQuery) *entities.PhoneResponseBody {
	var ret entities.PhoneResponseBody

	ret.PhoneID = phone.PhoneID
	ret.Description = phone.Description
	ret.PhoneNumber = phone.PhoneNumber

	return &ret
}

func updateContact(contacts map[int]entities.ContactResponseBody, contact *entities.ContactResponseBody) {
	contacts[contact.ContactID] = *contact
}

func (contactService *contactService) Search(firstName string, lastName string, pageNum int) ([]entities.ContactResponseBody, error) {
	db := setup.GetDBConn()
	searchQuery, _ := serverutils.GetQuery(sqlQueryGetAll)
	whereQuery, _ := serverutils.GetQuery(sqlQueryWhere)
	andQuery, _ := serverutils.GetQuery(sqlQueryAnd)
	isFirstNameSearch := false

	if firstName != "" {
		searchQuery += whereQuery
		searchQuery += serverutils.AddValuesToQuery(firstNameFieldInDB, firstName)
	}
	if lastName != "" {
		if isFirstNameSearch {
			searchQuery += andQuery
		} else {
			searchQuery += whereQuery
		}
		searchQuery += serverutils.AddValuesToQuery(lastNameFieldInDB, lastName)
	}
	searchQuery += serverutils.GetLimitQuery(pageNum*retrieveResultLimit, retrieveResultLimit)

	contactRows, err := db.Query(searchQuery)
	if err != nil {
		return nil, err
	}
	defer contactRows.Close()

	var contacts []entities.ContactResponseBody

	for contactRows.Next() {
		var contact entities.ContactResponseBody

		if err := contactRows.Scan(&contact.ContactID, &contact.FirstName, &contact.LastName); err != nil {
			return nil, err
		}

		addressService := NewAddressService()
		phoneService := NewPhoneService()

		addressQuery, err := addressService.FindContactAddresses(contact.ContactID)
		if err != nil {
			return nil, err
		}
		for _, address := range addressQuery {
			addressResponse := parseAddressQueryToResponse(&address)
			contact.AddressRes = append(contact.AddressRes, *addressResponse)
		}
		phoneQuery, err := phoneService.FindContactPhones(contact.ContactID)
		if err != nil {
			return nil, err
		}
		for _, phone := range phoneQuery {
			phoneResponse := parsePhoneQueryToResponse(&phone)
			contact.PhoneRes = append(contact.PhoneRes, *phoneResponse)
		}

		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (contactService *contactService) Save(newContact *entities.ContactRequestBody) (int, error) {
	db := setup.GetDBConn()

	insertContactQuery, _ := serverutils.GetQuery(sqlQueryInsertContact)

	result, err := db.Exec(insertContactQuery, newContact.FirstName, newContact.LastName)
	if err != nil {
		return 0, err
	}

	contactID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	addressService := NewAddressService()

	err = addressService.SaveBulk(int(contactID), newContact.AddressReq)
	if err != nil {
		return 0, err
	}

	phoneService := NewPhoneService()
	err = phoneService.SaveBulk(int(contactID), newContact.PhoneReq)
	if err != nil {
		return 0, err
	}

	return int(contactID), nil
}

func (contactService *contactService) Delete(contactID int) error {
	db := setup.GetDBConn()

	deleteContactQuery, _ := serverutils.GetQuery(sqlQueryDeleteContact)

	_, err := db.Exec(deleteContactQuery, contactID)
	if err != nil {
		return err
	}

	return nil
}

func (contactService *contactService) Edit(updateContact *entities.ContactRequestBody, contactID int) error {
	db := setup.GetDBConn()

	editContactQuery, _ := serverutils.GetQuery(sqlQueryEditContact)
	editContactQuery += prepareContactUpdateQuery(updateContact)
	editContactQuery += getWhereCond(contactIdFieldInDB, strconv.FormatInt(int64(contactID), 10))

	db.Exec(editContactQuery)

	phoneService := NewPhoneService()
	for _, phone := range updateContact.PhoneReq {
		if !isPhoneExist(&phone) {
			if err := phoneService.Save(contactID, &phone); err != nil {
				return err
			}

		} else {
			if err := phoneService.Edit(&phone); err != nil {
				return err
			}
		}
	}

	addressService := NewAddressService()
	for _, address := range updateContact.AddressReq {
		if !isAddressExist(&address) {
			if err := addressService.Save(contactID, &address); err != nil {
				return err
			}
		} else {
			if err := addressService.Edit(&address); err != nil {
				return err
			}
		}
	}
	return nil
}

func prepareContactUpdateQuery(contact *entities.ContactRequestBody) string {
	ret := ""
	isSeparatorNeeded := false
	if contact.FirstName != "" {
		ret += serverutils.AddValuesToQuery(firstNameFieldInDB, contact.FirstName)
		isSeparatorNeeded = true
	}

	if contact.LastName != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
		}
		ret += serverutils.AddValuesToQuery(lastNameFieldInDB, contact.LastName)
		isSeparatorNeeded = true
	}

	return ret
}

func isPhoneExist(phone *entities.PhoneRequestBody) bool {
	return phone.PhoneID != 0
}

func isAddressExist(address *entities.AddressRequestBody) bool {
	return address.AddressID != 0
}

func getWhereCond(fieldName string, id string) string {
	var ret string
	where, _ := serverutils.GetQuery(sqlQueryWhere)
	ret += where

	ret += serverutils.AddValuesToQuery(fieldName, id)

	return ret
}
