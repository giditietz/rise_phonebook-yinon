package server

import (
	"fmt"
	"net/http"
	"phonebook/server/entities"
	serverutils "phonebook/server/server_utils"
	"phonebook/setup"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	sqlQueryGetAll        = "getAll"
	sqlQueryWhere         = "where"
	sqlQueryAnd           = "and"
	sqlQueryInsertContact = "insertContact"
	sqlQueryInsertAddress = "insertAddress"
	sqlQueryInsertPhone   = "insertPhone"
	sqlQueryDeleteContact = "deleteContact"
	sqlQueryEditContact   = "editContact"
	sqlQueryEditPhone     = "editPhone"
	sqlQueryEditAddress   = "editAddress"
	sqlSeparatorValues    = ", "
)

const (
	ginQueryPage        = "page"
	ginDefaultPageStart = "0"
	ginQueryFirstName   = "first_name"
	ginQueryLastName    = "last_name"
	ginParamId          = "id"
)

const (
	retrieveResultLimit = 10
)

const (
	firstNameFieldInDB   = "first_name"
	lastNameFieldInDB    = "last_name"
	contactIdFieldInDB   = "contact_id"
	phoneIdFieldInDB     = "phone_id"
	addressIdFieldInDB   = "address_id"
	descriptionFieldInDB = "description"
	phoneNumberFieldInDB = "phone_number"
	cityFieldInDB        = "city"
	streetFieldInDB      = "street"
	homeNumberFieldInDB  = "home_number"
	apartmentFieldInDB   = "apartment"
)

func GetAllContacts(c *gin.Context) {
	db := setup.GetDBConn()
	getAllQuery, _ := serverutils.GetQuery(sqlQueryGetAll)
	pageNum, _ := strconv.Atoi(c.DefaultQuery(ginQueryPage, ginDefaultPageStart))

	getAllQuery += serverutils.GetLimitQuery(pageNum*retrieveResultLimit, retrieveResultLimit)
	rows, err := db.Query(getAllQuery)
	defer rows.Close()

	contacts := make(map[int]entities.ContactResponseBody)
	phones := make(map[int]bool)
	addresses := make(map[int]bool)

	for rows.Next() {
		var contact entities.ContactResponseBody
		var address entities.AddressQuery
		var phone entities.PhoneQuery

		if err := rows.Scan(&contact.ContactID, &contact.FirstName, &contact.LastName,
			&address.AddressID, &address.Description, &address.City, &address.Street,
			&address.HomeNumber, &address.Apartment,
			&phone.PhoneID, &phone.Description, &phone.PhoneNumber); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, contacts)
		}

		if val, ok := contacts[contact.ContactID]; ok {
			contact = val
		}

		if address.AddressID.Valid == true && !keyExist(addresses, int(address.AddressID.Int32)) {
			responseAddress := parseAddressQueryToResponse(&address)
			updateAddress(&contact, addresses, responseAddress)
		}
		if phone.PhoneID.Valid == true && !keyExist(phones, int(phone.PhoneID.Int32)) {
			responsePhone := parsePhoneQueryToResponse(&phone)
			updatePhone(&contact, phones, responsePhone)
		}
		updateContact(contacts, &contact)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, contacts)
	}

	c.IndentedJSON(http.StatusOK, contacts)
}

func SearchContact(c *gin.Context) {
	db := setup.GetDBConn()
	getSearchQuery, _ := serverutils.GetQuery(sqlQueryGetAll)
	whereQuery, _ := serverutils.GetQuery(sqlQueryWhere)
	andQuery, _ := serverutils.GetQuery(sqlQueryAnd)
	isFirstNameSearch := false
	isLastNameSearch := false

	firstName, isFirstNameSearch := c.GetQuery(ginQueryFirstName)
	if isFirstNameSearch {
		getSearchQuery += whereQuery
		getSearchQuery += serverutils.AddValuesToQuery(firstNameFieldInDB, firstName)
	}
	lastName, isLastNameSearch := c.GetQuery(ginQueryLastName)
	if isLastNameSearch {
		if isFirstNameSearch {
			getSearchQuery += andQuery
		} else {
			getSearchQuery += whereQuery
		}
		getSearchQuery += serverutils.AddValuesToQuery(lastNameFieldInDB, lastName)
	}

	pageNum, _ := strconv.Atoi(c.DefaultQuery(ginQueryPage, ginDefaultPageStart))
	getSearchQuery += serverutils.GetLimitQuery(pageNum*retrieveResultLimit, retrieveResultLimit)

	rows, err := db.Query(getSearchQuery)
	defer rows.Close()

	contacts := make(map[int]entities.ContactResponseBody)
	phones := make(map[int]bool)
	addresses := make(map[int]bool)

	for rows.Next() {
		var contact entities.ContactResponseBody
		var address entities.AddressQuery
		var phone entities.PhoneQuery

		if err := rows.Scan(&contact.ContactID, &contact.FirstName, &contact.LastName,
			&address.AddressID, &address.Description, &address.City, &address.Street,
			&address.HomeNumber, &address.Apartment,
			&phone.PhoneID, &phone.Description, &phone.PhoneNumber); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, contacts)
		}

		if val, ok := contacts[contact.ContactID]; ok {
			contact = val
		}

		if address.AddressID.Valid == true && !keyExist(addresses, int(address.AddressID.Int32)) {
			responseAddress := parseAddressQueryToResponse(&address)
			updateAddress(&contact, addresses, responseAddress)
		}
		if phone.PhoneID.Valid == true && !keyExist(phones, int(phone.PhoneID.Int32)) {
			responsePhone := parsePhoneQueryToResponse(&phone)
			updatePhone(&contact, phones, responsePhone)
		}
		updateContact(contacts, &contact)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, contacts)
	}

	c.IndentedJSON(http.StatusOK, contacts)
}

func keyExist(m map[int]bool, key int) bool {
	return m[key]
}

func updatePhone(contact *entities.ContactResponseBody, phones map[int]bool, phone *entities.PhoneResponseBody) {
	*&contact.PhoneRes = append(contact.PhoneRes, *phone)
	updateRecordExist(phones, phone.PhoneID)
}

func updateAddress(contact *entities.ContactResponseBody, addresses map[int]bool, address *entities.AddressResponseBody) {
	*&contact.AddressRes = append(contact.AddressRes, *address)
	updateRecordExist(addresses, address.AddressID)
}

func updateContact(contacts map[int]entities.ContactResponseBody, contact *entities.ContactResponseBody) {
	contacts[contact.ContactID] = *contact
}

func updateRecordExist(recordMap map[int]bool, key int) {
	recordMap[key] = true
}

func parseAddressQueryToResponse(address *entities.AddressQuery) *entities.AddressResponseBody {
	var ret entities.AddressResponseBody

	ret.AddressID = int(address.AddressID.Int32)
	ret.Description = address.Description.String
	ret.City = address.City.String
	ret.Street = address.Street.String
	ret.HomeNumber = address.HomeNumber.String
	ret.Apartment = address.Apartment.String

	return &ret
}

func parsePhoneQueryToResponse(phone *entities.PhoneQuery) *entities.PhoneResponseBody {
	var ret entities.PhoneResponseBody

	ret.PhoneID = int(phone.PhoneID.Int32)
	ret.Description = phone.Description.String
	ret.PhoneNumber = phone.PhoneNumber.String

	return &ret
}

func CreateContact(c *gin.Context) {
	db := setup.GetDBConn()
	var newContact entities.ContactRequestBody

	if err := c.BindJSON(&newContact); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	insertContactQuery, _ := serverutils.GetQuery(sqlQueryInsertContact)

	result, err := db.Exec(insertContactQuery, newContact.FirstName, newContact.LastName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	newAddresses := newContact.AddressReq

	fmt.Println(newAddresses)

	contactID, err := result.LastInsertId()

	insertAddressQuery, _ := serverutils.GetQuery(sqlQueryInsertAddress)
	for _, address := range newAddresses {
		_, err = db.Exec(insertAddressQuery, contactID, address.Description, address.City,
			address.Street, address.HomeNumber, address.Apartment)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	}

	newPhones := newContact.PhoneReq

	insertPhoneQuery, _ := serverutils.GetQuery(sqlQueryInsertPhone)
	for _, phone := range newPhones {
		_, err = db.Exec(insertPhoneQuery, contactID, phone.Description, phone.PhoneNumber)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, contactID)
}

func DeleteContact(c *gin.Context) {
	db := setup.GetDBConn()

	id := c.Param(ginParamId)

	deleteContactQuery, _ := serverutils.GetQuery(sqlQueryDeleteContact)

	_, err := db.Exec(deleteContactQuery, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func EditContact(c *gin.Context) {
	db := setup.GetDBConn()
	contactID := c.Param(ginParamId)

	var editContact entities.ContactRequestBody

	if err := c.BindJSON(&editContact); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	editContactQuery, _ := serverutils.GetQuery(sqlQueryEditContact)
	editContactQuery += prepareContactUpdateQuery(&editContact)
	editContactQuery += getWhereCond(contactIdFieldInDB, contactID)

	db.Exec(editContactQuery)

	for _, phone := range editContact.PhoneReq {
		if !isPhoneExist(&phone) {
			insertPhoneQuery, _ := serverutils.GetQuery(sqlQueryInsertPhone)
			_, err := db.Exec(insertPhoneQuery, contactID, phone.Description, phone.PhoneNumber)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		} else {
			editPhoneQuery := preparePhoneUpdateQuery(&phone)
			editPhoneQuery += getWhereCond(phoneIdFieldInDB, fmt.Sprintf("%d", phone.PhoneID))
			_, err := db.Exec(editPhoneQuery)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		}
	}

	for _, address := range editContact.AddressReq {
		if !isAddressExist(&address) {
			insertAddressQuery, _ := serverutils.GetQuery(sqlQueryInsertAddress)
			_, err := db.Exec(insertAddressQuery, contactID, address.Description,
				address.City, address.Street,
				address.HomeNumber, address.Apartment)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		} else {
			editAddressQuery := prepareAddressUpdateQuery(&address)
			editAddressQuery += getWhereCond(addressIdFieldInDB, fmt.Sprintf("%d", address.AddressID))
			_, err := db.Exec(editAddressQuery)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		}
	}

	c.IndentedJSON(http.StatusCreated, contactID)
}

func prepareContactUpdateQuery(contact *entities.ContactRequestBody) string {
	var ret string
	var isSeparatorNeeded bool
	if contact.FirstName != "" {
		ret += serverutils.AddValuesToQuery(firstNameFieldInDB, contact.FirstName)
		isSeparatorNeeded = true
	}

	if contact.LastName != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
			isSeparatorNeeded = false
		}
		ret += serverutils.AddValuesToQuery(lastNameFieldInDB, contact.LastName)
	}

	return ret
}

func isPhoneExist(phone *entities.PhoneRequestBody) bool {
	return phone.PhoneID != 0
}

func isAddressExist(address *entities.AddressRequestBody) bool {
	return address.AddressID != 0
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
	}
	return ret
}

func prepareAddressUpdateQuery(address *entities.AddressRequestBody) string {
	ret, _ := serverutils.GetQuery(sqlQueryEditAddress)
	var isSeparatorNeeded bool
	if address.Description != "" {
		ret += serverutils.AddValuesToQuery(descriptionFieldInDB, address.Description)
		isSeparatorNeeded = true
	}
	if address.City != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
		}
		ret += serverutils.AddValuesToQuery(cityFieldInDB, address.City)
	}
	if address.Street != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
		}
		ret += serverutils.AddValuesToQuery(streetFieldInDB, address.Street)
	}
	if address.HomeNumber != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
		}
		ret += serverutils.AddValuesToQuery(homeNumberFieldInDB, address.HomeNumber)
	}
	if address.Apartment != "" {
		if isSeparatorNeeded {
			ret += sqlSeparatorValues
		}
		ret += serverutils.AddValuesToQuery(apartmentFieldInDB, address.Apartment)
	}

	return ret
}

func getWhereCond(fieldName string, id string) string {
	var ret string
	where, _ := serverutils.GetQuery(sqlQueryWhere)
	ret += where

	ret += serverutils.AddValuesToQuery(fieldName, id)

	return ret
}
