package server

import (
	"database/sql"
	"net/http"
	"phonebook/setup"

	"github.com/gin-gonic/gin"
)

type Address struct {
	AddressID   sql.NullInt32  `json:"addressId"`
	Description sql.NullString `json:"description"`
	City        sql.NullString `json:"city"`
	Street      sql.NullString `json:"street"`
	HomeNumber  sql.NullString `json:"homeNumber"`
	Apartment   sql.NullString `json:"apartment"`
}

type Phone struct {
	PhoneId     sql.NullInt32  `json:"phoneId"`
	Description sql.NullString `json:"description"`
	PhoneNumber sql.NullString `json:"PhoneNumber"`
}

type Contact struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Address   []Address `json:"Address"`
	Phone     []Phone   `json:"Phone"`
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

type ContactRequestBody struct {
	FirstName  string             `json:"first_name"`
	LastName   string             `json:"last_name"`
	AddressReq AddressRequestBody `json:"address"`
	PhoneReq   PhoneRequestBody   `json:"phone"`
}

func GetAllContacts(c *gin.Context) {
	db := setup.GetDBConn()
	const query string = "SELECT * FROM contacts LEFT JOIN addresses USING (contact_id) LEFT JOIN phones USING (contact_id)"

	rows, err := db.Query(query)
	defer rows.Close()

	contacts := make(map[int]Contact)
	phones := make(map[sql.NullInt32]bool)
	addresses := make(map[sql.NullInt32]bool)

	for rows.Next() {
		var contact Contact
		var address Address
		var phone Phone

		if err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName,
			&address.AddressID, &address.Description, &address.City, &address.Street,
			&address.HomeNumber, &address.Apartment,
			&phone.PhoneId, &phone.Description, &phone.PhoneNumber); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, contacts)
		}
		if val, ok := contacts[contact.ID]; ok {
			if _, ok := phones[phone.PhoneId]; !ok {
				val.Phone = append(val.Phone, phone)
				phones[phone.PhoneId] = true
			}
			if _, ok := addresses[address.AddressID]; !ok {
				val.Address = append(val.Address, address)
				addresses[address.AddressID] = true
			}
			contacts[contact.ID] = val
		} else {
			if address.AddressID.Valid == true {
				contact.Address = append(contact.Address, address)
			}
			if phone.PhoneId.Valid == true {
				contact.Phone = append(contact.Phone, phone)
			}
			contacts[contact.ID] = contact
			phones[phone.PhoneId] = true
			addresses[address.AddressID] = true
		}
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, contacts)
	}

	c.IndentedJSON(http.StatusOK, contacts)
}

func CreateContact(c *gin.Context) {
	db := setup.GetDBConn()
	var newContact ContactRequestBody

	if err := c.BindJSON(&newContact); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	const createContactQuery string = "INSERT INTO contacts(first_name, last_name) VALUES (?, ?);"
	result, err := db.Exec(createContactQuery, newContact.FirstName, newContact.LastName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	contactId, err := result.LastInsertId()

	const addAddressQuery string = "INSERT INTO addresses(contact_id, description, city, street, home_number, apartment) VALUES (?, ?, ?, ?, ?, ?)"

	_, err = db.Exec(addAddressQuery, contactId, newContact.AddressReq.Description, newContact.AddressReq.City,
		newContact.AddressReq.Street, newContact.AddressReq.HomeNumber, newContact.AddressReq.Apartment)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	const addPhoneQuery string = "INSERT INTO phones(contact_id, description, phone_number) VALUES (?, ?, ?)"

	_, err = db.Exec(addPhoneQuery, contactId, newContact.PhoneReq.Description, newContact.PhoneReq.PhoneNumber)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, contactId)
}

func DeleteContact(c *gin.Context) {
	db := setup.GetDBConn()

	id := c.Param("id")

	const query string = "DELETE FROM contacts WHERE contact_id = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
