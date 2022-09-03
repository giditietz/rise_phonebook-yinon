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
	ContactID int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AddressQuery struct {
	AddressID   sql.NullInt32  `json:"addressId"`
	Description sql.NullString `json:"description"`
	City        sql.NullString `json:"city"`
	Street      sql.NullString `json:"street"`
	HomeNumber  sql.NullString `json:"home_number"`
	Apartment   sql.NullString `json:"apartment"`
}

type PhoneQuery struct {
	PhoneId     sql.NullInt32  `json:"phoneId"`
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
	ContactId   int    `json:"contact_id"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
}

type PhoneRequestBody struct {
	ContactId   int    `json:"contact_id"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
}

type ContactResponseBody struct {
	ContactID  int                   `json:"id"`
	FirstName  string                `json:"first_name"`
	LastName   string                `json:"last_name"`
	AddressRes []AddressResponseBody `json:"address"`
	PhoneRes   []PhoneResponseBody   `json:"phone"`
}

type AddressResponseBody struct {
	AddressID   int    `json:"addressId"`
	Description string `json:"description"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HomeNumber  string `json:"home_number"`
	Apartment   string `json:"apartment"`
}

type PhoneResponseBody struct {
	PhoneId     int    `json:"phoneId"`
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
			&phone.PhoneId, &phone.Description, &phone.PhoneNumber); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, contacts)
		}

		if val, ok := contacts[contact.ContactID]; ok {
			contact = val
		}

		if address.AddressID.Valid == true && !keyExist(addresses, int(address.AddressID.Int32)) {
			responseAddress := parseAddressQueryToResponse(&address)
			updateAddress(&contact, addresses, responseAddress)
		}
		if phone.PhoneId.Valid == true && !keyExist(phones, int(phone.PhoneId.Int32)) {
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
	updateRecordExist(phones, phone.PhoneId)
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

	ret.PhoneId = int(phone.PhoneId.Int32)
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

	contactId, err := result.LastInsertId()

	insertAddressQuery, _ := serverutils.GetQuery("insertAddress")
	for _, address := range newAddresses {
		_, err = db.Exec(insertAddressQuery, contactId, address.Description, address.City,
			address.Street, address.HomeNumber, address.Apartment)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	}

	newPhones := newContact.PhoneReq

	insertPhoneQuery, _ := serverutils.GetQuery("insertPhone")
	for _, phone := range newPhones {
		_, err = db.Exec(insertPhoneQuery, contactId, phone.Description, phone.PhoneNumber)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, contactId)
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
	// db := setup.GetDBConn()
	// var newContact ContactRequestBody

	// if err := c.BindJSON(&newContact); err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, err)
	// }

	// const createContactQuery string = "REPLACE INTO contacts(first_name, last_name) VALUES (?, ?);"
	// result, err := db.Exec(createContactQuery, newContact.FirstName, newContact.LastName)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// contactId, err := result.LastInsertId()

	// const addAddressQuery string = "REPLACE INTO addresses(contact_id, description, city, street, home_number, apartment) VALUES (?, ?, ?, ?, ?, ?)"

	// _, err = db.Exec(addAddressQuery, contactId, newContact.AddressReq.Description, newContact.AddressReq.City,
	// 	newContact.AddressReq.Street, newContact.AddressReq.HomeNumber, newContact.AddressReq.Apartment)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// const addPhoneQuery string = "REPLACE INTO phones(contact_id, description, phone_number) VALUES (?, ?, ?)"

	// _, err = db.Exec(addPhoneQuery, contactId, newContact.PhoneReq.Description, newContact.PhoneReq.PhoneNumber)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// c.IndentedJSON(http.StatusCreated, contactId)
}
