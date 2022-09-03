package server

import (
	"database/sql"
	"fmt"
	"net/http"
	serverutils "phonebook/server/server_utils"
	"phonebook/setup"

	"github.com/gin-gonic/gin"
)

type ContactQuery struct {
	ContactID int    `json:"ContactID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AddressQuery struct {
	AddressID   sql.NullInt32  `json:"addressID"`
	Description sql.NullString `json:"description"`
	City        sql.NullString `json:"city"`
	Street      sql.NullString `json:"street"`
	HomeNumber  sql.NullString `json:"home_number"`
	Apartment   sql.NullString `json:"apartment"`
}

type PhoneQuery struct {
	PhoneID     sql.NullInt32  `json:"phoneID"`
	Description sql.NullString `json:"description"`
	PhoneNumber sql.NullString `json:"PhoneNumber"`
}

type ContactRequestBody struct {
	FirstName  string               `json:"first_name"`
	LastName   string               `json:"last_name"`
	AddressReq []AddressRequestBody `json:"address"`
	PhoneReq   []PhoneRequestBody   `json:"phone"`
}

type AddressRequestBody struct {
	ContactID   int    `json:"contactID"`
	AddressID   int    `json:"addressID"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
}

type PhoneRequestBody struct {
	ContactID   int    `json:"contactID"`
	PhoneID     int    `json:"phoneID"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
}

type ContactResponseBody struct {
	ContactID  int                   `json:"contactID"`
	FirstName  string                `json:"firstName"`
	LastName   string                `json:"lastName"`
	AddressRes []AddressResponseBody `json:"address"`
	PhoneRes   []PhoneResponseBody   `json:"phone"`
}

type AddressResponseBody struct {
	AddressID   int    `json:"AddressID"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
}

type PhoneResponseBody struct {
	PhoneID     int    `json:"PhoneID"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
}

func GetAllContacts(c *gin.Context) {
	db := setup.GetDBConn()
	getAllQuery, _ := serverutils.GetQuery("getAll")

	rows, err := db.Query(getAllQuery)
	defer rows.Close()

	contacts := make(map[int]ContactResponseBody)
	phones := make(map[int]bool)
	addresses := make(map[int]bool)

	for rows.Next() {
		var contact ContactResponseBody
		var address AddressQuery
		var phone PhoneQuery

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

func updatePhone(contact *ContactResponseBody, phones map[int]bool, phone *PhoneResponseBody) {
	*&contact.PhoneRes = append(contact.PhoneRes, *phone)
	updateRecordExist(phones, phone.PhoneID)
}

func updateAddress(contact *ContactResponseBody, addresses map[int]bool, address *AddressResponseBody) {
	*&contact.AddressRes = append(contact.AddressRes, *address)
	updateRecordExist(addresses, address.AddressID)
}

func updateContact(contacts map[int]ContactResponseBody, contact *ContactResponseBody) {
	contacts[contact.ContactID] = *contact
}

func updateRecordExist(recordMap map[int]bool, key int) {
	recordMap[key] = true
}

func parseAddressQueryToResponse(address *AddressQuery) *AddressResponseBody {
	var ret AddressResponseBody

	ret.AddressID = int(address.AddressID.Int32)
	ret.Description = address.Description.String
	ret.City = address.City.String
	ret.Street = address.Street.String
	ret.HomeNumber = address.HomeNumber.String
	ret.Apartment = address.Apartment.String

	return &ret
}

func parsePhoneQueryToResponse(phone *PhoneQuery) *PhoneResponseBody {
	var ret PhoneResponseBody

	ret.PhoneID = int(phone.PhoneID.Int32)
	ret.Description = phone.Description.String
	ret.PhoneNumber = phone.PhoneNumber.String

	return &ret
}

func CreateContact(c *gin.Context) {
	db := setup.GetDBConn()
	var newContact ContactRequestBody

	if err := c.BindJSON(&newContact); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	insertContactQuery, _ := serverutils.GetQuery("insertContact")

	result, err := db.Exec(insertContactQuery, newContact.FirstName, newContact.LastName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	newAddresses := newContact.AddressReq

	fmt.Println(newAddresses)

	contactID, err := result.LastInsertId()

	insertAddressQuery, _ := serverutils.GetQuery("insertAddress")
	for _, address := range newAddresses {
		_, err = db.Exec(insertAddressQuery, contactID, address.Description, address.City,
			address.Street, address.HomeNumber, address.Apartment)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	}

	newPhones := newContact.PhoneReq

	insertPhoneQuery, _ := serverutils.GetQuery("insertPhone")
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

	id := c.Param("id")

	deleteContactQuery, _ := serverutils.GetQuery("deleteContact")

	_, err := db.Exec(deleteContactQuery, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func EditContact(c *gin.Context) {
	db := setup.GetDBConn()
	contactID := c.Param("id")

	var editContact ContactRequestBody

	if err := c.BindJSON(&editContact); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	editContactQuery, _ := serverutils.GetQuery("editContact")
	editContactQuery += prepareContactUpdateQuery(&editContact)
	editContactQuery += getWhereCond("contact_id", contactID)

	db.Exec(editContactQuery)

	for _, phone := range editContact.PhoneReq {
		if !isPhoneExist(&phone) {
			query, _ := serverutils.GetQuery("insertPhone")
			_, err := db.Exec(query, contactID, phone.Description, phone.PhoneNumber)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		} else {
			editPhoneQuery := preparePhoneUpdateQuery(&phone)
			editPhoneQuery += getWhereCond("contact_id", contactID)
			fmt.Println("query: ", editPhoneQuery)
			_, err := db.Exec(editPhoneQuery)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		}
	}

	for _, address := range editContact.AddressReq {
		if !isPhoneExist(&phone) {
			query, _ := serverutils.GetQuery("insertPhone")
			_, err := db.Exec(query, contactID, phone.Description, phone.PhoneNumber)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		} else {
			editPhoneQuery := preparePhoneUpdateQuery(&phone)
			editPhoneQuery += getWhereCond("contact_id", contactID)
			fmt.Println("query: ", editPhoneQuery)
			_, err := db.Exec(editPhoneQuery)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		}
	}

	c.IndentedJSON(http.StatusCreated, contactID)
}

func prepareContactUpdateQuery(contact *ContactRequestBody) string {
	var ret string
	if contact.FirstName != "" {
		ret += serverutils.AddValuesToQuery("first_name", contact.FirstName)
	}

	if contact.LastName != "" {
		ret += serverutils.AddValuesToQuery(", last_name", contact.LastName)
	}

	return ret
}

func isPhoneExist(phone *PhoneRequestBody) bool {
	return phone.PhoneID != 0
}

func isAddressExist(address *AddressRequestBody) bool {
	return address.AddressID != 0
}

func preparePhoneUpdateQuery(phone *PhoneRequestBody) string {
	ret, _ := serverutils.GetQuery("editPhone")
	if phone.Description != "" {
		ret += serverutils.AddValuesToQuery("description", phone.Description)
	}
	if phone.PhoneNumber != "" {
		ret += serverutils.AddValuesToQuery(", phone_number", phone.PhoneNumber)
	}
	return ret
}

func getWhereCond(fieldName string, id string) string {
	var ret string
	where, _ := serverutils.GetQuery("where")
	ret += where

	ret += serverutils.AddValuesToQuery(fieldName, id)

	return ret
}
