package service

import (
	"database/sql"
	"fmt"
	"phonebook/server/entities"
	serverutils "phonebook/server/server-utils"
	"strconv"

	"phonebook/server/setup"
)

type ContactService interface {
	FindAll(offset int, limit int) ([]entities.ContactResponseBody, error)
	Save(newContact *entities.ContactRequestBody) (int, error)
	Delete(contactID int) error
	Edit(updateContact *entities.ContactRequestBody, contactID int) error
	Search(firstName string, lastName string, pageNum int) ([]entities.ContactResponseBody, error)
	GetContactNum() (int, error)
}

type contactService struct {
}

func NewContactService() *contactService {
	return &contactService{}
}

func (contactService *contactService) FindAll(offset int, limit int) ([]entities.ContactResponseBody, error) {
	db := setup.GetDBConn()
	getAllQuery, err := serverutils.GetQuery(sqlQueryGetAll)
	if err != nil {
		return nil, err
	}

	limitQuery, err := serverutils.GetLimitQuery(offset, limit)
	if err != nil {
		return nil, err
	}

	getAllQuery += limitQuery
	contactRows, err := db.Query(getAllQuery)
	if err != nil {
		return nil, err
	}

	defer contactRows.Close()

	var contacts []entities.ContactResponseBody

	for contactRows.Next() {
		contacts, err = getContactFromQuery(contactRows, contacts)
		if err != nil {

			return nil, err
		}
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

func (contactService *contactService) Search(firstName string, lastName string, pageNum int) ([]entities.ContactResponseBody, error) {
	db := setup.GetDBConn()
	searchQuery, err := serverutils.GetQuery(sqlQueryGetAll)
	if err != nil {
		return nil, err
	}

	parameterQuery, err := prepareSearchParameterQuery(firstName, lastName)
	if err != nil {
		return nil, err
	}
	searchQuery += parameterQuery

	limitQuery, err := serverutils.GetLimitQuery(pageNum*retrieveResultLimit, retrieveResultLimit)
	if err != nil {
		return nil, err
	}
	searchQuery += limitQuery

	contactRows, err := db.Query(searchQuery)
	if err != nil {
		return nil, err
	}
	defer contactRows.Close()

	var contacts []entities.ContactResponseBody

	for contactRows.Next() {
		contacts, err = getContactFromQuery(contactRows, contacts)
		if err != nil {
			return nil, err
		}
	}
	return contacts, nil
}

func prepareSearchParameterQuery(firstName string, lastName string) (string, error) {
	var ret string
	whereQuery, err := serverutils.GetQuery(sqlQueryWhere)
	if err != nil {
		return "", err
	}
	orQuery, err := serverutils.GetQuery(sqlQueryOr)
	if err != nil {
		return "", err
	}
	isFirstNameSearch := false
	if firstName != "" {
		ret += whereQuery
		firstNameRegexQuery := fmt.Sprintf("'%s'", firstName)
		ret += serverutils.RegexQuery(firstNameFieldInDB, firstNameRegexQuery)
	}
	if lastName != "" {
		if isFirstNameSearch {
			ret += orQuery
		} else {
			ret += whereQuery
		}
		lastNameRegexQuery := fmt.Sprintf("'%s'", lastName)
		ret += serverutils.RegexQuery(lastNameFieldInDB, lastNameRegexQuery)
	}
	return ret, nil
}

func updateAddresses(contact *entities.ContactResponseBody) error {
	addressService := NewAddressService()

	addressQuery, err := addressService.FindContactAddresses(contact.ContactID)
	if err != nil {
		return err
	}
	for _, address := range addressQuery {
		addressResponse := parseAddressQueryToResponse(&address)
		contact.AddressRes = append(contact.AddressRes, *addressResponse)
	}
	return nil
}

func updatePhones(contact *entities.ContactResponseBody) error {
	phoneService := NewPhoneService()

	phoneQuery, err := phoneService.FindContactPhones(contact.ContactID)
	if err != nil {
		return err
	}
	for _, phone := range phoneQuery {
		phoneResponse := parsePhoneQueryToResponse(&phone)
		contact.PhoneRes = append(contact.PhoneRes, *phoneResponse)
	}
	return nil
}

func getContactFromQuery(row *sql.Rows, contacts []entities.ContactResponseBody) ([]entities.ContactResponseBody, error) {
	var contact entities.ContactResponseBody

	if err := row.Scan(&contact.ContactID, &contact.FirstName, &contact.LastName); err != nil {
		return nil, err
	}

	if err := updateAddresses(&contact); err != nil {
		return nil, err
	}

	if err := updatePhones(&contact); err != nil {
		return nil, err
	}

	contacts = append(contacts, contact)
	return contacts, nil
}

func (contactService *contactService) Save(newContact *entities.ContactRequestBody) (int, error) {
	db := setup.GetDBConn()

	insertContactQuery, err := serverutils.GetQuery(sqlQueryInsertContact)
	if err != nil {
		fmt.Println("f: ", err)
		return 0, err
	}

	result, err := db.Exec(insertContactQuery, newContact.FirstName, newContact.LastName)
	if err != nil {
		fmt.Println("s: ", err)
		return 0, err
	}

	contactID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("t: ", err)
		return 0, err
	}
	addressService := NewAddressService()

	err = addressService.SaveBulk(int(contactID), newContact.AddressReq)
	if err != nil {
		fmt.Println("f: ", err)
		return 0, err
	}

	phoneService := NewPhoneService()
	err = phoneService.SaveBulk(int(contactID), newContact.PhoneReq)
	if err != nil {
		fmt.Println("f2: ", err)
		return 0, err
	}

	return int(contactID), nil
}

func (contactService *contactService) Delete(contactID int) error {
	db := setup.GetDBConn()

	deleteContactQuery, err := serverutils.GetQuery(sqlQueryDeleteContact)
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteContactQuery, contactID)
	if err != nil {
		return err
	}

	return nil
}

func (contactService *contactService) Edit(updateContact *entities.ContactRequestBody, contactID int) error {
	db := setup.GetDBConn()

	editContactQuery, err := serverutils.GetQuery(sqlQueryEditContact)
	if err != nil {
		return err
	}

	editContactQuery += prepareContactUpdateQuery(updateContact)

	whereQuery, err := getWhereCond(contactIdFieldInDB, strconv.FormatInt(int64(contactID), 10))
	if err != nil {
		return err
	}

	editContactQuery += whereQuery

	if _, err = db.Exec(editContactQuery); err != nil {
		return err
	}

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

func (contactService *contactService) GetContactNum() (int, error) {
	db := setup.GetDBConn()
	var ret int
	queryGetNumOfContacts, err := serverutils.GetQuery(sqlQueryGetNumOfContacts)
	if err != nil {
		return 0, err
	}

	row := db.QueryRow(queryGetNumOfContacts)
	err = row.Scan(&ret)
	if err != nil {
		return 0, err
	}
	return ret, nil

}
